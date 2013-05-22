#include <mqueue.h>
#include <cstdio>
#include <cstdlib>

int
main(int argc,char* argv[])
{
  if(argc!=2)
  {
    printf("usage:mqunlink <name>\n");
    return 0;
  }
  if(mq_unlink(argv[1])<0)
  {
	  perror("mq_unlink");
	  exit(-1);
  }
  exit(0);
}

