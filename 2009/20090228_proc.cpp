//============================================================================
// Name        : process_thread.cpp
// Author      : bonly
// Version     :
// Copyright   : bonly's copyright notice
// Description : Hello World in C++, Ansi-style
//============================================================================
//#define _WIN32_WINNT 0x0501
//#define __USE_W32_SOCKETS
#include <boost/asio.hpp>
#include <boost/array.hpp>
#include <boost/bind.hpp>
#include <boost/thread.hpp>
#include <iostream>
#include <list>
#include <unistd.h>
using namespace std;
using namespace boost;
using namespace boost::asio;

class Application
{
  public:
   void init_child()
   {
      for(int i=0; i<3; ++i)
      {
        pid_t pid;
        if (fork()==0)
        {
           execlp("/bin/ls","/bin/ls");
        }
        _process.push_back(pid);
      }
   }
   void talk()
   {
      io_service io;
      ip::udp::socket socket(io,ip::udp::endpoint(ip::udp::v4(),2897));
      for(;;)
      {
        char buf[244];
        ip::udp::endpoint ep;
        boost::system::error_code error;
        socket.receive_from(buffer(buf),ep,0,error);
        if (error&&error!=error::message_size)
                exit(0);
        cerr << "recv from "<<ep.address()
             << ":" << ep.port()<<" : "<<buf<<endl;
        socket.send_to(buffer("job"),ep,0,error);
      }
      io.run();
   }
  public:
    list<string> _job;
    list<pid_t>  _process;
};

int main()
{
  Application app;
  thread thr1(bind(&Application::talk,ref(app)));
  thread thr2(bind(&Application::init_child,ref(app)));
  thr1.join();
  thr2.join();

  //if(fork()==0)
  //{
  //  execlp("/bin/ls","ls");
  //}
  wait(NULL);
  return 0;
}

/*
aCC -AA mainproc.cpp -o bp -L ~/boost_1_37_0/stage/lib -lboost_thread-mt -mt +DD64 -lboost_system-mt -g +W2461 +W2236 +W2111
*/

