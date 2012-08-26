#include <sys/types.h>
#include <sys/socket.h>
#include <stdio.h>
#include <arpa/inet.h>
#include <string.h>
#include <poll.h>
#include <unistd.h>

#define BUFSIZE 400

int main(int argc, char** argv)
{
    struct sockaddr_in srvadd;
    int sfd;
    char buf[BUFSIZE]="";
    int cnt;

    bzero(&srvadd, sizeof(srvadd));
    srvadd.sin_family = AF_INET;
    srvadd.sin_port = ntohs(2235);
    if (inet_aton(argv[1], &srvadd.sin_addr) == 0)
    {
        perror("Fail to get the server address by argument");
        return -1;
    }

    if ((sfd = socket(PF_INET, SOCK_STREAM, 0)) == -1)
    {
        perror("Fail to make the socket fd");
        return -1;
    }

    if (connect(sfd, (struct sockaddr*) (&srvadd), sizeof(srvadd)) == -1)
    {
        perror("Fail to connect to the server");
        return -1;
    }


    while((    cnt=read(STDIN_FILENO,buf,BUFSIZE)) >0)
    {
        write(sfd,buf,cnt);
        bzero(buf,BUFSIZE);
    }

    close(sfd);

    return 0;
}