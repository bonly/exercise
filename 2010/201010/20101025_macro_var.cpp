#include <iostream>
#define PR(X) #X

int main(){
   std::clog << PR(PRO) << std::endl;
   std::clog << PRO << std::endl;
   //std::clog << OTH << std::endl;
   return 0;
}

/*
g++ 20101025_macro_var.cpp -DPRO=\"DDDAR\"
*/
