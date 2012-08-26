//============================================================================
// Name        : interplat.cpp
// Author      : bonly
// Version     :
// Copyright   : bonly's copyright notice
// Description : Hello World in C++, Ansi-style
//============================================================================
#define BOOST_SYSTEM_NO_LIB
#define BOOST_THREAD_NO_LIB
#include <boost/asio.hpp>
#include <boost/thread.hpp>
#include <iostream>
using namespace std;
using namespace boost::asio;
using namespace boost;

class Stream
{
  public:
    Stream(io_service &io):_io(io),_stream(io)
    {

    }
    ip::tcp::socket& stream(){return _stream;}
  public:
    io_service &_io;
    ip::tcp::socket _stream;
};
class Acceptor
{
  public:
    Acceptor(io_service &io):_io(io),_accept(io)
    {
      ip::tcp::endpoint lo(ip::address::from_string("127.0.0.1"),9800);
      _accept.open(lo.protocol());
      _accept.bind(lo);
      _accept.listen();

      Stream stream(io);
      _accept.async_accept(stream.stream(),
            bind(&Acceptor::on_accp,this,placeholders::error));
    }
    void on_accp(const boost::system::error_code& error)
    {
      if(!error)
      {
        cerr << "accept a client"<<endl;
      }
    }

  private:
    io_service &_io;
    ip::tcp::acceptor _accept;
};
int main()
{
  io_service io;
  Acceptor accp(io);
  thread thr(bind(&io_service::run,ref(io)));
  thr.join();
	return 0;
}
