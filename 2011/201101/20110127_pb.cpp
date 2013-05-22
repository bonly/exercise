#include <iostream>
class A{
public:
  int a;
  virtual void update(int b){
    a = b;
  }
  virtual void testB()=0;
  virtual void updateB(int d){}
};

class B : virtual public A{
//class B : public A{
public:
  virtual void update(int d){
     A::a =d +2 ;
  }
  virtual void updateB(int d){
     A::a =d ;
  }
};

class C : virtual public A{  ///virtual 继承,只有一个a产生
//class C : public A{
public:
  virtual void update(int d){
     A::a =d +1 ;
  }
  virtual void updateA(int d){
     A::a =d +3;
  }
};

class D : virtual public A, B , C{
public:
  virtual void update(int d){
     A::a = d ;
  }
  virtual void testB(){
     A::updateB(4);
  }
};
   

int main(){
    A  *cc = new D;
    cc->testB();
    cc->update(56);
    delete cc;
    return 0;
}
