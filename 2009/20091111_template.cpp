#include <vector>
#include <stdexcept>

template <typename T>
class Stack
{
  private:
    std::vector<T> elems;    ///元素
  public:
    Stack();                 ///建构式
    void push (T const&);   ///push元素
    T pop ();                ///pop元素
    T top() const;          ///传回最顶端元素
    bool empty () const    ///stack是否为空
    { return elems.empty();}  ///注意是否需要拷贝及=运算子
};

template <typename T>
void Stack<T>::push (T const& elem)
{
  elems.push_back(elem);     ///追加(附于尾)
}

template <typename T>
Stack<T>::Stack()   ///@note 无论何时以这个class说明变量或函数时,都应写成Stack<T>
{

}
/**
 * @note 只有被呼叫到的成员函数式,才会被具现化,对class templates而言,只有
 * 当某个成员函数被使用时,才会进行具现化.
 */
template <typename T>
T Stack<T>::pop ()
{
  if (elems.empty())
  {
    throw std::out_of_range ("Stack<>::pop: empty stack");
  }
  T elem = elems.back();     ///保存最后元素的拷贝
  elems.pop_back();          ///移除最后一个元素
  return elem;              ///传回先前保存的最后元素
}

template <typename T>
T Stack<T>::top () const
{
  if (elems.empty())
  {
    throw std::out_of_range("Stack<>::top: empty stack");
  }
  return elems.back();       ///传回最后一个元素的拷贝
}

#include <iostream>
#include <string>
#include <cstdlib>

int main()
{
  try
  {
    Stack<int>  intStack;
    Stack<std::string> stringStack;

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
