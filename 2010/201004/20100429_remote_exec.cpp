// ExecProm.cpp : 定义控制台应用程序的入口点。
//

//#include "stdafx.h"

/**
 *  @file 20100429_remote_exec.cpp
 *
 *  @date 2012-3-19
 *  @Author: Bonly
 */

//#include "stdafx.h"
//#define _WIN32_WINNT 0x0501 //winxp后的版本
#include <windows.h>
#include <Winbase.h> ///for WTSGetActiveConsoleSessionId
#include <stdio.h>
#include <process.h>
#include <Tlhelp32.h>
#include <tchar.h>
#include <psapi.h>
#include <stdio.h>
#include <STDLIB.H>
#include <tlhelp32.h>
#include <WtsApi32.h>

#pragma comment(lib, "WtsApi32.lib")
//#pragma comment (lib,"psapi")

//#define LPSTR LPWSTR
//提升进程访问权限
bool EnableDebugPriv(HANDLE &hToken)
{
    //HANDLE hToken;
    LUID sedebugnameValue;
    TOKEN_PRIVILEGES tkp;
    if (!OpenProcessToken(GetCurrentProcess(),
                TOKEN_ADJUST_PRIVILEGES | TOKEN_QUERY, &hToken))
        return false;
    if (!LookupPrivilegeValue(NULL, SE_DEBUG_NAME, &sedebugnameValue))
    {
        CloseHandle(hToken);
        return false;
    }
    tkp.PrivilegeCount = 1;
    tkp.Privileges[0].Luid = sedebugnameValue;
    tkp.Privileges[0].Attributes = SE_PRIVILEGE_ENABLED;
    if (!AdjustTokenPrivileges(hToken, FALSE, &tkp, sizeof(tkp), NULL, NULL))
    {
        CloseHandle(hToken);
        return false;
    }
    return true;
}
// Get username from session id
bool GetSessionUserName(DWORD dwSessionId, char username[256])
{
    LPTSTR pBuffer = NULL;
    DWORD dwBufferLen;

    BOOL bRes = WTSQuerySessionInformation(WTS_CURRENT_SERVER_HANDLE,
                dwSessionId, WTSUserName, &pBuffer, &dwBufferLen);

    if (bRes == FALSE)
        return false;

    lstrcpy((LPSTR) username, pBuffer);
    WTSFreeMemory(pBuffer);

    return true;
}

// Get domain name from session id
bool GetSessionDomain(DWORD dwSessionId, char domain[256])
{
    LPTSTR pBuffer = NULL;
    DWORD dwBufferLen;

    BOOL bRes = WTSQuerySessionInformation(WTS_CURRENT_SERVER_HANDLE,
                dwSessionId, WTSDomainName, &pBuffer, &dwBufferLen);

    if (bRes == FALSE)
    {
        printf("WTSQuerySessionInformation Fail!\n");
        return false;
    }

    lstrcpy((LPSTR) domain, pBuffer);
    WTSFreeMemory(pBuffer);

    return true;
}

HANDLE GetProcessHandle(LPSTR szExeName) //遍历进程PID
{
    PROCESSENTRY32 Pc = { sizeof(PROCESSENTRY32) };
    HANDLE hSnapshot = CreateToolhelp32Snapshot(TH32CS_SNAPALL, 0);

    if (Process32First(hSnapshot, &Pc))
    {
        do
        {
            if (!stricmp((const char*) Pc.szExeFile, (const char*)szExeName))
            { //返回explorer.exe进程的PID
                //printf("PID=%d\n", Pc.th32ProcessID);
                return OpenProcess(PROCESS_ALL_ACCESS, TRUE, Pc.th32ProcessID);
            }
        } while (Process32Next(hSnapshot, &Pc));
    }

    return NULL;
}

//输出帮助的典型方法:
void Usage(void)
{
    fprintf(
                stderr,
                "===============================================================================\n"
                            "\t名称：在任意的远程桌面的session中运行指定的程序,需要具有system权限\n"
                            "\t环境：Win2003 + Visual C++ 6.0\n"
                            "\n"
                            "\t使用方法：\n"
                            "\tsession 1 c:\\win2003\\system32\\svchosts.exe //在会话1里面运行程序!\n"
                            "===============================================================================\n");
}

