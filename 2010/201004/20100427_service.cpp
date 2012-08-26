/**
 *  @file 20100427_service.cpp
 *
 *  @date 2012-3-14
 *  @Author: Bonly
 */

//////////////////////////////////////////////////////////////////////
// NT Service Stub Code (For XYROOT )
//////////////////////////////////////////////////////////////////////

#include <stdio.h>
#include <windows.h>
#include <winbase.h>
#include <winsvc.h>
#include <process.h>


const int nBufferSize = 500;
char pServiceName[nBufferSize+1];
char pExeFile[nBufferSize+1];
char pInitFile[nBufferSize+1];
char pLogFile[nBufferSize+1];
int nProcCount = 0;
PROCESS_INFORMATION* pProcInfo = 0;

SERVICE_STATUS          serviceStatus;
SERVICE_STATUS_HANDLE   hServiceStatusHandle;

VOID WINAPI XYNTServiceMain( DWORD dwArgc, LPTSTR *lpszArgv );
VOID WINAPI XYNTServiceHandler( DWORD fdwControl );

CRITICAL_SECTION myCS;

void WriteLog(char* pFile, char* pMsg)
{
    ::EnterCriticalSection(&myCS);
    try
    {
        FILE* pLog = fopen(pFile,"a");
        fprintf(pLog,pMsg);
        fclose(pLog);
    } catch(...) {}
    ::LeaveCriticalSection(&myCS);
}

//////////////////////////////////////////////////////////////////////
//
// Configuration Data and Tables
//

SERVICE_TABLE_ENTRY   DispatchTable[] =
{
    {pServiceName, XYNTServiceMain},
    {NULL, NULL}
};


// helper functions

BOOL StartProcess(int nIndex)
{
    STARTUPINFO startUpInfo = { sizeof(STARTUPINFO),NULL,"",NULL,0,0,0,0,0,0,0,STARTF_USESHOWWINDOW,0,0,NULL,0,0,0};

    char pItem[nBufferSize+1];
    sprintf(pItem,"Process%d\0",nIndex);
    char pCommandLine[nBufferSize+1];
    GetPrivateProfileString(pItem,"CommandLine","",pCommandLine,nBufferSize,pInitFile);
    char pUserInterface[nBufferSize+1];
    GetPrivateProfileString(pItem,"UserInterface","N",pUserInterface,nBufferSize,pInitFile);
    BOOL bUserInterface = (pUserInterface[0]=='y'||pUserInterface[0]=='Y'||pUserInterface[0]=='1')?TRUE:FALSE;
    if(bUserInterface)
    {
        startUpInfo.wShowWindow = SW_SHOW;
        startUpInfo.lpDesktop = NULL;
    }
    else
    {
        startUpInfo.wShowWindow = SW_HIDE;
        startUpInfo.lpDesktop = "";
    }
    char pWorkingDir[nBufferSize+1];
    GetPrivateProfileString(pItem,"WorkingDir","",pWorkingDir,nBufferSize,pInitFile);
    if(CreateProcess(NULL,pCommandLine,NULL,NULL,TRUE,NORMAL_PRIORITY_CLASS,NULL,strlen(pWorkingDir)==0?NULL:pWorkingDir,&startUpInfo,&pProcInfo[nIndex]))
    {
        char pPause[nBufferSize+1];
        GetPrivateProfileString(pItem,"PauseStart","100",pPause,nBufferSize,pInitFile);
        Sleep(atoi(pPause));
        return TRUE;
    }
    else
    {
        long nError = GetLastError();
        char pTemp[121];
        sprintf(pTemp,"Failed to start program '%s', error code = %d\n", pCommandLine, nError);
        WriteLog(pLogFile, pTemp);
        return FALSE;
    }
}

void EndProcess(int nIndex)
{
    char pItem[nBufferSize+1];
    sprintf(pItem,"Process%d\0",nIndex);
    char pPause[nBufferSize+1];
    GetPrivateProfileString(pItem,"PauseEnd","100",pPause,nBufferSize,pInitFile);
    int nPauseEnd = atoi(pPause);
    if(nIndex>=0&&nIndex<nProcCount)
    {
        if(pProcInfo[nIndex].hProcess)
        {
            if(nPauseEnd>0)
            {
                PostThreadMessage(pProcInfo[nIndex].dwThreadId,WM_QUIT,0,0);
                Sleep(nPauseEnd);
            }
            TerminateProcess(pProcInfo[nIndex].hProcess,0);
        }
    }
}

