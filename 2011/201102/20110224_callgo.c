#include <stdio.h>
extern int Myfunc() __asm__("main.Myfunc");

int main(){
  int tmp;
  printf("calling the function\n");
  tmp = Myfunc();
  printf("called the function, got %d\n", tmp);
  return 0;
}

/*
gcc -g 20110224_callgo.c -o a.out -L. -lmy
没有-g也行...但新版中go中必须有main.main
*/

