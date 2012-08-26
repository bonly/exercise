#include <mqueue.h>
#include <cstdlib>
#include <cstdio>
#include <fcntl.h> //O_WRONLY
#include <sys/stat.h>
#include <sys/types.h>
#include <cstring>
class Base
{
  public:
      Base():_msgtype(0){}
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
main(int argc,char* argv[])
{          
  mqd_t mqd;
  AMsg  *ptr;
  size_t len;
  uint_t prlo;

  mqd = mq_open(argv[1],O_WRONLY);
  ptr = new AMsg();
  memset(ptr->DataA,0,20);
  strcpy(ptr->DataA,"class AMsg");
  if(mq_send(mqd,(const char*)ptr,sizeof(AMsg),2)<0)
  {
    perror("mq_send");
    return -1;
  }
  delete ptr;
  exit(0);
}
/*
aCC -AA try_mqsend.cpp -o mqsend -lrt
*/


