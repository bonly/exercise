//#include <windows.h>
#include <winsock2.h>
#include <mswsock.h>
//#include <winsock.h>
#include "detours.h"
#include <stdio.h>
#include <fstream>
using namespace std;
#pragma warning(disable:4127)  

VOID NullExport()
{
}


int (WINAPI * Real_WSASend)(SOCKET a0,
                            LPWSABUF a1,
                            DWORD a2,
                            LPDWORD a3,
                            DWORD a4,
                            LPWSAOVERLAPPED a5,
                            LPWSAOVERLAPPED_COMPLETION_ROUTINE a6)
    = WSASend;

int (WINAPI * Real_WSASendTo)(SOCKET a0,
                              LPWSABUF a1,
                              DWORD a2,
                              LPDWORD a3,
                              DWORD a4,
                              CONST sockaddr* a5,
                              int a6,
                              LPWSAOVERLAPPED a7,
                              LPWSAOVERLAPPED_COMPLETION_ROUTINE a8)
    = WSASendTo;

int WINAPI Mine_WSASend(SOCKET a0,
                        LPWSABUF a1,
                        DWORD a2,
                        LPDWORD a3,
                        DWORD a4,
                        LPWSAOVERLAPPED a5,
                        LPWSAOVERLAPPED_COMPLETION_ROUTINE a6)
{
    printf("wsa send\n");
    int rv = 0;
    //__try {
        rv = Real_WSASend(a0, a1, a2, a3, a4, a5, a6);
    //} __finally {
        //printf("%p: WSASend(,,,,,,) -> %x\n", a0, rv);
    //};
    return rv;
}

int WINAPI Mine_WSASendTo(SOCKET a0,
                          LPWSABUF a1,
                          DWORD a2,
                          LPDWORD a3,
                          DWORD a4,
                          sockaddr* a5,
                          int a6,
                          LPWSAOVERLAPPED a7,
                          LPWSAOVERLAPPED_COMPLETION_ROUTINE a8)
{
    int rv = 0;
    ofstream fd("c:\\abc.log",ios_base::app);
    fd << "send: "<<a1->buf<<endl;
    fd.close();
    //printf("send\n");

    //__try {
        rv = Real_WSASendTo(a0, a1, a2, a3, a4, a5, a6, a7, a8);
    //} __finally {
        //printf("%p: WSASendTo(,,,,,,,,) -> %x\n", a0, rv);
    //};
    return rv;
}

BOOL APIENTRY DllMain(HINSTANCE hModule, DWORD dwReason, PVOID lpReserved)
{
    (void)lpReserved;
    (void)hModule;
    if (dwReason == DLL_PROCESS_ATTACH) {
        DetourTransactionBegin();
        DetourUpdateThread(GetCurrentThread());
        DetourAttach(&(PVOID&)Real_WSASendTo, Mine_WSASendTo);
        DetourTransactionCommit();
    }
    else if (dwReason == DLL_PROCESS_DETACH) {
        DetourTransactionBegin();
        DetourUpdateThread(GetCurrentThread());
        DetourDetach(&(PVOID&)Real_WSASendTo, Mine_WSASendTo);
        DetourTransactionCommit();
    }
    return TRUE;

}


