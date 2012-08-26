#include <cstdio>
#include <cpthread>

void run()
{
    pthread_exit(0);
}

int
main()
{
    pthread_t thread;
    long count = 0;
    while(1)
    {
        if (rc = pthread_create(&thread,0,run,0))
        {
            print("error, rc is %d, so far %ld threads created\n", rc, count);
            perror("fail:");
            return -1;
        }
        count ++;
    }
    return 0;
}

/*
 * 此线程创建后并没有pthread_join，因此泄漏内存
 */

