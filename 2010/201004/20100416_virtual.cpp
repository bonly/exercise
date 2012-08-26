/**
 *  @file 20100416_virtual.cpp
 *
 *  @date 2012-3-1
 *  @Author: Bonly
 */

#include <iostream>
using namespace std;

class A
{
  public:
    virtual void funa()
    {
      clog << "A::funa" << endl;
    }

    void funb()
    {
      clog << "A::funb" << endl;
    }
};

class B :public A
{
  public:
    virtual void funa()
    {
      clog << "B::funa" << endl;
    }

    void funb()
    {
      clog << "B::funb" << endl;
    }
};

int main()
{
  A* k = new B;
  k->funa();
  k->funb();

  delete k;

  return 0;
}
