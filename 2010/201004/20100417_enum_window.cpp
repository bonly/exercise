/**
 *  @file 20100417_enum_window.cpp
 *
 *  @date 2012-3-11
 *  @Author: Bonly
 */
#include <windows.h>
#include <cstdio>
BOOL CALLBACK EnumWindowsProc(
  HWND hwnd,      // handle to parent window
  LPARAM lParam   // application-defined value
)
{
  char a[125];
  ::GetWindowText(hwnd,a,124);
  /*
  if(strcmp(a,"C:\\WINDOWS\\system32\\cmd.exe")==0)
  {
    ::SetWindowText(hwnd,"Hello，World");
    return FALSE;
  }
  */
  printf("title: %s\n", a);

  return TRUE;
}
int main()
{
  ::EnumWindows(EnumWindowsProc,NULL);

  return 0;
}


/*
函数功能

　　该函数枚举所有屏幕上的顶层窗口，并将窗口句柄传送给应用程序定义的回调函数。回调函数返回FALSE将停止枚举，否则EnumWindows函数继续到所有顶层窗口枚举完为止。
函数原型
　　BOOL EnumWindows（WNDENUMPROC lpEnumFunc，LPARAM lParam）；
　　参数：
　　lpEnumFunc：指向一个应用程序定义的回调函数指针，请参看EnumWindowsProc。
　　lPararm：指定一个传递给回调函数的应用程序定义值。
　　回调函数原型
　　BOOL CALLBACK EnumWindowsProc(HWND hwnd,LPARAM lParam);
　　参数：
　　hwnd：顶层窗口的句柄
　　lparam：应用程序定义的一个值(即EnumWindows中lParam)

返回值

　　如果函数成功，返回值为非零；如果函数失败，返回值为零。若想获得更多错误信息，请调用GetLastError函数。
*/
 */
