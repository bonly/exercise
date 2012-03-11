/**
 *  @file 20100418_Enum_process.cpp
 *
 *  @date 2012-3-11
 *  @Author: Bonly
 */
#include <cstdlib>
#include <cstdio>
#include <iostream>
#include <windows.h>
#include <TlHelp32.h>
using namespace std;


int main(int argc, char *argv[])
{
    HANDLE  hSnapshot;
  int count = 0;
  hSnapshot = ::CreateToolhelp32Snapshot(TH32CS_SNAPPROCESS,0);  //创建快照
  if(INVALID_HANDLE_VALUE == hSnapshot)
  {
    printf("CreateToolhelp32Snapshot call faild。");
    return 1;
  }
  PROCESSENTRY32 process;
  process.dwSize = sizeof(PROCESSENTRY32);   //注意 必不可少
  BOOL first = ::Process32First(hSnapshot,&process);
  printf("process name:\t\tprocess ID\n");

  while(first)     //循环 列出进程信息
  {
    ++count;
    printf("%s\t\t%ld\n",process.szExeFile,process.th32ProcessID);
    first = ::Process32Next(hSnapshot,&process);
  }
  printf("total process: %d\n",count);


  ::CloseHandle(hSnapshot);
    system("PAUSE");
    return EXIT_SUCCESS;
}


