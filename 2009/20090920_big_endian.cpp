#include <iostream>
using namespace std;

bool check()
{
  unsigned short test = 0x1122;
  if(*((unsigned char*)&test)==0x11)
      return true;  //大端big-endian
  else
      return false; //小端little-endian
}

int main()
{
  if(check())
      clog << "字节顺序是大端的\n";
  else
      clog << "字节顺序是小端的\n";
  return 0;
}

/*
 * 网络字节顺序是大端的
 * 所有htol类函数都是把小端字节转换为大端
 */
