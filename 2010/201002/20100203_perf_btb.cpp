#include <stdio.h> 
#include <stdlib.h> 

void foo() 
{ 
    int i,j; 
    for(i=0; i< 20; i++) 
        j+=2; 
} 
int main(void) 
{ 
    int i; 
    for(i = 0; i< 100000000; i++) 
        foo(); 
    return 0; 
} 
