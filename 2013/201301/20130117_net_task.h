#include <boost/asio/ip/tcp.hpp>
#include <boost/asio/placeholders.hpp>
#include <boost/asio/write.hpp>
#include <boost/asio/read.hpp>
#include <boost/shared_ptr.hpp>
#include <boost/function.hpp>
#include <boost/enable_shared_from_this.hpp>
#include <boost/make_shared.hpp>
#include <boost/lexical_cast.hpp>
#include <boost/bind.hpp>
#include <iostream>

#include "20130116_timer_task.h"

class tcp_connection_ptr{
	boost::shared_ptr<boost::asio::ip::tcp::socket> _socket;

public:
	explicit tcp_connection_ptr(
		boost::shared_ptr<boost::asio::ip::tcp::socket> socket)
	: _socket(socket){}

	explicit tcp_connection_ptr(
		boost::asio::io_service& ios,
		const boost::asio::ip::tcp::endpoint& endpoint)
	: _socket(boost::make_shared<boost::asio::ip::tcp::socket>(boost::ref(ios))){
		_socket->connect(endpoint);
	}

	template<typename Functor>
	void async_write(
		const boost::asio::const_buffers_1& buf, const Functor& func)const{
		boost::asio::async_write(*_socket, buf, func);
	}

	template<typename Functor>
	void async_write(
		const boost::asio::mutable_buffers_1& buf, const Functor& func)const{
		boost::asio::async_write(*_socket, buf, func);
	}

	template<typename Functor>
	void async_read(
		const boost::asio::mutable_buffers_1& buf,
		const Functor& func,
		std::size_t at_least_bytes) const{
		boost::asio::async_read(
			*_socket, buf, boost::asio::transfer_at_least(at_least_bytes), func);
	}

	void shutdown() const{
		_socket->shutdown(boost::asio::ip::tcp::socket::shutdown_both);
		_socket->close();
	}
};

namespace detail{
class tcp_listener : public boost::enable_shared_from_this<tcp_listener>{
	typedef boost::asio::ip::tcp::acceptor acceptor_t;
	acceptor_t _acceptor;
	boost::function<void(tcp_connection_ptr)> _func;

public:
	template<class Functor>
	tcp_listener(
		boost::asio::io_service& io_service,
		unsigned short port,
		const Functor& task_unwrapped)
	: _acceptor(io_service, boost::asio::ip::tcp::endpoint(
		boost::asio::ip::tcp::v4(), port)),
	  _func(task_unwrapped){}

	void stop(){
		_acceptor.close();
	}

	void push_task(){
		if (!_acceptor.is_open()){
			return;
		}

    	typedef boost::asio::ip::tcp::socket socket_t;
		boost::shared_ptr<socket_t> socket = boost::make_shared<socket_t>(
			boost::ref(_acceptor.get_io_service()));

		_acceptor.async_accept(*socket, boost::bind(
			&tcp_listener::handle_accept,
			this->shared_from_this(),
			tcp_connection_ptr(socket),
			boost::asio::placeholders::error));
	}

private:
	void handle_accept(
		const tcp_connection_ptr& new_connection,
		const boost::system::error_code& error){
		push_task(); // start another acceptor
		if (!error){
			make_task_wrapped(boost::bind(_func, new_connection))();//run task
		}else{
			std::cerr << error << std::endl;
		}
	}
};
}


namespace tp_network{
class tasks_processor : public tp_timers::tasks_processor{
	typedef std::map<
	    unsigned short,
	    boost::shared_ptr<detail::tcp_listener>
	    > listeners_map_t;
	listeners_map_t _listeners;

public:
	static tasks_processor& get(){
		static tasks_processor proc;
		return proc;
	}

	template<typename Functor>
	void add_listener(unsigned short port_num, const Functor& func){
		listeners_map_t::const_iterator it = _listeners.find(port_num);
		if (it != _listeners.end()){
			throw std::logic_error(
				"Such listener for port '" 
				+ boost::lexical_cast<std::string>(port_num)
				+ "' already created");
		}
		_listeners[port_num] = boost::make_shared<detail::tcp_listener>(boost::ref(_ios), port_num, func);
		_listeners[port_num]->push_task(); //begin accepting
	}

	void remove_listener(unsigned short port_num){
		listeners_map_t::iterator it = _listeners.find(port_num);
		if (it == _listeners.end()){
			throw std::logic_error(
				"No listener for port '"
			   + boost::lexical_cast<std::string>(port_num)
			   + "' created");
		}
		(*it).second->stop();
		_listeners.erase(it);
	}

	tcp_connection_ptr create_connection(const char* addr, const unsigned short port_num){
		return tcp_connection_ptr(_ios, boost::asio::ip::tcp::endpoint(
			boost::asio::ip::address_v4::from_string(addr), port_num));
	}
};
}
