#include <stdio.h>
#include <cstdlib>
#include <iostream>
#include <boost/thread.hpp>
#include <boost/aligned_storage.hpp>
#include <boost/array.hpp>
#include <boost/bind.hpp>
#include <boost/enable_shared_from_this.hpp>
#include <boost/noncopyable.hpp>
#include <boost/shared_ptr.hpp>
#include <boost/asio.hpp>

using boost::asio::ip::tcp;

class handler_allocator
	: private boost::noncopyable
{
public:
	handler_allocator()
		: in_use_(false)
	{
	}

	void* allocate(std::size_t size)
	{
		if (!in_use_ && size < storage_.size)
		{
			in_use_ = true;
			return storage_.address();
		}
		else
		{
			return ::operator new(size);
		}
	}

	void deallocate(void* pointer)
	{
		if (pointer == storage_.address())
		{
			in_use_ = false;
		}
		else
		{
			::operator delete(pointer);
		}
	}

private:
	// Storage space used for handler-based custom memory allocation.
	boost::aligned_storage<1024> storage_;

	// Whether the handler-based custom allocation storage has been used.
	bool in_use_;
};

template <typename Handler>
class custom_alloc_handler
{
public:
	custom_alloc_handler(handler_allocator& a, Handler h)
		: allocator_(a),
		handler_(h)
	{
	}

	template <typename Arg1>
	void operator()(Arg1 arg1)
	{
		handler_(arg1);
	}

	template <typename Arg1, typename Arg2>
	void operator()(Arg1 arg1, Arg2 arg2)
	{
		handler_(arg1, arg2);
	}

	friend void* asio_handler_allocate(std::size_t size,
		custom_alloc_handler<Handler>* this_handler)
	{
		return this_handler->allocator_.allocate(size);
	}

	friend void asio_handler_deallocate(void* pointer, std::size_t /*size*/,
		custom_alloc_handler<Handler>* this_handler)
	{
		this_handler->allocator_.deallocate(pointer);
	}

private:
	handler_allocator& allocator_;
	Handler handler_;
};

// Helper function to wrap a handler object to add custom allocation.
template <typename Handler>
inline custom_alloc_handler<Handler> make_custom_alloc_handler(
	handler_allocator& a, Handler h)
{
	return custom_alloc_handler<Handler>(a, h);
}

/// A pool of io_service objects.
class io_service_pool
	: private boost::noncopyable
{
public:
	/// Construct the io_service pool.
	explicit io_service_pool(std::size_t pool_size) : next_io_service_(0) 
	{
		if (pool_size == 0)
			throw std::runtime_error("io_service_pool size is 0");

		// Give all the io_services work to do so that their run() functions will not
		// exit until they are explicitly stopped.
		for (std::size_t i = 0; i < pool_size; ++i)
		{
			io_service_ptr io_service(new boost::asio::io_service);
			work_ptr work(new boost::asio::io_service::work(*io_service));
			io_services_.push_back(io_service);
			work_.push_back(work);
		}
	}

	/// Run all io_service objects in the pool.
	void run()
	{
		// Create a pool of threads to run all of the io_services.
		std::vector<boost::shared_ptr<boost::thread> > threads;
		for (std::size_t i = 0; i < io_services_.size(); ++i)
		{
			boost::shared_ptr<boost::thread> thread(new boost::thread(
				boost::bind(&boost::asio::io_service::run, io_services_[i])));
			threads.push_back(thread);
		}

		// Wait for all threads in the pool to exit.
		for (std::size_t i = 0; i < threads.size(); ++i)
			threads[i]->join();
	}

	/// Stop all io_service objects in the pool.
	void stop()
	{
		// Explicitly stop all io_services.
		for (std::size_t i = 0; i < io_services_.size(); ++i)
			io_services_[i]->stop();
	}

	/// Get an io_service to use.
	boost::asio::io_service& get_io_service()
	{
		// Use a round-robin scheme to choose the next io_service to use.
		boost::asio::io_service& io_service = *io_services_[next_io_service_];
		++next_io_service_;
		if (next_io_service_ == io_services_.size())
			next_io_service_ = 0;
		return io_service;
	}

private:
	typedef boost::shared_ptr<boost::asio::io_service> io_service_ptr;
	typedef boost::shared_ptr<boost::asio::io_service::work> work_ptr;

	/// The pool of io_services.
	std::vector<io_service_ptr> io_services_;

	/// The work that keeps the io_services running.
	std::vector<work_ptr> work_;

	/// The next io_service to use for a connection.
	std::size_t next_io_service_;
};