int main(int argc, char **argv)
{
    if (argc == 1) //遍历所有的session
    {
        // 函数的句柄
        HMODULE hInstKernel32 = NULL;
        HMODULE hInstWtsapi32 = NULL;

        // 这里的代码用的是VC6，新版的SDK已经包括此函数，无需LoadLibrary了。
        /*
        typedef DWORD (WINAPI *WTSGetActiveConsoleSessionIdPROC)();

        WTSGetActiveConsoleSessionIdPROC WTSGetActiveConsoleSessionId = NULL;
        hInstKernel32 = LoadLibrary((LPSTR) "Kernel32.dll");

        if (!hInstKernel32)
        {
            return FALSE;
        }

        WTSGetActiveConsoleSessionId = (WTSGetActiveConsoleSessionIdPROC) GetProcAddress(
                                hInstKernel32, "WTSGetActiveConsoleSessionId");

        if (!WTSGetActiveConsoleSessionId)
        {
            return FALSE;
        }

        // WTSQueryUserToken 函数，通过会话ID得到令牌
        typedef BOOL (WINAPI *WTSQueryUserTokenPROC)(ULONG SessionId,
                    PHANDLE phToken);

        WTSQueryUserTokenPROC WTSQueryUserToken = NULL;
        hInstWtsapi32 = LoadLibrary((LPSTR) "Wtsapi32.dll");
        if (!hInstWtsapi32)
        {
            return FALSE;
        }

        WTSQueryUserToken = (WTSQueryUserTokenPROC) GetProcAddress(
                    hInstWtsapi32, "WTSQueryUserToken");

        if (!WTSQueryUserToken)
        {
            return FALSE;
        }
        */

        //遍历3389登录的session:
        /*
         typedef struct _WTS_SESSION_INFO {
         DWORD SessionId;
         LPTSTR pWinStationName;
         WTS_CONNECTSTATE_CLASS State;
         }WTS_SESSION_INFO, *PWTS_SESSION_INFO;
         */
        WTS_SESSION_INFO *sessionInfo = NULL;
        DWORD sessionInfoCount;
        char domain1[256];
        char username1[256];
        BOOL result = WTSEnumerateSessions(WTS_CURRENT_SERVER_HANDLE, 0, 1,
                    &sessionInfo, &sessionInfoCount);

        unsigned int userCount(0);
        int num = 0;
        for (unsigned int i = 0; i < sessionInfoCount; ++i)
        {
            if ((sessionInfo[i].State == WTSActive)
                        || (sessionInfo[i].State == WTSDisconnected))
            {
                printf("session %d information:\n", num++);
                printf("\tsessionInfo.SessionId=%d\n",
                            sessionInfo[i].SessionId);
                GetSessionDomain(sessionInfo[i].SessionId, domain1); //获得Session Domain
                printf("\tSession Domain = %s\n", domain1);

                GetSessionUserName(sessionInfo[i].SessionId, username1);
                printf("\tSession user's name = %s\n", username1);

                userCount++;
            }
        }
        printf("session's number:%d\n\n", userCount);
        Usage();
        //printf("example:\n\tsession 1 c:\\win2003\\system32\\svchosts.exe //在会话1里面运行程序!\n");
        //printf("程序说明:在其它session中(如任意的远程桌面的session中)运行指定的程序,需要具有system权限\n");
        WTSFreeMemory(sessionInfo); //释放
    }
    else if (argc == 3) //session 1 c:\win2003\temp\klog.exe
    {
        // 得到当前登录用户的令牌
        /*
         HANDLE hTokenDup = NULL;
         bRes = WTSQueryUserToken(dwSessionId, &hTokenDup);

         if (!bRes)
         {
         printf("WTSQueryUserToken Failed!%d\n",GetLastError());
         return FALSE;
         }
         */

        /*
         bRes = ImpersonateLoggedOnUser(hTokenDup);
         if (!bRes)
         {
         printf("ImpersonateLoggedOnUser!%d\n",GetLastError());
         return FALSE;
         }
         */

        //MessageBox(NULL,"test2","test1",MB_OK);
        //system("winver.exe");

        HANDLE hThisProcess = GetCurrentProcess(); // 获取当前进程句柄
        //HANDLE hThisProcess = GetProcessHandle("winlogon.exe");
        if(hThisProcess == NULL)
          return 0;

        HANDLE hTokenThis = NULL;
        HANDLE hTokenDup = NULL;
        // 打开当前进程令牌;或用LogonUser取得hTokenThis
        /*
        if(!LogonUser(user, domain, password, LOGON32_LOGON_INTERACTIVE, LOGON32_PROVIDER_DEFAULT, &handle))
        {
             printf( "Logon User failed:%d\n",GetLastError());
             return false;
        }
        //*/

        if(!OpenProcessToken(hThisProcess, TOKEN_ALL_ACCESS, &hTokenThis))
        {
             printf( "Dup failed:%d\n",GetLastError());
             return false;
        }

        /*
         if(!EnableDebugPriv(hTokenThis)) ///提权
         {
         printf("get System Priv faild!%d\n", GetLastError());
         return FALSE;
         }
         */
        // 复制一个进程令牌，目的是为了修改session id属性，以便在其它session中创建进程;
        if(!DuplicateTokenEx(hTokenThis, STANDARD_RIGHTS_ALL, NULL, //MAXIMUM_ALLOWED
                    SecurityIdentification, TokenPrimary, &hTokenDup))
        {
             printf( "Dup failed:%d\n",GetLastError());
             return false;
        }

        /** @note DuplicateTokenEx:
         这个函数的主要功能就是根据现有的Token 来创建一个新的Token，
         我们需要的就是这个新的Token，如何得到这个新的Token，看来我们还得往下追踪哦，
         第二个参数是新建Token的权限，这个参数直接设置成MAXIMUM_ALLOWED就可以，
         第三个参数是 安全描述符的指针，这里直接设置成NULL，分配成一个默认的安全描述符，
         第四个参数设置成SecurityIdentification  的OK,(安全描述符的级别)
         第五个参数 Token类型 直接设置成TokenPrimary OK  因为它就是用于CreateProcessAsUser的函数的。
         最后一个参数就是我们想要的参数了。我们主要就是得到他，要想得到这个参数 我们必须得到第一个参数。

         第一个参数是一个已经存在的Token，数要想得到这个Token,必须找到以个进程，把这个进程的Token 拿过来为我们所用，
         得到这个Token 要运行函数OpenPrcessToken通过这个函数得到Token的句柄，
         但是要得到这个Token的句柄，必须要通过函数OpenProcess得到进程的句柄，
         要得到进程的句柄 必须通过函数ProcessIdToSessionId得到进程的ID。
         因为我们系统每次登入的时候都有一个winlogon.exe，
         我们就拿这个进程的ID(这里只用了当前进程)来得到CreateProcessAsUser函数的第一个参数。
         */
        //获取活动session id，这里要注意，如果服务器还没有被登录而使用了远程桌面，这样用是可以的，如果有多个session存在，
        //不能简单使用此函数，需要枚举所有session并确定你需要的一个，或者干脆使用循环，针对每个session都执行后面的代码
        //SetTokenInformation(hTokenDup, TokenSessionId, &dwSessionId, sizeof(DWORD)); //把session id设置到备份的令牌中
        //DWORD dwSessionId = WTSGetActiveConsoleSessionId(); ///winxp后的版本才支持此函数
        DWORD dwSessionId = atoi(argv[1]); //与会话进行连接
        //WTSQueryUserToken(dwSessionId, &hTokenDup);  //可不写?
        if(0 == SetTokenInformation(hTokenDup, TokenSessionId, &dwSessionId,
                    sizeof(DWORD)))
        {
            printf("SetTokenInformation!%d\n", GetLastError());
            return FALSE;
        }

        //好了，现在要用新的令牌来创建一个服务进程。注意：是“服务”进程！如果需要以用户身份运行，必须在前面执行LogonUser来获取用户令牌
        STARTUPINFO si;
        PROCESS_INFORMATION pi;

        ZeroMemory(&si, sizeof(STARTUPINFO));
        ZeroMemory(&pi, sizeof(PROCESS_INFORMATION));

        si.cb = sizeof(STARTUPINFO);
        si.lpDesktop = (LPSTR) "WinSta0\\Default";

        LPVOID pEnv = NULL;
        DWORD dwCreationFlag = NORMAL_PRIORITY_CLASS | CREATE_NEW_CONSOLE; // 注意标志

        //CreateEnvironmentBlock(&pEnv, hTokenDup, FALSE); // 创建环境块
        // 创建新的进程，这个进程就是你要弹出窗口的进程，它将工作在新的session中
        char path[MAX_PATH];
        lstrcpy((LPSTR) path, (LPSTR) argv[2]);
        CreateProcessAsUser(hTokenDup, NULL, (LPSTR) path, NULL, NULL, FALSE,
                    dwCreationFlag, pEnv, NULL, &si, &pi);
    }
    return 0;
}

