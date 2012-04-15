/**
 *  @file 20100423_keep_chrome_open.cpp
 *
 *  @date 2012-3-12
 *  @author Bonly
 */
#include <windows.h>
#include <dir.h>
#include <TlHelp32.h>
#include <stdio.h>
#include <process.h> ///for execl
#include <getopt.h>

#define SCNAME "chrome_mon"

enum{SC_RUN,INSTALL,RUN,STOP,DEL};

#define SLEEP_TIME 5000
#define LOGFILE "C:/keep_chrome_open.log"
void ExecuteProcess(char *filename/*process name to execute*/);
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

void usage(char *pname)
{
    printf("usage : %s [-h] [-i] [-r] [-s] [-d]\n", pname);
    printf("        -h show help, optional\n");
    printf("        -i install service\n");
    printf("        -r run service\n");
    printf("        -s stop service\n");
    printf("        -d remove service\n");
    printf("-----------------------------------------------\n");
    printf("example:\n");
    printf("%s -i\n", pname);
}

int read_arg(int argc, char* argv[])
{
    int optch;
    const char optstring[] = "hirsd";

    while((optch = getopt(argc, argv, optstring)) != -1)
    {
        switch(optch)
        {
            case 'h':   /* 打印帮助信息 */
                usage(argv[0]);
                exit(0);

            case 's':  /* 停止*/
                return STOP;
                break;

            case 'i':  /* 安装 */
                return INSTALL;
                break;

            case 'r': /* 运行*/
                return RUN;
                break;

            case 'd':
                return DEL;
                break;

            default:
                return SC_RUN;
                break;
        }
    }
    return -1;
}

int sc_run()
{
    WriteToLog("begin scrun\n");
    SERVICE_TABLE_ENTRY ServiceTable[2];
    ServiceTable[0].lpServiceName = (LPSTR) "Keep_chrome_open";
    ServiceTable[0].lpServiceProc = (LPSERVICE_MAIN_FUNCTION) ServiceMain;

    ServiceTable[1].lpServiceName = NULL;
    ServiceTable[1].lpServiceProc = NULL;

    StartServiceCtrlDispatcher(ServiceTable); ///启动服务的控制分派机线程
    return 0;
}

int install(char *pname)
{
    char cmd[255]="";
    char path[100]="";
    getcwd(path, 100);
    sprintf(cmd, "sc create %s binpath= \"%s\\%s\"", SCNAME, path, pname);
    WriteToLog(cmd);
    return system(cmd);
}

int stop()
{
    char cmd[255]="";
    sprintf(cmd, "sc stop %s",SCNAME);
    return system(cmd);
}

int start()
{
    char cmd[255]="";
    sprintf(cmd, "sc start %s",SCNAME);
    return system(cmd);
}

int del()
{
    char cmd[255]="";
    sprintf(cmd, "sc delete %s",SCNAME);
    return system(cmd);
}

