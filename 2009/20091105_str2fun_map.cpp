#include <map>
#include <iostream>
using namespace std;

class BaseFun;
void kfun()
{
    cerr << "this is in kfun\n";
}

class BaseFun
{
  public:
      typedef void (*FUN)(BaseFun*, int &addition);
      map<string,FUN> fun;

      BaseFun()
      {
      }

      void init()
      { 
          //fun.insert(make_pair("kfun",kfun));
          fun.insert(make_pair("chk",BaseFun::chk));
          ///每个子类把自己要特殊处理的子段(即函数指针加进去)
      }

      static void chk(BaseFun *obj, int &addition)
      {
        cerr << "in fun BaseFun()\n";
        obj->real_chk(addition);
      }

      void real_chk(int &addition)
      {
        ///这里才是每个子类实现不同操作的地方
        ///这里可以考虑不操作,只把另外要操作的函数指针按一定的顺序加到一个列表中,
        ///最后等所有数据都收集齐了再遍历列表处理所有调用
        cerr << "in real_chk, addition is " << addition <<endl;
      }
};


int main()
{
   int ki=38;
   BaseFun  bf;
   bf.init();
   BaseFun::FUN k=bf.fun["kfun"];

   if(k==NULL)
       cerr << "these is no fun name kfun\n";
   else 
       bf.fun["kfun"](&bf,ki);

   bf.fun["chk"](&bf,ki);
   return 0;
}
