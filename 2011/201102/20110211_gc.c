#include "20110211_gc.h"
#include <stdio.h>

void mprints(char* str){
    printf("%s\n", str);
}

/*
gcc -shared -fPIC -O2 -o libmgc.so 20110211_gc.c
*/