BOOL BounceProcess(char* pName, int nIndex)
{
    SC_HANDLE schSCManager = OpenSCManager( NULL, NULL, SC_MANAGER_ALL_ACCESS);
    if (schSCManager==0)
    {
        long nError = GetLastError();
        char pTemp[121];
        sprintf(pTemp, "OpenSCManager failed, error code = %d\n", nError);
        WriteLog(pLogFile, pTemp);
    }
    else
    {
        SC_HANDLE schService = OpenService( schSCManager, pName, SERVICE_ALL_ACCESS);
        if (schService==0)
        {
            long nError = GetLastError();
            char pTemp[121];
            sprintf(pTemp, "OpenService failed, error code = %d\n", nError);
            WriteLog(pLogFile, pTemp);
        }
        else
        {
            SERVICE_STATUS status;
            if(nIndex>=0&&nIndex<128)
            {
                if(ControlService(schService,(nIndex|0x80),&status))
                {
                    CloseServiceHandle(schService);
                    CloseServiceHandle(schSCManager);
                    return TRUE;
                }
                long nError = GetLastError();
                char pTemp[121];
                sprintf(pTemp, "ControlService failed, error code = %d\n", nError);
                WriteLog(pLogFile, pTemp);
            }
            else
            {
                char pTemp[121];
                sprintf(pTemp, "Invalid argument to BounceProcess: %d\n", nIndex);
                WriteLog(pLogFile, pTemp);
            }
            CloseServiceHandle(schService);
        }
        CloseServiceHandle(schSCManager);
    }
    return FALSE;
}

BOOL KillService(char* pName)
{
    SC_HANDLE schSCManager = OpenSCManager( NULL, NULL, SC_MANAGER_ALL_ACCESS);
    if (schSCManager==0)
    {
        long nError = GetLastError();
        char pTemp[121];
        sprintf(pTemp, "OpenSCManager failed, error code = %d\n", nError);
        WriteLog(pLogFile, pTemp);
    }
    else
    {
        SC_HANDLE schService = OpenService( schSCManager, pName, SERVICE_ALL_ACCESS);
        if (schService==0)
        {
            long nError = GetLastError();
            char pTemp[121];
            sprintf(pTemp, "OpenService failed, error code = %d\n", nError);
            WriteLog(pLogFile, pTemp);
        }
        else
        {
            SERVICE_STATUS status;
            if(ControlService(schService,SERVICE_CONTROL_STOP,&status))
            {
                CloseServiceHandle(schService);
                CloseServiceHandle(schSCManager);
                return TRUE;
            }
            else
            {
                long nError = GetLastError();
                char pTemp[121];
                sprintf(pTemp, "ControlService failed, error code = %d\n", nError);
                WriteLog(pLogFile, pTemp);
            }
            CloseServiceHandle(schService);
        }
        CloseServiceHandle(schSCManager);
    }
    return FALSE;
}

BOOL RunService(char* pName, int nArg, char** pArg)
{
    SC_HANDLE schSCManager = OpenSCManager( NULL, NULL, SC_MANAGER_ALL_ACCESS);
    if (schSCManager==0)
    {
        long nError = GetLastError();
        char pTemp[121];
        sprintf(pTemp, "OpenSCManager failed, error code = %d\n", nError);
        WriteLog(pLogFile, pTemp);
    }
    else
    {
        SC_HANDLE schService = OpenService( schSCManager, pName, SERVICE_ALL_ACCESS);
        if (schService==0)
        {
            long nError = GetLastError();
            char pTemp[121];
            sprintf(pTemp, "OpenService failed, error code = %d\n", nError);
            WriteLog(pLogFile, pTemp);
        }
        else
        {
            if(StartService(schService,nArg,(const char**)pArg))
            {
                CloseServiceHandle(schService);
                CloseServiceHandle(schSCManager);
                return TRUE;
            }
            else
            {
                long nError = GetLastError();
                char pTemp[121];
                sprintf(pTemp, "StartService failed, error code = %d\n", nError);
                WriteLog(pLogFile, pTemp);
            }
            CloseServiceHandle(schService);
        }
        CloseServiceHandle(schSCManager);
    }
    return FALSE;
}

