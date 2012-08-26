//============================================================================
// Name        : process_thread.cpp
// Author      : bonly
// Version     :
// Copyright   : bonly's copyright notice
// Description : Hello World in C++, Ansi-style
//============================================================================
#define _WIN32_WINNT 0x0501
#define __USE_W32_SOCKETS
#include <boost/asio.hpp>
#include <boost/array.hpp>
#include <boost/bind.hpp>
#include <iostream>
using namespace std;
using namespace boost;
using namespace boost::asio;

class Srv
{
	public:
		Srv(io_service &io)
		:_io(io),_socket(_io,
				       ip::udp::endpoint(
				    		 ip::udp::v4(),9837))
    {start_recive();}
	private:
		void start_recive()
		{
			_socket.async_receive_from(
					buffer(_recv_buffer),_remote_endpoint,
					bind(&Srv::handle_receive,this,
							placeholders::error,
							placeholders::bytes_transferred));
		}
		void handle_receive(const boost::system::error_code& error,size_t)
		{
			if (!error || error == error::message_size)
			{
				cerr << "recv msg: "<<_recv_buffer<<endl;
				shared_ptr<string> message(new string("this is a test"));
				_socket.async_send_to(buffer(*message),_remote_endpoint,
				bind(&Srv::handle_send,this,message,
						placeholders::error,
						placeholders::bytes_transferred));
				start_recive();
			}
		}
		void handle_send(shared_ptr<string> msg,
				             const boost::system::error_code& ,
				             size_t)
		{
		}

	private:
		io_service &_io;
		ip::udp::socket _socket;
		ip::udp::endpoint _remote_endpoint;
		char _recv_buffer[244];
};
void run()
{
	io_service io;
	ip::udp::socket socket(io,ip::udp::endpoint(ip::udp::v4(),9837));
	for (;;)
	{
		char buf[244];
		ip::udp::endpoint ep;
		boost::system::error_code error;
		socket.receive_from(buffer(buf),ep,0,error);
		if (error&&error!=error::message_size)
			throw boost::system::system_error(error);
		cerr << "recv: "<<buf <<endl;

		boost::system::error_code ignor_err;
		socket.send_to(buffer("from server"),ep,0,ignor_err);
	}
}
int main() {
	try
	{
		io_service io;
		io_service::work work(io);
		Srv srv(io);
		io.run();

    //run();
	}
	catch(std::exception& e)
	{
		std::cerr << e.what() << std::endl;
	}
	return 0;
}

