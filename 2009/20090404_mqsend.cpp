//============================================================================
// Name        :
// Author      : bonly
// Version     :
// Copyright   :
// Description : 生命周期数据结构
//============================================================================
#ifndef __DATA_HPP__
#define __DATA_HPP__

//用户
struct BF_SUBSCRIBER
{
  long long          SUBS_ID;
  long long          ACC_NBR;
  char               BRAND_ID[2+1];
  char               STATUS[2+1];
  char               BELONG_DISTRICT[6+1];
  int                SUBS_TYPE;
  int                DEF_ACCT_ID;
  BF_SUBSCRIBER()
  {memset(this,0,sizeof(BF_SUBSCRIBER));}
};

#endif

#include "message.hpp"
#include <unistd.h>
#include "data.hpp"
#include <cstdlib>
#include <boost/shared_ptr.hpp>
using namespace boost;

bool stop=false;

int
main (int argc, char* argv[])
{
  MQ mq;
  while(true)
  {
    if (stop)
      break;

    int size = mq.get_room ("/tmp/test");
    if (size < 0)
    {
       sleep(atoi(argv[1]));
       continue;
    }
    shared_ptr<BF_SUBSCRIBER> ca (new BF_SUBSCRIBER);
    mq.send((const char*)(&(*ca)),sizeof(BF_SUBSCRIBER));
    sleep(atoi(argv[1]));
  }
  mq.unlink();
  return 0;
}

/*
aCC -AA +DD64 -lrt sendmain.cpp message.o -o msend -L/home/hejb/boost_1_37_0/stage/lib -lboost_system-mt
*/

