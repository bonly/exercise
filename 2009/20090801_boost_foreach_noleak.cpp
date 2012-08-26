#include <vector>
#include <boost/format.hpp>
#include <boost/foreach.hpp>
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
        clog << format ("delete ptr\n");
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
  BOOST_FOREACH(A* p, wp)
  {
    //DeleteObject<A> obj;
    //obj.operator()(p);
    DeleteObject<A>()(p);
  }
}

int main()
{
  doSomething();
  return 0;
}

