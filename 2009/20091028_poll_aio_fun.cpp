
#include <netinet/in.h>
#include <unistd.h>
#include <stdio.h>
#include <string.h>
#include <aio.h>
#include <stdlib.h>
#include <sys/poll.h>
#include <arpa/inet.h>
#include <errno.h>
char buf[200] = "";

void aio_sh(sigval_t sigval)
{
  struct aiocb *req;
  req = (struct aiocb*)sigval.sival_ptr;
  if(aio_error(req)==0)
  {
    fprintf(stderr,"get a que\n");
    fprintf(stderr,"data: %s\n",buf);
    int ret = aio_return(req); //无论返不返回,都不会再调第二次
  }
}


int main()
{
  //取得socket
  int srvsock = socket(AF_INET, SOCK_STREAM, 0);
  if (srvsock == -1)
  {
    perror("create socket fail\n");
    return -1;
  }

  //初始化地址
  socklen_t addr_len = sizeof(struct sockaddr_in);
  struct sockaddr_in srvadd;
  bzero(&srvadd, addr_len);
  srvadd.sin_family = AF_INET;
  srvadd.sin_port = ntohs(2235);
  srvadd.sin_addr.s_addr = ntohl(INADDR_ANY);

  if ((bind(srvsock, (struct sockaddr*) &srvadd, addr_len)) == -1)
  {
    perror("bind fail\n");
    return -1;
  }
  if (listen(srvsock, 2) == -1)
  {
    perror("listen fail\n");
    return -1;
  }

  struct pollfd pl;
  pl.fd = srvsock;
  pl.events = POLLIN;
  while (1)
  {
    int i = poll(&pl, 1, 3);
    if (i == -1 && errno == EINTR)
      continue;
    else if (i == -1)
    {
      perror("polling fail\n");
      return -1;
    }
    else if (i == 0)
      continue;

    if (pl.revents == POLLIN)
    {
      struct sockaddr_in cltadd;
      int ic = accept(srvsock, (struct sockaddr*) &cltadd, &addr_len);
      if (ic == -1)
      {
        perror("accept fail\n");
        continue;
      }

      struct aiocb iocb;
      iocb.aio_buf = buf;
      iocb.aio_fildes = ic;
      iocb.aio_nbytes = 10;
      iocb.aio_offset = 0;
      iocb.aio_sigevent.sigev_value.sival_ptr = &iocb;
      iocb.aio_sigevent.sigev_notify = SIGEV_THREAD;
      iocb.aio_sigevent.sigev_notify_function = aio_sh;
      iocb.aio_sigevent.sigev_notify_attributes = NULL;


      int ret = aio_read(&iocb);
    }
  }
}
