#include <iostream>
using namespace std;

int
main()
{
   const char* a="this is \
                  a test";
   cout << a <<endl;
   const char* b="this is"
       " a test";
   cout << b <<endl;
   return 0;
}

