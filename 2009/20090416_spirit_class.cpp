//============================================================================
// Name        : radom.cpp
// Author      : bonly
// Version     :
// Copyright   : bonly's copyright notice
// Description : Hello World in C++, Ansi-style
//============================================================================

#include <iostream>
#include <string>
#include <boost/spirit/include/classic_core.hpp>
using namespace BOOST_SPIRIT_CLASSIC_NS;
using namespace std;

int main()
{
	string sp("134,12,155");
	rule<> r = real_p >> *(',' >> real_p);
	bool ret = parse(sp.c_str(), r, space_p).full;
	if (ret) cerr << "³É¹¦\n";
	else cerr << "Ê§°Ü\n";

	string source("1:23;3:12;8:234;");
	string psrc(source);

	int pot = 0;
	while( -1 != (pot = psrc.find_first_of(';',0)))
	{
	  string msg(psrc,0,pot);
	  int sub=0;
	  string kind;
	  string money;
	  if (0!=(sub = msg.find_first_of(':',0)))
	  {
	  	kind=msg.substr(0,sub);
	  	money=msg.substr(sub+1);
	  }
	  else
	  {
	  	kind=msg;
	  	money="0";
	  }
	  cerr << kind << " is " << money << endl;
	  psrc.erase(0,pot+1);
	}


	return 0;
}