/*
每一个应用程序实例在运行起来后都会在当前系统下产生一个进程，大多数应用程序均拥有可视界面，用户可以通过标题栏上的关闭按钮关闭程序。但是也有为数不少的在后台运行的程序是没有可视界面的，对于这类应用程序用户只能通过CTRL+ALT+DEL热键呼出"关闭程序"对话框显示出当前系统进程列表，从中可以结束指定的任务。显然，该功能在一些系统监控类软件中还是非常必需的，其处理过程大致可以分为两步：借助系统快照实现对系统当前进程的枚举和根据枚举结果对进程进行管理。本文下面即将对此过程的实现进行介绍。
当前进程的枚举
要对当前系统所有已开启的进程进行枚举，就必须首先获得那些加载到内存的进程当前相关状态信息。在Windows操作系统下，这些进程的当前状态信息不能直接从进程本身获取，系统已为所有保存在系统内存中的进程、线程以及模块等的当前状态的信息制作了一个只读副本--系统快照，用户可以通过对系统快照的访问完成对进程当前状态的检测。在具体实现时，系统快照句柄的获取是通过Win32 API函数CreateToolhelp32Snapshot()来完成的，通过该函数不仅可以获取进程快照，而且对于堆、模块和线程的系统快照同样可以获取。
使用这个函数前必须在头文件里包含tlhelp32.h头文件。
CreateToolhelp32Snapshot函数为指定的进程、进程使用的堆[HEAP]、模块[MODULE]、线程[THREAD]）建立一个快照[snapshot]。
HANDLE WINAPI CreateToolhelp32Snapshot(
                                       DWORD dwFlags,
                                       DWORD th32ProcessID
                                       );

参数：其中，参数dwFlags:指定将要创建包含哪一类系统信息的快照句柄，本程序中只需要检索系统进程信息，因此可将其设置为TH32CS_SNAPPROCESS；函数第二个参数th32ProcessID`则指定了进程的标识号，当设置为0时指定当前进程。
dwFlags
[输入]指定快照中包含的系统内容，这个参数能够使用下列数值（变量）中的一个。
   TH32CS_INHERIT - 声明快照句柄是可继承的。
   TH32CS_SNAPALL - 在快照中包含系统中所有的进程和线程。
   TH32CS_SNAPHEAPLIST - 在快照中包含在th32ProcessID中指定的进程的所有的堆。
   TH32CS_SNAPMODULE - 在快照中包含在th32ProcessID中指定的进程的所有的模块。
   TH32CS_SNAPPROCESS - 在快照中包含系统中所有的进程。
   TH32CS_SNAPTHREAD - 在快照中包含系统中所有的线程。
th32ProcessID
[输入]指定将要快照的进程ID。如果该参数为0表示快照当前进程。该参数只有在设置了TH32CS_SNAPHEAPLIST或TH32CS_SNAPMOUDLE后才有效，在其他情况下该参数被忽略，所有的进程都会被快照。
返回值：
调用成功，返回快照的句柄，调用失败，返回INVAID_HANDLE_VALUE。
备注：
使用GetLastError函数查找该函数产生的错误状态码。
要删除快照，使用CloseHandle函数。
在得到快照句柄之后只能以只读的方式对其进行访问。至于对系统快照句柄的使用同普通对象句柄的使用并没有什么太大区别，在使用完之后也需要通过CloseHandle()函数将其销毁。

BOOL Process32First（）函数
参数：HANDLE hSnapshot 传入的Snapshot句柄
参数：LPPROCESSENTRY32 lppe 指向PROCESSENTRY32结构的指针
作用：从Snapshot得到第一个进程记录信息

BOOL Process32Next（）函数
参数：HANDLE hSnapshot 传入的Snapshot句柄
参数：LPPROCESSENTRY32 lppe 指向PROCESSENTRY32结构的指针
作用：从Snapshot得到下一个进程记录信息

BOOL Module32First（）函数

参数：HANDLE hSnapshot传入的Snapshot句柄
参数：LPMODULEENTRY3 lpme 指向一个 MODULEENTRY32结构的指针
作用：从Snapshot得到第一个Module记录信息

BOOL Module32Next（）函数
参数：HANDLE hSnapshot传入的Snapshot句柄
参数：LPMODULEENTRY3 lpme 指向一个 MODULEENTRY32结构的指针
作用：从Snapshot得到下一个Module记录信息

BOOL Thread32First（）函数
参数：HANDLE hSnapshot传入的Snapshot句柄
参数：LPTHREADENTRY32 lpte指向一个 THREADENTRY32结构的指针
作用：从Snapshot得到第一个Thread记录信息

BOOL Thread32Next（）函数
参数：HANDLE hSnapshot传入的Snapshot句柄
参数：LPTHREADENTRY32 lpte指向一个 THREADENTRY32结构的指针
作用：从Snapshot得到下一个Thread记录信息

HANDLE OpenProcess（）函数
 参数：DWORD dwDesiredAccess 权限描叙信息
这里我用到了PROCESS_ALL_ACCESS功能是具有所有权限
参数：BOOL bInheritHandle 确定该句柄是否可以被程继承
参数：dwPrcessID 进程ID号
作用：打开一个存在的进程对象

列举进程
在得到系统的快照句柄后，就可以对当前进程的标识号进行枚举了，通过这些枚举出的进程标识号可以很方便的对进程进行管理。进程标识号通过函数Process32First() 和 Process32Next()而得到，这两个函数可以枚举出系统当前所有开启的进程，并且可以得到相关的进程信息。 这两个函数原型声明如下：
      BOOL WINAPI Process32First(HANDLE hSnapshot, LPPROCESSENTRY32 lppe);
      BOOL WINAPI Process32Next(HANDLE hSnapshot,LPPROCESSENTRY32 lppe);
   以上两个函数分别用于获得系统快照中第一个和下一个进程的信息，并将获取得到的信息保存在指针lppe所指向的PROCESSENTRY32结构中。函数第一个参数hSnapshot为由CreateToolhelp32Snapshot()函数返回得到的系统快照句柄；第二个参数lppe为指向结构PROCESSENTRY32的指针，PROCESSENTRY32结构可对进程作一个较为全面的描述，其定义如下：
      typedef struct tagPROCESSENTRY32 {
      DWORD dwSize; // 结构大小；
      DWORD cntUsage; // 此进程的引用计数；
      DWORD th32ProcessID; // 进程ID;
      DWORD th32DefaultHeapID; // 进程默认堆ID；
      DWORD th32ModuleID; // 进程模块ID；
      DWORD cntThreads; // 此进程开启的线程计数；
      DWORD th32ParentProcessID; // 父进程ID；
      LONG pcPriClassBase; // 线程优先权；
      DWORD dwFlags; // 保留；
      char szExeFile[MAX_PATH]; // 进程全名；
      } PROCESSENTRY32;
    以上三个API函数均在头文件tlhelp32.h中声明，运行时需要有kernel32.lib库的支持。通过这三个函数可以枚举出当前系统已开启的所有进程，并可获取到进程的各相关信息。
 */
