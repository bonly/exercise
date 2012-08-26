#include <boost/asio.hpp>
#include <boost/bind.hpp>

class client
{
  public:
    client():_socket(_io){}
    void init()
    {
      boost::asio::ip::tcp::endpoint ep(
         boost::asio::ip::address::from_string("127.0.0.1"),38989);
      boost::system::error_code er;
      _socket.connect (ep,er);
      if (er)
      {
         perror("sock");
      }
      _socket.send (boost::asio::buffer("test",5));
      _socket.close();
    }
  private:
    boost::asio::io_service _io;
    boost::asio::ip::tcp::socket _socket;
};

class server
{
  public:
    server(boost::asio::io_service &io):_io(io),_socket(_io),_acceptor(_io)
    { }
    void init()
    {
       boost::asio::ip::tcp::endpoint ep(
           boost::asio::ip::address::from_string("127.0.0.1"),38989);
       _acceptor.open (ep.protocol());
       _acceptor.set_option (boost::asio::ip::tcp::acceptor::reuse_address( true ) );
       _acceptor.bind (ep);
       _acceptor.listen ();
       _acceptor.async_accept (_socket,
              boost::bind (&server::on_accepted, this, 
              boost::asio::placeholders::error));
    }
    void on_accepted( boost::system::error_code const &e )
    {
       if (!e)
         std::cerr << "accept a client\n";
    }
  private:
    boost::asio::io_service &_io;
    boost::asio::ip::tcp::socket _socket;
    boost::asio::ip::tcp::acceptor _acceptor;
};

int 
main (int argc, char * argv[])
{
  if(argc>=2)
  {
    boost::asio::io_service io;
    server sr(io);
    sr.init();
    io.run();
  }
  else
  {
    client sr;
    sr.init();
  }
  return 0; 
}

