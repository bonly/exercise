#include <iostream>
#include <boost/bind.hpp>
#include <boost/asio.hpp>
#include <boost/lexical_cast.hpp>
#include <boost/format.hpp>
#include <boost/thread.hpp>
#include <boost/bind.hpp>
#include <sys/wait.h>
using namespace std;
using namespace boost;
using namespace boost::asio;

class Srv
{
  public:
    Srv():_socket(_io),_acceptor(_io){}
    void init()
    {
       ip::tcp::endpoint ep(
           ip::address::from_string("127.0.0.1"),9837);
       _acceptor.open(ep.protocol());
       _acceptor.set_option(ip::tcp::acceptor::reuse_address(true));
       _acceptor.bind(ep);
       _acceptor.listen();
       _acceptor.async_accept(_socket,
            bind (&Srv::on_accepted, this,placeholders::error));
       myaccept();
    }

    void myaccept()
    {
       _acceptor.async_accept(_socket,
            bind (&Srv::on_accepted, this,placeholders::error));
    }
    void on_accepted(boost::system::error_code const &e)
    {
       if(!e)
       {
         cout << "Accept a client\n";
         pid_t pid;
         if ((pid=fork())==0)
         {
           _acceptor.close();
           char szSock[22];
           memset(szSock,0,22);
           strcpy(szSock,lexical_cast<string>(_socket.native()).c_str());
           cerr << "param is: "<<szSock<<endl;
           if (execlp("work","work",szSock)<0)
             cerr << "start work fail\n";
           cerr << "after fail exec\n"; //exec是代替程序，不出错是不会再执行后面的内容
           exit(0);
         }
       }
       printf("acceptor [%d] aliving\n",getpid());
       _socket.close();
       myaccept();
       //wait(NULL);
    }
    void run()
    {
       _io.run();
    }
  public:
    io_service _io;
    ip::tcp::socket _socket;
    ip::tcp::acceptor _acceptor;
};

void wait_job()
{
  while(true)
    wait(NULL);
}
int
main()
{
    
    printf("acceptor [%d] begin\n",getpid());
    Srv server;
    server.init();
    thread thr(bind(&Srv::run,ref(server)));
    thread wjob(bind(&wait_job));
    wjob.join();
    thr.join();
    return 0;
}

/*
g++ -g -L ~/boost_1_38_0/stage/lib/ -lboost_system-gcc32-mt-1_38 accp.cpp -o srv -pthread -l boost_thread-gcc32-mt
*/
