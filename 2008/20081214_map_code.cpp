//============================================================================
// Name        : cr.cpp
// Author      : bonly
// Version     :
// Copyright   : bonly's copyright notice
// Description : Hello World in C++, Ansi-style
//============================================================================

#include <iostream>
#include <map>
#include <string>
using namespace std;

map<int,string> bmap;
int main()
{
	bmap.insert(make_pair(2001,"�ɹ�"));
	bmap.insert(make_pair(1001,"�û������� | �ʻ������� | �ʱ�������"));

	map<int,string>::iterator p =
		bmap.find(1001);
	if (p!=bmap.end())
	{
		cout << "1001: " << p->second;
	}

	return 0;
}

