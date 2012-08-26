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

   cout << format ("%d\n")%myhandler;//��ʾ������ַ
   //ע��,���ﲻ����������ָ��,���Ƕ���һ������ָ�������,����������Լ������,������Ϊfp
   //typedef int (*fp)(int a);
   //fp fpi;//���������Լ������������fp������һ��fpi�ĺ���ָ��!

   //void (*fp)(int);
   //fp=ctrl_c_op;
   //s.sa_handler = fp;
   //s.sa_handler = ctrl_c_op; //��д��Ҳ��

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

