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
	pid = wait(&stat);
	printf("child %d terminated\n",pid);
	return;
}

int
main (int argc,char* argv[])
{
  int i,maxi,maxfd,listenfd,connfd,sockfd;
  int nready,client[FD_SETSIZE];
  ssize_t n;
  fd_set rset,allset;
  char line[MAXLINE];
  socklen_t clilen;
  struct sockaddr_in cliaddr,servaddr;

  listenfd = Socket(AF_INET,SOCK_STREAM,0);
  bzero(&servaddr,sizeof(servaddr));
  servaddr.sin_family=AF_INET;
  servaddr.sin_addr.s_addr = htonl(INADDR_ANY);
  servaddr.sin_port = htons(SERV_PORT);
  Bind(listenfd,(sockaddr*)&servaddr,sizeof(servaddr));
  Listen(listenfd,LISTENQ);
  maxfd = listenfd; //初始化为侦听的handler
  maxi = -1;  //client[]的索引
  for (i=0;i<FD_SETSIZE;++i)
  {
  	client[i] = -1; //表示空的
  }
  FD_ZERO(&allset);
  FD_SET(listenfd,&allset);

  for(;;)
  {
  	rset = allset; //复制
  	nready = select(maxfd +1,&rset,NULL,NULL,NULL);
  	if (FD_ISSET(listenfd,&rset))//new client connection
  	{
  		clilen = sizeof(cliaddr);
  		connfd = Accept(listenfd,(sockaddr*)&cliaddr,&clilen);
  		for(i=0;i<FD_SETSIZE;++i)
  		{
  			if(client[i]<0)
  			{
  				client[i]=connfd; //save descriptor
  				break;
  			}
  		}
  		if(i==FD_SETSIZE)
  			err_sys("too many clients");
  		FD_SET(connfd,&allset); //add new descriptor to set
  		if(connfd>maxfd)
  			maxfd = connfd; //for select
  		if(i>maxi)
  			maxi = i; //max index in client[]
  		if(--nready<=0)
  			continue; //no more readable descriptors
  	}

  	for (i=0;i<=maxi;++i) //check all clients for data
  	{
  		if((sockfd = client[i])<0)
  			continue;
  		if(FD_ISSET(sockfd,&rset))
  		{
  			if((n = readline(sockfd,line,MAXLINE))==0)
  			{
  				//connection closed by client
  				Close(sockfd);
  				FD_CLR(sockfd,&allset);
  				client[i]=-1;
  			}
  			else
  				writen(sockfd,line,n);
  			if(--nready<=0)
  				break; //no more readable descriptors
  		}
  	}
  }
}


