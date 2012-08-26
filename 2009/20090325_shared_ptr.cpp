/*֤��shared_ptr=��ʱ����ԭ��û���õ�ɾ����
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