int main(int argc, char* argv[])
{
    switch(read_arg(argc,argv))
    {
        case INSTALL:
            install(argv[0]);
            break;
        case RUN:
            start();
            break;
        case STOP:
            stop();
            break;
        case DEL:
            del();
            break;
        case SC_RUN:
        default:
            sc_run();
            break;
    }

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
    ServiceStatus.dwControlsAccepted = SERVICE_ACCEPT_STOP
                | SERVICE_ACCEPT_SHUTDOWN; ///接受的操作请求
    ServiceStatus.dwWin32ExitCode = 0; ///退出时带出的返回值
    ServiceStatus.dwServiceSpecificExitCode = 0; ///退出时带出的返回值
    ServiceStatus.dwCheckPoint = 0; ///初始化服务需30秒以上时用于主线程提示，<30时为0
    ServiceStatus.dwWaitHint = 0; ///初始化服务需30秒以上时的提示

    hStatus = RegisterServiceCtrlHandler( ///注册处理函数
                "CheckChrome", (LPHANDLER_FUNCTION) ControlHandler);
    if (hStatus == (SERVICE_STATUS_HANDLE) 0)
    {
        return; ///注册失败
    }

    error = InitService(); ///初始化服务
    if (error)
    {
        WriteToLog("Service Init failed\n");
        ServiceStatus.dwCurrentState = SERVICE_STOPPED; ///初始化失败
        ServiceStatus.dwWin32ExitCode = -1;
        SetServiceStatus(hStatus, &ServiceStatus);
        return;
    }

    ServiceStatus.dwCurrentState = SERVICE_RUNNING; ///向SCM报告服务状态
    SetServiceStatus(hStatus, &ServiceStatus);

    while (ServiceStatus.dwCurrentState == SERVICE_RUNNING)
    {
        int result = KeepChromeOpen();
        if (result != 0)
        {
            ServiceStatus.dwCurrentState = SERVICE_STOPPED;
            ServiceStatus.dwWin32ExitCode = -1;
            SetServiceStatus(hStatus, &ServiceStatus);
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
    switch (dwCode)
    {
        case SERVICE_CONTROL_STOP:
            ServiceStatus.dwCurrentState = SERVICE_STOPPED;
            ServiceStatus.dwWin32ExitCode = 0;
            SetServiceStatus(hStatus, &ServiceStatus);
            WriteToLog("Stopped service\n");
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

    HANDLE hSnapshot;
    int count = 0;
    hSnapshot = ::CreateToolhelp32Snapshot(TH32CS_SNAPPROCESS, 0); //创建快照
    if (INVALID_HANDLE_VALUE == hSnapshot)
    {
        WriteToLog("CreateToolhelp32Snapshot call faild!\n");
        return 1;
    }

    PROCESSENTRY32 process;
    process.dwSize = sizeof(PROCESSENTRY32); //注意 必不可少
    BOOL first = ::Process32First(hSnapshot, &process);

    bool found = false;
    while (first) //循环 列出进程信息
    {
        ++count;
        if (strstr(process.szExeFile, "chrome.exe") != 0)
        {
            found = true;
            WriteToLog("found chrome\n");
            break;
        }
        first = ::Process32Next(hSnapshot, &process);
    }

    ::CloseHandle(hSnapshot);
    if (!found)
    {
        WriteToLog("not found chrome\n");
        //system("Chrome --kiosk");
        //spawnl(_P_OVERLAY, "cmd","/c", "start","D:\\Chrome\\Chrome\\GoogleChromePortable.exe","--kiosk",NULL); ///_P_OVERLAY结束原来的程序  _P_WAIT结束后返回到原来的程序
        //execl("notepad", "notepad", "test.txt", NULL);
        //if (0 != system("D:\\Chrome\\Chrome\\GoogleChromePortable.exe --kiosk"))
        //system("echo ""cmd /c start D:\\Chrome\\Chrome\\GoogleChromePortable.exe --kiosk"" > c:\\kk.bat");
        /*
        spawnl(_P_OVERLAY,"c:\\kk.bat",NULL);
        {
            char buff[255] = "";
            strncpy(buff, strerror(errno), 255);
            WriteToLog(buff);
        }
        */
        //ExecuteProcess("c:\\windows\\notepad.exe");
        ExecuteProcess("chrome --kiosk");
    }
    return 0;
}

typedef DWORD (*WTSGETACTIVECONSOLESESSIONID)(void);
typedef BOOL (WINAPI *WTSQUERYUSERTOKEN)(ULONG,PHANDLE);
WTSQUERYUSERTOKEN WTSProc=NULL;
HMODULE hModWTS=NULL;

void ExecuteProcess(char *filename/*process name to execute*/)
{
    PROCESS_INFORMATION pInfo;
    STARTUPINFO sInfo;

    ZeroMemory(&sInfo,sizeof(sInfo));
    ZeroMemory(&pInfo,sizeof(pInfo));

    sInfo.cb=sizeof(STARTUPINFO);
    sInfo.cbReserved2=0;
    sInfo.dwFillAttribute=0;
    sInfo.dwFlags=STARTF_USESHOWWINDOW;
    //sInfo.wShowWindow=SW_HIDE;
    sInfo.wShowWindow=SW_SHOW;
    sInfo.lpDesktop = "winsta0\\default";
    sInfo.lpDesktop=NULL;
    sInfo.lpReserved=NULL;
    sInfo.lpReserved2=NULL;

    sInfo.lpTitle = NULL;
    sInfo.dwX = 0;
    sInfo.dwY = 0;
    sInfo.dwFillAttribute = 0;


    HANDLE hToken=NULL;
    HMODULE hmod = LoadLibrary("kernel32.dll");
    if(hmod==NULL)
    {
        return;
    }

    WTSGETACTIVECONSOLESESSIONID lpfnWTSGetActiveConsoleSessionId =
                (WTSGETACTIVECONSOLESESSIONID)GetProcAddress(hmod,"WTSGetActiveConsoleSessionId");
    DWORD dwSessionId=-1;

    // Get the Session ID of active user using
    // API - WTSGetActiveConsoleSessionId

    dwSessionId = lpfnWTSGetActiveConsoleSessionId();
    FreeLibrary(hmod);

    hModWTS = LoadLibrary("wtsapi32.dll");

    if(hModWTS !=NULL)
    {
        WTSProc = (WTSQUERYUSERTOKEN) GetProcAddress(hModWTS,"WTSQueryUserToken");
        if(WTSProc!=NULL)
        {
            if(WTSProc(dwSessionId,&hToken)!=0)
            {
                if(CreateProcessAsUser(hToken,NULL,filename,NULL,NULL,FALSE,NORMAL_PRIORITY_CLASS | CREATE_NEW_CONSOLE,NULL,NULL,&sInfo,&pInfo))
                {
                    CloseHandle(pInfo.hThread);
                    CloseHandle(pInfo.hProcess);
                }
                else
                {
                    //Fail
                }
                CloseHandle(hToken);
            }
        }
        FreeLibrary(hModWTS);
        hModWTS = NULL;
    }

return;
}

/*
 * sc create test1 binpath= e:\workspace\exercise\Debug\20100423_keep_chrome_open.exe
 * sc delete test1
 */
