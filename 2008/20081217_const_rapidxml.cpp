//============================================================================
// Name        : long_check.cpp
// Author      : bonly
// Version     :
// Copyright   : Bonly's copyright notice
// Description : Hello World in C++, Ansi-style
//============================================================================

#include <iostream>
#include <string>
#include <rapidxml.hpp>
using namespace std;

int main()
{
  //char text[]="<?xml version=\"1.0\"?><flash:stream to=\"bonly\" xmlns=\"jabber:client\" xmlns:flash=\"http://www.jabber.com/streams/flash\" version=\"1.0\" />";
	string str("<?xml version=\"1.0\"?><flash:stream to=\"bonly\" xmlns=\"jabber:client\" xmlns:flash=\"http://www.jabber.com/streams/flash\" version=\"1.0\" />");
	using namespace rapidxml;
	xml_document<> doc;    // character type defaults to char
	try
	{
	  doc.parse<0>(
	  		//直接用字符串，即const 失败或加(char*)强制转换也失败
	  		//const_cast<char*>("<?xml version=\"1.0\"?><flash:stream to=\"bonly\" xmlns=\"jabber:client\" xmlns:flash=\"http://www.jabber.com/streams/flash\" version=\"1.0\" />")
	  		const_cast<char*>(str.c_str()) //强制转换string得到的const char*成功
	  		);    // 0 means default parse flags
	}
	catch(rapidxml::parse_error & e)
	{
		cout << e.what() << endl;
	}

	cout << "Name of my first node is: " << doc.first_node()->name() << "\n";

	return 0;
}

