#include <pthread.h>
#include <stdio.h>
#include <string.h>
#include <sys/types.h>
#include <unistd.h>
#include "20100130_th_thread.h"
using namespace std;
int var = 0;
pthread_mutex_t mut_thread=PTHREAD_MUTEX_INITIALIZER;
pthread_mutex_t new_thread_push_mut = PTHREAD_MUTEX_INITIALIZER;
void *child_fn ( void* arg ) {
   pthread_mutex_lock(&mut_thread);
   var++; /* Unprotected relative to parent */ /* this is line 6 */
   printf("var is %d",var);
   pthread_mutex_unlock(&mut_thread);
   return NULL;
}

int main ( void ) {
   pthread_t child;
   pthread_mutex_init(&mut_thread,NULL);
   for(int i=0; i<3; i++){
         pthread_create(&child, NULL, child_fn, NULL);
          pthread_mutex_lock(&mut_thread);
         var++; /* Unprotected relative to child */ /* this is line 13 */
         pthread_mutex_unlock(&mut_thread);
         pthread_join(child, NULL);
//     printf("var is %d",var);
   }
   return 0;
}
/*
 g++ 20100129_thread.cpp -ldl -lpthread -pg
 */
