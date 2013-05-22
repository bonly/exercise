#include <vector>
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
  vector<A*> wp;
  for (int i=0; i<3; ++i)
  {
   wp.push_back(new A);
  }
}

int main()
{
  doSomething();
  return 0;
}

