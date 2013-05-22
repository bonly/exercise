#include <boost/shared_ptr.hpp>
#include <boost/format.hpp>
#include <iostream>
using namespace std;
using namespace boost;

class A
{
 public:
  A(){ clog << format("A::A\n");}
  virtual ~A(){clog << format("A::~A\n");}
  virtual int
  fun(){clog << format("A::fun\n");}
};

class B : public A
{
 public:
  B(){clog << format("B::B\n");}
  virtual ~B(){clog << format("B::~B\n");}
  virtual int
  fun(){clog << format("B::fun\n");}
};

int
gfun(shared_ptr<A> &obj)
{
   shared_ptr<B> c(new B);
   obj=c;
   return 0;
}


int
main()
{
  shared_ptr<A> e;
  gfun(e);
  e->fun();
  return 0;
}
