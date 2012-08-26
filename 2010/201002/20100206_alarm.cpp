#include <unistd.h>


int fun()
{
  alarm(3);
  while(true)
  {
    sleep(10);
  }
  return 0;
}


int main()
{
  fun();
  return 0;
}
