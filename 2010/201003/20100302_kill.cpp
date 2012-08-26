#include <sys/types.h>
#include <signal.h>
#include <cstdlib>
#include <cerrno>
#include <iostream>
using namespace std;

int main(int argc, char* argv[])
{
    char *pchar = argv[1];
    char *pchar_end = pchar+10;
    errno = 0;
    int pid = strtol(pchar, &pchar_end, 10);
    if(errno!=0)
    {
        perror(strerror(errno));
        return 0;
    }
    errno = 0;
    if(0!=kill (pid,SIGTERM))
    {
        cerr << strerror(errno);
    }
    return 0;


}
