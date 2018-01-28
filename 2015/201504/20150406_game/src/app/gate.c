#ifndef __GATE_C__
#define __GATE_C__
#include "gate.h"
#include <stdio.h>

Callback fn;

void bridge_callback(unsigned int sn, char *buffer){
    if (fn){
        printf("call back in %0x...\n", fn);
        return fn(sn, buffer);
    }
    return;
}

void SetCallBack(Callback pt){
    printf("in set call back %0x..\n", pt);
    if (pt){
        fn = pt;
    }
    return;
}
#endif
