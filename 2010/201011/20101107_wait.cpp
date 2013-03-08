/**
多线程编程中，线程A循环计算，然后sleep一会接着计算（目的是减少CPU利用率）；存在的问题是，如果要关闭程序，通常选择join线程A等待线程A退出，可是我们必须等到sleep函数返回，该线程A才能正常退出，这无疑减慢了程序退出的速度。当然，你可以terminate线程A，但这样做很不优雅，且会存在一些未知问题。采用pthread_cond_timedwait(pthread_cond_t * cond, pthread_mutex_t *mutex, const struct timespec * abstime)可以优雅的解决该问题，设置等待条件变量cond，如果超时，则返回；如果等待到条件变量cond，也返回
*/
#include <stdio.h>
#include <sys/time.h>
#include <unistd.h>
#include <pthread.h>
#include <errno.h>
 
pthread_t thread;
pthread_cond_t cond;
pthread_mutex_t mutex;
bool flag = true;
 
void * thr_fn(void * arg) {
  struct timeval now;
  struct timespec outtime;
  pthread_mutex_lock(&mutex);
  while (flag) {
    printf(".\n");
    gettimeofday(&now, NULL);
    outtime.tv_sec = now.tv_sec + 5;
    outtime.tv_nsec = now.tv_usec * 1000;
    pthread_cond_timedwait(&cond, &mutex, &outtime);
  }
  pthread_mutex_unlock(&mutex);
  printf("thread exit\n");
}
 
int main() {
  pthread_mutex_init(&mutex, NULL);
  pthread_cond_init(&cond, NULL);
  if (0 != pthread_create(&thread, NULL, thr_fn, NULL)) {
    printf("error when create pthread,%d\n", errno);
    return 1;
  }
  char c ;
  while ((c = getchar()) != 'q');
  printf("Now terminate the thread!\n");
  flag = false;
  pthread_mutex_lock(&mutex);
  pthread_cond_signal(&cond);
  pthread_mutex_unlock(&mutex);
  printf("Wait for thread to exit\n");
  pthread_join(thread, NULL);
  printf("Bye\n");
  return 0;
}
