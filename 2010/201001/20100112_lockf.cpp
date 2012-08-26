#include <sys/types.h>
#include <unistd.h>
#include <stdio.h>
#include <sys/stat.h>
#include <fcntl.h>

int main()
{
    int fd=open("./testfd.txt",O_RDWR);
    lockf(fd, F_TLOCK, 0);
    sleep(10);
    close(fd);
    return 0;
}
