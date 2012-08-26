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
class helloworld
{
	public:
	 helloworld(std::string k):_name(k){}
	 void operator()()
	 {
		static int i=0;
		while(i<3)
		{
			cout << "hello " << _name << endl;
			boost::thread::sleep(
					boost::get_system_time()+boost::posix_time::milliseconds(3000)
					);

			++i;
		}
	 }
	private:
		std::string _name;
};
int main()
{
	//helloworld ts;
	//boost::thread thr(ts);
	boost::thread thr(helloworld("test"));
	thr.join();
	return 0;
}
