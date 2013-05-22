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
#include "data.hpp"
#include <boost/shared_ptr.hpp>
#include <unistd.h>

using namespace boost;

int
main(int argc, char* argv[])
{
  MQ mq;
  for(;;)
  {
    shared_ptr<BF_SUBSCRIBER> sub(new BF_SUBSCRIBER);
    mq.get_bf_subscriber("/tmp/test",
                     &(*sub));
    sleep(atoi(argv[1]));
  }
  return 0;
}

/*
aCC -AA +DD64 recvmain.cpp message.o -o mrecv -lrt -L/home/hejb/boost_1_37_0/stage/lib -lboost_system-mt
*/

