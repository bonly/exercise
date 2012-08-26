#include <sys/types.h>
#include <unistd.h>
#include <stdio.h>
#include <sys/stat.h>
#include <fcntl.h>
#include <sys/file.h>

int main()
{
    int fd=open("./testfd.txt",O_RDWR);
    flock(fd, LOCK_EX);
    sleep(10);
    close(fd);
    return 0;
}
