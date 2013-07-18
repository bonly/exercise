#include <cstdio>

int main(){
   unsigned int ary[3]={3, 4, 5};
   unsigned int *p = ary;
   printf("p:%d\n", *p);
   printf("p1: %d\n", ++*p);
   p = ary;
   printf("p2: %d\n", *(p+2));
   return 0;
}
   
