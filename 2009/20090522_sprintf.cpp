#include <cstdio>
#include <cstring>
#include <strings.h>

int
main ()
{
  char buf[255];
  bzero(buf,255);
  strcpy(buf,"this is a test!");
  sprintf(buf,"%s%s","[13412]",buf);
  printf("string is: %s\n",buf);
  return 0;
}

/*
 * HP-UX上结果是:
 * string is: [13412][13412][13412][
 */

