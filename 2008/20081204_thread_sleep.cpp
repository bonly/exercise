//============================================================================
// Name        : thread.cpp
// Author      : bonly
// Version     :
// Copyright   : bonly's copyright notice
// Description : Hello World in C++, Ansi-style
//============================================================================
#include <boost/thread.hpp>
#include <iostream>
using namespace std;
void
helloworld()
{
	static int i=0;
	while(i<10000)
	{
		cout << "hello world" << endl;
		boost::thread::sleep(
				boost::get_system_time()+boost::posix_time::milliseconds(3000)
				);

		++i;
	}
}

int main()
{
	boost::thread thr(&helloworld);
	thr.join();
	return 0;
}
