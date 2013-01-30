#include <iostream>

#define DEFPRINT(PR) \
template <typename T> void print_var(const char* k, T v) \
{ \
  PR << k << " : " << v ; \
}
#define PRINT_VAR(k) print_var(#k, k)

class MyC{
public:
   int cint;
};

DEFPRINT(std::clog);
int main (){
   int myint = 20;
   PRINT_VAR(myint);
   std::clog << std::endl;

   MyC ac;
   ac.cint = 30;
   PRINT_VAR(ac.cint);
   std::clog << std::endl;
   return 0;
}

