#include <iostream>
#include <cstdio>

char* get_str(const int k)
{
    char my[]="the number is: ";
    char num[10]={0};
    sprintf(num,"%d",k);
    char const *p = my ;
    return p;
}

int main()
{
   for(int i = 0; i < 100; ++i)
   {
      char *k = get_str(i);
      std::cerr << k << std::endl;
   }
   return 0;
}
