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

template< class Obj,template<class Obj> class CreationPolicy>
class WidgetManager : public CreationPolicy<Obj>
{
  public:
   void wprint()
   {
     Obj* k = CreationPolicy<Obj>::Create(); //如果没加域需加-fpermissive,且有警告
     k->print();
   }

};

int main()
{
  WidgetManager<Widget, OpNewCreator > wig; //不要写OpNewCreate<Widget> ,因此处需要的是一个模板,而不是类
  wig.wprint();
  return 0;
}

/*
g++ 20090810_policy.cpp 
*/
