#include <sys/socket.h>
#include <netinet/in.h>
#include <stdlib.h>
#include <stdio.h>
#include <unistd.h>

void err_sys(const char *errmsg);

int main(void)
{
        int sockfd[2];
        pid_t pid;

        if ((socketpair(AF_LOCAL, SOCK_STREAM, 0, sockfd)) == -1)
                err_sys("socketpair");

        if ((pid = fork()) == -1)
                err_sys("fork");
        else if (pid == 0) {            /* child process */
                char buf[] = "hello china", s[BUFSIZ];
                ssize_t n;

                close(sockfd[1]);

                write(sockfd[0], buf, sizeof(buf));
                if ((n = read(sockfd[0], s, sizeof(s))) == -1)
                        err_sys("read");
                write(STDOUT_FILENO, s, n);

                exit(0);
        } else if (pid > 0) {           /* parent process */
                char buf[BUFSIZ];
                ssize_t n;

                close(sockfd[0]);
                n = read(sockfd[1], buf, sizeof(buf));
                write(sockfd[1], buf, n);

                exit(0);
        }
}

void err_sys(const char *errmsg)
{
        perror(errmsg);
        exit(1);
}

