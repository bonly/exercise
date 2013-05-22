#include <stdio.h>
#include <signal.h>
#include <errno.h>
#include <cstdlib>
//#include <ucontext.h>
#include <iostream>
#include <boost/format.hpp>
using namespace boost;
using namespace std;
void myhandler (int sn , siginfo_t*  si , void* sc)
{
   //unsigned int mnip;
   //int i;
   printf(" signal number = %d, signal errno = %d, signal code = %d\n",
                                   si->si_signo,si->si_errno,si->si_code);
   printf(" senders' pid = %x, sender's uid = %d, \n",si->si_pid,si->si_uid);
}
void ctrl_c_op(int signo)
{
  cout << "oK"<< endl;
}
int main()
{
   struct sigaction s;

   cout << format ("%d\n")%myhandler;//显示函数地址
   //注意,这里不是生命函数指针,而是定义一个函数指针的类型,这个类型是自己定义的,类型名为fp
   //typedef int (*fp)(int a);
   //fp fpi;//这里利用自己定义的类型名fp定义了一个fpi的函数指针!

   //void (*fp)(int);
   //fp=ctrl_c_op;
   //s.sa_handler = fp;
   //s.sa_handler = ctrl_c_op; //此写法也可

   s.sa_flags = SA_SIGINFO;
   s.sa_sigaction = myhandler;
   if(sigaction (SIGTERM,&s,(struct sigaction *)NULL)) {
      printf("Sigaction returned error = %d\n", errno);
      exit(0);
   }
   while(1){
  	 sleep(3);
   }
   return 0;
}

