package main

import "C"
import "fmt"

//export Sum
func Sum(arg1, arg2 int32) int32 {
	return arg1 + arg2
}

//export Hello
func Hello() {
	fmt.Println("hello world from go dll")
}

func main() {
}

/*
http://qiita.com/y-okubo/items/29357a057c65fa8fd96f

mydll.def:
LIBRARY mydllDef
EXPORTS
Sum @ 1
Hello @ 2

go build -buildmode=c-archive mydll.go

gcc -m64 -shared -o mydll.dll mydll.def mydll.a -Wl,--allow-multiple-definition -static -lstdc++ -lwinmm -lntdll -lWs2_32
C:\Program Files (x86)\Microsoft Visual Studio 14.0\VC\bin\x86_amd64\dumpbin.exe /headers rest.dll | findstr machine


#if UNITY_STANDALONE_WIN
	[DllImport ("mydll", EntryPoint = "Hello")]
	extern static System.IntPtr Foo();
#endif

UNITY_EDITOR 编辑器调用。

UNITY_STANDALONE_OSX 专门为Mac OS（包括Universal，PPC和Intelarchitectures）平台的定义。

UNITY_DASHBOARD_WIDGET Mac OS Dashboard widget (Mac OS仪表板小部件)。

UNITY_STANDALONE_WIN Windows 操作系统。

UNITY_STANDALONE_LINUX Linux的独立的应用程序。

UNITY_STANDALONE 独立的平台（Mac，Windows或Linux）。

UNITY_WEBPLAYER 网页播放器（包括Windows和Mac Web播放器可执行文件）。

UNITY_WII Wii游戏机平台。

UNITY_IPHONE iPhone平台。

UNITY_ANDROID Android平台。

UNITY_PS3 PlayStation 3。

UNITY_XBOX360 Xbox 360。

UNITY_NACL 谷歌原生客户端（使用这个必须另外使用UNITY_WEBPLAYER）。

UNITY_FLASH Adobe Flash。

判断Unity版本，目前支持的版本

UNITY_2_6 平台定义为主要版本的Unity 2.6。

UNITY_2_6_1 平台定义的特定版本1的主要版本2.6。

UNITY_3_0 平台定义为主要版本的Unity 3.0。

UNITY_3_0_0 平台定义的特定版本的Unity 3.0 0。

UNITY_3_1 平台定义为主要版本的Unity 3.1。

UNITY_3_2 平台定义为主要版本的Unity 3.2。

UNITY_3_3 平台定义为主要版本的Unity 3.3。

UNITY_3_4 平台定义为主要版本的Unity 3.4。

UNITY_3_5 平台定义为主要版本的Unity 3.5。

UNITY_4_0 平台定义为主要版本的Unity 4.0。

UNITY_4_0_1 主要版本4.0.1统一的平台定义。

UNITY_4_1 平台定义为主要版本的Unity 4.1。

UNITY_4_2 平台定义为主要版本的Unity 4.2。


Assets/../Editor	Plugin will be set only compatible with Editor, and won’t be used when building to platform.
Assets/../Editor/(x86 or x86_64 or x64)	Plugin will be set only compatible with Editor, CPU value will be assigned depending on the subfolder.
Assets/../Plugins/(x86_64 or x64)	x64 Standalone plugins will be set as compatible.
Assets/../Plugins/x86	x86 Standalone plugins will be set as compatible.
Assets/Plugins/Android/(x86 or armeabi or armeabi-v7a)	Plugin will be set only compatible with Android, if CPU subfolder is present, CPU value will be set as well.
Assets/Plugins/iOS	Plugin will be set only compatible with iOS.
Assets/Plugins/WSA/(x86 or ARM)	Plugin will be set only compatible with Windows Store apps and Windows Phone 8.1, if subfolder is CPU present, CPU value will be set as well. Metro keyword can be used instead of WSA.
Assets/Plugins/WSA/(SDK80 or SDK81 or PhoneSDK81)	Same as above, additionally SDK value will be set, you can also add CPU subfolder afterwards. For compatibility reasons, SDK81 - Win81, PhoneSDK81 - WindowsPhone81.
Assets/Plugins/Tizen	Plugin will be set only compatible with Tizen.
Assets/Plugins/PSVita	Plugin will be set only compatible with Playstation Vita.
Assets/Plugins/PS4	Plugin will be set only compatible with Playstation 4.
Assets/Plugins/SamsungTV	Plugin will be set only compatible with Samsung TV.
*/
