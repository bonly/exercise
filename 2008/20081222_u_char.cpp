//============================================================================
// Name        : exch.cpp
// Author      : bonly
// Version     :
// Copyright   : bonly's copyright notice
// Description : Hello World in C++, Ansi-style
//============================================================================

#include <iostream>
unsigned int FourByteToInt(unsigned char* pszInput)
{
  unsigned int nLen = (pszInput[0] << 24) + (pszInput[1] << 16) +
    (pszInput[2] << 8) + pszInput[3];
//  unsigned int nLen = (pszInput[0] << 24) | (pszInput[1] << 16) |
//    (pszInput[2] << 8) | pszInput[3];
  return nLen;
}
int
main ()
{
  unsigned int i = 3214;
  printf("i:%x\n",i);

  unsigned char* p = (unsigned char*)&i;
  printf("p[0]:%x\t",p[0]);
  printf("p[1]:%x\t",p[1]);
  printf("p[2]:%x\t",p[2]);
  printf("p[3]:%x\t\n",p[3]);

  unsigned int af = (p[0]<<24)+(p[1]<<16)+(p[2]<<8)+(p[3]);
  unsigned int kf = ((p[0]<<24)&0xffffff) | ((p[1]<<16)&0xffffff) | ((p[2]<<8)&0xffffff) | ((p[3])&0xffffff);
  unsigned int ef = (p[0]<<24) | (p[1]<<16) | (p[2]<<8) | (p[3]);

  printf("af:%x\n",af);
  printf("af(lu):%lu\n",af);
  printf("kf:%x\n",kf);
  printf("ef:%x\n",ef);

  {
  	printf("2th test:\n");
    unsigned char p[4];
    memset(p,0,4);
    p[0] = (i>>24)&0xffffffff ;
    p[1] = (i>>16)&0xffffffff ;
    p[2] = (i>>8 )&0xffffffff ;
    p[3] = i      &0xffffffff ;

    printf("x:%x\n",(unsigned int)*(reinterpret_cast<unsigned int*>(p)));
    //printf("lu:%lu\n",(unsigned int)reinterpret_cast<unsigned int*>(p));
    printf("p[0]:%x\t",p[0]);
    printf("p[1]:%x\t",p[1]);
    printf("p[2]:%x\t",p[2]);
    printf("p[3]:%x\t\n",p[3]);
    printf("p:%x\n",(unsigned int)*(reinterpret_cast<unsigned int*>(p)));
  }
  {
  	printf("new decode:\n");
  	unsigned char ch[4]={0x8e,0x0c,0x00,0x00};
  	printf("ch:%x\n",(unsigned int)*(reinterpret_cast<unsigned int*>(ch)));
  	unsigned char chr[4];
  	chr[0]=ch[3];
  	chr[1]=ch[2];
  	chr[2]=ch[1];
  	chr[3]=ch[0];
  	printf("chr:%x\n",(unsigned int)*(reinterpret_cast<unsigned int*>(chr)));
  	unsigned int ifd=FourByteToInt(chr);
  	printf("ifd:%x\n",ifd);
  }

}
