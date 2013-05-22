// simp.cpp : Defines the entry point for the DLL application.
//

#include <stdio.h>
#include <windows.h>
#include "detours.h"

static LONG dwSlept = 0;
static VOID (WINAPI * TrueSleep)(DWORD dwMilliseconds) = Sleep;

__declspec(dllexport) VOID WINAPI TimedSleep(DWORD dwMilliseconds)
{
    DWORD dwBeg = GetTickCount();
    TrueSleep(dwMilliseconds);
    DWORD dwEnd = GetTickCount();

    InterlockedExchangeAdd(&dwSlept, dwEnd - dwBeg);
}

BOOL WINAPI DllMain(HINSTANCE hinst, DWORD dwReason, LPVOID reserved)
{
    LONG error;
    (void)hinst;
    (void)reserved;

    if (dwReason == DLL_PROCESS_ATTACH) {
        printf("simple.dll: Starting.\n");
        fflush(stdout);

        DetourRestoreAfterWith();

        DetourTransactionBegin();
        DetourUpdateThread(GetCurrentThread());
        DetourAttach(&(PVOID&)TrueSleep, TimedSleep);
        error = DetourTransactionCommit();

        if (error == NO_ERROR) {
            printf("simple.dll: Detoured Sleep().\n");
        }
        else {
            printf("simple.dll: Error detouring Sleep(): %d\n", error);
        }
    }
    else if (dwReason == DLL_PROCESS_DETACH) {
        DetourTransactionBegin();
        DetourUpdateThread(GetCurrentThread());
        DetourDetach(&(PVOID&)TrueSleep, TimedSleep);
        error = DetourTransactionCommit();

        printf("simple.dll: Removed Sleep() (result=%d), slept %d ticks.\n",
               error, dwSlept);
        fflush(stdout);
    }
    return TRUE;
}


