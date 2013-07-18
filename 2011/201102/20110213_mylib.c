#include "20110213_mylib.h"
int sum(int a, int b){
   return a+b;
}

/*
gcc -shared 20110213_mylib.c -o libmyl.so
*/
