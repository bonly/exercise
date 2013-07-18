#include <iostream>

class W{
  public:
    static int getI(){
      return 100;
    }
};

class V{
  public:
    template<class K> void PrintK(){
        std::clog << K::getI() << std::endl;
    }
};

void m(){
   V v;
   //v.template PrintK<W>(); ///vc 下可成功，gcc不认此方式
   v.PrintK<W>(); ///vc 下可成功，gcc不认此方式
}

int main(){
   m();
   return 0;
}

