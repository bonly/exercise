#include <cstdio>
#include <sys/ipc.h>

int
main (int argc, char* argv[])
{
      key_t key = ftok(argv[1], 1);
      printf("%x\n", key);
      return 0;
}
