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
        cerr << _buf ;
      }
      _stream.async_read_some(buffer(_buf,sizeof(_buf)),
          boost::bind(&OCS_Stream::on_recv,this,placeholders::error));
    }

    int recv()
    {
      bzero (_buf,2048);
      _stream.async_read_some(buffer(_buf,sizeof(_buf)),
          boost::bind(&OCS_Stream::on_recv,this,placeholders::error));
      _timer.expires_from_now(posix_time::seconds(3));
      _timer.async_wait(bind(&OCS_Stream::timeout,this));
      return 0;
    }

  private:
    io_service&     _io;
    ip::tcp::socket _stream;
    deadline_timer  _timer;
    char            _buf[2048];
};

class Deal
{
  public:
    int run(int sock);
  private:
    io_service _io;

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
  cerr << "begin\n";
  OCS_Stream ocs(sock,_io);
  ocs.recv();
  _io.run();
  cerr << "after run\n";
  return 0;
}

int
main(int argc, char* argv[])
{
  Deal deal;
  deal.run(atoi(argv[1]));
  return 0;
}
