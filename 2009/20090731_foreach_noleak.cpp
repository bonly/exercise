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

template<typename T>
struct DeleteObject
{
  public:
    void operator()(const T* ptr) const
    {
        delete ptr;
    }
};

void doSomething()
{
  vector<A*> wp;
  for (int i=0; i<3; ++i)
  {
   wp.push_back(new A);
  }
  for_each(wp.begin(), wp.end(), DeleteObject<A>());
}

int main()
{
  doSomething();
  return 0;
}

