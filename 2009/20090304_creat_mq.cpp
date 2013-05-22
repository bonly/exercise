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
  flags = O_RDWR|O_CREAT;
  while((c=getopt(argc,argv,"e"))!=-1)
  {
    switch(c)
    {
      case 'e':
        flags |= O_EXCL;
        break;
    }
  }
  if (optind !=argc -1)
  {
    cerr << "usage:mqcreate [-e]<name>"<<endl;
    return -1;
  }
  mqd = mq_open(argv[optind],flags,FILE_MODE,NULL);
  mq_close(mqd);
  exit(0);
}


/*
aCC -AA try_mq.cpp -lrt
*/

