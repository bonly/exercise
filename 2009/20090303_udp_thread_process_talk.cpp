//============================================================================
// Name        : process_thread.cpp
// Author      : bonly
// Version     :
// Copyright   : bonly's copyright notice
// Description : Hello World in C++, Ansi-style
//============================================================================
//此程序是线程中创建进程
#define _WIN32_WINNT 0x0501
#define __USE_W32_SOCKETS
#include <boost/asio.hpp>
#include <boost/array.hpp>
#include <boost/bind.hpp>
#include <boost/thread.hpp>
#include <iostream>
#include <list>
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
      	if (pid=vfork()==0)
      	{
      		execlp("process_thread_cli.exe","process_thread_cli.exe",(char*)0);
      		cerr << "start child fail\n";
      		exit(0);
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
  //app.init_child();
  //thread thr1(bind(&Application::talk,ref(app)));
  thread thr2(bind(&Application::init_child,ref(app)));
  //thr1.join();
  thr2.join();
  //wait(NULL);
	return 0;
}

