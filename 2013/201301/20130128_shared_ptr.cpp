#include <boost/shared_ptr.hpp>
#include <iostream>
#include <boost/enable_shared_from_this.hpp>
using namespace boost;

class A;
// void DoSomething(shared_ptr<A>&) {  // same err
void DoSomething(shared_ptr<A> ap) { //shared_ptr as param. did not call A()creator
	std::clog << "count: " << ap.use_count() << std::endl; //2
	std::clog << "in Do" << std::endl;  
    //do something
}
 
// class A : public enable_shared_from_this<A> {  //it is the same err
class A {
public:
         void doSomething() {
                 shared_ptr<A> ptr_a(this);
             	 std::clog << "count: " << ptr_a.use_count() << std::endl;   //1
                 // shared_ptr<A> ptr_a(this->shared_from_this()); //tr1::bad_weak_ptr
                 DoSomething(ptr_a);
             	 std::clog << "count: " << ptr_a.use_count() << std::endl;  //1
         }//函数结束的时候，xxx ptr_a被销毁的同时也销毁了自己 A
         void interface(){
         	std::clog << "begin" << std::endl;
         	doSomething();
         	std::clog << "end" << std::endl;   //can not get here
         }
         A(){
         	std::cerr << "create" << std::endl;
         }
         ~A(){
         	std::cerr << "del" << std::endl;
         }
};
 
int main() {
         A a;
         a.interface();
         //continue do something with a, but it was already destory
         std::clog << "in main" << std::endl;  //can't get here
}


/*
http://blog.csdn.net/rogeryi/article/details/1444525
不要在属于类型A的接口的一部分的非成员函数或者跟A有紧密联系的辅助函数里面使用xxx_ptr<A>作为函数的参数类型。

A* pa = new A;
xxx_ptr<A> ptr_a_1(pa);
xxx_ptr<A> ptr_a_2(pa);
 
很明显，在ptr_a_1和ptr_a_2生命周期结束的时候都会去删除pa，pa被删除了两次，这肯定会引起你程序的崩溃
*/
