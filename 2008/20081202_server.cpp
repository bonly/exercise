/*
 * App.h
 *
 *  Created on: 2008-11-11
 *      Author: Bonly
 */
#define __USE_W32_SOCKETS
#define _WIN32_WINNT 0x0501
#ifndef APP_H_
#define APP_H_
#include <boost/smart_ptr.hpp>
#include <boost/asio.hpp>
#include <boost/bind.hpp>
#include <iostream>
#include <boost/date_time/c_time.hpp>
#include <boost/lexical_cast.hpp>
#include <boost/format.hpp>
#include <boost/algorithm/string.hpp>
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
		memset(this,0x20,sizeof(head)); //0是结束符,空格是十六进制0x20
	}
	void init()
	{
		memcpy(blockMark,"{1:",3);
		memcpy(verID,"0",1);
		memcpy(mesgPRI,"9",1);
		memcpy(finalMark,"}",1);

		time_t now=time(0);
		tm * t_now=localtime(&now);

		char cdate[14+1]; //struct中可以在函数中加入新的变量,不影响直接发送出去,因成员数据的位置不会变化
		//strftime(cdate,14+1,"%Y%m%d%H%M%S",gmtime(&now));
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
	{
	}
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
  	memset(this,' ',sizeof(tail)); //空格字符赋值也就是0x20
  	//memset(this,'\40',sizeof(tail)); //单引号中用的是八进制空格为40
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
    decode(string icode):finish(false),code(icode)
    {
    	try
    	{
    		code.copy((char*)&hd,sizeof(hd));
    		code.copy((char*)&tl,sizeof(tl),code.length()-sizeof(tl));
    		//std::cout << "\t" << (char*)&hd << "\n";
    	}
    	catch(...)
    	{

    	}
    }
	public:
		head hd;
		body bd;
		tail tl;
		bool finish;
		string code;
};
class bss_decode:public decode
{
	public:
		bss_decode(string code):decode(code)
		{
			//std::cout << "bss_decode"<< std::endl;
		}
		bool parse()
		{
			bool head=false;
			if (strncmp(hd.blockMark,"{1:",sizeof(hd.blockMark))==0)
			{
				//std::cout << "recive head, begin...\n";
				//std::cout << "\t" << (char*)&hd << "\n";
				if (strncmp(hd.finalMark,"}",sizeof(hd.finalMark))==0)
				{
					//std::cout << "receive head, end.\n";
					head=true;
				}
			}

			bool tail=false;
			if (head)
			{
				//cout << "tail is: "<<(char*)&tl << endl;
				if (strncmp(tl.blockMark,"{C:",sizeof(tl.blockMark))==0)
				{
					//std::cout << "recive tail, begin...\n";
					if (strncmp(tl.fin,"}",sizeof(tl.fin))==0)
					{
						//std::cout << "receive tail, end.\n";
						tail =true;
					}
				}
			}

			bool result = false;
			if (tail)
			{
				std::cout << "receive head: " << (char*)&hd << "\n";

				//TODO parse body
				bd.blockMark=code.substr (sizeof(hd), bd.blockMark.size());
				bd.TAG=code.substr (sizeof(hd)+bd.blockMark.size(),bd.TAG.size());
				try
				{
					std::string c_len(hd.mesgLen,6);
					boost::trim(c_len);
					int msg_leng=boost::lexical_cast<int>(c_len);

					bd.msg=code.substr(sizeof(hd)+bd.blockMark.size()+bd.TAG.size(),
						msg_leng-
							(sizeof(hd)+bd.blockMark.size()+bd.TAG.size()+bd.fin.size()+sizeof(tl))
					);
					bd.fin=code.substr (msg_leng-bd.fin.size(), bd.fin.size());
				}
				catch(...)
				{
					std::cerr << "信息长度位不正确\n" ;
				}

				std::cout << "receive body: " << bd.blockMark << bd.TAG << bd.msg << bd.fin << "\n";
				std::cout << "receive msg: " << bd.msg << "\n";
				std::cout << "receive tail: " << (char*)&tl << "\n";

				result = true;
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
					bd->msg="";
				}
				catch(std::exception &e)
				{
					std::cerr << e.what();
				}

				std::string msg=bd->blockMark + bd->TAG + bd->msg + bd->fin;
				int leng=sizeof(hd)+msg.size()+sizeof(tl);
				strncpy(hd.mesgLen,(boost::format("%1$6s")%leng).str().c_str(),sizeof(hd.mesgLen));
//				strncpy(hd.mesgLen,
//						(boost::lexical_cast<string>(leng)).c_str(),
//						sizeof(hd.mesgLen));

				std::cerr << "send head: "<< (char*)&hd << "\n";
				std::cerr << "send msg: "<<msg;
				std::cerr << "send tail: "<< (char*)&tl << "\n";

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
			//_socket.async_receive (
			boost::asio::async_read(
					_socket,
					boost::asio::buffer(str,sizeof(str)),
					boost::bind(&Stream::handle_read,shared_from_this(),boost::asio::placeholders::error)
			);
			return false;
		}


		void handle_read(const boost::system::error_code &error)
		{
			if (!error)
			{
//				if (memcmp(str,"",1)==0)
//				  pack.append(" ");
//				else  //如果发送时是用(\20)空格符,不是用结束符(\0),则可以直接append到字符串中
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
		char str[1];
		std::string pack;
};

class Listen
{
	public:
		Listen (boost::asio::io_service &io_svc,boost::asio::ip::tcp::endpoint &endpoint)
		  :io(io_svc),acc(io_svc,endpoint)
		{
			std::cout << "\t create listener\n";
		}

		int run(boost::shared_ptr<Stream> stream=NULL,const boost::system::error_code & error)
		{
			if (error) std::cerr << error << std::endl;

			if (stream!=NULL)
				stream->start();

			stream = boost::shared_ptr<Stream>(new Stream(io));
			acc.async_accept(
					stream->socket(),
					boost::bind(&Listen::run,this,stream,boost::asio::placeholders::error)
					);
			return 0;
		}

	private:
		boost::asio::io_service &io;
		boost::asio::ip::tcp::acceptor acc;
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
		boost::shared_ptr<Listen> acceptor;
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

int
App::listen()
{
	boost::asio::ip::tcp::endpoint ep(
			boost::asio::ip::address::from_string("192.168.0.1"),8989
			);
  acceptor = boost::shared_ptr<Listen>(new Listen(io,ep));
  boost::shared_ptr<Stream> p;
  boost::system::error_code e;
  acceptor->run(p,e);
  return 0;
}

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
	std::cout << "start...\n";
  app = boost::shared_ptr<jmb_test::App>(new jmb_test::App());
  app->listen();
  //app->connect();
  app->run();
	return EXIT_SUCCESS;
}

