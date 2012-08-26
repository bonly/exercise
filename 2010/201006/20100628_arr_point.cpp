#include <iostream>
using namespace std;

char *st[3];
void fun(char* a=0)
{
   st[0]=a;
   clog << st[0] << endl;
}

int main()
{
  fun("test");
  return 0;
}

