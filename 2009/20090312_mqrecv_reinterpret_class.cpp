#include <mqueue.h>
#include <cstdio>
#include <fcntl.h>
#include <cstdlib>

class Base
{   
      public:
         Base():_msgtype(0){}
         int msgtype(){return _msgtype;}
      protected:
         int _msgtype;
};
class AMsg:public Base
{        
        public:
           AMsg(){_msgtype=1;}
           char DataA[20];
};    
class BMsg:public Base
{   
     public:
         BMsg(){_msgtype=2;}
         char DataB[12];
         int  intb;
}; 
      
int 
main (int argc,char* argv[])
{ 
  int c,flags;
  mqd_t mqd;
  ssize_t n;
  uint_t prio;
  Base *buff;
  struct mq_attr attr;
  flags = O_RDONLY;
  while((c=getopt(argc,argv,"n"))!=-1)
  {
    switch(c)
    {
      case 'n':
    flags |= O_NONBLOCK;
    break;
    }
  }
  if(optind != argc -1)
  {
      printf("usage:mqrecv [-n] <name>");
      exit(0);
  }
  mqd = mq_open(argv[optind],flags);
  if(mqd<0)
  {
      perror("mq_open");
      exit(-1);
  }
  mq_getattr(mqd,&attr);
  //buff = malloc(attr.mq_msgsize);
  buff = new Base;
  n = mq_receive(mqd,(char*)buff,attr.mq_msgsize,&prio);
  printf("read %ld bytes, priority = %u\n",(long)n,prio);
  printf("msgtype:%d \n",buff->msgtype());
  AMsg *buf = reinterpret_cast<AMsg*>(buff);
  printf("msgtype: %d \nData: %s\n",buf->msgtype(),buf->DataA);
  exit(0);
}
/*
 * aCC -AA try_mqrecv.cpp -o recv -lrt
 */

