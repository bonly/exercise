#include <cstdio>
#include <cerrno>
#include <cstdarg>
#include <cstring>
#include <cstdlib>
#include <unistd.h> //read
#include <arpa/inet.h>

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
void
err_quit(const char* fmt,...)
{
	va_list ap;
	va_start(ap,fmt);
//	err_doit(0,LOG_ERR,fmt,ap);
	err_doit(0,1,fmt,ap);
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
Connect(int sockfd,const struct sockaddr* servaddr,socklen_t addrlen)
{
	int n;
	if ( (n=connect(sockfd,servaddr,addrlen))<0)
		err_sys("connect error");
	return n;
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
ssize_t
readn(int fd,void*vptr,size_t n)
{
	ssize_t nleft;
	ssize_t nread;
	char *ptr;
	ptr=(char*)vptr;
	nleft=n;
	while(nleft>0)
	{
		if((nread=read(fd,ptr,nleft))<0)
		{
			if(errno==EINTR)
				nread=0;
			else
				return(-1);
		}
		nleft-=nread;
		ptr+=nread;
	}
	return(n-nleft);
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
void
str_cli(FILE *fp,int sockfd)
{
	char sendline[MAXLINE],recvline[MAXLINE];
	while(fgets(sendline,MAXLINE,fp)!=NULL)
	{
		writen(sockfd,sendline,strlen(sendline));
		if(readline(sockfd,recvline,MAXLINE)==0)
			err_quit("str_cli:server terminated prematurely");
		fputs(recvline,stdout);
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

int
main(int argc,char* argv[])
{
	int    i,sockfd[5];
	struct sockaddr_in servaddr;
	if(argc!=2)
		err_quit("usage:tcpcli<IPaddress>");
	for (i =0; i<5; ++i)
	{
		sockfd[i]=Socket(AF_INET,SOCK_STREAM,0);
		bzero(&servaddr,sizeof(servaddr));
		servaddr.sin_family=AF_INET;
		servaddr.sin_port=htons(SERV_PORT);
		inet_pton(AF_INET,argv[1],&servaddr.sin_addr);
		Connect(sockfd[i],(sockaddr*)&servaddr,sizeof(servaddr));
	}
	str_cli(stdin,sockfd[0]);
	exit(0);
}

