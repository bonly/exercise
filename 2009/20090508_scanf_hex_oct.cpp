#include <iostream>
#include <string>
using namespace std;
/* 十六进表示的字符串转数值
#include<stdio.h>
#include<string.h>
#include<stdlib.h>
char x[30]="0x32,0x33,0x35,0x36";
int i,j,len,d;
unsigned char y[10],c;
int main()
{
    len=strlen(x);
    for (i=j=0;j<len;i++)
    {
        sscanf(x+j,"%X",&d);
        y[i]=d;
        j+=5;
    }
    return 0;
}
*/
/*  //十六进转十进
int  Hex2oct（unsigned char  *dest， unsigned char  *src）
{
     int n;

     sscanf(src, "%x", &n);
     sprintf(dest, "%d", n);

     return n;
}
 */



int main(int argc, char argv[])
{
	char a[21]="";
	char b[22]="";
	char e[22]="";
	int  c=-1;
	int  d=-1;
	int  f=-1;
  //sscanf("1|2|3|20090702234215|||600|","%d|%*d|%d|%[^|]|%[^|]||%s",&c,&d,a,e,b);
  sscanf ("3|2|ab|de|","%d|%d|%[^|]|%[^|]|",&c,&d,a,b);
  cout << "a: " << a <<endl;
  cout << "b: " << b <<endl;
  cout << "c: " << c <<endl;
  cout << "d: " << d <<endl;
  cout << "e: " << e <<endl;
  cout << "f: " << f <<endl;
  return 0;

}


