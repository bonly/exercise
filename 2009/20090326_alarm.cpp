//============================================================================
// Name        : bdb.cpp
// Author      : bonly
// Version     :
// Copyright   : bonly's copyright notice
// Description : debug info
// fun(const unsigned char*):�����е�ָ��ֻ��һ��copy,���ǵ�ַ��ֵ
//============================================================================

#include <iostream>
#include <string>
#include <unistd.h>
#include <boost/bind.hpp>
using namespace boost;

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
	CA ab;
	signal(SIGALRM, CA::sigalrm_fn);
	alarm(1);
	//while(1) pause();
	while(1) sleep(3);

	return 0;
}

