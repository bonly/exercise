/**
 *  @file 20100502_loguser.cpp
 *
 *  @date 2012-3-20
 *  @author Bonly
 */
#include <windows.h>
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

    WTSGETACTIVECONSOLESESSIONID lpfnWTSGetActiveConsoleSessionId = (WTSGETACTIVECONSOLESESSIONID)GetProcAddress(hmod,"WTSGetActiveConsoleSessionId");
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

int main()
{
    ExecuteProcess("c:\windows\notepad.exe");
    return 0;
}

