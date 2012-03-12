/**
 *  @file 20100423_keep_chrome_open.cpp
 *
 *  @date 2012-3-12
 *  @Author: Bonly
 */

/**
 *  @file 20100421_service.cpp
 *
 *  @date 2012-3-12
 *  @Author: Bonly
 */

#include <windows.h>
#include <TlHelp32.h>
#include <stdio.h>

#define SLEEP_TIME 5000
#define LOGFILE "C:/keep_chrome_open.log"

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
int KeepChromeOpen();

int main()
{
  SERVICE_TABLE_ENTRY ServiceTable[2];
  ServiceTable[0].lpServiceName = (LPSTR)"Keep_chrome_open";
  ServiceTable[0].lpServiceProc = (LPSERVICE_MAIN_FUNCTION)ServiceMain;

  ServiceTable[1].lpServiceName = NULL;
  ServiceTable[1].lpServiceProc = NULL;

  StartServiceCtrlDispatcher (ServiceTable); ///启动服务的控制分派机线程
  return 0;
}

/**
 * @brief 服务线程入口
 */
void ServiceMain(int argc, char* argv[])
{
  int error;
  ServiceStatus.dwServiceType = SERVICE_WIN32; ///服务类型，固定
  ServiceStatus.dwCurrentState = SERVICE_START_PENDING; ///当前状态
  ServiceStatus.dwControlsAccepted = SERVICE_ACCEPT_STOP | SERVICE_ACCEPT_SHUTDOWN; ///接受的操作请求
  ServiceStatus.dwWin32ExitCode = 0; ///退出时带出的返回值
  ServiceStatus.dwServiceSpecificExitCode = 0; ///退出时带出的返回值
  ServiceStatus.dwCheckPoint = 0; ///初始化服务需30秒以上时用于主线程提示，<30时为0
  ServiceStatus.dwWaitHint = 0; ///初始化服务需30秒以上时的提示

  hStatus = RegisterServiceCtrlHandler( ///注册处理函数
      "CheckChrome",
      (LPHANDLER_FUNCTION) ControlHandler);
  if (hStatus == (SERVICE_STATUS_HANDLE)0)
  {
    return; ///注册失败
  }

  error = InitService(); ///初始化服务
  if (error)
  {
    WriteToLog("Service Init failed\n");
    ServiceStatus.dwCurrentState = SERVICE_STOPPED; ///初始化失败
    ServiceStatus.dwWin32ExitCode = -1;
    SetServiceStatus (hStatus, &ServiceStatus);
    return ;
  }

  ServiceStatus.dwCurrentState = SERVICE_RUNNING; ///向SCM报告服务状态
  SetServiceStatus (hStatus, &ServiceStatus);

  while (ServiceStatus.dwCurrentState == SERVICE_RUNNING)
  {
    int result = KeepChromeOpen();
    if (result!=0)
    {
      ServiceStatus.dwCurrentState = SERVICE_STOPPED;
      ServiceStatus.dwWin32ExitCode = -1;
      SetServiceStatus (hStatus, &ServiceStatus);
      return;
    }
    Sleep(SLEEP_TIME);
  }
  return;
}

/**
 * @brief 服务控制实现
 */
void ControlHandler(DWORD dwCode)
{
  switch(dwCode)
  {
    case SERVICE_CONTROL_STOP:
      WriteToLog("Stoping service\n");
      ServiceStatus.dwCurrentState = SERVICE_STOP;
      ServiceStatus.dwWin32ExitCode = 0;
      SetServiceStatus (hStatus, &ServiceStatus);
      break;
  }
}

int InitService()
{
  WriteToLog("Service init\n");
  return 0;
}

int KeepChromeOpen()
{

  HANDLE  hSnapshot;
  int count = 0;
  hSnapshot = ::CreateToolhelp32Snapshot(TH32CS_SNAPPROCESS,0);  //创建快照
  if(INVALID_HANDLE_VALUE == hSnapshot)
  {
    WriteToLog("CreateToolhelp32Snapshot call faild!\n");
    return 1;
  }

  PROCESSENTRY32 process;
  process.dwSize = sizeof(PROCESSENTRY32);   //注意 必不可少
  BOOL first = ::Process32First(hSnapshot,&process);

  bool found = false;
  while(first)     //循环 列出进程信息
  {
    ++count;
    if(strstr(process.szExeFile, "chrome.exe")!=0)
    {
      found = true;
      break;
    }
    first = ::Process32Next(hSnapshot,&process);
  }

  ::CloseHandle(hSnapshot);
  if (!found)
  {
  //system("Chrome --kiosk");
    if(0!=system("D:\\Chrome\\Chrome\\GoogleChromePortable.exe --kiosk"))
    {
      char buff[255]="";
      strncpy(buff,strerror(errno),255);
      WriteToLog(buff);
    }
  }
  return 0;
}


