#include <mqueue.h>
#include <cstdlib>
#include <cstdio>
#include <fcntl.h> //O_WRONLY
#include <sys/stat.h>
#include <sys/types.h>

int
main(int argc,char* argv[])
{
  mqd_t mqd;
  void  *ptr;
  size_t len;
  uint_t prlo;
  if(argc!=4)
  {
    printf ("usage: mqsend <name> <#bytes> <priority>");
    exit(0);
  }
  len = atoi(argv[2]);
  prlo = atoi(argv[3]);
  mqd = mq_open(argv[1],O_WRONLY);
  if (mqd<0)
  {
	  perror("mq_open");
	  exit(-1);
  }
  ptr = calloc(len,sizeof(char));
  if(mq_send(mqd,(const char*)ptr,len,prlo)<0)
  {
    perror("mq_send");
    return -1;
  }
  
  exit(0);
}
/*
aCC -AA try_mqsend.cpp -o mqsend -lrt
*/