class session
	: public boost::enable_shared_from_this<session>
{
public:
	session(boost::asio::io_service& work_service, boost::asio::io_service& io_service)
		: socket_(io_service)
		, io_work_service(work_service)
	{
	}

	tcp::socket& socket()
	{
		return socket_;
	}

	void start()
	{
		socket_.async_read_some(boost::asio::buffer(data_),
			make_custom_alloc_handler(allocator_,
			boost::bind(&session::handle_read,
			shared_from_this(),
			boost::asio::placeholders::error,
			boost::asio::placeholders::bytes_transferred)));
	}

	void handle_read(const boost::system::error_code& error,
		size_t bytes_transferred)
	{
		if (!error)
		{
			boost::shared_ptr<std::vector<char> > buf(new std::vector<char>);

			buf->resize(bytes_transferred);
			std::copy(data_.begin(), data_.begin() + bytes_transferred, buf->begin());
			io_work_service.post(boost::bind(&session::on_receive, shared_from_this(), buf, bytes_transferred));

			socket_.async_read_some(boost::asio::buffer(data_),
				make_custom_alloc_handler(allocator_,
				boost::bind(&session::handle_read,
				shared_from_this(),
				boost::asio::placeholders::error,
				boost::asio::placeholders::bytes_transferred)));
		}
	}

	void handle_write(const boost::system::error_code& error)
	{
		if (!error)
		{
		}
	}

	void on_receive(boost::shared_ptr<std::vector<char> > buffers, size_t bytes_transferred)
	{
		char* data_stream = &(*buffers->begin());
		// in here finish the work.
		std::cout << "receive :" << bytes_transferred << " bytes." << "message :" << data_stream << std::endl;
	}

private:
	// The io_service used to finish the work.
	boost::asio::io_service& io_work_service;

	// The socket used to communicate with the client.
	tcp::socket socket_;

	// Buffer used to store data received from the client.
	boost::array<char, 1024> data_;

	// The allocator to use for handler-based custom memory allocation.
	handler_allocator allocator_;
};

typedef boost::shared_ptr<session> session_ptr;

class server
{
public:
	server(short port, std::size_t io_service_pool_size)
		: io_service_pool_(io_service_pool_size)
		, io_service_work_pool_(io_service_pool_size)
		, acceptor_(io_service_pool_.get_io_service(), tcp::endpoint(tcp::v4(), port))
	{
		session_ptr new_session(new session(io_service_work_pool_.get_io_service(), io_service_pool_.get_io_service()));
		acceptor_.async_accept(new_session->socket(),
			boost::bind(&server::handle_accept, this, new_session,
			boost::asio::placeholders::error));
	}

	void handle_accept(session_ptr new_session,
		const boost::system::error_code& error)
	{
		if (!error)
		{
			new_session->start();
			new_session.reset(new session(io_service_work_pool_.get_io_service(), io_service_pool_.get_io_service()));
			acceptor_.async_accept(new_session->socket(),
				boost::bind(&server::handle_accept, this, new_session,
				boost::asio::placeholders::error));
		}
	}

	void run()
	{
		io_thread_.reset(new boost::thread(boost::bind(&io_service_pool::run, &io_service_pool_)));
		work_thread_.reset(new boost::thread(boost::bind(&io_service_pool::run, &io_service_work_pool_)));
	}

	void stop()
	{
		io_service_pool_.stop();
		io_service_work_pool_.stop();

		io_thread_->join();
		work_thread_->join();
	}

private:
	boost::shared_ptr<boost::thread> io_thread_;
	boost::shared_ptr<boost::thread> work_thread_;
	io_service_pool io_service_pool_;
	io_service_pool io_service_work_pool_;
	tcp::acceptor acceptor_;
};

int main(int argc, char* argv[])
{
	try
	{
		if (argc != 2)
		{
			std::cerr << "Usage: server <port>\n";
			return 1;
		}

		using namespace std; // For atoi.
		server s(atoi(argv[1]), 10);

		s.run();

		getchar();

		s.stop();
	}
	catch (std::exception& e)
	{
		std::cerr << "Exception: " << e.what() << "\n";
	}

	return 0;
}

/*
使用io_service作为处理工作的work pool，可以看到，就是通过io_service.post投递一个Handler到io_service的队列,Handler在这个io_service.run内部得到执行，有可能你会发现，io_services.dispatch的接口也和io_service.post一样，但不同的是它是直接调用而不是经过push到队列然后在io_services.run中执行，而在这个示例当中，显然我们需要把工作交到另一个线程去完成，这样才不会影响网络接收线程池的工作以达到高效率的接收数据，这种设计与前面的netsever其实相同，这就是典型的Half Sync/Half Async。二者的区别就是netsever自己实现了工作队列，而不是直接使用io_service，这种设计实际上在win下是使用了iocp作为工作队列。

 不过我更倾向于前一种设计，因为那样做，代码一切都在自己的掌握中，而io_service则是经过许多封装代码，并且本身设计只是用于处理网络完成事件的。

无论如何使用，都能感觉到使用boost.asio实现服务器，不尽是一件非常轻松的事，而且代码很漂亮，逻辑也相当清晰，这点上很不同于ACE。


*/