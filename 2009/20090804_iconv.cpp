#include <iostream>
#include <boost/format.hpp>
#include <iconv.h>

using namespace std;
using namespace boost;

int main(int argc, char* argv[])
{
   clog << "need to conv: " << argv[1] << endl;

   //iconv_t cd = iconv_open("VISCII","GB18030");
   iconv_t cd = iconv_open("UCS-2","GB18030");  //��iconv -l�鿴֧����Щת��
   if (cd == (iconv_t)-1)  //������ô��ֵļ�鷽��
   {
        clog << strerror(errno) << endl;
   }

   char in[255];
   memset(in,0,255);
   //strcpy(in,argv[1]);
   strcpy(in,"#1431 erts");  //������������ʱ��Ҫ��"������,#�ַ��Ų���������

   char* pin = in;
   size_t inbytesleft=strlen(in);
   size_t outbytesleft=255;

   char  out[255];
   memset(out,0,255);
   char* pout = out;

   clog << "before:\n";
   unsigned char* p = (unsigned char*)pin;
   for (int i=0; i<inbytesleft; ++i) fprintf(stderr,"[%02X]",p[i]);

   //inbytesleft ��Ҫת���ĳ���,����ִ�������ʣ��δת���ĳ���
   //outbytesleft ת����ŵĿ��ÿռ�,ִ������Ǵ�ſռ�ʣ��ĳ���
   //pin �� pout ��ת��ʱ���ᱻ�޸�,Ӧʹ��ԭʼ����ǰ��ָ��
   int len = iconv(cd, &pin, &inbytesleft, &pout, &outbytesleft); 
   if(len<0) clog <<  strerror(errno) << endl;

   //clog << format("old: %s\t conved: %s\n")%argv[1]%out;
   //fprintf(stderr,"old: %s\t conved: %s\n",argv[1],out);  
   clog << endl << "after:\n";
   p = (unsigned char*) out; //pin �� pout ��ת��ʱ���ᱻ�޸�,Ӧʹ��ԭʼ����ǰ��ָ��
   for (i=0; i<255-outbytesleft; ++i) fprintf(stderr,"[%02X]",p[i]);
   clog << endl;



   int ret = iconv_close(cd);
   return 0;
}