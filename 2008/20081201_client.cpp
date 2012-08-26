/*
 * App.h
 *
 *  Created on: 2008-11-11
 *      Author: Bonly
 */
//#define __USE_W32_SOCKETS
//#define _WIN32_WINNT 0x0501
#ifndef APP_H_
#define APP_H_
#include <boost/smart_ptr.hpp>
#include <boost/asio.hpp>
#include <boost/bind.hpp>
#include <iostream>
#include <boost/date_time/c_time.hpp>
#include <boost/lexical_cast.hpp>
#include <boost/format.hpp>
namespace jmb_test
{
using namespace std;
struct head
{
  char blockMark[3];
  char verID[1];
  char mesgLen[6];
  char appTradeCode[8];
  char starAddr[12];
  char destAddr[12];
  char mesgPurp[1];
  char outForm[1];
  char mesgID[20];
  char mesgReqNo[20];
  char workDate[8];
  char sentTime[14];
  char expTime[4];
  char deliTime[6];
  char mesgPRI[1];
  char reserve[20];
  char finalMark[1];
  head()
  {
    memset(this,0x20,sizeof(head));
  }
  void init()
  {
    memcpy(blockMark,"{1:",3);
    memcpy(verID,"0",1);
    memcpy(mesgPRI,"9",1);
    memcpy(finalMark,"}",1);

    time_t now=time(0);
    tm * t_now=localtime(&now);
    char cdate[14+1];
    strftime(cdate,14+1,"%Y%m%d%H%M%S",t_now);
    strncpy(sentTime,cdate,14);
    strncpy(workDate,cdate,8);
  }
};
struct body
{
  string blockMark;
  string TAG;
  string fin;
  string msg;

        body():blockMark(3,' '),TAG(1,' '),fin(1,' ')
        {}
  void init()
  {
    blockMark="{2:";
    TAG=":";
    fin="}";
  }
};
struct tail
{
  char blockMark[3];
  char MAC[32];
  char fin[1];
  tail()
  {
    memset(this,' ',sizeof(tail));
  }
  void init()
  {
    strncpy(blockMark,"{C:",sizeof(blockMark));
    strncpy(fin,"}",sizeof(fin));
  }
};

class decode
{
  public:
    decode(string code):finish(false)
    {
      code.copy((char*)&hd,sizeof(hd));
      code.copy((char*)&tl,sizeof(tl),code.length()-sizeof(tl));
    }
  public:
    head hd;
    body bd;
    tail tl;
    bool finish;
};
class bss_decode:public decode
{
  public:
    bss_decode(string code):decode(code)
    {
      std::cout << "bss_decode"<< std::endl;
    }
    bool parse()
    {
      bool head=false;
      if (strcmp(hd.blockMark,"{1:")==0)
      {
        std::cout << "recive head, begin...\n";
        if (strcmp(hd.finalMark,"}")==0)
        {
          std::cout << "recive head, end.\n";
          head=true;
        }
      }

      bool tail=false;
      if (head)
      {
        if (strcmp(tl.blockMark,"{C:")==0)
        {
          std::cout << "recive tail, begin...\n";
          if (strcmp(tl.fin,"}")==0)
          {
            std::cout << "recive tail, end.\n";
            tail =true;
          }
        }
      }

      bool result = false;
      if (tail)
      {
        //TODO parse body
        return true;
      }
      return result;
    }
};

class Connect
{
  public:
    Connect(boost::asio::io_service &io_svc, boost::asio::ip::tcp::endpoint & ep)
      :io(io_svc),_socket(io)
    {
      try
      {
        boost::system::error_code er;
        _socket.connect(ep,er);
        if(er)
        {
          std::cout << "connect failure, errcode: "<< er.message() << "\n";
          return;
        }

        std::cout << "\t connect ok\n";
        head hd;hd.init();
        boost::shared_ptr<body> bd(new body);bd->init();
        tail tl;tl.init();

        try
        {
          bd->msg.append(":ECD:收付费企业代码");
          bd->msg.append("ECD:收付费企业代码");
          bd->msg.append(":JBR:经办人编号");
          bd->msg.append(":CLZ:业务流水");
          bd->msg.append(":WD0:工作日期");
          bd->msg.append(":CCH:场次号");
          bd->msg.append(":8ED:费项代码");
          bd->msg.append(":EBN:企业开户行行号");
          bd->msg.append(":EAC:企业托收账号");
          bd->msg.append(":JFH:缴费号/合同号");
        }
        catch(std::exception &e)
        {
          std::cerr << e.what();
        }

        std::string msg=bd->blockMark + bd->TAG + bd->msg + bd->fin;
                                int leng=sizeof(hd)+msg.size()+sizeof(tl);
        strncpy(hd.mesgLen,(boost::format("%1$6s")%leng).str().c_str(),sizeof(hd.mesgLen));

                                std::cout << "head: " << (char*)&hd << std::endl;
        std::cout << "send msg: "<<msg << std::endl;
                                std::cout << "tail: " << (char*)&tl << std::endl;

        _socket.send(boost::asio::buffer(&hd,sizeof(hd)));
        _socket.send(boost::asio::buffer(msg,msg.size()));
        _socket.send(boost::asio::buffer(&tl,sizeof(tl)));
        _socket.close();
      }
      catch(boost::system::error_code & e)
      {
        std::cerr << e << std::endl;
      }
    }
    boost::asio::ip::tcp::socket& socket(){return _socket;}
  private:
    boost::asio::io_service & io;
    boost::asio::ip::tcp::socket _socket;
};

class Stream : public boost::enable_shared_from_this<Stream>
{
  public:
    Stream(boost::asio::io_service & io):_socket(io)
    {
      std::cout << "Stream had been created\n";
    }
    ~Stream()
    {
      std::cout << "destroy Stream\n";
      _socket.close();
    }
    boost::asio::ip::tcp::socket& socket(){return _socket;}

