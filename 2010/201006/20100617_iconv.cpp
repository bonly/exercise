#include <unistd.h>
#include <stdio.h>
#include <sys/types.h>
#include <iconv.h> //convert function
#include <sys/stat.h>
#include <fcntl.h>
#include <string.h>

#define S 2000

void convert(const char *fromset,const char *toset,char *from,int from_len,char *to,int to_len)
{
    printf("%s is to be converted!\n",from);
    iconv_t cd,cdd;
    cd=iconv_open(toset,fromset);
    char **from2=&from;
    char **to2=&to;
    if(iconv(cd,from2,(size_t*)&from_len,to2,(size_t*)&to_len)==-1)
        printf("Convert fail!\n");
    else
        printf("Convert success!\n");
    iconv_close(cd);
    return ;
}

int main()
{
    char from[]="你好";  
    char to[S];
    //convert("GB2312","BIG5",from,strlen(from),to,S);  //把gb2312转换成big5
    convert("UTF-8","GBK",from,strlen(from),to,S);  //把gb2312转换成big5
    printf("%s\n",to);
    return 0;
}