//////////////////////////////////////////////////////////////////////
//
// This routine gets used to start your service
//
VOID WINAPI XYNTServiceMain( DWORD dwArgc, LPTSTR *lpszArgv )
{
    DWORD   status = 0;
    DWORD   specificError = 0xfffffff;

    serviceStatus.dwServiceType        = SERVICE_WIN32;
    serviceStatus.dwCurrentState       = SERVICE_START_PENDING;
    serviceStatus.dwControlsAccepted   = SERVICE_ACCEPT_STOP | SERVICE_ACCEPT_SHUTDOWN | SERVICE_ACCEPT_PAUSE_CONTINUE;
    serviceStatus.dwWin32ExitCode      = 0;
    serviceStatus.dwServiceSpecificExitCode = 0;
    serviceStatus.dwCheckPoint         = 0;
    serviceStatus.dwWaitHint           = 0;

    hServiceStatusHandle = RegisterServiceCtrlHandler(pServiceName, XYNTServiceHandler);
    if (hServiceStatusHandle==0)
    {
        long nError = GetLastError();
        char pTemp[121];
        sprintf(pTemp, "RegisterServiceCtrlHandler failed, error code = %d\n", nError);
        WriteLog(pLogFile, pTemp);
        return;
    }

    // Handle error condition
    status = GetLastError();
    if (status!=NO_ERROR)
    {
        serviceStatus.dwCurrentState       = SERVICE_STOPPED;
        serviceStatus.dwCheckPoint         = 0;
        serviceStatus.dwWaitHint           = 0;
        serviceStatus.dwWin32ExitCode      = status;
        serviceStatus.dwServiceSpecificExitCode = specificError;
        SetServiceStatus(hServiceStatusHandle, &serviceStatus);
        return;
    }

    // Initialization complete - report running status
    serviceStatus.dwCurrentState       = SERVICE_RUNNING;
    serviceStatus.dwCheckPoint         = 0;
    serviceStatus.dwWaitHint           = 0;
    if(!SetServiceStatus(hServiceStatusHandle, &serviceStatus))
    {
        long nError = GetLastError();
        char pTemp[121];
        sprintf(pTemp, "SetServiceStatus failed, error code = %d\n", nError);
        WriteLog(pLogFile, pTemp);
    }

    for(int i=0;i<nProcCount;i++)
    {
        pProcInfo[i].hProcess = 0;
        StartProcess(i);
    }
}

//////////////////////////////////////////////////////////////////////
//
// This routine responds to events concerning your service, like start/stop
//
VOID WINAPI XYNTServiceHandler(DWORD fdwControl)
{
    switch(fdwControl)
    {
        case SERVICE_CONTROL_STOP:
        case SERVICE_CONTROL_SHUTDOWN:
            serviceStatus.dwWin32ExitCode = 0;
            serviceStatus.dwCurrentState  = SERVICE_STOPPED;
            serviceStatus.dwCheckPoint    = 0;
            serviceStatus.dwWaitHint      = 0;
            {
                for(int i=nProcCount-1;i>=0;i--)
                {
                    EndProcess(i);
                }
                if (!SetServiceStatus(hServiceStatusHandle, &serviceStatus))
                {
                    long nError = GetLastError();
                    char pTemp[121];
                    sprintf(pTemp, "SetServiceStatus failed, error code = %d\n", nError);
                    WriteLog(pLogFile, pTemp);
                }
            }
            return;
        case SERVICE_CONTROL_PAUSE:
            serviceStatus.dwCurrentState = SERVICE_PAUSED;
            break;
        case SERVICE_CONTROL_CONTINUE:
            serviceStatus.dwCurrentState = SERVICE_RUNNING;
            break;
        case SERVICE_CONTROL_INTERROGATE:
            break;
        default:
            if(fdwControl>=128&&fdwControl<256)
            {
                int nIndex = fdwControl&0x7F;
                if(nIndex>=0&&nIndex<nProcCount)
                {
                    EndProcess(nIndex);
                    StartProcess(nIndex);
                }
                else if(nIndex==127)
                {
                    for(int i=nProcCount-1;i>=0;i--)
                    {
                        EndProcess(i);
                    }
                    for(int i=0; i<nProcCount; i++)
                    {
                        StartProcess(i);
                    }
                }
            }
            else
            {
                long nError = GetLastError();
                char pTemp[121];
                sprintf(pTemp,  "Unrecognized opcode %d\n", fdwControl);
                WriteLog(pLogFile, pTemp);
            }
    };
    if (!SetServiceStatus(hServiceStatusHandle,  &serviceStatus))
    {
        long nError = GetLastError();
        char pTemp[121];
        sprintf(pTemp, "SetServiceStatus failed, error code = %d\n", nError);
        WriteLog(pLogFile, pTemp);
    }
}


