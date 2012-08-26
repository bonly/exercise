#include <unistd.h>
#include <ctype.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <fcntl.h>
#include <limits.h>
#include <sys/types.h>
#include <sys/stat.h>

#define SERVER_FIFO_NAME "/tmp/server_fifo"
#define CLIENT_FIFO_NAME "/tmp/client_%d_fifo"

#define BUFFER_SIZE PIPE_BUF
#define MESSAGE_SIZE 20
#define NAME_SIZE 256

typedef struct message
{
    pid_t client_pid;
    char data[MESSAGE_SIZE + 1];
}message;


//#include "client.h"

int main()
{
    int server_fifo_fd;
    int client_fifo_fd;

    int res;
    char client_fifo_name[NAME_SIZE];

    message msg;

    char *p;

    
    /*
    if (mkfifo(SERVER_FIFO_NAME, 0777) == -1)
    {
        fprintf(stderr, "Sorry, create server fifo failure!/n");
        exit(EXIT_FAILURE);
    }
    */
    //明确设置umask,因为不知道谁会读写管道
    umask(0);
    if (mkfifo(SERVER_FIFO_NAME, S_IRUSR|S_IWUSR|S_IRGRP | S_IWGRP))
    {
        perror("mkfifo");
        exit(1);
    }
    /*
    //在Linux下不推荐使用mknod，因为其中有许多臭虫在NFS下工作更要小心，能使用mkfifo就不要用mknod，因为mkfifo()是POSIX.1 标准。
    umask (0);
    if (mknod(SERVER_FIFO_NAME,S_IFIFO | S_IRUSR | S_IWUSR | S_IRGRP | S_IWGRP,0))
    {
        perror("mknod");
        exit(1);
    }
    */
    server_fifo_fd = open(SERVER_FIFO_NAME, O_RDONLY);
    if (server_fifo_fd == -1)
    {
        fprintf(stderr, "Sorry, server fifo open failure!/n");
        exit(EXIT_FAILURE);
    }

    sleep(5);

   while (res = read(server_fifo_fd, &msg, sizeof(msg)) > 0)
   {
        p = msg.data;
        while (*p)
        {
        *p = toupper(*p);
        ++p;
        }

        sprintf(client_fifo_name, CLIENT_FIFO_NAME, msg.client_pid);
        client_fifo_fd = open(client_fifo_name, O_WRONLY);
        if (client_fifo_fd == -1)
        {
        fprintf(stderr, "Sorry, client fifo open failure!/n");
        exit(EXIT_FAILURE);
        }

        write(client_fifo_fd, &msg, sizeof(msg));
        close(client_fifo_fd);
    }

    close(server_fifo_fd);
    unlink(SERVER_FIFO_NAME);
    exit(EXIT_SUCCESS);
}
/*
  如果你同时用读写方式(O_RDWR)方式打开，则不会引起阻塞。
  如果你用只读方式(O_RDONLY)方式打开，则open()会阻塞一直到有写方打开管道， 除非你指定了O_NONBLOCK，来保证打开成功
  同样以写方式(O_WRONLY)打开也会阻塞到有读方打开管道，不同的是如果 O_NONBLOCK被指定open()会以失败告终。
  */

/*
 * 如果每次写入的数据少于PIPE_BUF的大小，那么就不会出现数据交叉的情况。但由于对写入的多少没有限制而read()操作会读取尽可能多的数据，因此你不能知道数据到底是谁写入的,PIPE_BUF的大小根据POSIX标准不能小于512，一些系统里在<limits.h>中被定义，[译者注：Linux中不,是，其值是4096。]这可以通过pathconf()或fpathconf()
对单独管道进行咨询得到
*/
