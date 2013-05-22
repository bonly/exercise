//#include <sys/socket.h>
//#include <sys/types.h>
#include <netinet/in.h>  //sockaddr_in
#include <cstdarg> //va_list
#include <cerrno> //errno
#include <cstdio> //vsprintf
#include <cstring> //strlen
#include <cstdlib> //exit
#include <unistd.h> //fork
#include <sys/wait.h> //wait

enum
{MAXLINE=255,SERV_PORT=2456,LISTENQ=5};
static void
err_doit(int errnoflag, int level, const char* fmt, va_list ap)
{
	int errno_save,n;
	char buf[MAXLINE+1];
	errno_save = errno;
#ifdef HAVE_VSNPRINTF
	vsnprintf(buf,MAXLINE,fmt,ap);
#else
	vsprintf(buf,fmt,ap);
#endif
	n=strlen(buf);
	if(errnoflag)
		snprintf(buf+n,MAXLINE-n,":%s",strerror(errno_save));
	strcat(buf,"\n");
	//if(daemon_proc){
	//	syslog(level,buf);}
	//else
	{
		fflush(stdout);
		fputs(buf,stderr);
		fflush(stderr);
	}
	return;
}
void
err_sys(const char* fmt,...)
{
	va_list ap;
	va_start(ap,fmt);
	//err_doit(1,LOG_ERR,fmt,ap);
	err_doit(1,1,fmt,ap);
	va_end(ap);
	exit(1);
}
int
Socket(int family,int type,int protocol)
{
	int n;
	if ( (n=socket(family,type,protocol))<0)
		err_sys("socket error");
	return n;
}
int
Bind(int sockfd,const struct sockaddr*myaddr,socklen_t addrlen)
{
	int n;
	if ( (n=bind(sockfd,myaddr,addrlen))<0)
		err_sys("bind error");
	return n;
}
void
Listen(int sockfd,int backlog)
{
  char *ptr;
  //can override 2nd argument with environment variable
  if((ptr=getenv("LISTENQ"))!=NULL)
  	backlog=atoi(ptr);
  if(listen(sockfd,backlog)<0)
  	err_sys("listen error");
}
int
Accept(int sockfd,struct sockaddr*clladdr,socklen_t *addrlen)
{
	int n;
	if ( (n=accept(sockfd,clladdr,addrlen))<0)
		err_sys("accept error");
	return n;
}
pid_t
Fork(void)
{
	pid_t n;
	if ( (n=fork())==1)
		err_sys("fork error");
	return n;
}
int
Close(int sockfd)
{
	int n;
	if ((n=close(sockfd)==1))
		err_sys("close error");
	return n;
}
ssize_t
readline(int fd,void*vptr,ssize_t maxlen)
{
	ssize_t n,rc;
	char c,*ptr;
	ptr=(char*)vptr;
	for (n=1;n<maxlen;n++)
	{
		again:
		  if((rc=read(fd,&c,1))==1)
		  {
		  	*ptr++=c;
		  	if(c=='\n')
		  		break;
		  	else if(rc==0)
		  	{
		  		if(n==1)
		  			return(0);
		  		else
		  			break;
		  	}
		  }
		  else
		  {
		  		if(errno==EINTR)
		  			goto again;
		  		return(-1);
		  }
	}
  *ptr=0;
  return(n);
}
ssize_t
writen(int fd,const void*vptr,size_t n)
{
	size_t nleft;
	ssize_t nwritten;
	const char *ptr;
	ptr=(char*)vptr;
	nleft=n;
	while(nleft>0)
	{
		if((nwritten = write(fd,ptr,nleft))<=0)
		{
			if(errno==EINTR)
				nwritten=0;
			else
				return(-1);
		}
		nleft-=nwritten;
		ptr+=nwritten;
	}
	return(n);
}
void
str_echo(int sockfd)
{
	ssize_t n;
	char line[MAXLINE];
	for(;;)
	{
		if((n=readline(sockfd,line,MAXLINE))==0)//客户端用^D送FIN信号,服务端以ACK响应关闭
			return;
		writen(sockfd,line,n);
	}
}

typedef void Sigfunc(int);
Sigfunc*
Signal(int signo, Sigfunc *func)
{
	struct sigaction act,oact;
	act.sa_handler = func;
	sigemptyset(&act.sa_mask);
	act.sa_flags = 0;
	if (signo == SIGALRM)
	{
#ifdef SA_INTERRUPT
		act.sa_flags |= SA_INTERRUPT; //SunOS 4.x
#endif
	}
	else
	{
#ifdef SA_RESTART
		act.sa_flags |= SA_RESTART; //SVR4,4.4BSD
#endif
	}
	if(sigaction(signo,&act,&oact)<0)
		return (SIG_ERR);
	return (oact.sa_handler);
}
void
sig_chld(int signo)
{
	pid_t pid;
	int stat;
	//所有5个信号都在信号处理程序执行之前产生,
	//而此信号处理程序又只执行一次,因为UNIX信号一般是不排队的.
	//如果客户与服务器在同一主机上,信号处理程序执行一次,留下四个僵尸进程,
	//若不同的主机上运行客户和服务器,信号处理程序一般执行两次:
	//一次作为第一个产生的信号的结果,由于另外4个信号在信号处理程序执行时发生,
	//所以处理程序一般情况下会再被调用一次,这就留下三个僵尸进程.
	//但有时,可能依赖于FIN到达服务器主机的时机,信号处理程序执行三次或四次
	//正确解决方法是用waitpid而不是wait
	pid = wait(&stat);
	printf("child %d terminated\n",pid);
	return;
}

int
main (int argc,char* argv[])
{
  int listenfd,connfd;
  pid_t childpid;
  socklen_t chllen;
  struct sockaddr_in chladdr,servaddr;
  listenfd = Socket(AF_INET,SOCK_STREAM,0);
  bzero(&servaddr,sizeof(servaddr));
  servaddr.sin_family=AF_INET;
  servaddr.sin_addr.s_addr=htonl(INADDR_ANY);
  servaddr.sin_port = htons(SERV_PORT);
  Bind(listenfd,(sockaddr*)&servaddr,sizeof(servaddr));//通用地址结构指针转换
  Listen(listenfd,LISTENQ);
  Signal(SIGCHLD,sig_chld);
  for(;;)
  {
  	chllen = sizeof(chladdr);
  	if (connfd = Accept(listenfd,(sockaddr*)&chladdr,&chllen) < 0)
  	{//慢系统调用(slow system call)阻塞,
  	 //处理SIGCHLD中断后如果系统设计不自动重启select/accept...
  	 //等阻塞调用,将有EINTR错误返回
  		if(errno == EINTR)
  			continue; //back to for()
  		else
  			err_sys("accept error");
  	}
  	if((childpid=Fork())==0)
  	{
  		Close(listenfd);
  		str_echo(connfd);
  		exit(0);
  	}
  	Close(connfd);
  }
}


