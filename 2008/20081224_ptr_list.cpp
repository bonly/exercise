#include <iostream>
#include <vector>
#include <boost/shared_ptr.hpp>
using namespace std;
using namespace boost;

class Me
{
  public:
   Me(){ std::cerr << "Me()\n";}
   ~Me(){ cerr << "~Me()\n";}
};

int
main()
{
  vector< shared_ptr<Me> > vec;
  shared_ptr<Me> f1(new Me);
  vec.push_back (f1);

  shared_ptr<Me> f2;
  f2 = shared_ptr<Me>(new Me);
  vec.push_back (f2);
  return 0;
}
  


