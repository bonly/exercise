#include <iostream>
using namespace std;
template<typename T>
struct OpNewCreator
{
    static T* Create()
    {
      clog << "use new to create\n";
      return new T;
    }
};

class Widget
{
  public:
   Widget(){clog << "create widget\n";}
   void print(){ clog << "now widget print\n";}
};

template< template<typename Created> class CreationPolicy>
class WidgetManager : public CreationPolicy<Widget>
{
  public:
   void wprint()
   {
     Widget* k = OpNewCreator<Widget>::Create(); //如果没加域需加-fpermissive,且有警告
     k->print();
   }

};

int main()
{
  WidgetManager<OpNewCreator> wig;
  wig.wprint();
  return 0;
}

/*
g++ 20090809_policy.cpp -fpermissive
*/
