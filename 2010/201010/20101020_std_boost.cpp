#include <boost/shared_ptr.hpp>
#include <boost/make_shared.hpp>
#include <iostream>
///避免用using namesapce 
//using namespace std;
//using namespace boost;
namespace shared_ptr = boost::shared_ptr;
namespace b = boost; 
//using make_shared = boost::make_shared;

struct A{
  int a;
};

int main() {
  //boost::shared_ptr<A> aint = boost::make_shared<A>();
  //b::shared_ptr<A> aint = b::make_shared<A>();
  //boost::shared_ptr<A> aint(new A);
  shared_ptr<A> aint(new A);
  std::clog << aint << std::endl;
  return 0;
}

