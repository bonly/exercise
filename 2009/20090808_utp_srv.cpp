#include <iostream>
#include <boost/asio.hpp>
#include <boost/bind.hpp>  //必须加这个以保证用的是boost的bind
using namespace boost;
using namespace boost::asio;
using namespace std;
char buf[255];

void on_recv(const boost::system::error_code& error,size_t)
{
  if(!error||error==error::message_size)
  {
     cerr << "recv msg: " << buf << endl;
  }
}

int
main()
{
  io_service io;
  ip::udp::socket stream(io,ip::udp::endpoint(ip::udp::v4(),9837));
  ip::udp::endpoint remote;

  stream.async_receive_from(buffer(buf),
                            remote,
                            bind(&on_recv,placeholders::error,
                                          placeholders::bytes_transferred));
  io.run();
  return 0;
}