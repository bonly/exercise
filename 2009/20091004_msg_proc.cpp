//============================================================================
// Name        : msg_proc.cpp
// Author      : bonly
// Version     :
// Copyright   : bonly's copyright notice
// Description : Hello World in C++, Ansi-style
//============================================================================
#include <boost/asio.hpp>
#include <boost/shared_ptr.hpp>
#include <boost/thread.hpp>
#include <boost/bind.hpp>
#include <iostream>
#include <queue>
using namespace std;
using namespace boost;

asio::io_service io;
asio::ip::tcp::socket  stream(io);

enum {MAX_MSG_LEN=10};
queue<shared_ptr<char[MAX_MSG_LEN]> > msg;
char tmp[MAX_MSG_LEN]={0};
void getdata();
int connect();
void data();

int connect()
{
  asio::ip::tcp::endpoint ep(asio::ip::address::from_string("127.0.0.1"),20201);
  stream.connect(ep);
  return 0;
}
void data()
{
  char *od = new char[MAX_MSG_LEN];
  memcpy(od,tmp,MAX_MSG_LEN);
  std::clog << tmp << std::endl;
  getdata();
  delete od;
}
void getdata()
{
  //stream.async_read_some(asio::buffer(tmp),bind(data));
  stream.async_receive(asio::buffer(tmp),bind(data));
}


int main()
{
  connect ();
  getdata();
  thread tio(bind(&asio::io_service::run, &io));
  tio.join();
  return 0;
}
