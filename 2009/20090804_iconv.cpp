#include <iostream>
#include <boost/format.hpp>
#include <iconv.h>

using namespace std;
using namespace boost;

int main(int argc, char* argv[])
{
   clog << "need to conv: " << argv[1] << endl;

   //iconv_t cd = iconv_open("VISCII","GB18030");
   iconv_t cd = iconv_open("UCS-2","GB18030");  //用iconv -l查看支持哪些转换
   if (cd == (iconv_t)-1)  //就是这么奇怪的检查方法
   {
        clog << strerror(errno) << endl;
   }

   char in[255];
   memset(in,0,255);
   //strcpy(in,argv[1]);
   strcpy(in,"#1431 erts");  //在命令行输入时需要用"包起来,#字符才不会有意外

   char* pin = in;
   size_t inbytesleft=strlen(in);
   size_t outbytesleft=255;

   char  out[255];
   memset(out,0,255);
   char* pout = out;

   clog << "before:\n";
   unsigned char* p = (unsigned char*)pin;
   for (int i=0; i<inbytesleft; ++i) fprintf(stderr,"[%02X]",p[i]);

   //inbytesleft 需要转换的长度,函数执行完后是剩余未转换的长度
   //outbytesleft 转换存放的可用空间,执行完后是存放空间剩余的长度
   //pin 及 pout 在转换时都会被修改,应使用原始传入前的指针
   int len = iconv(cd, &pin, &inbytesleft, &pout, &outbytesleft); 
   if(len<0) clog <<  strerror(errno) << endl;

   //clog << format("old: %s\t conved: %s\n")%argv[1]%out;
   //fprintf(stderr,"old: %s\t conved: %s\n",argv[1],out);  
   clog << endl << "after:\n";
   p = (unsigned char*) out; //pin 及 pout 在转换时都会被修改,应使用原始传入前的指针
   for (i=0; i<255-outbytesleft; ++i) fprintf(stderr,"[%02X]",p[i]);
   clog << endl;



   int ret = iconv_close(cd);
   return 0;
}