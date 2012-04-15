/**
 *  @file 20100501_createprocesswithlog.cpp
 *
 *  @date 2012-3-20
 *  @Author: Bonly
 */

//#define _WIN32_WINNT    0x0500
#include <Windows.h>
#include <WinBase.h>
#include <stdio.h>
#include <lm.h>
#include <Lmaccess.h>

//#pragma comment(lib,"Netapi32.lib")
//#pragma comment(lib,"Advapi32.lib")

int APIENTRY WinMain(HINSTANCE hInstance, HINSTANCE hPrevInstance,
            LPSTR lpCmdLine, int nCmdShow = SW_SHOW)
{
    STARTUPINFOW starinfo = { 0 };
    PROCESS_INFORMATION proinfo = { 0 };

    USER_INFO_1 user;
    DWORD dwLevel = 1;
    DWORD dwError = 0;
    NET_API_STATUS nStatus;

    user.usri1_name = L"abc";
    user.usri1_password = L"1212";
    user.usri1_priv = USER_PRIV_USER;
    user.usri1_home_dir = NULL;
    user.usri1_comment = NULL;
    user.usri1_flags = UF_SCRIPT;
    user.usri1_script_path = NULL;

    nStatus = NetUserAdd(NULL, dwLevel, (LPBYTE) &user, &dwError);

    //LogonUser("Bonly",".",NULL,LOGON32_LOGON_INTERACTIVE,LOGON32_PROVIDER_DEFAULT,&tokena);

    if (nStatus == NERR_Success)
    {
        Sleep(100);
        MessageBoxW(NULL, L"创建成功", L"OK", NULL);

        starinfo.cb = sizeof(starinfo);
        CreateProcessWithLogonW(L"bonly", NULL, NULL,
                    LOGON_WITH_PROFILE, NULL, L"c:\\windows\\notepad.EXE", DEBUG_PROCESS,
                    NULL, NULL, &starinfo, &proinfo);

        nStatus = NetUserDel(NULL, L"guest");
        MessageBoxW(NULL, L"删除成功", L"OK", NULL);
    }

    return 0;

}

