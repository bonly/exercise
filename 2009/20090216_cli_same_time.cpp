#define _WIN32_WINNT 0x0501
#define __USE_W32_SOCKETS
#include <iostream>
#include <string>
#include <boost/array.hpp>
#include <boost/asio.hpp>
#include <boost/thread.hpp>
using namespace std;
using namespace boost;
using namespace boost::asio;
//不断向服务器发送数据的同时
//通过本地输入同时发送数据到服务器

void run()
{
	while(true)
	{
		io_service io;
		ip::udp::endpoint ep(
				ip::address::from_string("127.0.0.1"),9837);

		ip::udp::socket socket(io);
		socket.open(ip::udp::v4());
		socket.send_to(buffer("from client"),ep);

		ip::udp::endpoint rec_ep;
		char buf[244];
		bzero (buf,244);
		socket.receive_from(buffer(buf),rec_ep);
		cout << "recv: "<<buf<<endl;
		sleep(1);
	}
}
void send(char* pch)
{
	io_service io;
	ip::udp::endpoint ep(
			ip::address::from_string("127.0.0.1"),9837);

	ip::udp::socket socket(io);
	socket.open(ip::udp::v4());
	socket.send_to(buffer(pch),ep);

	char buf[244];
	bzero (buf,244);
	socket.receive_from(buffer(buf),ep);
	cout << "recv form "<<ep.address() << " :"<<buf<<endl;
}
void control()
{
	while(true)
	{
		fd_set rset;
		FD_ZERO(&rset);
		for(;;)
		{
			FD_SET(fileno(stdin),&rset);//cin不支持
			select(fileno(stdin)+1,&rset,NULL,NULL,NULL);
			if(FD_ISSET(fileno(stdin),&rset))
			{
				char cmd[255];
				bzero(cmd,255);
				cin.getline(cmd,255);
				send(cmd);
			}
		}

		sleep(2);
	}
}
int
main ()
{
  thread thr1(run);
  //thread thr2(run);
  thread thr3(control);
  thr1.join();
  //thr2.join();
  thr3.join();

	return 0;

}