//////////////////////////////////////////////////////////////////////
//
// Uninstall
//
VOID UnInstall(char* pName)
{
    SC_HANDLE schSCManager = OpenSCManager( NULL, NULL, SC_MANAGER_ALL_ACCESS);
    if (schSCManager==0)
    {
        long nError = GetLastError();
        char pTemp[121];
        sprintf(pTemp, "OpenSCManager failed, error code = %d\n", nError);
        WriteLog(pLogFile, pTemp);
    }
    else
    {
        SC_HANDLE schService = OpenService( schSCManager, pName, SERVICE_ALL_ACCESS);
        if (schService==0)
        {
            long nError = GetLastError();
            char pTemp[121];
            sprintf(pTemp, "OpenService failed, error code = %d\n", nError);
            WriteLog(pLogFile, pTemp);
        }
        else
        {
            if(!DeleteService(schService))
            {
                char pTemp[121];
                sprintf(pTemp, "Failed to delete service %s\n", pName);
                WriteLog(pLogFile, pTemp);
            }
            else
            {
                char pTemp[121];
                sprintf(pTemp, "Service %s removed\n",pName);
                WriteLog(pLogFile, pTemp);
            }
            CloseServiceHandle(schService);
        }
        CloseServiceHandle(schSCManager);
    }
}

//////////////////////////////////////////////////////////////////////
//
// Install
//
VOID Install(char* pPath, char* pName)
{
    SC_HANDLE schSCManager = OpenSCManager( NULL, NULL, SC_MANAGER_CREATE_SERVICE);
    if (schSCManager==0)
    {
        long nError = GetLastError();
        char pTemp[121];
        sprintf(pTemp, "OpenSCManager failed, error code = %d\n", nError);
        WriteLog(pLogFile, pTemp);
    }
    else
    {
        SC_HANDLE schService = CreateService
        (
            schSCManager,   /* SCManager database      */
            pName,          /* name of service         */
            pName,          /* service name to display */
            SERVICE_ALL_ACCESS,        /* desired access          */
            SERVICE_WIN32_OWN_PROCESS|SERVICE_INTERACTIVE_PROCESS , /* service type            */
            SERVICE_AUTO_START,      /* start type              */
            SERVICE_ERROR_NORMAL,      /* error control type      */
            pPath,          /* service's binary        */
            NULL,                      /* no load ordering group  */
            NULL,                      /* no tag identifier       */
            NULL,                      /* no dependencies         */
            NULL,                      /* LocalSystem account     */
            NULL
        );                     /* no password             */
        if (schService==0)
        {
            long nError =  GetLastError();
            char pTemp[121];
            sprintf(pTemp, "Failed to create service %s, error code = %d\n", pName, nError);
            WriteLog(pLogFile, pTemp);
        }
        else
        {
            char pTemp[121];
            sprintf(pTemp, "Service %s installed\n", pName);
            WriteLog(pLogFile, pTemp);
            CloseServiceHandle(schService);
        }
        CloseServiceHandle(schSCManager);
    }
}

void WorkerProc(void* pParam)
{
    char pCheckProcess[nBufferSize+1];
    GetPrivateProfileString("Settings","CheckProcess","60",pCheckProcess, nBufferSize,pInitFile);
    int nCheckProcess = atoi(pCheckProcess);
    while(nCheckProcess>0&&nProcCount>0)
    {
        ::Sleep(1000*60*nCheckProcess);
        for(int i=0;i<nProcCount;i++)
        {
            char pItem[nBufferSize+1];
            sprintf(pItem,"Process%d\0",i);
            char pRestart[nBufferSize+1];
            GetPrivateProfileString(pItem,"Restart","No",pRestart,nBufferSize,pInitFile);
            if(pRestart[0]=='Y'||pRestart[0]=='y'||pRestart[0]=='1')
            {
                DWORD dwCode;
                if(::GetExitCodeProcess(pProcInfo[i].hProcess, &dwCode))
                {
                    if(dwCode!=STILL_ACTIVE)
                    {
                        if(StartProcess(i))
                        {
                            char pTemp[121];
                            sprintf(pTemp, "Restarted process %d\n", i);
                            WriteLog(pLogFile, pTemp);
                        }
                    }
                }
                else
                {
                    long nError = GetLastError();
                    char pTemp[121];
                    sprintf(pTemp, "GetExitCodeProcess failed, error code = %d\n", nError);
                    WriteLog(pLogFile, pTemp);
                }
            }
        }
    }
}

