//============================================================================
// Name        : try_process.cpp
// Author      : bonly
// Version     :
// Copyright   : bonly's copyright notice
// Description : Hello World in C++, Ansi-style
//============================================================================
//#define __USE_W32_SOCKETS
//#define _WIN32_WINNT 0x0501
#include <iostream>
#include <string>
#include <sys/wait.h>
#include <boost/format.hpp>
#include <boost/asio.hpp>
//#include <boost/thread.hpp>

using namespace std;
using namespace boost;
using namespace boost::asio;
class Child
{
  public:
    Child():_stream(_io)
    {
       _stream.connect (
                ip::tcp::endpoint(
                    ip::address::from_string("127.0.0.1"),
                    2988));
     }
     ~Child(){_stream.close();exit(0);}
     void run()
     {
       _stream.send(buffer("this is from child\n",20));
     }
  private:
    io_service _io;
    ip::tcp::socket _stream;
};
int main(int argc, char* argv[])
{
  io_service pio;
  ip::tcp::acceptor acceptor(pio,ip::tcp::endpoint(ip::tcp::v4(),2988));
  ip::tcp::socket socket(pio);

  int chpid;
  if((chpid=fork())==0)
  {
     Child ch;
     ch.run();
  }
  acceptor.accept(socket);
  char buf[1024];
  socket.read_some(buffer(buf,1024));
  cout << format("get message: %s\n")%buf;
  socket.close();

  int status;
  wait(&status);

  return 0;
}

/*
aCC -AA +W2236 proc.cpp  -L /home/hejb/boost_1_37_0/stage/lib/ -l boost_system-mt-1_37 -mt
*/

