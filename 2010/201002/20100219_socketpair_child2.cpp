#include <iostream>
#include <cstdio>
#include <cstdlib>
using std::cerr;
using std::endl;

int main(int argc, char* argv[])
{
    //FILE* fdr = fdopen(atoi(argv[1]),"r"); ///没有用
    FILE* fdw = fdopen(atoi(argv[1]),"w");

    char buf[]="hello from child2";
    write (fileno(fdw), buf, sizeof(buf));
    while(true)
    {
        sleep(5);
    }

    //fclose(fdr);
    fclose(fdw);
    return 0;
}

