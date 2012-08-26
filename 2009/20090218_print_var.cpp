#include <cstdio>
#include <iostream>

template<typename T> void _my_prt(const char* k, T v)
{
  std::cout << k <<": " << v << std::endl;
}
  
#define prt(v) _my_prt(#v,v)

#define prt2(v) printf(#v)

int
main()
{
   int myvar=12;
   prt(myvar);
   prt2(myvar);
   std::cout << std::endl;
   return 0;

}

