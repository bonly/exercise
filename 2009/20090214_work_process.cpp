#include <iostream>
#include <boost/bind.hpp>
#include <boost/asio.hpp>
#include <boost/lexical_cast.hpp>
#include <boost/format.hpp>
#include <sys/wait.h>
using namespace std;
using namespace boost;
using namespace boost::asio;

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
         //cerr << "error_code: "<<e<<endl;
	 if(strlen(buf)>0)
         {
           _socket.send(buffer(buf));
           bzero(buf,244);
           on_recv();
	 }
       }
       else
       {  
          _socket.send(buffer("recv err message!\n"));
          _socket.close();
       }
    }
    void run(){_io.run();}
   public:
    io_service _io;
    ip::tcp::socket _socket;
    char buf[244];
};

int
main (int argc, char* argv[])
{
    printf("worker [%d] begin \n",getpid());
    Client client;
    int sock = lexical_cast<int>(argv[1]);
    client.init(sock);
    client.run();
    printf("worker [%d] end\n",getpid());

    return 0;
}

