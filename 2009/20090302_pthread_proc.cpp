//�������߳��д�������,
//��cygwin�ɹ�
//HP-UX�гɹ�
//exec�������Ĳ���Ϊ(char*)0,���Ҳ���ֻ����char*
//һ�����̾��൱��һ�����̡߳�
//fork���̣��ӽ��̸��Ƹ����̵Ľ��̻����������̽�������Ӱ���ӽ��̵����С������л����ƽ��̻�����
//create�̣߳����̹߳����̵߳��̻߳������̣߳�һ�����߳��µĶ���̣߳��л������ƻ����������߳����п죬ʡȥ�˸��ƻ�����ʱ�䡣���߳����н��������̵߳����о��������ˡ�
//���̷߳������̵߳ķ�����
//1.���̵߳ȴ����߳����н���
//2.���źţ�����˵�����߳���һ��ȫ�ֱ��� p=1,���߳���һ while(p) һֱ���������У��������߳� p = 0,while(p) ���߳��˳���
//һ������ create�˼����̣߳�����fork()������fork�����ӽ��̲��ܸ��Ƹ��̵߳��̣߳�Ҳ����˵��fork�������ӽ���ֻ���Ƹ��ֳɵ�ִ�л�����
//�߳���һ��ִ���塣
//���������л���+ִ���塣
//ulimit -a ��ʾ���̻��������� 
//
#include <cstdio>
#include <pthread.h>
#include <cstring>
#include <sys/types.h>
#include <sys/wait.h>
#include <unistd.h>
#include <cstdlib>
pthread_t ntid;
void printids(const char*s)
{
  pid_t pid;
  pthread_t tid;
  pid = getpid();
  tid = pthread_self();
  printf ("%s pid %u tid %u (0x%x)\n",
          s,(unsigned int)pid,(unsigned int)tid,(unsigned int)tid);
}

void create_process()
{
  pid_t pid;
  if ((pid=vfork())==0)
  {
	  //printids("in vfork");// ���в������,����
	  execlp("/cygdrive/d/cygwin/bin/ls.exe","/cygdrive/d/cygwin/bin/ls.exe","/",(char*)0);
	  printf("start chld fail\n");
	  exit(0);
  }
  printids("out vfork");
  printf("end create_process\n");
}
void create_process_fun()
{
  pid_t pid;
  if ((pid=fork())==0)
  {
	  for (int i=0;i<20;++i)
		  printf("%d\n",i);
	  _exit(0);
  }
}
void* thr_fn(void *arg)
{
  printids("new thread: ");
  create_process();
  //create_process_fun();
  wait(NULL);
  return ((void*)0);
}

int main()
{
  int err;
  err = pthread_create(&ntid,NULL,thr_fn,NULL);
  if (err !=0)
  {
    printf("can't create thread: %s\n",strerror(err));
    return 1;
  }
  printids("main thread: ");
  sleep(1);
  return 0;
}

