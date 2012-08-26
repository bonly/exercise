#include <unistd.h>
#include <stdio.h>

int main()
{
    FILE *fl=fopen("./testfd.txt","ra+");
    sleep(10);
    fclose(fl);
    return 0;
}
