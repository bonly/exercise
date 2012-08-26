/*
我们知道父进程在子进程被fork出来之前打开的文件描述符是能被子进程继承下来的，但是一旦子进程已经创建后，父进程打开的文件描述符要怎样才能传递给子进程呢？Unix提供相应的技术来满足这一需求，这就是同一台主机上进程间的文件描述符传递，很美妙而且强大的技术。
想象一下我们试图实现一个服务器，接收多个客户端的连接，我们欲采用多个子进程并发的形式来处理多客户端的同时连接，这时候我们可能有两种想法：
1、客户端每建立一条连接，我们fork出一个子进程负责处理该连接；
2、预先创建一个进程池，客户端每建立一条链接，服务器就从该池中选出一个空闲(Idle)子进程来处理该连接。
后者显然更高效，因为减少了子进程创建的性能损耗，反应的及时性大大增强。这里恰恰就出现了我们前面提到的问题，所有子进程都是在服务器Listen到一条连接以前就已经fork出来了，也就是说新的连接描述符子进程是不知道的，需要父进程传递给它，它接收到相应的连接描述符后，才能与相应的客户端进行通信处理。这里我们就可以使用'传递文件描述符'的方式来实现。
在'UNIX网络编程第1卷'的14.7小节中对这种技术有详细的阐述，实际上这种技术就是利用sendmsg和recvmsg在一定的UNIX域套接口(或者是某种管道)上发送和接收一种特殊的消息，这种消息可以承载'文件描述符'罢了，当然操作系统内核对这种消息作了特殊的处理。在具体一点儿'文件描述符'是作为辅助数据(Ancillary Data)通过msghdr结构中的成员msg_control(老版本中称为msg_accrights)发送和接收的。值得一提的是发送进程在将'文件描述符'发送出去后，即使立即关闭该文件描述符，该文件描述符对应的文件设备也没有被真正的关闭，其引用计数仍然大于一，直到接收进程成功接收后，再关闭该文件描述符，如果这时文件设备的引用计数为0，那么才真正关闭该文件设备。
OK，下面是一个简单的文件描述符传递的例子，该例子实现这样一个功能：即子进程负责在父进程传递给它的文件描述符对应的文件尾加上特定的'LOGO'字符串。例子环境为Solaris 9 + GCC 3.2
*/

/* test_fdpass.c */
#define HAVE_MSGHDR_MSG_CONTROL
#include <stdio.h>
#include <string.h>
#include <stdlib.h>

#include <unistd.h>
#include <sys/wait.h>
#include <sys/types.h>
#include <sys/stat.h>
#include <fcntl.h>
#include <errno.h>

#include <sys/socket.h> /* for socketpair */

#define MY_LOGO         "-- Tony Bai"

static int send_fd(int fd, int fd_to_send)
{
        struct iovec    iov[1];
        struct msghdr   msg;
        char            buf[1];

        if (fd_to_send >= 0) {
                msg.msg_accrights       = (caddr_t)&fd_to_send;
                msg.msg_accrightslen    = sizeof(int);
        } else {
                msg.msg_accrights       = (caddr_t)NULL;
                msg.msg_accrightslen    = 0;
        }

        msg.msg_name    = NULL;
        msg.msg_namelen = 0;

        iov[0].iov_base = buf;
        iov[0].iov_len  = 1;
        msg.msg_iov     = iov;
        msg.msg_iovlen  = 1;

        if(sendmsg(fd, &msg, 0) < 0) {
                printf("sendmsg error, errno is %d\n", errno);
                return errno;
        }

        return 0;
}

static int recv_fd(int fd, int *fd_to_recv)
{
        struct iovec    iov[1];
        struct msghdr   msg;
        char            buf[1];

        msg.msg_accrights       = (caddr_t)fd_to_recv;
        msg.msg_accrightslen    = sizeof(int);

        msg.msg_name    = NULL;
        msg.msg_namelen = 0;

        iov[0].iov_base = buf;
        iov[0].iov_len  = 1;
        msg.msg_iov     = iov;
        msg.msg_iovlen  = 1;

        if (recvmsg(fd, &msg, 0) < 0) {
                return errno;
        }

        if(msg.msg_accrightslen != sizeof(int)) {
                *fd_to_recv = -1;
        }

        return 0;
}

int x_sock_set_block(int sock, int on)
{
        int             val;
        int             rv;

        val = fcntl(sock, F_GETFL, 0);
        if (on) {
                rv = fcntl(sock, F_SETFL, ~O_NONBLOCK&val);
        } else {
                rv = fcntl(sock, F_SETFL, O_NONBLOCK|val);
        }

        if (rv) {
                return errno;
        }

        return 0;
}

int main() {
        pid_t   pid;
        int     sockpair[2];
        int     rv;
        char    fname[256];
        int     fd;

        rv = socketpair(AF_UNIX, SOCK_STREAM, 0, sockpair);
        if (rv < 0) {
                printf("Call socketpair error, errno is %d\n", errno);
                return errno;
        }

        pid = fork();
        if (pid == 0) {
                /* in child */
                close(sockpair[1]);

                for ( ; ; ) {
                        rv = x_sock_set_block(sockpair[0], 1);
                        if (rv != 0) {
                                printf("[CHILD]: x_sock_set_block error, errno is %d\n", rv);
                                break;
                        }

                        rv = recv_fd(sockpair[0], &fd);
                        if (rv < 0) {
                                printf("[CHILD]: recv_fd error, errno is %d\n", rv);
                                break;
                        }

                        if (fd < 0) {
                                printf("[CHILD]: child process exit normally!\n");
                                break;
                        }

                       /* 处理fd描述符对应的文件 */
                        rv = write(fd, MY_LOGO, strlen(MY_LOGO));
                        if (rv < 0) {
                                printf("[CHILD]: write error, errno is %d\n", rv);
                        } else {
                                printf("[CHILD]: append logo successfully\n");
                        }
                        close(fd);
                }

                exit(0);
        }

        /* in parent */
        for ( ; ; ) {
                memset(fname, 0, sizeof(fname));
                printf("[PARENT]: please enter filename:\n");
                scanf("%s", fname);

                if (strcmp(fname, "exit") == 0) {
                        rv = send_fd(sockpair[1], -1);
                        if (rv < 0) {
                                printf("[PARENT]: send_fd error, errno is %d\n", rv);
                        }
                        break;
                }

                fd = open(fname, O_RDWR | O_APPEND);
                if (fd < 0) {
                        if (errno == ENOENT) {
                                printf("[PARENT]: can't find file '%s'\n", fname);
                                continue;
                        }
                        printf("[PARENT]: open file error, errno is %d\n", errno);
                }

                rv = send_fd(sockpair[1], fd);
                if (rv != 0) {
                        printf("[PARENT]: send_fd error, errno is %d\n", rv);
                }

                close(fd);
        }

        wait(NULL);
        return 0;
}

/*
编译：gcc -o test_fdpass -lsocket -lnsl test_fdpass.c
执行：test_fdpass(事先在同一目录下创建一个文件kk.log)

[PARENT]: please enter filename:
kk.log
[CHILD]: append logo successfully
[PARENT]: please enter filename:
cc.log
[PARENT]: can't find file 'cc.log'
exit
[CHILD]: child process exit normally!

你可以发现kk.log内容的末尾已经加上了我的独特LOGO '-- Tony Bai'。^_^
关于文件描述符传递的更多细节， W. Richard Stevens的'UNIX网络编程第1卷'和'UNIX环境高级编程'两本书中都有详细说明，参读即可。
*/
