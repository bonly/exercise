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
  //���������ӹ����ڴ�
  return _shm_proc.map_data("/tmp/test.cast");
}

int //ע��
Process_Ctrl::regedit()
{
  
}

int //��ȡ����
Process_Ctrl::get_task()
{
  Process* proc = &(_shm_proc.data()); //ȡ��Process
  for (int i=0; i<2; ++i)
  {//����Process�е�task
    cerr << proc->task[i].server_ip << endl;
  }
  return -1;
}
  
  