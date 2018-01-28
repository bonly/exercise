using System;
using System.Runtime.InteropServices; //for DllImport
using System.Threading;
using System.IO;

public class Net{
    const string lib = "./libtechappen.so";

    [DllImport (lib, EntryPoint = "Run", CallingConvention = CallingConvention.StdCall, CharSet = CharSet.Ansi)]
    private static extern int Srv();

    [UnmanagedFunctionPointer(CallingConvention.StdCall)]
    delegate void MsgCallback(uint sn, string str);  // 定义供外部调用的回调函数
    
    [DllImport (lib, EntryPoint = "SetCallBack", CallingConvention = CallingConvention.StdCall)]
    extern static System.IntPtr set_callback([MarshalAs(UnmanagedType.FunctionPtr)] MsgCallback cb_fn);

    [DllImport (lib, EntryPoint = "Cb", CallingConvention = CallingConvention.StdCall, CharSet = CharSet.Ansi)]
    private static extern void cb();  

    public static void Main(string[] args){
        System.Threading.Thread newThread = new System.Threading.Thread(recv);
        newThread.Start();
        set_callback(MsgCallback_impl);
        Console.ReadKey();
        cb();
        Console.ReadKey();
        newThread.Abort();
    }

    public static void recv(){
        Srv();
        while(true){
            // Console.WriteLine("recv loop");
        }
    }

    public static void MsgCallback_impl(uint sn, string str){
		Console.WriteLine("ok, recv... {0}", str);
	}
};

/*
mcs main.cs
mono main.exe
*/