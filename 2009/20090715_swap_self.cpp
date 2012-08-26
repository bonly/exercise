#include <cstdlib>
#include <iostream>
using namespace std;

int
main(int argc, char* argv[])
{
  int i = 3; int j=4;
  clog << i << "\t" << j << endl;
  swap(i,j);
  clog << i << "\t" << j << endl;

  swap(i,i);
  clog << i << "\t" << j << endl;
  return 0;
}

