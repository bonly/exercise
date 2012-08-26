/*
 * shm.h
 *
 *  Created on: 2009-12-3
 *      Author: bonly
 */

#ifndef SHM_H_
#define SHM_H_
#include <sys/ipc.h>
#include <sys/shm.h>
#include <cstring>
#include <cstdio>

template<typename Data>
class SHM
{
  public:
    static key_t ftok(const char* filename)
    {
      key_t key = ::ftok(filename,0);
      if (-1==key)
        perror("ftok error");
      return key;
    }
    static int shmget(key_t key,size_t size,int shmflg)
    {
      int shm_id = ::shmget(key,size,shmflg);
      if(-1==shm_id)
        perror("shmget error");
      return shm_id;
    }
    int shmdt()
    {
      return ::shmdt(_data);  //(void *shmaddr)
    }

    int map_data(const char* rp_file)
    {
      _key = ftok(rp_file);
      if(_key==-1) return -1;

      _shmid = shmget(_key,sizeof(Data),IPC_CREAT|0600);
      if(_shmid==-1) return -1;

      _data = (Data*) shmat(_shmid,NULL,0);
      if(_data==NULL) return -1;

      return 0;
    }

  public:
    Data&    data(){return *_data;}
  private:
    char*   _filename;
    Data*   _data;
    key_t   _key;
    int     _shmid;
};

struct Task
{
  pid_t pid[2];
  int   sock;
  char  server_ip[16];
  int   server_port;
  char  qry[255];
  char  ret[1024];  
};

struct Process
{
  Task task[2];
};

class Process_Ctrl
{
  public:
    static Process_Ctrl& instance();
    int  regedit();
    int  map_data();
    int  get_task();

  public:
    static Process_Ctrl* process_ctrl;
    
  private:
    SHM<Process> _shm_proc;
};


#endif /* SHM_H_ */
