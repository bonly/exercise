#include <boost/format.hpp>
#include <boost/shared_ptr.hpp>
#include <iostream>
using namespace std;
using namespace boost;
class Csck
{
   public:
     /*static shared_ptr<Csck> operator()()  //operator cann't be static
     {
          cout<<"from Csck::operator()\n";
          return shared_ptr<Csck>(new Csck);
     }*/
     //Csck(){cout<<"from Csck::Csck()\n";}
     ~Csck(){cout<<"from Csck::~Csck()\n";}
};

static Csck ck=Csck();

int main()
{
   cout << "this is main\n";
   return 0;
}

