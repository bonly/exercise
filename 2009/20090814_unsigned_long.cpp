#include <cstdio>
#include <cstdlib>

int
main()
{
  char telno[]="13719360007";
  unsigned long tel = atol(telno);
  printf("no is: %ld\n",tel);

  long long msisdn = atoll(telno);
  printf("num is: %lld\n",msisdn);
  return 0; 
}
