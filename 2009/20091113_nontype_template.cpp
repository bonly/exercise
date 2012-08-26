#include <vector>
#include <stdexcept>

template <typename T, int MAXSIZE > /// 使用非型别模板参数（常规值：常数整数，包括enum或指向外部数据的指针。）
///不能是浮点数、字串(因是内部联结物)、常量浮点运算式,
///但外部连结物可以,如: extern char const s[]="hello" ; MyClass<s> x;
class Stack
{
  private:
    T elems[MAXSIZE];    ///元素
    int numElems;         ///当前的元素个数
  public:
    Stack();///建构式
    void push (T const& elem)   ///push元素
    {
      if (numElems == MAXSIZE)
        throw std::out_of_range("Stack<>::push():stack is full.");
      elems[numElems] = elem;
      ++numElems;
    }
    void pop ();                ///pop元素
    T top() const;          ///传回最顶端元素
    bool empty () const    ///stack是否为空
    {
      return numElems == 0;
    }
    bool full() const
    {
      return numElems == MAXSIZE;
    }
};

template <typename T,int MAXSIZE>
Stack<T,MAXSIZE>::Stack():numElems(0)   ///@note 无论何时以这个class说明变量或函数时,都应写成Stack<T,MAXSIZE>
{
}

template <typename T,int MAXSIZE>
void Stack<T,MAXSIZE>::pop ()
{
  if (numElems <= 0)
  {
    throw std::out_of_range ("Stack<>::pop: empty stack");
  }
  --numElems;
}

template <typename T,int MAXSIZE>
T Stack<T,MAXSIZE>::top () const
{
  if (numElems <= 0)
  {
    throw std::out_of_range("Stack<>::top: empty stack");
  }
  return elems[numElems -1];       ///传回最后一个元素的拷贝
}

#include <iostream>
#include <string>
#include <cstdlib>

int main()
{
  try
  {
    Stack<int,3>  intStack;
    Stack<std::string,4> stringStack;

    intStack.push(7);
    std::cout << intStack.top() << std::endl;

    stringStack.push("hello");
    std::cout << stringStack.top() << std::endl;
    intStack.pop();
    stringStack.pop();
  }
  catch (std::exception const& ex)
  {
    std::cerr << "Exception: " << ex.what() << std::endl;
    return EXIT_FAILURE;
  }
}

/**
 * 自己定的概念模板参数都想像为非指针...这样才可以在定义偏特化时使用指针作参数变量
 */

