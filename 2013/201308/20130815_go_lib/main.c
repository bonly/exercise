#include <stdio.h>
#include "20130815_go_lib.h"

int main()
{
    ListenAndServe(":8000");
    return 0;
}

/*
in mac os:
gcc -o gohttp-c main.c 20130815_go_lib.a \
      -framework CoreFoundation -framework Security -lpthread
*/
/*
gcc -o gohttp-c main.c 20130815_go_lib.a -lpthread
*/
