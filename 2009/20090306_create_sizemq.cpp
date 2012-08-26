#include <mqueue.h>
#include <unistd.h>
#include <sys/types.h>
#include <sys/stat.h>
#include <fcntl.h> //O_RDWR ...
#include <iostream>
using namespace std;

#define FILE_MODE (S_IRUSR|S_IWUSR|S_IRGRP|S_IROTH)
int
main (int argc,char* argv[])
{
  int c,flags;
  mqd_t mqd;
  struct mq_attr attr;
  flags = O_RDWR|O_CREAT;
  while((c=getopt(argc,argv,"em:z:"))!=-1)
  {
    switch(c)
    {
      case 'e':
        flags |= O_EXCL;
        break;
      case 'm':
        attr.mq_maxmsg = atol(optarg); //队列中的最大消息数
        break;
      case 'z':
        attr.mq_msgsize = atol(optarg); //任意给定消息的最大字节数
        break;
    }
  }
  if (optind !=argc -1)
  {
    cerr << "usage:mqcreate [-e] [-m maxmsg -z msgsize] <name>"<<endl;
    return -1;
  }
  if ((attr.mq_maxmsg !=0 && attr.mq_msgsize==0)||
      (attr.mq_maxmsg ==0 && attr.mq_msgsize!=0))
  {
    cerr << "must specify both -m maxmsg and -z msgsize"<< endl;
    return -1;
  }

  mqd = mq_open(argv[optind],flags,FILE_MODE,
                (attr.mq_maxmsg!=0)?&attr:NULL);
  if(mqd<0)
  {    
      perror("mq_open");
      exit(-1);
  }   
  
  mq_close(mqd);
  exit(0);
}

/*
aCC -AA try_mq.cpp -lrt
*/

