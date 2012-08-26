#include<stdio.h>
#include<stdlib.h>
#include<errno.h>
#include<string.h>
#include<sys/types.h>
#include<netinet/in.h>
#include<sys/socket.h>
#include<sys/wait.h>
#include<arpa/inet.h>
#include<unistd.h>

#define SERVER_PORT 3002
#define MAX_CONNECT 5

int main()
{
    int serverSock, clientSock;
    struct sockaddr_in serverAddr;
    struct sockaddr_in clientAddr;

    serverSock = socket(AF_INET, SOCK_STREAM, 0);
    if (serverSock == -1)
    {
        printf("error when create socket/n");
        exit(1);
    }

    serverAddr.sin_family = AF_INET;
    serverAddr.sin_port = htons(SERVER_PORT);
    serverAddr.sin_addr.s_addr = INADDR_ANY;

    bzero(&(serverAddr.sin_zero), 8);
    if (bind(serverSock, (struct sockaddr *)&serverAddr, sizeof(struct sockaddr)) == -1)
    {
        perror("bind error/n");
        exit(1);
    }

    if (listen(serverSock, MAX_CONNECT) == -1)
    {
        perror("listen error/n");
        exit(1);
    }

    while(true)
    {
        size_t sin_size = sizeof(struct sockaddr_in);
        if ((clientSock = accept(serverSock, (struct sockaddr *)&clientAddr, &sin_size)) == -1)
        {
            perror("accept error");
            continue;
        }

        printf("receive a connection from %s/n", inet_ntoa(clientAddr.sin_addr));
        if (!fork())
        {
            if (send(clientSock, "Hello", 5, 0) == -1)
            {
                perror("send error/n");
                close(clientSock);
                exit(0);
            }
        }
        close(clientSock);
    }

    return 0;
}

