/*
 * shm.cpp
 *
 *  Created on: 2009-12-3
 *      Author: bonly
 */
#include <iostream>
using namespace std;
#include "shm.hpp"

Process_Ctrl* Process_Ctrl::process_ctrl = 0;

Process_Ctrl&
Process_Ctrl::instance()
{
  if (Process_Ctrl::process_ctrl == 0)
  {
    //todo: create shm
    Process_Ctrl::process_ctrl = new Process_Ctrl;
  }
  return *Process_Ctrl::process_ctrl;
}

int
Process_Ctrl::map_data()
{
  //创建或连接共享内存
  return _shm_proc.map_data("/tmp/test.cast");
}

int //注册
Process_Ctrl::regedit()
{
  
}

int //获取任务
Process_Ctrl::get_task()
{
  Process* proc = &(_shm_proc.data()); //取得Process
  for (int i=0; i<2; ++i)
  {//遍历Process中的task
    cerr << proc->task[i].server_ip << endl;
  }
  return -1;
}
  
  