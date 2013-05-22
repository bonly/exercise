#include <map>
#include <cstdio>
#include <boost/format.hpp>
#include <iostream>

using namespace boost;
using namespace std;

class A
{
 public:
  A(){ clog << format("A::A\n");}
  ~A(){ clog << format("A::~A\n");}
};

void doSomething()
{
  map<string,A*> wp;
  for (int i=0; i<3; ++i)
  {
   char ckey[20];
   sprintf(ckey,"key%d",i);
   wp.insert(make_pair<string,A*>(string(ckey),new A));
  }
  for (map<string,A*>::iterator i= wp.begin(); i!=wp.end(); ++i)
  {
   delete i->second;
  }
}

int main()
{
  doSomething();
  return 0;
}

