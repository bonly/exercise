//测试在线程中创建进程,
//在cygwin成功
//HP-UX中成功
//exec函数最后的参数为(char*)0,而且参数只能是char*
//一个进程就相当于一个主线程。
//fork进程：子进程复制父进程的进程环境。父进程结束不会影响子进程的运行。进程切换复制进程环境。
//create线程：子线程共享父线程的线程环境。线程（一个主线程下的多个线程）切换不复制环境，所以线程运行快，省去了复制环境的时间。主线程运行结束，子线程的运行就无意义了。
//子线程返回主线程的方法：
//1.主线程等待子线程运行结束
//2.用信号，比如说用主线程有一个全局变量 p=1,子线程有一 while(p) 一直在无限运行，这是主线程 p = 0,while(p) 子线程退出。
//一个进程 create了几个线程，进程fork()，这是fork出的子进程不能复制父线程的线程，也就是说，fork出来的子进程只复制父现成的执行环境。
//线程是一个执行体。
//进程是运行环境+执行体。
//ulimit -a 显示进程环境的属性 
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
	  //printids("in vfork");// 运行不了这句,出错
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

