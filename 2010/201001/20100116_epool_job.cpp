/* 
 几个用到的类型定义以及全局变量bq 
 */
char smtp_cmd_format;
struct epoll_event ev, events[MAX_EPOLL_SIZE];
int kdpfd, nfds;
struct block_queue
{
      int queue[THREADS_COUNT];
      long size;
      pthread_cond_t cond;
      pthread_mutex_t mutex;
} block_queue_t;
block_queue_t bq;
struct block_queue_param
{
      void* func;
      void* queue;
} block_queue_param_t;

void *block_queue(void * param)
{
   void(* func)(void*);
   int fd, i;
   /* 
    由于block_queue是pthread_create的回调方法, 
    所以block_queue的参数必须是void*类型 
    */
   block_queue_t* bque = (block_queue_param_t*) param->queue;
   /* 
    param->func是block_queue解锁时需要调用的函数, 
    而这个函数的参数是一个int fd, 
    该fd是消息队列中刚刚插入的一个元素. 
    */
   func = (block_queue_param_t*) param->func;

   for (;;)
   {
      /* 
       lock->wait->unlock 
       这是经典的模式, 
       切记: 
       pthread_cond_wait的方法自带了先解锁,再等待,最后再加锁的代码 
       */
      pthread_mutex_lock(bque->mutex);
      /* 
       线程在pthread_cond_wait这里被block住了, 
       当听到pthread_cond_signal通知的时候, 
       内核会从阻塞队列里面通过先进先出的原则唤醒一个线程, 
       这个线程会执行pthread_cond_wait之后的代码. 
       */
      pthread_cond_wait(bque->cond, bque->mutex);

      if (bque->size == 0)
      {
         /* 
          啥也不做 
          */
      }
      else
      {
         fd = bque->queue[0];
         /* 
          移动队列, 
          由于该队列是简单的用数组保存fd, 
          所以移动这个操作必不可少,但肯定性能比链表差很多, 
          这是懒惰的代价 
          */
         for (i = 0; i < bque->size; ++i)
            bque->queue[i] = bque->queue[i + 1];
         bque->queue[bque->size - 1] = 0;
         bque->size--;
         /* 
          执行被唤醒后的方法,参数是刚刚插入到队列中的一个fd 
          */
         func(fd);
      }

      pthread_mutex_unlock(bque->mutex);
   }
}

void insert_queue(struct block_queue *bque, int fd)
{
   /* 
    加锁->通知->解锁 
    
    将元素插入队列之前需要先加锁 
    */
   pthread_mutex_lock(bque->mutex);
   /* 
    检查队列目前的大小, 
    检查1: 
    当大小已经达到定义的数组大小THREADS_COUNT时, 
    抛弃该fd,说明服务器忙不过来了,消息队列已经满了 
    
    检查2: 
    当大小超过数组定义的大小THREADS_COUNT时, 
    肯定发生了异常,那就直接退出服务吧. 
    */
   if (bque->size == THREADS_COUNT)
      return;
   /* 
    bque->size其实也是队列末尾位置指针, 
    当插入一个元素后,这个指针自然也要向后移动一位. 
    */
   bque->queue[bque->size + 1] = fd;
   if (++bque->size > THREADS_COUNT)
   {
      fprintf(stderr, "Queue size over folow.%d", bque->size);
      exit(1);
   }
   /* 
    当元素插入bque队列时, 
    该通过pthread_cond_signal通知内核去调度wait的线程了 
    */
   pthread_cond_signal(bque->cond);
   pthread_mutex_unlock(bque->mutex);

}

/* 
 init_threads代码是初始化线程组的, 
 随便写写的,大家知道怎么实现就行 
 */
