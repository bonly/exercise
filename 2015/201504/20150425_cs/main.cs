using System;
using System.Runtime.InteropServices; //for DllImport
using System.Threading;
using System.IO;

public class Net{
    const string lib = "./libtechappen.so";

    [DllImport (lib, EntryPoint = "Init", CallingConvention = CallingConvention.StdCall, CharSet = CharSet.Ansi)]
    private static extern int Init();

    [DllImport (lib, EntryPoint = "Run", CallingConvention = CallingConvention.StdCall, CharSet = CharSet.Ansi)]
    private static extern int Run();

    [DllImport (lib, EntryPoint = "Put_pack", CallingConvention = CallingConvention.StdCall, CharSet = CharSet.Ansi)]
    private static extern int Put_pack();

    [UnmanagedFunctionPointer(CallingConvention.StdCall)]
    delegate void MsgCallback(uint sn, string str);  // 定义供外部调用的回调函数

    [DllImport (lib, EntryPoint = "SetCallBack", CallingConvention = CallingConvention.StdCall)]
    extern static System.IntPtr set_callback([MarshalAs(UnmanagedType.FunctionPtr)] MsgCallback cb_fn);

    public static void Main(string[] args){
        //信号处理
        Console.CancelKeyPress += delegate(object sender, ConsoleCancelEventArgs e) {
            Console.WriteLine("recv ctrl+c ...");
        };
        Init();

         set_callback(MsgCallback_impl);

        System.Threading.Thread newThread = new System.Threading.Thread(srv);
        newThread.Start();
        Console.ReadKey();

        Put_pack();
        Console.ReadKey();
        newThread.Abort();
    }

    public static void srv(){
        Console.WriteLine("wait result packet");
        Run(); //此处不会返回，是跟随线程的
        Console.WriteLine("end packet recv");
    }

	public static void MsgCallback_impl(uint sn, string str){
		Console.WriteLine("ok, recv... {0}", str);
	}    

};

/*
mcs main.cs
mono main.exe
*/