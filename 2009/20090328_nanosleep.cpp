//============================================================================
// Name        : bdb.cpp
// Author      : bonly
// Version     :
// Copyright   : bonly's copyright notice
// Description : debug info
// fun(const unsigned char*):参数中的指针只是一份copy,不是地址传值
//============================================================================

#include <iostream>
#include <string>
#include <unistd.h>
#include <sys/select.h>


class CA
{
	public:
	static void sigalrm_fn(int sig)
	{
		using namespace std;
		string data("2009-04-28 14:00:32");
		data = data.erase(10);
		data = data.erase(7,1);
		data = data.erase(4,1);
		cout << data << endl;
		printf("alarm!\n");
		alarm(2);
		return;
	}
};



int main(int argc, char* argv[])
{
	while(true)
	{

    struct timespec tm={.tv_sec = atol(argv[2]),.tv_nsec = atol(argv[1])};
    nanosleep(&tm,NULL);
		CA::sigalrm_fn(3);
	}
	return 0;
}

