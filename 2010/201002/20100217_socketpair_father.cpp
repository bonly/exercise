#include <cstdlib>
#include <cstdio>
#include <errno.h>
#include <string.h>
#include <sys/types.h>          /* See NOTES */
#include <sys/socket.h>
#include <iostream>
using std::cerr;
using std::endl;

int fd[2];
int make_child1()
{
   if(socketpair(AF_UNIX, SOCK_STREAM, 0, fd)!=0)
   {
       cerr << "create socket pair fail! " << endl;
       return -1;
   }
   pid_t pid = fork();
   switch(pid)
   {
       case -1:
       {
           cerr << "fork: " << strerror(errno) << endl;
           break;
       }
       case 0: ///子进程
       {
           close(fd[0]);
           char arg1[10]="";
           char arg2[10]="";
           //snprintf(arg1,10,"%d",fd[0]);  ///没有用
           snprintf(arg2,10,"%d",fd[1]);
           int ret = execl("c1","c1", arg2, (char*)NULL);
           if (ret == -1)
           {
               cerr <<"execl exec error\n";
               exit(0);
           }
           else
           {

           }
           break;
       }
       default:  ///父进程
       {
           break;
       }
   }
   return 0;     
}

int make_child2()
{
   pid_t pid = fork();
   switch(pid)
   {
       case -1:
       {
           cerr << "fork: " << strerror(errno) << endl;
           break;
       }
       case 0: ///子进程
       {
           close(fd[0]);
           char arg1[10]="";
           char arg2[10]="";
           //snprintf(arg1,10,"%d",fd[0]);///没用
           snprintf(arg2,10,"%d",fd[1]);
           int ret = execl("c2","c2", arg2, (char*)NULL);
           if (ret == -1)
           {
               cerr <<"execl exec error\n";
               exit(0);
           }
           else
           {

           }
           break;
       }
       default:  ///父进程
       {
           close(fd[1]);
           char str[]="hello from father";
           write (fd[0], str, strlen(str));
	   
	   char str1[255]="";
	   read (fd[0], str1, 255);
	   cerr << "read from child: " << str1 << endl;
           break;
       }
   }
   return 0;     
}
int main()
{
    make_child1();
    make_child2();
    while(1)
        sleep(10);

    return 0;
}

