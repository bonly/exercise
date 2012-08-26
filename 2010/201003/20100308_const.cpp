#include <iostream>
using namespace std;

/**
 *  定义是从右往左读,左面所有的括起来
 */

int main()
{
   const int a = 10; /// a 不可变
   //a = 0; 失败
   int b = 2;

   const int* c = &b; //c指针所在的内容不可变 == (const int)(*)
   //*c = 3; //指针所在的内容不可变
   c = &b; //指针指向可变

   int* const d = &b;  // == (int*) (const)
   *d = 3;  //指针所在的内容可变
   //d = &b;  //指针的指向不可变

   return 0;
}

