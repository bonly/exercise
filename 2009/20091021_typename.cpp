#include<iostream>
namespace Debug
{
    template<typename Type>
        void print(const Type& v)
        {
            std::cerr << v << std::endl;
        }
}
#define debug_print(x) Debug::print(x)

int main()
{
  using namespace Debug;
  using namespace std;
  int a=10;
  char b='a';
  float c=12.0;
  debug_print (a);
  clog << endl;
  debug_print (b);
  clog << endl;
  debug_print (c);
  clog << endl;
  return 0;

}
