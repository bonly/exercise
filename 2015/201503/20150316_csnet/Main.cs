using System;
using System.Runtime.InteropServices; // for DllImport
using System.Threading;

public class Net {
    static Net nt;
    static MsgCallback callback; // 回调

    [DllImport("libtechappen.so", EntryPoint = "Srv", CallingConvention = CallingConvention.StdCall)]
    extern static System.IntPtr srv();

    [DllImport("libtechappen.so", EntryPoint = "Send", CallingConvention = CallingConvention.StdCall)]
    extern static System.IntPtr send();

    [UnmanagedFunctionPointer(CallingConvention.StdCall)]
    delegate void MsgCallback(string str);  // 定义供外部调用的回调函数

    [DllImport ("libtechappen.so", EntryPoint = "SetCallBack", CallingConvention = CallingConvention.StdCall)]
    extern static System.IntPtr set_callback([MarshalAs(UnmanagedType.FunctionPtr)] MsgCallback cb_fn);

    static void Main(string[] args){
        Console.WriteLine("begin...");
        
        srv();
        nt = new Net();

        callback = nt.MsgCallback_impl;

        Thread.Sleep(1000);
        nt.Loop();

        Console.WriteLine("end.");
    }

    public void MsgCallback_impl(string str){
		// Debug.LogFormat("ok, recv... {0}", str);
    }

    void Loop(){
        for (;;){
            // Console.WriteLine("running...");
            send();
        }
    }
};
