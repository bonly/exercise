#include <mqueue.h>
#include <cstdio>
#include <fcntl.h>
#include <cstdlib>

int
main (int argc,char* argv[])
{
  int c,flags;
  mqd_t mqd;
  ssize_t n;
  uint_t prio;
  void *buff;
  struct mq_attr attr;
  flags = O_RDONLY;
  while((c=getopt(argc,argv,"n"))!=-1)
  {
    switch(c)
    {
      case 'n':
      flags |= O_NONBLOCK;  //非阻塞方式,无消息时返回消息长度为-1
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
  buff=malloc(attr.mq_msgsize);
  n = mq_receive(mqd,(char*)buff,attr.mq_msgsize,&prio);
  printf("read %ld bytes, priority = %u\n",(long)n,prio);
  exit(0);
}
/*
 * aCC -AA try_mqrecv.cpp -o recv -lrt
 */



