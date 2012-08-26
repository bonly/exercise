/**
 *  @file 20100422_scm.cpp
 *
 *  @date 2012-3-12
 *  @Author: Bonly
 */

#include "windows.h"
#include "stdio.h"

//全局变量
SERVICE_STATUS ServiceStatus;
SERVICE_STATUS_HANDLE ServiceStatusHandle;

//函数声明
void InstallService(char *szServicePath);
void WINAPI ServiceStart(DWORD dwArgc, LPTSTR *lpArgv);
void WINAPI ServiceControl(DWORD dwCode);
DWORD WINAPI Service(LPVOID lpvThread); //服务功能函数

int main(int argc, char* argv[])
{
  //定义SERVICE_TABLE_ENTRY DispatchTable[] 结构
  SERVICE_TABLE_ENTRY DispatchTable[2] =
  {
  { (LPSTR) "mixu_yy", ServiceStart },
  { NULL, NULL } };

  //安装或者打开服务
  StartServiceCtrlDispatcher(DispatchTable);
  InstallService(argv[0]);

  return 0;
}
//函数定义
void InstallService(char *szServicePath)
{
  SC_HANDLE schSCManager;
  SC_HANDLE schService;
  SERVICE_STATUS InstallServiceStatus;
  DWORD dwErrorCode;

  //打开服务管理数据库
  schSCManager = OpenSCManager(NULL, NULL, SC_MANAGER_ALL_ACCESS);
  if (schSCManager == NULL)
  {
    //Open Service Control Manager Database Failed!;
    return;
  }
  //创建服务
  schService = CreateService(schSCManager, "mixu_yy", "mixu_yy",
      SERVICE_ALL_ACCESS, SERVICE_WIN32_OWN_PROCESS, SERVICE_AUTO_START,
      SERVICE_ERROR_IGNORE, szServicePath, NULL, NULL, NULL, NULL, NULL);
  if (schService == NULL)
  {
    dwErrorCode = GetLastError();
    if (dwErrorCode != ERROR_SERVICE_EXISTS)
    {
      //创建服务失败
      CloseServiceHandle(schSCManager);
      return;
    }
    else
    {
      //要创建的服务已经存在
      schService = OpenService(schSCManager, "mixu_yy", SERVICE_START);
      if (schService == NULL)
      {
        //Open Service Failed!;
        CloseServiceHandle(schSCManager);
        return;
      }
    }
  }
  else
  {
    //Create Service Success!;
  }
  //启动服务
  if (StartService(schService, 0, NULL) == 0)
  {
    //启动失败
    dwErrorCode = GetLastError();
    if (dwErrorCode == ERROR_SERVICE_ALREADY_RUNNING)
    {
      //Service already run!;
      CloseServiceHandle(schSCManager);
      CloseServiceHandle(schService);
      return;
    }
  }
  else
  {
    //Service pending
  }

  while (QueryServiceStatus(schService, &InstallServiceStatus) != 0)
  {
    if (InstallServiceStatus.dwCurrentState == SERVICE_START_PENDING)
    {
      //Sleep(100);
    }
    else
    {
      break;
    }
  }
  if (InstallServiceStatus.dwCurrentState != SERVICE_RUNNING)
  {
    //Failure!
  }
  else
  {
    //Sucess!
  }

  //擦屁股
  CloseServiceHandle(schSCManager);
  CloseServiceHandle(schService);
  return;
}

void WINAPI ServiceStart(DWORD dwArgc, LPTSTR *lpArgv)
{
  HANDLE hThread;

  ServiceStatus.dwServiceType = SERVICE_WIN32;
  ServiceStatus.dwCurrentState = SERVICE_START_PENDING;
  ServiceStatus.dwControlsAccepted = SERVICE_ACCEPT_STOP
      | SERVICE_ACCEPT_PAUSE_CONTINUE;

  ServiceStatus.dwServiceSpecificExitCode = 0;
  ServiceStatus.dwWin32ExitCode = 0;
  ServiceStatus.dwCheckPoint = 0;
  ServiceStatus.dwWaitHint = 0;

  ServiceStatusHandle = RegisterServiceCtrlHandler("mixu_yy", ServiceControl);
  if (ServiceStatusHandle == 0)
  {
    //error
    return;
  }

  ServiceStatus.dwCurrentState = SERVICE_RUNNING;
  ServiceStatus.dwCheckPoint = 0;
  ServiceStatus.dwWaitHint = 0;

  if (SetServiceStatus(ServiceStatusHandle, &ServiceStatus) == 0)
  {
    //SetServiceStatus error!
    return;
  }

  //创建服务线程   服务完成的功能在这里调用
  hThread = CreateThread(NULL, 0, Service, NULL, 0, NULL);
  if (hThread == NULL)
  {
    //CreateThread error!
    return;
  }
  CloseHandle(hThread);
  return;
}

//服务控制模块

void WINAPI ServiceControl(DWORD dwCode)
{
  switch (dwCode)
  {
    case SERVICE_CONTROL_PAUSE:
      ServiceStatus.dwCurrentState = SERVICE_PAUSED;
      break;

    case SERVICE_CONTROL_CONTINUE:
      ServiceStatus.dwCurrentState = SERVICE_RUNNING;
      break;

    case SERVICE_CONTROL_STOP:
      ServiceStatus.dwCurrentState = SERVICE_STOPPED;
      ServiceStatus.dwWin32ExitCode = 0;
      ServiceStatus.dwCheckPoint = 0;
      ServiceStatus.dwWaitHint = 0;
      if (SetServiceStatus(ServiceStatusHandle, &ServiceStatus) == 0)
      {
        //SetServiceStatus error!
      }
      return;
    case SERVICE_CONTROL_INTERROGATE:
      break;
    default:
      break;
  }

  if (SetServiceStatus(ServiceStatusHandle, &ServiceStatus) == 0)
  {
    //SetServiceStatus error!
  }

  return;
}

//服务线程函数
DWORD WINAPI Service(LPVOID lpvThread)
{
  //实现函数功能的地方。
  return 1;
}

