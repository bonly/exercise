/**
 *  @file 20100414_enum.cpp
 *
 *  @date 2012-2-27
 *  @Author: Bonly
 */

#include <iostream>

#define AN(X) _##X

using namespace std;
class Obj
{
  public:
    enum EM{begin=0}step;
    int k;
};

class cat : public Obj
{
  public:
    enum EM{end=10};
};

class dog : public Obj
{
  public:
    enum EM{start=10};
};

struct fish : public Obj
{
    int IDX;
};

struct bird : public Obj
{
    float DX;
};

enum Animal{
#define MYENUM(e) _##e,
#include "20100414_enum.h"
#undef MYENUM
};

const char* AnimalDescription[]={
#define MYENUM(e) "n_"#e,
#include "20100414_enum.h"
#undef MYENUM
};

enum idx{
#define MYENUM(e) e_##e,
#include "20100414_enum.h"
#undef MYENUM
};

///定义变量
#define MYENUM(e) e* var_##e;
#include "20100414_enum.h"
#undef MYENUM

///使用变量
#define MYENUM(e) var_##e,
Obj* lst[]={
#include "20100414_enum.h"
};
#undef MYENUM


int main()
{
  dog b;
  b.k = 0;
  b.k = 10;
  switch (b.k)
  {
    case dog::begin:
      clog << "begin" << endl;
      break;
    case dog::start:
      clog << "end" << endl;
      break;
  }

  clog << AnimalDescription[_dog] << endl;

  var_bird = new bird;
  var_bird->DX = 30;
  clog << var_bird->DX << endl;
  delete var_bird;
  return 0;
}

/**
考虑下面的枚举类型：

enum Animal { dog, cat, fish, bird };

现在dog在0的位置。编译器本身是强类型安全的，但是宏却不是。对于枚举类型，VS的调试器会显示出具体的类型而不是一个个整数长了。那么用于输出枚举常量的函数需要一个更人性的格式转换。下面的代码或许有帮助：

wchar_t* AnimalDiscription[] = { L"dog", L"cat", L"fish", L"bird" };

有了这个数组，在调试的时候就能够输出友好的字符串了，而我们所需要的只是根据枚举量查找数组下标。

使用宏自动生成枚举类型我就只需要维护一份实体，这里不仅仅是说枚举常量还包括对应的字符串。 考虑文件animal.inc:

MYENUM(dog)

MYENUM(cat)

MYENUM(fish)

MYENUM(bird)

和对应的C++ 代码：

enum Animal {

#define MYENUM(e) _##e,

#include "animal.inc"

#undef MYENUM

};

wchar_t* AnimalDescription[] = {

#define MYENUM(e) L"_" L#e,

#include "animal.inc"

#undef MYENUM

};

现在只要编辑animal.inc就能同时更新枚举类型和对应的描述。我在常量前面加了下划线以使宏工作正常。这是因为token-pasting操作符##不能出现在宏首。通过#我们可以产生字符串。在#前面加一个L可以将字符串声明为宽字符。

这些宏可以通过编译器的参数/P和/EP来调试。编译器会产生如下的预处理文件：

enum Animal {

_dog,

_cat,

_fish,

_bird,

};

wchar_t* AnimalDescription[] = {

L"_" L"dog",

L"_" L"cat",

L"_" L"fish",

L"_" L"bird",

};

宏字符串替换技巧也可以用于产生代码。下面是一个用字符串替换产生函数原型的例子：

#define MYENUM(e) void Order_##e();

#include "animal.inc"

#undef MYENUM

This expands to:

void Order_dog();

void Order_cat();

void Order_fish();

void Order_bird();

你可能会想根据动物的种类做一些事情。下面有一个switch的例子:

#define MYENUM(e) case _##e:\

Order_##e();\

break;

#include "animal.inc"

#undef MYENUM

This expands to:

case _dog: Order_dog(); break;

case _cat: Order_cat(); break;

case _fish: Order_fish(); break;

case _bird: Order_bird(); break;

在这个例子里面还需要给每一个函数写个定义包括 Order_dog()、Order_cat()、等等。如果你给animal.inc加了一种新的动物，记得要写个新的Order_开头的函数的定义。不过，链接器报错会提醒你这个事情的。

宏字符串替换是一个非常好用的工具， 特别是可以将内部数据集中保存，大幅度的减少出错的机会。

 */

/**
1、# （stringizing）字符串化操作符。其作用是：将宏定义中的传入参数名转换成用一对双引号括起来参数名字符串。其只能用于有传入参数的宏定义中，且必须置于宏定义体中的参数名前。
如：
#define example(instr) printf("the input string is:\t%s\n",#instr)
#define example1(instr) #instr
当使用该宏定义时：
example(abc)； 在编译时将会展开成：printf("the input string is:\t%s\n","abc");
string str=example1(abc)； 将会展成：string str="abc"；
注意：
对空格的处理
a。忽略传入参数名前面和后面的空格。
   如：str=example1(   abc )； 将会被扩展成 str="abc"；
b.当传入参数名间存在空格时，编译器将会自动连接各个子字符串，用每个子字符串中只以一个空格连接，忽略其中多余一个的空格。
   如：str=exapme( abc    def); 将会被扩展成 str="abc def"；

2、## （token-pasting）符号连接操作符
宏定义中：参数名，即为形参，如#define sum(a,b) (a+b)；中a和b均为某一参数的代表符号，即形式参数。
而##的作用则是将宏定义的多个形参成一个实际参数名。
如：
#define exampleNum(n) num##n
int num9=9;
使用：
int num=exampleNum(9); 将会扩展成 int num=num9;
注意：
1.当用##连接形参时，##前后的空格可有可无。
如：#define exampleNum(n) num ## n 相当于 #define exampleNum(n) num##n
2.连接后的实际参数名，必须为实际存在的参数名或是编译器已知的宏定义

// preprocessor_token_pasting.cpp
#include <stdio.h>
#define paster( n ) printf_s( "token" #n " = %d", token##n )
int token9 = 9;

int main()
{
   paster(9);
}
运行结果：
token9 = 9

3、@# （charizing）字符化操作符。
只能用于有传入参数的宏定义中，且必须置于宏定义体中的参数名前。作用是将传的单字符参数名转换成字符，以一对单引用括起来。
#define makechar(x)  #@x
a = makechar(b);
展开后变成了：
a= 'b';

4、\ 行继续操作符
当定义的宏不能用一行表达完整时，可以用"\"表示下一行继续此宏的定义。
 */
