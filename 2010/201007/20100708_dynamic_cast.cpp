#include <iostream>
using namespace std;
class CBase 
{ 
   public:
     virtual ~CBase(){} //source type is not polymorphic,必须有这个才不报错
};
class CDerived: public CBase { };

int main()
{
	CBase *k = new CDerived;
	CBase b; CBase* pb;
	CDerived d; CDerived* pd;

	pb = dynamic_cast<CBase*>(&d);     // ok: derived-to-base
	pd = dynamic_cast<CDerived*>(&b);  // wrong: base-to-derived 
	clog << pb << endl;
	clog << pd << endl; ///=0 因为不安全(父类变子类) 但如果是要变为void*则不受检测
	pd = dynamic_cast<CDerived*>(k);
	clog << pd << endl;  /// new的对象可以识别出来
	return 0;
}
