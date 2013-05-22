//============================================================================
// Name        : radom.cpp
// Author      : bonly
// Version     :
// Copyright   : bonly's copyright notice
// Description : Hello World in C++, Ansi-style
//============================================================================

#include <iostream>
using namespace std;

int main() {
	cout << ":bonly^_^" << endl; // prints :bonly^_^
	srand(time(NULL));//只能初始化一次
	for (int i=0; i<=10; ++i)
	{
		cout << random() << endl;
		cout << rand() << endl;
	}
	return 0;
}

