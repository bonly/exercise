/**
 *  @file 20100421_service.cpp
 *
 *  @date 2012-3-12
 *  @Author: Bonly
 */

#include <windows.h>
#include <stdio.h>

#define SLEEP_TIME 5000
#define LOGFILE "C:/Temp/emestatus.txt"

int WriteToLog(const char* str)
{
  FILE* log;
  log = fopen(LOGFILE, "a+");
  if (log == NULL)
    return -1;
  fprintf(log, "%s\n", str);
  fclose(log);
  return 0;
}

SERVICE_STATUS ServiceStatus;
SERVICE_STATUS_HANDLE hStatus;

void ServiceMain(int argc, char* argv[]);
void ControlHandler(DWORD request);
int InitService();

int main()
{
  WriteToLog("main\n");
  SERVICE_TABLE_ENTRY ServiceTable[2];
  ServiceTable[0].lpServiceName = (LPSTR)"MemoryStatus";
  ServiceTable[0].lpServiceProc = (LPSERVICE_MAIN_FUNCTION)ServiceMain;

  ServiceTable[1].lpServiceName = NULL;
  ServiceTable[1].lpServiceProc = NULL;

  StartServiceCtrlDispatcher (ServiceTable); ///启动服务的控制分派机线程
  WriteToLog("main end\n");
  return 0;
}

/**
 * @brief 服务线程入口
 */
void ServiceMain(int argc, char* argv[])
{
  WriteToLog("Service main\n");
  int error;
  ServiceStatus.dwServiceType = SERVICE_WIN32; ///服务类型，固定
  ServiceStatus.dwCurrentState = SERVICE_START_PENDING; ///当前状态
  ServiceStatus.dwControlsAccepted = SERVICE_ACCEPT_STOP | SERVICE_ACCEPT_SHUTDOWN; ///接受的操作请求
  ServiceStatus.dwWin32ExitCode = 0; ///退出时带出的返回值
  ServiceStatus.dwServiceSpecificExitCode = 0; ///退出时带出的返回值
  ServiceStatus.dwCheckPoint = 0; ///初始化服务需30秒以上时用于主线程提示，<30时为0
  ServiceStatus.dwWaitHint = 0; ///初始化服务需30秒以上时的提示

  hStatus = RegisterServiceCtrlHandler( ///注册处理函数
      "MemoryStatus",
      (LPHANDLER_FUNCTION) ControlHandler);
  if (hStatus == (SERVICE_STATUS_HANDLE)0)
  {
    return; ///注册失败
  }

  error = InitService(); ///初始化服务
  if (error)
  {
    ServiceStatus.dwCurrentState = SERVICE_STOPPED; ///初始化失败
    ServiceStatus.dwWin32ExitCode = -1;
    SetServiceStatus (hStatus, &ServiceStatus);
    return ;
  }

  ServiceStatus.dwCurrentState = SERVICE_RUNNING; ///向SCM报告服务状态
  SetServiceStatus (hStatus, &ServiceStatus);

  MEMORYSTATUS memory;
  while (ServiceStatus.dwCurrentState == SERVICE_RUNNING)
  {
    char buffer[16];
    GlobalMemoryStatus(&memory);
    sprintf(buffer, "%ld", memory.dwAvailPhys);
    int result = WriteToLog(buffer);
    if (result)
    {
      ServiceStatus.dwCurrentState = SERVICE_STOPPED;
      ServiceStatus.dwWin32ExitCode = -1;
      SetServiceStatus (hStatus, &ServiceStatus);
      return;
    }
    Sleep(SLEEP_TIME);
    WriteToLog("while\n");
  }
  return;
}

/**
 * @brief 服务控制实现
 */
void ControlHandler(DWORD dwCode)
{
  WriteToLog("Service control\n");
}

int InitService()
{
  WriteToLog("Service init\n");
  return 0;
}
