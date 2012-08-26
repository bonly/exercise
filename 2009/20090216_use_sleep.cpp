// app_sleep.cpp : Defines the entry point for the console application.
//

#include <windows.h>
#include <stdio.h>

int __cdecl main(int argc, char ** argv)
{
    if (argc == 2) {
        Sleep(atoi(argv[1]) * 1000);
    }
    else {
        printf("sleep5.exe: Starting.\n");

        Sleep(5000);

        printf("sleep5.exe: Done sleeping.\n");
    }
    return 0;
}

