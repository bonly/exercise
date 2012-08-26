#include <iostream>
using namespace std;


struct A
{
    bool operator==(int k)
    {
        cerr << "compare a with " << k << endl;
        return a==k;
    }
    int a;
};

struct B
{
    bool operator==(int k)
    {
        cerr << "compare b with " << k << endl;
        return b==k;
    }
    int b;
};

int main()
{
  A a;
  a.a=1;
  B b;
  b.b=2;

  if(!(a == 2 && b == 2))
  {
      cerr << "&& 操作从左至右,左边失败则右边不计算" << endl << endl;
  }

  if(b == 2 && a == 1)
  {
      cerr << "&& 操作从左至右,左边成功则右边计算" << endl << endl;
  }

  if(a == 2 || b == 2)
  {
      cerr << "|| 操作从左至右,左边失败右边也需计算" << endl << endl;
  }

  if(b == 2 || a == 2)
  {
      cerr << "|| 操作从左至右,左边成功右边不计算" << endl << endl;
  }
  return 0;
}