//////////////////////////////////////////////////////////////////////
//
// Standard C Main
//
int main(int argc, char *argv[] )
{
    ::InitializeCriticalSection(&myCS);
    char pModuleFile[nBufferSize+1];
    DWORD dwSize = GetModuleFileName(NULL,pModuleFile,nBufferSize);
    pModuleFile[dwSize] = 0;
    if(dwSize>4&&pModuleFile[dwSize-4]=='.')
    {
        sprintf(pExeFile,"%s",pModuleFile);
        pModuleFile[dwSize-4] = 0;
        sprintf(pInitFile,"%s.ini",pModuleFile);
        sprintf(pLogFile,"%s.log",pModuleFile);
    }
    else
    {
        sprintf(pExeFile,"%s",argv[0]);
        sprintf(pInitFile,"%s","XYNTService.ini");
        sprintf(pLogFile,"%s","XYNTService.log");
    }

    GetPrivateProfileString("Settings","ServiceName","XYNTService",pServiceName,nBufferSize,pInitFile);
    char pCount[nBufferSize+1];
    GetPrivateProfileString("Settings","ProcCount","",pCount,nBufferSize,pInitFile);
    nProcCount = atoi(pCount);
    if(nProcCount>0)
    {
        pProcInfo = new PROCESS_INFORMATION[nProcCount];
    }
    if(argc==2&&_stricmp("-u",argv[1])==0)
    {
        UnInstall(pServiceName);
    }
    else if(argc==2&&_stricmp("-i",argv[1])==0)
    {
        Install(pExeFile, pServiceName);
    }
    else if(argc==2&&_stricmp("-b",argv[1])==0)
    {
        KillService(pServiceName);
        RunService(pServiceName,0,NULL);
    }
    else if(argc==3&&_stricmp("-b",argv[1])==0)
    {
        int nIndex = atoi(argv[2]);
        if(BounceProcess(pServiceName, nIndex))
        {
            char pTemp[121];
            sprintf(pTemp, "Bounced process %d.\n", nIndex);
            WriteLog(pLogFile, pTemp);
        }
        else
        {
            char pTemp[121];
            sprintf(pTemp, "Failed to bounce process %d.\n", nIndex);
            WriteLog(pLogFile, pTemp);
        }
    }
    else if(argc==3&&_stricmp("-k",argv[1])==0)
    {
        if(KillService(argv[2]))
        {
            char pTemp[121];
            sprintf(pTemp, "Killed service %s.\n", argv[2]);
            WriteLog(pLogFile, pTemp);
        }
        else
        {
            char pTemp[121];
            sprintf(pTemp, "Failed to kill service %s.\n", argv[2]);
            WriteLog(pLogFile, pTemp);
        }
    }
    else if(argc>=3&&_stricmp("-r",argv[1])==0)
    {
        if(RunService(argv[2], argc>3?(argc-3):0,argc>3?(&(argv[3])):NULL))
        {
            char pTemp[121];
            sprintf(pTemp, "Ran service %s.\n", argv[2]);
            WriteLog(pLogFile, pTemp);
        }
        else
        {
            char pTemp[121];
            sprintf(pTemp, "Failed to run service %s.\n", argv[2]);
            WriteLog(pLogFile, pTemp);
        }
    }
    else
    {
        if(_beginthread(WorkerProc, 0, NULL)==-1)
        {
            long nError = GetLastError();
            char pTemp[121];
            sprintf(pTemp, "_beginthread failed, error code = %d\n", nError);
            WriteLog(pLogFile, pTemp);
        }
        if(!StartServiceCtrlDispatcher(DispatchTable))
        {
            long nError = GetLastError();
            char pTemp[121];
            sprintf(pTemp, "StartServiceCtrlDispatcher failed, error code = %d\n", nError);
            WriteLog(pLogFile, pTemp);
        }
    }
    delete []pProcInfo;
    ::DeleteCriticalSection(&myCS);
    return 0;
}

/* 此服务按ini配置的顺序起动指定程序
 * -i 安装 -u 卸载
程序同一个目录下放xyntservice.ini
 [Settings]
  ServiceName = XYNTService
  ProcCount = 3
  CheckProcess = 30

 [Process0]
  CommandLine = notepad.exe
  WorkingDir = c:\MyDir
  PauseStart = 1000
  PauseEnd = 1000
  UserInterface = Yes
  Restart = Yes

 [Process1]
  CommandLine = c:\MyDir\XYDataManager.exe
  WorkingDir = c:\MyDir
  PauseStart = 1000
  PauseEnd = 1000
  UserInterface = Yes
  Restart = Yes

 [Process2]
  CommandLine= java XYRoot.XYRoot XYRootJava.ini
  UserInterface = No
  Restart = No


 */



