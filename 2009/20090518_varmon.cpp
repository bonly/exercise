#include <string>
#include <iostream>
#include <boost/asio.hpp>
#include <boost/thread.hpp>
using namespace std;

#define BOOST_ASIO_FILE_MONITOR_ERROR				0x0000
#define BOOST_ASIO_FILE_MONITOR_CREATE				0x0040
#define BOOST_ASIO_FILE_MONITOR_CHANGE_NAME			0x0001
#define BOOST_ASIO_FILE_MONITOR_CHANGE_SIZE			0x0008
#define BOOST_ASIO_FILE_MONITOR_DELETE				0x0040

class var_monitor_impl
{
public:
	void destroy()
	{
		// When destory() is called watch() must return immediately. The current implementation
		// doesn't do it as ReadDirectoryChangesW() will block and won't return before a file
		// change is detected. This means that the core I/O service the file monitor is attached to
		// must not be destroyed when there is an outstanding asynchronous operation (a program
		// will seem to be unresponsive as ~basic_file_monitor_service() joins the thread and waits
		// for watch() to return).
	}

	template <typename Char>
	int watch(const std::basic_string<Char> &filename, int event_mask, boost::system::error_code &ec)
	{
     return 0;
	}

private:
};

template <typename VarMonitorImplementation = var_monitor_impl>
class basic_var_monitor_service : public boost::asio::io_service::service
{
public:
	static boost::asio::io_service::id id;

	explicit basic_var_monitor_service(boost::asio::io_service &io_service)
		: boost::asio::io_service::service(io_service), work_(new boost::asio::io_service::work(work_io_service_)),
		work_thread_(boost::bind(&boost::asio::io_service::run, &work_io_service_))
	{
	}

	~basic_var_monitor_service()
	{
		// The worker thread will finish when work_ is reset as all asynchronous operations
		// have been aborted and were discarded before (in destroy).
		work_.reset();

		// Event processing is stopped to discard queued operations.
		work_io_service_.stop();

		// The worker thread is joined to make sure the file monitor service is destroyed
		// _after_ the thread is finished (in case the last blocked asynchronous operation
		// has not yet been aborted completely).
		work_thread_.join();
	}

	typedef boost::shared_ptr<VarMonitorImplementation> implementation_type;

	void construct(implementation_type &impl)
	{
		impl.reset(new VarMonitorImplementation());
	}

	void destroy(implementation_type &impl)
	{
		// Currently blocked asynchronous operations are aborted.
		impl->destroy();

		// Queued asynchronous operations are discarded.
		impl.reset();
	}

	template <typename Char>
	int watch(implementation_type &impl, const std::basic_string<Char> &filename, int event_mask, boost::system::error_code &ec)
	{
		return impl->watch(filename, event_mask, ec);
	}

	template <typename Char, typename Handler>
	class watch_operation
	{
	public:
		watch_operation(implementation_type &impl, const std::basic_string<Char> &filename, int event_mask, boost::asio::io_service &io_service, Handler handler)
			: impl_(impl), filename_(filename), event_mask_(event_mask), io_service_(io_service), work_(io_service), handler_(handler)
		{
		}

		void operator()() const
		{
			implementation_type impl = impl_.lock();
			if (!impl)
			{
				this->io_service_.post(boost::asio::detail::bind_handler(handler_, boost::asio::error::operation_aborted, 0));
				return;
			}

			boost::system::error_code ec;
			int file_monitor_event = impl->watch(filename_, event_mask_, ec);
			this->io_service_.post(boost::asio::detail::bind_handler(handler_, ec, file_monitor_event));
		}

	private:
		boost::weak_ptr<VarMonitorImplementation> impl_;
		std::basic_string<Char> filename_;
		int event_mask_;
		boost::asio::io_service &io_service_;
		boost::asio::io_service::work work_;
		Handler handler_;
	};

	template <typename Char, typename Handler>
	void async_watch(implementation_type &impl, const std::basic_string<Char> &filename, int event_mask, Handler handler)
	{
		work_io_service_.post(watch_operation<Char, Handler>(impl, filename, event_mask, this->get_io_service(), handler));
	}

private:
	void shutdown_service()
	{
	}

	boost::asio::io_service work_io_service_;
	boost::scoped_ptr<boost::asio::io_service::work> work_;
	boost::thread work_thread_;
};

template <typename VarMonitorImplementation>
boost::asio::io_service::id basic_var_monitor_service<VarMonitorImplementation>::id;


template <typename Service>
class var_monitor : public boost::asio::basic_io_object<Service>
{
public:
	explicit var_monitor(boost::asio::io_service &io_service)
		: boost::asio::basic_io_object<Service>(io_service)
	{
	}

	template <typename Char>
	int watch(std::basic_string<Char> filename, int event_mask)
	{
		return this->service.watch(this->implementation, filename, event_mask);
	}

	template <typename Char, typename Handler>
	void async_watch(std::basic_string<Char> filename, int event_mask, Handler handler)
	{
		this->service.async_watch(this->implementation, filename, event_mask, handler);
	}
};

typedef var_monitor<basic_var_monitor_service<> > vm;

#define BOOST_TEST_MODULE native_socket_test
#include <boost/test/included/unit_test.hpp>

void watch_handle(const boost::system::error_code &ec, int file_monitor_event)
{
	BOOST_CHECK_EQUAL(ec, boost::system::error_code());
	BOOST_CHECK_EQUAL(file_monitor_event, BOOST_ASIO_FILE_MONITOR_CHANGE_SIZE);
}

BOOST_AUTO_TEST_CASE (hostname1)
{
	boost::asio::io_service io_service;

	vm mon(io_service);
	mon.async_watch<char>("C:\\test.txt", BOOST_ASIO_FILE_MONITOR_CHANGE_SIZE, boost::bind(watch_handle, boost::asio::placeholders::error, _2));

	io_service.run();
}

/*
 * aCC -AA +DD64 -L ~/boost_1_37_0/stage/lib/ -l boost_thread-mt-1_37 -l boost_system-mt-1_37 mymon.cpp -o mon -mt
 */

