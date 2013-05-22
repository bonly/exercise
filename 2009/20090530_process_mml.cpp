/*
 * head.hpp
 *
 *  Created on: Aug 31, 2009
 *      Author: bonly
 */

#ifndef HEAD_HPP_
#define HEAD_HPP_
#include <iostream>
using namespace std;

#include <boost/bind.hpp>
using namespace boost;
#include <boost/asio.hpp>
using namespace boost::asio;


#endif /* HEAD_HPP_ */

/*
 * deal.hpp
 *
 *  Created on: Aug 31, 2009
 *      Author: bonly
 */

#ifndef DEAL_HPP_
#define DEAL_HPP_
#include "head.hpp"

class OCS_Stream
{
  public:
    OCS_Stream (ip::tcp::socket::native_type sock, io_service &io)
      :_io(io),_stream(io,ip::tcp::v4(),sock),_timer(io)
    {
    }
    ~OCS_Stream()
    {
      close();
    }

    void close(){ _stream.close();}
    void timeout(){cerr << "cecv timeout!\n";_stream.close();}

    int send(const char* buf, const int len)
    {
      _stream.async_send(buffer(buf,len),boost::bind(&OCS_Stream::on_recv,this,placeholders::error));
      return 0;
    }

    void on_recv(const system::error_code &error)
    {
      if(error)
      {
        cerr << "recv error: "<< error.message() << "\n";
        close();
        return;
      }
      else
      {
        std::istream is(&_buf);
        string line;
        getline(is, line);
        _timer.cancel(); //cancel timeout after recveive cmd
        //cerr << line ;
      }
    }

    int recv()
    {
      async_read_until(_stream,_buf,'\n',
          boost::bind(&OCS_Stream::on_recv,this,placeholders::error));

      _timer.expires_from_now(posix_time::seconds(3));
      _timer.async_wait(bind(&OCS_Stream::timeout,this));
      return 0;
    }

    int buf_str(string &str)
    {
      std::istream is(&_buf);
      is >> str;
      //cerr <<"str is: "<< str << "\n";
      return str.length();
    }
  private:
    io_service&      _io;
    ip::tcp::socket  _stream;
    asio::streambuf  _buf;
    deadline_timer   _timer;
};

class MML_Stream
{
  public:
    MML_Stream (io_service &io)
     :_io(io),_stream(io)
     {}

    int connect(const char* ip, const int port)
    {
      try
      {
        ip::tcp::endpoint ep(
            ip::address::address::from_string(ip),port);
        system::error_code error;
        _stream.connect(ep,error);
        if (error)
        {
          cerr << "connect failure: " << error.message() << "\n";
          return -1;
        }
      }
      catch(...)
      {
        cerr << "connect failure\n";
      }
      return 0;
    }

    bool is_open(){return _stream.is_open();}
    int send(const char* buf)
    {
      return asio::write(_stream,buffer(buf,strlen(buf)));
    }
  private:
    io_service&     _io;
    ip::tcp::socket _stream;
};

class Deal
{
  public:
    Deal():_mml(_io){}
    int run(int sock);
  private:
    io_service _io;
    MML_Stream _mml;

};

#endif /* DEAL_HPP_ */
/*
 * deal.cpp
 *
 *  Created on: Aug 31, 2009
 *      Author: bonly
 */
#include "deal.hpp"

int
Deal::run(int sock)
{
  OCS_Stream ocs(sock,_io);

  while(true)
  {
    ocs.recv();
    //OCS::parse();

    if(!_mml.is_open())
      _mml.connect("127.0.0.1",9999);

    string msg;
    ocs.buf_str(msg);
    cerr << "send: "<< msg <<endl;
    _mml.send("klkj;q");

    _io.run();
  }
  return 0;
}

int
main(int argc, char* argv[])
{
  Deal deal;
  deal.run(atoi(argv[1]));
  return 0;
}

