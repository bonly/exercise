#include <iostream>
#include <string.h>

class A{
//struct A{
public:
  int ia;
  char cb[12];
  int ic;
  //virtual 
  int proc(){ //一旦开启virtual就不能当作struct
     std::clog << "in A" <<  ia << cb << std::endl;
     return 0;
   }
};


class B : public A{
public:
   B(const A a):A(a){}
   B(){}
   int proc(){
     std::clog << "in B" << ia << cb << std::endl;
     return 0;
   }
};

int main(){
  A ds={12,"",33}; //作为class是不能以这种方式初始化
  memcpy(&ds.cb, "hello", sizeof(ds.cb));
  ds.proc();
  
  //B cs={13, "ok", 34}; //有了继承就会当作class,不能此方式初始化
  //B cs(A);
  A* cs = new B(ds);
  memcpy(cs->cb, "world", sizeof(cs->cb));
  cs->proc(); //没有virtual,调是的A的

  return 0;
}
  
