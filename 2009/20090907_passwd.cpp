#include <fstream>
using namespace std;

int
main (int argc, char* argv[])
{
  ofstream of(argv[1]);
  for (int i = 0; i<100000000; ++i)
  {
    char num[9]={0};
    sprintf(num,"%08d",i);
    of << num << endl;
  }
  return 0;
}

