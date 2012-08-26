#include <netinet/in.h>
#include <unistd.h>
#include <stdio.h>
#include <string.h>
#include <aio.h>
#include <stdlib.h>
#include <sys/poll.h>
#include <arpa/inet.h>
#include <errno.h>

#define MAXCLIENT 6

#ifndef INFTIM
#define INFTIM -1
#endif/*INFTIM*/

struct clt_node/*client node*/
{
    struct aiocb iocb;
    char buf[BUFSIZ];
    struct clt_node *next, *previous;
};

#define LOCATELAST(h,ptr) do\
{\
    ptr=h;\
    while(ptr != NULL && ptr->next != NULL)\
    {\
        ptr=ptr->next;\
    }\
} while (0)

void sh_aio(int signo, siginfo_t * info, void *contest);
void destroy_node(struct clt_node* ptr);

struct clt_node* head = NULL;
int client_cnt = 0;
int main()
{
    struct sigaction sa;
    int srvsock;
    struct sockaddr_in srvadd, cltadd;
    socklen_t addr_len;
    struct pollfd pl;
    int i;/*for arbitrary use*/

    /*------------------------initialize the server------------------------*/
    if ((srvsock = socket(AF_INET, SOCK_STREAM, 0)) == -1)
    {
        perror("Fail while making the socket fd");
        return -1;
    }

    /*initializing the address*/
    addr_len = sizeof(struct sockaddr_in);
    bzero(&srvadd, addr_len);
    srvadd.sin_family = AF_INET;
    srvadd.sin_port = ntohs(2235);
    srvadd.sin_addr.s_addr = ntohl(INADDR_ANY);

    if ((bind(srvsock, (struct sockaddr*) &srvadd, addr_len)) == -1)
    {
        perror("Fail while binding the address to the socket");
        return -1;
    }
    if (listen(srvsock, MAXCLIENT) == -1)
    {
        perror("Fail while setting up listening");
        return -1;
    }
    /*------------------------initialize the server finished---------------*/

    /*------------------------install the signal handler-------------------*/
    sa.sa_flags = SA_RESTART | SA_SIGINFO;
    sa.sa_sigaction = sh_aio;
    sigfillset(&sa.sa_mask);
    sigdelset(&sa.sa_mask, SIGIO);
    sigdelset(&sa.sa_mask, SIGINT);
    if (sigaction(SIGIO, &sa, NULL) == -1)
    {
        perror("Fail while installing the signal handler");
        return -1;
    }
    /*------------------------install the signal handler finished----------*/
    sigprocmask(SIG_SETMASK, &sa.sa_mask, NULL);/*setting up the current global siganl mask set*/

    /*initializing for poll*/
    pl.fd = srvsock;
    pl.events = POLLIN;
    while (1)
    {
        i = poll(&pl, 1, INFTIM);
        if (i == -1 && errno == EINTR)
            continue;
        else if (i == -1)
        {
            perror("Fail while polling");
            return -1;
        }
        else if (i == 0)
            continue;

        if (pl.revents & POLLIN)/*new incoming client*/
        {
            i = accept(srvsock, (struct sockaddr*) &cltadd, &addr_len);
            if (i == -1)
            {
                perror("Fail while accepting the new client");
                continue;
            }
            if (client_cnt == MAXCLIENT)
            {
                fprintf(stderr, "Max client count %d reached.\n", client_cnt
                        + 1);
                close(i);
                continue;
            }

            {
                struct clt_node *ptr, *tmp;

                printf("Accepted client from %s.\n", inet_ntoa(cltadd.sin_addr));
                if ((tmp = (clt_node*)calloc(1, sizeof(struct clt_node))) == NULL)
                //if(0==(tmp=new clt_node))
                {
                    perror("Fail while getting the memory for the new client");
                    return -1;
                }
                tmp->next = NULL;
                tmp->previous = NULL;

                LOCATELAST(head,ptr);

                if (ptr == NULL)/*single node*/
                {
                    head = tmp;
                    ptr = tmp;
                }
                else
                {
                    ptr->next = tmp;
                    tmp->previous = ptr;
                    ptr = ptr->next;
                }

                ptr->iocb.aio_buf = ptr->buf;
                ptr->iocb.aio_fildes = i;
                ptr->iocb.aio_nbytes = 10;
                ptr->iocb.aio_offset = 0;
                ptr->iocb.aio_sigevent.sigev_notify = SIGEV_SIGNAL;
                ptr->iocb.aio_sigevent.sigev_signo = SIGIO;
                ptr->iocb.aio_sigevent.sigev_value.sival_ptr = &ptr->iocb;
                if (aio_read(&ptr->iocb) == -1)
                {
                    perror("Fail to install aio read");
                    destroy_node(ptr);/*destroy this node*/
                    continue;
                }

                client_cnt++;
            }
        }

    }

    return 0;
}

void destroy_node(struct clt_node* ptr)
{
    if (ptr == NULL)
        return;

    aio_cancel(ptr->iocb.aio_fildes, &ptr->iocb);/*canceling it*/
    close(ptr->iocb.aio_fildes);

    if (ptr->previous == NULL)/*first node*/
    {
        if (ptr->next != NULL)/*not single node*/
        {
            head = ptr->next;
            head->previous = NULL;
        }
        else/*single node*/
        {
            head = NULL;
        }
        goto DOJOB;
    }
    if (ptr->next == NULL)/*last node, would never single node*/
    {
        ptr->previous->next = NULL;
        goto DOJOB;
    }
    else
    {
        struct clt_node *p, *n;
        p = ptr->previous;
        n = ptr->next;
        p->next = n;
        n->previous = p;
        goto DOJOB;
    }

    DOJOB: free(ptr);
    client_cnt--;
    return;
}

void sh_aio(int signo, siginfo_t * info, void *contest)
{
    struct aiocb* ref;

    if (info->si_signo == SIGIO)
    {
        ref = (aiocb*)info->si_value.sival_ptr;/*reading the aiocb*/
        if (aio_error(ref) == 0)
        {
            int cnt;
            cnt = aio_return(ref);
            if (cnt == 0)
                goto DESTROY;
            printf("Socket: %d says:\n", ref->aio_fildes);
            write(STDOUT_FILENO, (char*) ref->aio_buf, cnt);
            bzero((char*) ref->aio_buf, ref->aio_nbytes);
            if (aio_read(ref) == -1)
            {
                perror("Fail to install aio read");

                {
                    DESTROY: ;
                    struct clt_node* ptr = head;
                    while ((ptr != NULL) && &ptr->iocb != ref)
                        ptr = ptr->next;
                    if (ptr == NULL)/*reaching the last node but could not find*/
                    {
                        fprintf(stderr, "Fatal error occurs!\n");
                        exit(EXIT_FAILURE);
                    }

                    destroy_node(ptr);/*destroy this node*/
                }

            }
        }

    }
}

