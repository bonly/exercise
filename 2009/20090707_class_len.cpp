#include <iostream>

class base
{
  public:
    char a[4];
};

class ch_base : public base
{
  public:
   char a[4];
};


int main()
{
  ch_base ch;
  std::cout << "base len: " << sizeof(base) << std::endl;
  std::cout << "ch_base len: " << sizeof(ch_base) << std::endl;
  std::cout << "ch_base var len: "  << sizeof(ch) << std::endl;
  return 0;
}
  
