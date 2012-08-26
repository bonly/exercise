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



int main()
{
	while(true)
	{
		struct timeval tm;
		tm.tv_sec  = 1;
		tm.tv_usec = 0;
		select (0,NULL,NULL,NULL,&tm);

		CA::sigalrm_fn(3);
	}
	return 0;
}

