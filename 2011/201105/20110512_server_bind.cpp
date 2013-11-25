/**
 * @file 20110512_server_bind.cpp
 * @brief 
 * @author bonly
 * @date 2013年10月29日 bonly Created
 */
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

void call_back(const boost::system::error_code& ec){
    std::clog << "accept a client" << std::endl;
}

int main(int argc, char* argv[]){
    srv mysrv;

    server sv(mysrv.Io());

    Bus::CB back = call_back;
    sv.Accept(back);

    mysrv.Start();

    return 0;
}

