#include <iostream>
#include <boost/bind.hpp>
#include <boost/asio.hpp>
#include <boost/lexical_cast.hpp>
#include <boost/format.hpp>
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
       //wait(); 这里也不行
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
           if (execlp("srv","srv",lexical_cast<string>(_socket.native()).c_str())<0)
             cerr << "start srv fail\n";
           exit(0); //exec系列函数是直接代替进程，因此这里的代码在成功时是不会再被执行的
         }
         else
         {
           cerr << "Father aliving\n";
           _socket.close();
           //_acceptor.async_accept(_socket,
           // bind (&Srv::on_accepted, this,placeholders::error));
	   // wait();//放这里不行
           //waitpid(pid,NULL,0); //放这里，第二个连接会不正常
           myaccept();
         }
         //wait(); //放这里时子进程结束后在Father结束前是杀不死的僵尸
	 //waitpid(pid,NULL,0); //放这里，第二个连接会不正常
       }
    }
    void run()
    {
       //wait(); //不行
       _io.run();
       //wait(); //放这里时子时子进程是可杀的僵尸 
    }
  public:
    io_service _io;
    ip::tcp::socket _socket;
    ip::tcp::acceptor _acceptor;
};

class Client
{
   public:
    Client():_socket(_io)
    { bzero(buf,244);}
    void init(int sock)
    {
       _socket.assign(ip::tcp::v4(),sock);
       on_recv();
    }
    void on_recv()
    {
       _socket.async_receive(buffer(buf),
            bind (&Client::echo, this,placeholders::error));
    }
    void echo(boost::system::error_code const &e)
    {
       if(!e)
       {
         _socket.send(buffer(buf));
         bzero(buf,244);
         on_recv();
       }
    }
    void run(){_io.run();}
   public:
    io_service _io;
    ip::tcp::socket _socket;
    char buf[244];
};

int
main(int argc, char* argv[])
{
  if (argc <=1)
  {
    Srv server;
    server.init();
    server.run();
    cerr << "this is father end\n"; //这里后面的都没到达
    //wait();
  }
  else
  {
    cerr << "this is child being\n";
    Client client;
    int sock = lexical_cast<int>(argv[1]);
    cout << format("param is %s\n")%argv[1];
    client.init(sock);
    client.run();
    cerr << "this is child end\n";
    _exit(0);
  }
  return 0;
}

