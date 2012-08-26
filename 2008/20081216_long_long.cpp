//============================================================================
// Name        : long_check.cpp
// Author      : bonly
// Version     :
// Copyright   : Bonly's copyright notice
// Description : Hello World in C++, Ansi-style
//============================================================================

#include <iostream>
#include <boost/cast.hpp>
using namespace std;

int main()
{
	char ch='k';
	int  it=254;
	cout << "sizeof(char): " << sizeof(ch) << endl;
	cout << "sizeof(int): " << sizeof(it) << endl;
	cout << "max of int: " << std::numeric_limits<int>::max() << endl;
	cout << "max of long long: " << std::numeric_limits<long long>::max() << endl;
	cout << "max of long: " << std::numeric_limits<long>::max() << endl;
	cout << "min of float: " << std::numeric_limits<float>::min() << endl;
	try
	{
	  unsigned long long lk=boost::numeric_cast<long long int>(13719360007uLL);
		//long long lk=13719360007LL;
	  cout << "my phone number is: " << lk << endl;
	}
	catch(boost::bad_numeric_cast & e)
	{
		std::cout << e.what() << '\n';
	}
	return 0;
}

