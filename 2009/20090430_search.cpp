#include <iostream>
#include <string>
#include <vector>
#include <boost/foreach.hpp>
#include <boost/format.hpp>
using namespace std;
using namespace boost;

vector<int> exclude;
int setup_exclude_msg(const char* str)
{
  string cnt(str);
  int pot = 0;
  while (-1 != (pot = cnt.find_first_of('|',0)))
  {
  	int sms_type = -1;
  	string msg(cnt,0,pot);
  	if (1!=sscanf (cnt.c_str(),"%d|",&sms_type))
  	{
  		cerr << "设置不正确\n";
  		exit (EXIT_FAILURE);
  	}
  	cnt.erase(0,pot+1);
  	exclude.push_back(sms_type);
  }
  return 0;
}

bool search_exclude_sms(const int kind)
{
	vector<int>::iterator p;
	p = find (exclude.begin(),exclude.end(),kind);
	return p==exclude.end()?false:true;
}

int
main (int argc, char* argv[])
{
	string tm(argv[1]);
  setup_exclude_msg(tm.c_str());
  BOOST_FOREACH(int i,exclude)
  {
  	cout << format("setup: %d\n")%i;
  }

  int kind = -1;
  do
  {
  	cout << "sms type for search: " ;
  	cin >> kind;
  	if(search_exclude_sms(kind))
  		cout << format("found %d\n")%kind;
  	else
  		cout << format("not found %d\n")%kind;
  }while(kind!=-1);

	return 0;
}

