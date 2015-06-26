#include<stdio.h>
#include<stdlib.h>
int main()
{
	        int *ptr = NULL;
	        *ptr = 0;
		return 0;
}


/*
 *访问不存在的地址
dmesg:
 [88829.585759] a.out[1008]: segfault at 0 ip 080484ab sp bfa1e138 error 6 in a.out[8048000+1000]
 使用objdump生成二进制的相关信息，重定向到文件中:
  objdump -d ./a.out > segfault3Dump
在segfault3Dump文件中查找发生段错误的地址:
grep -n -A 10 -B 10 "080484ab" ./segfault3Dump
段错误发生main函数，对应的汇编指令是 80484ab所在的行

catchsegv ./segfault3 
*/

