#include <iostream>

void fun()
{
  int i = 10;
  //void iner_fun() ///C++中只允许内嵌block/类,不允许内嵌函数.
  {
     std::clog << "i value: " << i;
  }
  class a
  {
      int k;
  };
}

int main()
{
    fun();
    return 0;
}
