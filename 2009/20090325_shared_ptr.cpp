/*证明shared_ptr=的时候会把原来没有用的删除掉
 *
 */
#include <boost/shared_ptr.hpp>
#include <iostream>

using namespace std;
using namespace boost;

class TA
{
  public:
   TA(){cerr << "TA::TA\n" ;}
   ~TA(){cerr << "TA::~TA\n";}
};

int
main()
{
   shared_ptr<TA> ta(new TA);
   shared_ptr<TA> ba(new TA);
   cerr <<"begin\n";
   ba = ta;
   cerr << "end\n";
   return 0;
}

