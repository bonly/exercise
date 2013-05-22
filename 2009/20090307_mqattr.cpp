#include <mqueue.h>
#include <cstdio>
#include <cstdlib>
#include <fcntl.h>

int
main(int argc,char* argv[])
{
  mqd_t mqd;
  struct mq_attr attr;
  if (argc!=2)
  {
    printf("usage: mqgetattr <name>\n");
    exit(0);
  }
  mqd = mq_open(argv[1],O_RDONLY);
  if(mqd<0)
  {
      perror("mq_open");
      return -1;
  }   
  if(mq_getattr(mqd,&attr)<0)
  {
      perror("mq_getattr");
      return -1;
  }   
  printf("max #msgs = %ld, max #bytes/msg = %ld,"
         "#currently on queue = %ld\n",
         attr.mq_maxmsg,attr.mq_msgsize,attr.mq_curmsgs);
  mq_close(mqd);
  exit(0);
}

