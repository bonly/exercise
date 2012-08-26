/**
 * @file 20100318_mutex.cpp
 * @brief
 *
 * @author bonly
 * @date 2011-11-22 bonly created
 */
//filelock.h
#ifndef __FILE_LOCK_HPP__
#define __FILE_LOCK_HPP__


#ifdef __cplusplus
extern "C" {
#endif


int file_lock(int fd);
int file_unlock(int fd);

#ifdef __cplusplus
}
#endif

#endif //__FILE_LOCK_HPP__



//filelock.cpp

#include <fcntl.h>
#include <unistd.h>

//#include "filelock.h";

int file_lock(int fd){
  struct flock s_flock;
  s_flock.l_type = F_WRLCK;

  s_flock.l_whence = SEEK_SET;
  s_flock.l_start = 0;
  s_flock.l_len = 0;
  s_flock.l_pid = getpid();

  //F_SETLKW对加锁操作进行阻塞，
  //F_SETLK不对加锁操作进行阻塞，立即返回
  return fcntl(fd, F_SETLKW, &s_flock);
}


int file_unlock(int fd){
  return fcntl(fd, F_SETLKW, F_UNLCK);
}


//test.cpp

//#include "filelock.h"

#include <sys/types.h>
#include <sys/stat.h>
#include <fcntl.h>

#include <unistd.h>
#include <stdio.h>
#include <stdlib.h>
#include <sys/wait.h>
#include <errno.h>


int main(int argc,char *argv[]) {
  int pid = -1;

  int i = 5;
  while(i--){
    if ((pid = fork()) < 0) { //fork出错
      puts("fork1 error");
    } else if (pid > 0) {//父进程
      sleep(5);

      if (waitpid(pid, NULL, 0) < 0)
        puts("waitpid error");

    } else {//子进程
      sleep(1);
      int li_file = 0;
      int li_lck_st = -1;
      li_file = open("tt.txt", O_WRONLY|O_CREAT, 0777);
      if( li_file < 0 ) {
        printf("file open error\n");
      }
      printf("li_file=[%d] pid=[%d]\n", li_file , getpid() );
      li_lck_st = file_lock(li_file);
      sleep(5);
      printf("li_lck_st=%d pid =%d\n", li_lck_st, getpid() );
      file_unlock(li_file);
      close(li_file);
      printf("close file pid=%d unlock\n", getpid());
      return 0;
    }
  }
  return 0;
}





