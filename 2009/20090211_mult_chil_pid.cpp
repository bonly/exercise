#include <sys/wait.h>
#include <sys/types.h>
#include <unistd.h>
#include <iostream>
#include <boost/format.hpp>
using namespace boost;
using namespace std;

int
main(int argc, char* argv[])
{
  pid_t pid;
  pid_t pid1;
  pid_t pid2;
  if(argc==1)
  {
    if((pid=fork())==0)
    {
      if (execlp("mypid","mypid1","1")<0)
       cerr << "start srv fail\n";
    }
    else
    {
      for(int i=0;i<2;++i)
      {
        sleep (2);
        cout << "main:i am running..." << getpid() << endl;;
      }
      //cout <<format("main[%d] finished\n")%getpid();
      /*
      cout << "main finished\n";
      wait(&pid);
      cout << "main end 1\n";
      wait(&pid1);
      cout << "main end 2\n";
      wait(&pid2);
      cout << "main end 3\n";
      //����,��Ϊ��ʱ����������ʵ��֪���ӽ��̵��ӽ��̵Ľ��̺�PID*/
      /*
      for(int i=1;i<4;++i)
      {
        wait(NULL); //Ҳ����,0��1�Ž��̻�����
        cout <<"main end " << i << endl;
      }
      */
    }
  }
  else { //����������Ľ���
  if (strcmp(argv[1],"1")==0)
  {
    if((pid1=fork())==0)
    {
      if (execlp("mypid","mypid2","2")<0)
        cerr<< "start 2srv fail\n";
    }
    else{
    for(int i=0;i<4;++i)
    {
      sleep (2);
      cout << "P1:i am running..." << getpid()<< endl;
    }
    //cout <<format("P1[%d] finished\n")%getpid();
    cout << "P1 finished\n";
    //wait(NULL); 
    //wait(&pid1);
    }
  }
  if (strcmp(argv[1],"2")==0)
  {
    if((pid2=fork())==0)
    {
      if (execlp("mypid","mypid3","3")<0)
        cerr<< "start 3srv fail\n";
    }
    else {
    for(int i=0;i<6;++i)
    {
      sleep (2);
      cout << "P2:i am running..." << getpid()<< endl;
    }
    //cout <<format("P2[%d] finished\n")%getpid();
    cout << "P2 finished\n";
    //wait(NULL);
    //wait(&pid2);
    }
  }
  if (strcmp(argv[1],"3")==0)
  {
    for(int i=0;i<8;++i)
    {
      sleep(2);
      cout << "P3:i am running... " << getpid()<< endl;
    }
    //cout <<format("P3[%d] finished\n")%getpid(); 
    cout << "P3 finished\n";
    //_exit(0);
    //exit(0);
    //return 0;
  }
  }
  cout << "PID: "<<getpid()<< " wait\n";
  wait(NULL);
  return 0;
  //exit(0);
}

