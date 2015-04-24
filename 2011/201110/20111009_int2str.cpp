#include <cstdio>
#include <iostream>
using namespace std;

int main(){
  char str[25]="";
  sprintf(str, "%x", 100); //将100转为16进制字符串
  cout << str << endl;
  return 0;
}
