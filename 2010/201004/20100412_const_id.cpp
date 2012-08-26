/**
 *  @file 20100412_const_id.cpp
 *
 *  @date 2012-2-27
 *  @Author: Bonly
 */
#include <iostream>
using namespace std;

class ObjBase
{
  public:
    const int id;
    virtual void test()=0;
    ObjBase(int i):id(i){}
};

class Bird : public ObjBase
{
  public:
    virtual void test()
    {
      cout << id << endl;
    }
    Bird(int k):ObjBase(k){};

};

int main()
{
  Bird a(4);

  return 0;
}
