#include <signal.h>
#include <sys/time.h>
#include <unistd.h>
#include <sys/types.h>
#include <cstdlib>
#include <cstdio>
main(int argc,char**argv)
{
	pid_t pid;
	int signum;
	union sigval mysigval;
	signum=atoi(argv[1]);
	pid=(pid_t)atoi(argv[2]);
	mysigval.sival_int=8;//��������庬�壬ֻ����˵������
	if(sigqueue(pid,signum,mysigval)==-1)
		printf("send error\n");
	sleep(2);
}

