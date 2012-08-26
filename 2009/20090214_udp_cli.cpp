#define _WIN32_WINNT 0x0501
#define __USE_W32_SOCKETS
#include <iostream>
#include <boost/array.hpp>
#include <boost/asio.hpp>
using namespace std;
using namespace boost;
using namespace boost::asio;

int
main ()
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

	return 0;

}

