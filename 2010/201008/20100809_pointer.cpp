#include <iostream>
using namespace std;

struct Ac{
  int  a;
  int  b;
  int  c;
};

int main(){
  Ac k;
  k.a = 11;
  k.b = 24;
  k.c = 32;
  int *p = &k.a;

  for (int i=0; i<3; ++i, ++p) {
     clog << i << ": " << *p << endl;
  }
}
