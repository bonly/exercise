/*
����֪�����������ӽ��̱�fork����֮ǰ�򿪵��ļ����������ܱ��ӽ��̼̳������ģ�����һ���ӽ����Ѿ������󣬸����̴򿪵��ļ�������Ҫ�������ܴ��ݸ��ӽ����أ�Unix�ṩ��Ӧ�ļ�����������һ���������ͬһ̨�����Ͻ��̼���ļ����������ݣ����������ǿ��ļ�����
����һ��������ͼʵ��һ�������������ն���ͻ��˵����ӣ����������ö���ӽ��̲�������ʽ�������ͻ��˵�ͬʱ���ӣ���ʱ�����ǿ����������뷨��
1���ͻ���ÿ����һ�����ӣ�����fork��һ���ӽ��̸���������ӣ�
2��Ԥ�ȴ���һ�����̳أ��ͻ���ÿ����һ�����ӣ��������ʹӸó���ѡ��һ������(Idle)�ӽ�������������ӡ�
������Ȼ����Ч����Ϊ�������ӽ��̴�����������ģ���Ӧ�ļ�ʱ�Դ����ǿ������ǡǡ�ͳ���������ǰ���ᵽ�����⣬�����ӽ��̶����ڷ�����Listen��һ��������ǰ���Ѿ�fork�����ˣ�Ҳ����˵�µ������������ӽ����ǲ�֪���ģ���Ҫ�����̴��ݸ����������յ���Ӧ�������������󣬲�������Ӧ�Ŀͻ��˽���ͨ�Ŵ����������ǾͿ���ʹ��'�����ļ�������'�ķ�ʽ��ʵ�֡�
��'UNIX�����̵�1��'��14.7С���ж����ּ�������ϸ�Ĳ�����ʵ�������ּ�����������sendmsg��recvmsg��һ����UNIX���׽ӿ�(������ĳ�ֹܵ�)�Ϸ��ͺͽ���һ���������Ϣ��������Ϣ���Գ���'�ļ�������'���ˣ���Ȼ����ϵͳ�ں˶�������Ϣ��������Ĵ����ھ���һ���'�ļ�������'����Ϊ��������(Ancillary Data)ͨ��msghdr�ṹ�еĳ�Աmsg_control(�ϰ汾�г�Ϊmsg_accrights)���ͺͽ��յġ�ֵ��һ����Ƿ��ͽ����ڽ�'�ļ�������'���ͳ�ȥ�󣬼�ʹ�����رո��ļ������������ļ���������Ӧ���ļ��豸Ҳû�б������Ĺرգ������ü�����Ȼ����һ��ֱ�����ս��̳ɹ����պ��ٹرո��ļ��������������ʱ�ļ��豸�����ü���Ϊ0����ô�������رո��ļ��豸��
OK��������һ���򵥵��ļ����������ݵ����ӣ�������ʵ������һ�����ܣ����ӽ��̸����ڸ����̴��ݸ������ļ���������Ӧ���ļ�β�����ض���'LOGO'�ַ��������ӻ���ΪSolaris 9 + GCC 3.2
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

                       /* ����fd��������Ӧ���ļ� */
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
���룺gcc -o test_fdpass -lsocket -lnsl test_fdpass.c
ִ�У�test_fdpass(������ͬһĿ¼�´���һ���ļ�kk.log)

[PARENT]: please enter filename:
kk.log
[CHILD]: append logo successfully
[PARENT]: please enter filename:
cc.log
[PARENT]: can't find file 'cc.log'
exit
[CHILD]: child process exit normally!

����Է���kk.log���ݵ�ĩβ�Ѿ��������ҵĶ���LOGO '-- Tony Bai'��^_^
�����ļ����������ݵĸ���ϸ�ڣ� W. Richard Stevens��'UNIX�����̵�1��'��'UNIX�����߼����'�������ж�����ϸ˵�����ζ����ɡ�
*/
