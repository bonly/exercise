/*
 * peek从流中预提数据，并不真正取走数据，可用来看是否为结束符eof
 *
 */
#include <fstream>
#include <iostream>
using namespace std;

int main(void)
{
  char ch, temp;
  while (cin.get(ch))
  {
       temp = cin.peek();
       cout<<temp;
  }
  return 0;
}

