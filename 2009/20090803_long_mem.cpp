#include <boost/format.hpp>
#include <iostream>
#include <cstdio>
using namespace std;

int
main()
{
  unsigned long org=0x1U;
  unsigned char *p = (unsigned char*) &org;
  for(int i=0; i<8; ++i) fprintf(stderr,"[%02X]",p[i]);
  cout << endl;
  
  unsigned long long lorg=0x1LLU;
  unsigned char *lp =(unsigned char*) &lorg;
  for(int i=0; i<8; ++i) fprintf(stderr,"[%02X]",lp[i]);
  cout << endl;

  return 0;
}

