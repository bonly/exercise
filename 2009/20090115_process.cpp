//============================================================================
// Name        : try_process.cpp
// Author      : bonly
// Version     :
// Copyright   : bonly's copyright notice
// Description : Hello World in C++, Ansi-style
//============================================================================

#include <iostream>
#include <boost/format.hpp>
#include <sys/wait.h>
#include <fcntl.h>
#include <errno.h>
using namespace std;
using namespace boost;
extern int errno;
int main()
{
	char buf[100];
	pid_t cld_pid;
	int fd,status;
	if ((fd=open("temp",O_CREAT|O_TRUNC|O_RDWR,S_IRWXU))==-1)
	{
		cout << format("open error %d\n")%errno;
		exit(1);
	}
	strcpy (buf,"this is parent process write\n");
	if((cld_pid=fork())==0)
	{
		strcpy(buf,"This is child process write\n");
		printf("tihs is child process\n");
		cout << format ("My PID(child) is %d\n")%getpid();
		cout << format ("My parent PID is %d\n")%getppid();
		write(fd,buf,strlen(buf));
		close(fd);
		exit(0);
	}
	else
	{
		cout << "This is parent process\n";
		cout << format ("My PID(parent)is %d\n")%getpid();
		cout << format ("My Child PID is %d\n")%cld_pid;
		write(fd,buf,strlen(buf));
		close(fd);
	}
	wait(&status);
	return 0;
}

