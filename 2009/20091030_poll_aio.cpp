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
  //在所有数据完成后调用此方法的,
  //即while (aio_error(&iocb) == EINPROGRESS)结束后
  fprintf(stderr, "sh_data: %s\n", buf);
  //struct aiocb *req;
  //req = (struct aiocb*) sigval.sival_ptr;
  //此函数在子线程中运行的
  //while (aio_error(req) == EINPROGRESS)
  //  ;
  //  int ret = -1;
  //  if (ret = aio_return(req))
  //  {
  //    fprintf(stderr, "get a que\n");
  //    fprintf(stderr, "data: %s\n", buf);
  //  }
  //  else
  //  {
  //    perror("aio_error");
  //  }
  //if (aio_error(req) == 0)
  //{
  /* Request completed successfully, get the return status */
  // int ret = aio_return(req);
  //}
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
  ///* //必须要先zero aiocb 否则reuse时就出现aio_read失败的现象,因值不定
  int flag = 1;
  if ((setsockopt(srvsock, SOL_SOCKET, SO_REUSEADDR, &flag, sizeof(flag))) < 0)
  {
    perror("setsocket");
    return -1;
  }
  //*/

  //初始化地址
  socklen_t addr_len = sizeof(struct sockaddr_in);
  struct sockaddr_in srvadd;
  bzero(&srvadd, addr_len);
  srvadd.sin_family = AF_INET;
  srvadd.sin_port = ntohs(2236);
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

  int ic = -1;
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

    if (pl.revents == POLLIN) //有连入,则生成客户socket
    {
      struct sockaddr_in cltadd;
      ic = accept(srvsock, (struct sockaddr*) &cltadd, &addr_len);
      if (ic == -1)
      {
        perror("accept fail\n");
        continue;
      }
      break;
    }
  }

  pl.fd = ic;
  while (1) //不能让aiocb设置放在此前,应先让poll检查是否有数据,有才使用aiocb去处理
  {
    int i = poll(&pl, 1, 3);
    if (i == -1 && errno == EINTR)
      continue;
    else if (i == -1)
    {
      perror("polling fail\n");
      close(ic);
      close(srvsock);
      return -1;
    }
    else if (i == 0)
    {
      fprintf(stderr, "超时\n");
      continue;
    }
    else if (i > 0)
      break;
  }

  if (pl.revents == POLLIN)
  {
    //异步读设置
    struct aiocb iocb;
    bzero(&iocb, sizeof(iocb));
    iocb.aio_buf = buf;
    iocb.aio_fildes = ic;
    iocb.aio_nbytes = 10;
    iocb.aio_offset = 0; //从源中开始读取时的偏移位
    //iocb.aio_sigevent.sigev_value.sival_ptr = &iocb;
    //用子线程操作要注意程序的生命期控制.或者子线程只作修改标志位的操作
    //iocb.aio_sigevent.sigev_notify = SIGEV_THREAD;
    //iocb.aio_sigevent.sigev_notify_function = aio_sh;
    //iocb.aio_sigevent.sigev_notify_attributes = NULL;

    int ret = aio_read(&iocb);
    if (-1 == ret)
    {
      perror("aio_read");
      close(ic);
      close(srvsock);
      return -1;
    }

    while (aio_error(&iocb) == EINPROGRESS)
      //等待操作完成
      ;

    if ((ret = aio_return(&iocb)) > 0) //取操作返回结果
    {
      fprintf(stderr, "get a que\n");
      fprintf(stderr, "data: %s\n", buf);
    }
    else
    {
      perror("aio_error");
    }
  }
  close(ic);
  close(srvsock);
  //等待子线程结束
}
