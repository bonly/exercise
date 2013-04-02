#include <iostream>
#include <vector>
#include <algorithm>
using namespace std;

class CTest {
public:
 CTest() : m_nData(20) { NULL; }
 void TestLambda() {
  vector<int> vctTemp;
  vctTemp.push_back(1);
  vctTemp.push_back(2);

  // 无函数对象参数，输出：1 2
  {
   for_each(vctTemp.begin(), vctTemp.end(), [](int v){ cout << v << endl; });
  }

  // 以值方式传递作用域内所有可见的局部变量（包括this），输出：11 12
  {
   int a = 10;
   for_each(vctTemp.begin(), vctTemp.end(), [=](int v){ cout << v+a << endl; });
  }

  // 以引用方式传递作用域内所有可见的局部变量（包括this），输出：11 13 12
  {
   int a = 10;
   for_each(vctTemp.begin(), vctTemp.end(), [&](int v)mutable{ cout << v+a << endl; a++; });
   cout << a << endl;
  }

  // 以值方式传递局部变量a，输出：11 13 10
  {
   int a = 10;
   for_each(vctTemp.begin(), vctTemp.end(), [a](int v)mutable{ cout << v+a << endl; a++; });
   cout << a << endl;
  }

  // 以引用方式传递局部变量a，输出：11 13 12
  {
   int a = 10;
   for_each(vctTemp.begin(), vctTemp.end(), [&a](int v){ cout << v+a << endl; a++; });
   cout << a << endl;
  }

  // 传递this，输出：21 22
  {
   for_each(vctTemp.begin(), vctTemp.end(), [this](int v){ cout << v+m_nData << endl; });
  }

  // 除b按引用传递外，其他均按值传递，输出：11 12 17
  {
   int a = 10;
   int b = 15;
   for_each(vctTemp.begin(), vctTemp.end(), [=, &b](int v){ cout << v+a << endl; b++; });
   cout << b << endl;
  }


  // 操作符重载函数参数按引用传递，输出：2 3
  {
   for_each(vctTemp.begin(), vctTemp.end(), [](int &v){ v++; });
   for_each(vctTemp.begin(), vctTemp.end(), [](int v){ cout << v << endl; });
  }

  // 空的Lambda表达式
  {
   [](){}();
   []{}();
  }
 }

private:
 int m_nData;
};


int main(){
   CTest test;
   test.TestLambda();
   return 0;
}

/*
       C++ 11中的Lambda表达式用于定义并创建匿名的函数对象，以简化编程工作。Lambda的语法形式如下：
              [函数对象参数] (操作符重载函数参数) mutable或exception声明 ->返回值类型 {函数体}
      可以看到，Lambda主要分为五个部分：[函数对象参数]、(操作符重载函数参数)、mutable或exception声明、->返回值类型、{函数体}。下面分别进行介绍。
      一、[函数对象参数]，标识一个Lambda的开始，这部分必须存在，不能省略。函数对象参数是传递给编译器自动生成的函数对象类的构造函数的。函数对象参数只能使用那些到定义Lambda为止时Lambda所在作用范围内可见的局部变量（包括Lambda所在类的this）。函数对象参数有以下形式：
           1、空。没有使用任何函数对象参数。
           2、=。函数体内可以使用Lambda所在作用范围内所有可见的局部变量（包括Lambda所在类的this），并且是值传递方式（相当于编译器自动为我们按值传递了所有局部变量）。
           3、&。函数体内可以使用Lambda所在作用范围内所有可见的局部变量（包括Lambda所在类的this），并且是引用传递方式（相当于编译器自动为我们按引用传递了所有局部变量）。
           4、this。函数体内可以使用Lambda所在类中的成员变量。
           5、a。将a按值进行传递。按值进行传递时，函数体内不能修改传递进来的a的拷贝，因为默认情况下函数是const的。要修改传递进来的a的拷贝，可以添加mutable修饰符。
           6、&a。将a按引用进行传递。
           7、a, &b。将a按值进行传递，b按引用进行传递。
           8、=，&a, &b。除a和b按引用进行传递外，其他参数都按值进行传递。
           9、&, a, b。除a和b按值进行传递外，其他参数都按引用进行传递。
      二、(操作符重载函数参数)，标识重载的()操作符的参数，没有参数时，这部分可以省略。参数可以通过按值（如：(a,b)）和按引用（如：(&a,&b)）两种方式进行传递。
      三、mutable或exception声明，这部分可以省略。按值传递函数对象参数时，加上mutable修饰符后，可以修改按值传递进来的拷贝（注意是能修改拷贝，而不是值本身）。exception声明用于指定函数抛出的异常，如抛出整数类型的异常，可以使用throw(int)。
      四、->返回值类型，标识函数返回值的类型，当返回值为void，或者函数体中只有一处return的地方（此时编译器可以自动推断出返回值类型）时，这部分可以省略。
      五、{函数体}，标识函数的实现，这部分不能省略，但函数体可以为空。
*/
