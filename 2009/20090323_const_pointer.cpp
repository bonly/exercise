//============================================================================
// Name        : bdb.cpp
// Author      : bonly
// Version     :
// Copyright   : bonly's copyright notice
// Description : debug info
// fun(const unsigned char*):�����е�ָ��ֻ��һ��copy,���ǵ�ַ��ֵ
//============================================================================

#include <cstdlib>
#include <cstring>
#include <csignal>
#include <cstdio>
#include <cerrno>

struct Data{
const unsigned char* gen_buf()
{
	bzero(buf,20);
	strcpy(buf,"a test");
	return (const unsigned char*)buf;
}
char buf[20];
};

int get_abuff(const unsigned char** p,Data& k)//ֻ�мӶ�һ��ָ�������޸�ԭֵ
{
  *p = k.gen_buf();
  printf("in get_abuff: %s\n",*p);
  return 0;
}

int main()
{
	Data data;
  const unsigned char* p=0;
  get_abuff(&p,data);
  printf("%s,\n",p);
	return 0;
}

