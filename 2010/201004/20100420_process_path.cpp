/**
 *  @file 20100420_process_path.cpp
 *
 *  @date 2012-3-11
 *  @Author: Bonly
 */
#include <windows.h>
void ShowProcess2()
{
  DWORD processPID[MAX_NUM]; //保存进程ID
  DWORD dwneed;
  ::EnumProcesses(processPID,sizeof(processPID),&dwneed);
  int count = dwneed/sizeof(DWORD);
  printf("total process: %d\n",count);
  HANDLE hProcess;
  HMODULE hModule;
  DWORD need;
  DWORD nSize = 0;
  wchar_t fileName[100] = {0};
  for(int i=0; i<count; ++i)
  {
    hProcess = ::OpenProcess(PROCESS_QUERY_INFORMATION | PROCESS_VM_READ,FALSE,processPID[i]);
    if(hProcess)
    {
      ::EnumProcessModules(hProcess,&hModule,sizeof(hModule),&need);
      ::GetModuleFileNameExA(hProcess,hModule, (LPSTR)fileName,sizeof(fileName));
      printf("%d     %s\n",i,fileName);
    }
  }
  CloseHandle(hProcess);
  CloseHandle(hModule);
}


/*
EnumProcessModules Function
获得指定进程中所有模块的句柄。
语法
BOOL WINAPI EnumProcessModules(in   HANDLE hProcess,out  HMODULE *lphModule,in   DWORD cb,out  LPDWORD lpcbNeeded);
参数
hProcess [传入]
指定进程的句柄。
lphModule [传出]
用来存放所有模块句柄的数组。
cb [传入]
lphModule参数所传入的数组的大小，单位是字节。
lpcbNeeded [传出]
要把所有模块的句柄存放进lphModule参数所传入的数组中，所需要的字节数。
返回值
如果函数执行成功，则返回值为非零。
如果函数执行失败，则返回值为零。可以调用 GetLastError函数来获得更多的错误信息。
备注
EnumProcessModules的设计，主要是为调试器和类似程序在必须获取其它进程的模块信息时使用的。如果目标进程的模块列表已经损坏，或者尚未初始化，那么，EnumProcessModules可能会执行失败，或者返回错误的信息。

建议使用数组来存放大批模块句柄的值，因为，难以确定在你调用EnumProcessModules时，有多少模块在当时的进程中。如果lphModule所传递的数组太小，以至于无法容纳进程中的所有的模块句柄，这点你可以从lpcbNeeded参数的值来判断。如果lpcbNeeded比传入的数组要大，请扩大数组并重新调用EnumProcessModules函数。

调用EnumProcessModules函数所枚举的模块数量，是lpcbNeeded参数的值除以HMODULE的大小，HMODULE的大小可以用sizeof来获得。

有LOAD_LIBRARY_AS_DATAFILE标志时，EnumProcessModules函数无法检索已加载的模块句柄。有关详细信息，请参阅LoadLibraryEx。

不要对本函数的任何返回值调用CloseHandle函数。这些返回的信息只是来自于一次快照，并没有资源需要释放。

如果在WOW64上运行的32位应用程序调用本函数，那它只能枚举32位进程中的模块。对于64位的进程，本函数将执行失败，错误代码是ERROR_PARTIAL_COPY (299)。

使用CreateToolhelp32Snapshot函数，对指定进程和堆栈、模块，以及这些进程所使用的线程采取快照。

从Windows 7和Windows Server 2008 R2开始，Psapi.h为PSAPI函数建立了版本号。PSAPI的版本号影响程序在必须加载库和调用本函数时所用的名字。

如果PSAPI_VERSION大于等于2，本函数将被定义为Psapi.h或者Kernel32.lib和Kernel32.dll中的K32EnumProcessModules。在PSAPI_VERSION为1时调用K32EnumProcessModules，本函数将被定义为在Psapi.h或者Psapi.lib和Psapi.dll封装中的EnumProcessModulesas函数。

程序必须运行在较早版本的Windows，或者Windows 7和以后的版本上时，你总是可以使用EnumProcessModules函数。为保证符号的正确识别，请添加Psapi.lib到TARGETLIBS宏并且使用参数–DPSAPI_VERSION=1来编译程序。在程序运行时动态加载Psapi.dll。
 */