int init_threads()
{
   size_t i = 0;
   block_queue_param_t bqp;
   /* 
    smtp_echo是处理epoll扔进队列中的fd的函数, 
    该方法实现了整个模型的业务逻辑, 
    整体代码的IO处理+消息队列以及业务处理分的很清晰, 
    三个模块每个只有一处代码和其它模块通讯,没有多少耦合. 
    */
   bqp.func = (void*) smtp_echo;
   bqp.queue = (void*) bq;
   pthread_cond_init(bqp.cond, NULL);
   pthread_mutex_init(bqp.mutex, NULL);
   for (i = 0; i < THREADS_COUNT; ++i)
   {
      pthread_t child_thread;
      pthread_attr_t child_thread_attr;
      pthread_attr_init(&child_thread_attr);
      pthread_attr_setdetachstate(&child_thread_attr, PTHREAD_CREATE_DETACHED);
      if (pthread_create(&child_thread, &child_thread_attr, block_queue,
               (void *) bqp) < 0)
      {
         printf("pthread_create Failed : %s\n", strerror(errno));
         return 1;
      }
      else
      {
         printf("pthread_create Success : %d\n", (int) child_thread);
         return 0;
      }
   }

}

/* 
 handler是主线程访问的方法, 
 主线程通过handler把一个fd扔到消息队列之后, 
 不再做任何事情就直接返回了. 
 
 在我的应用中,主线程是个epoll实现的服务器, 
 由于epoll被响应的时候会知道哪些fd已经就位, 
 于是直接把就位的fd扔到消息队列中就好了, 
 主线程在继续等待其它fd的响应,而不需要去关心fd如何被处理. 
 */

int handler(void* fd)
{
   printf("handler:fd => %d\n", *(int *) (fd));
   insert_queue(&bq, fd);
   return 0;
}

/* 
 main函数是整个程序的入口点, 
 也是epoll服务器的实现, 
 epoll的思想很精髓,用法很简单, 
 只要把man 4 epoll_ctl的例子copy出来,就可用了, 
 不过那个例子语法有点问题, 
 而且events数组是用指针,应该用[]实现,因为指针没有分配空间. 
 */
int main(int argc, char **argv)
{
   int server_socket = init_smtp();
   int n;

   if (init_threads() == 0)
      printf("Success full init_threads.");

   smtp_cmd_format = "^([a-zA-Z0-9]) (.*)$";
   kdpfd = epoll_create(MAX_EPOLL_SIZE);
   ev.events = EPOLLIN | EPOLLET;
   ev.data.fd = server_socket;
   if (epoll_ctl(kdpfd, EPOLL_CTL_ADD, server_socket, &ev) < 0)
   {
      fprintf(stderr, "epoll set insertion error: fd=%d < 0", server_socket);
      return -1;
   }

   /* 
    epoll的使用看这里 
    */
   for (;;)
   {
      struct sockaddr_in local;
      socklen_t length = sizeof(local);
      int client;

      nfds = epoll_wait(kdpfd, events, MAX_EPOLL_SIZE, -1);

      /* 
       当没有事件要处理时,epoll会阻塞住, 
       否则,内核会填充events数组,里面的每一个events[n].data.fd就是发生io时间的文件句柄 
       */

      for (n = 0; n < nfds; ++n)
      {
         /* 
          这里要判断一下请求的来源, 
          if(events[n].data.fd == server_socket) { 
          这里是新建的连接, 
          因为io发生在server_socket上 
          } 
          else{ 
          这里是已有的连接, 
          因为fd!= server_socket 
          那fd肯定是之前从server_socket接收到, 
          并且通过epoll_ctl(kdpfd, EPOLL_CTL_ADD, client, &ev) 
          加入到kdpfd所指向的内存空间中. 
          kdpfd其实是个文件句柄,在epoll_create(MAX_EPOLL_SIZE)时得到 
          } 
          */
         if (events[n].data.fd == server_socket)
         {
            client = accept(server_socket, (struct sockaddr *) &local, &length);
            if (client < 0)
            {
               perror("accept");
               continue;
            }
            setnonblocking(client);
            smtp_cmd("220", client);
            ev.events = EPOLLIN | EPOLLET;
            ev.data.fd = client;
            if (epoll_ctl(kdpfd, EPOLL_CTL_ADD, client, &ev) < 0)
            {
               fprintf(stderr, "epoll set insertion error: fd=%d < 0", client);
               return -1;
            }
         }
         else
         /* 
          当已有的fd发生io操作时, 
          执行如下代码.也就是把fd扔到消息队列中. 
          */
         if (handler((void *) &events[n].data.fd) != 0)
            perror("handler ret != 0");
      }
   }

   close(server_socket);
   return 0;
}
