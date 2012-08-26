#include <iostream>
#include <boost/asio.hpp>
#include <boost/bind.hpp>
using namespace boost;
using namespace boost::asio;
using namespace std;
char buf[255];

void on_accp(const boost::system::error_code& error)
{
  if(!error)
  {
     cerr << "accept a client: "  << endl;
  }
}

int
main()
{
  io_service io;
  ip::tcp::endpoint lo(ip::address::from_string("127.0.0.1"),9837);
  ip::tcp::acceptor accept(io);
  accept.open(lo.protocol());
  accept.bind(lo);
  accept.listen();

  ip::tcp::socket stream(io);

  accept.async_accept(stream,
                      bind(&on_accp,placeholders::error));
  io.run();
  return 0;
}