    bool start()
    {
      std::cout << "start to community with client\n";
      _socket.async_receive (
        boost::asio::buffer(str,sizeof(str)),
        boost::bind(&Stream::handle_read,shared_from_this(),boost::asio::placeholders::error)
      );

      return false;
    }


    void handle_read(const boost::system::error_code &error)
    {
      if (!error)
      {
        pack.append(str);
        bss_decode de(pack);
        bool finish=de.parse();
        if (!finish)
          _socket.async_receive (
          boost::asio::buffer(str,sizeof(str)),
          boost::bind(&Stream::handle_read,shared_from_this(),boost::asio::placeholders::error)
                );
        else
        {
          std::cout << "finish"<< endl;
        }
      }
    }

  private:
    boost::asio::ip::tcp::socket _socket;
    char str[1024];
    std::string pack;
};
class App
{
  public:
    App();
    virtual ~App();
    int listen();
    int connect();
    int run();
  private:
    boost::asio::io_service io;
    //boost::shared_ptr<Listen> acceptor;
    boost::shared_ptr<Connect> connector;
};

}

#endif /* APP_H_ */

/*
 * App.cpp
 *
 *  Created on: 2008-11-11
 *      Author: Bonly
 */

#include "App.h"
namespace jmb_test
{

App::App()
{
  // TODO Auto-generated constructor stub
  std::cout << "start test program \n";
}

App::~App()
{
  // TODO Auto-generated destructor stub
  std::cout << "end test program \n";
}
/*
int
App::listen()
{
  boost::asio::ip::tcp::endpoint ep(
      boost::asio::ip::address::from_string("192.168.0.1"),8989
      );
  acceptor = boost::shared_ptr<Listen>(new Listen(io,ep));
  boost::shared_ptr<Stream> p;
  boost::system::error_code e;
  acceptor->run(e,p);
  return 0;
}
*/
int
App::run()
{
  return io.run();
}
int
App::connect()
{
  boost::asio::ip::tcp::endpoint ep(
      boost::asio::ip::address::from_string("192.168.0.1"),8989
      );
  connector = boost::shared_ptr<Connect>(new Connect(io,ep));

  return 0;
}
}


boost::shared_ptr<jmb_test::App> app;

int
main ()
{
  app = boost::shared_ptr<jmb_test::App>(new jmb_test::App());
  //app->listen();
  app->connect();
  app->run();
  return EXIT_SUCCESS;
}


