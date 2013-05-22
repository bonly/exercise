#include <cstdio>
#include <cstdlib>
#include <unistd.h> //_SC_MQ_OPEN_MAX

int
main(int argc,char* argv[])
{
        printf("MQ_OPEN_MAX = %ld, MQ_PRIO_MAX = %ld\n", 
                        sysconf(_SC_MQ_OPEN_MAX),sysconf(_SC_MQ_PRIO_MAX));
        exit(0);
}

/*
 * MQ_OPEN_MAX 一个进程能同时拥有打开的消息队列的最大数目
 * MQ_PRIO_MAX 任意消息的最大优先级值加1
 */
