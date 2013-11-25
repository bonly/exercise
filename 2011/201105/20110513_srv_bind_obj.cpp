/**
 * @file 20110513_srv_bind_obj.cpp
 * @brief 
 * @author bonly
 * @date 2013年10月30日 bonly Created
 */
//忽略未使用定义的警告
#pragma GCC diagnostic ignored "-Wunused-local-typedefs"
#pragma GCC diagnostic ignored "-Wunused-variable"
#include <boost/asio.hpp>
#include <boost/thread.hpp>
#include <boost/bind.hpp>
#include <iostream>

namespace Bus{
using namespace boost;
using namespace boost::asio;
using namespace boost::asio::ip;
using boost::bind;
using boost::thread;
using namespace std;

typedef boost::function<void (const boost::system::error_code&)> CB;
class server{
public:
  server(io_service &io):_io(io),_socket(_io),_acceptor(_io){
  }

  void Accept(CB &cb){
      tcp::endpoint ep(address::from_string("127.0.0.1"), 8989);
      _acceptor.open(ep.protocol());
      _acceptor.bind(ep);
      _acceptor.listen();
      _acceptor.async_accept(_socket,
              cb
              );
  }
private:
  io_service &_io;
  tcp::socket _socket;
  tcp::acceptor _acceptor;
};

class srv{
public:
    io_service& Io(){
        return _io;
    }
    void Run(){
       while(true){
//           std::clog << "running...\n";
           _io.poll();
       }
    }
    void Start(){
        _thr = thread(bind(&srv::Run,this));
        _thr.join();
    }
private:
    io_service _io;
    thread _thr;
};
}
using Bus::srv;
using Bus::server;

class MyC{
public:
    void call_back(const boost::system::error_code& ec){
        std::clog << "accept a client" << std::endl;
    }
};

int main(int argc, char* argv[]){
    srv mysrv;

    server sv(mysrv.Io());

    MyC obj;
    Bus::CB back = (boost::bind(&MyC::call_back, &obj, boost::asio::placeholders::error)); ///@note 需要error的参数
    sv.Accept(back);

    mysrv.Start();

    return 0;
}

/**
 * asio中的handle定义如下:
A free function as an accept handler:

void accept_handler(
    const boost::system::error_code& ec)
{
  ...
}
An accept handler function object:

struct accept_handler
{
  ...
  void operator()(
      const boost::system::error_code& ec)
  {
    ...
  }
  ...
};
A non-static class member function adapted to an accept handler using bind():

void my_class::accept_handler(
    const boost::system::error_code& ec)
{
  ...
}
...
acceptor.async_accept(...,
    boost::bind(&my_class::accept_handler,
      this, boost::asio::placeholders::error));

 */


