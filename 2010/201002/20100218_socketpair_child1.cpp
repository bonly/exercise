#include <iostream>
#include <cstdio>
#include <cstdlib>
#include <string.h>
using std::cerr;
using std::endl;

int main(int argc, char* argv[])
{
    //FILE* fdr = fdopen(atoi(argv[1]),"r"); //没有用
    FILE* fdw = fdopen(atoi(argv[1]),"w");

    char buf[255]="";
    read (fileno(fdw), buf, 255);/// 0->1
    cerr << "read from father: " << buf << endl;
    
    char buf1[]="hello from child1";
    write (fileno(fdw), buf1, strlen(buf1)); /// 1->0
    while(true)
    {
        
        sleep(5);
    }

    //fclose(fdr);
    fclose(fdw);
    return 0;
}

