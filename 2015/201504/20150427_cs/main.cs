using System;
using System.Runtime.InteropServices; //for DllImport
using System.Threading;
using System.IO;

static class API{
    public const string lib = "./libtechappen.so";

    [DllImport (lib, EntryPoint = "Init", CallingConvention = CallingConvention.StdCall, CharSet = CharSet.Ansi)]
    public static extern int Init();

    [DllImport (lib, EntryPoint = "Run", CallingConvention = CallingConvention.StdCall, CharSet = CharSet.Ansi)]
    public static extern int Run();

    [DllImport (lib, EntryPoint = "Stop", CallingConvention = CallingConvention.StdCall, CharSet = CharSet.Ansi)]
    public static extern int Stop();

    [DllImport (lib, EntryPoint = "Put_pack", CallingConvention = CallingConvention.StdCall, CharSet = CharSet.Ansi)]
    public static extern int Put_pack();
};

public class Net{
    [DllImport (API.lib, EntryPoint = "SetCallBack", CallingConvention = CallingConvention.StdCall)]
    extern static System.IntPtr set_callback([MarshalAs(UnmanagedType.FunctionPtr)] MsgCallback cb_fn);

    [UnmanagedFunctionPointer(CallingConvention.StdCall)]
    delegate void MsgCallback(uint sn, string str);  // 定义供外部调用的回调函数

    public static void Main(string[] args){
        Srv server = new Srv();

        //信号处理
        Console.CancelKeyPress += delegate(object sender, ConsoleCancelEventArgs e) {
            Console.WriteLine("get sign ctrl+c ...");
            return;
        };

        if (server.Thr() == null){
            Console.WriteLine("init srv failed");
            return;
        }

        set_callback(MsgCallback_impl);

        Console.ReadKey();

        API.Put_pack();
        // Console.ReadKey();
    }

	public static void MsgCallback_impl(uint sn, string str){
		Console.WriteLine("ok, recv... {0}", str);
	}    

};


public class Srv{
    Thread newThread = null;

    public Thread Thr(){
        return newThread;
    }

    public Srv(){
        Console.WriteLine("init thread");
        if (0 != API.Init()) {
            return;
        }
        newThread = new Thread(new ThreadStart(srv));
        newThread.Start();
    }

    public static void srv(){
        Console.WriteLine("wait result packet");
        API.Run(); //此处不会返回，是跟随线程的
        Console.WriteLine("end packet recv");
    }

    ~Srv(){
        API.Stop();
        if (newThread != null){
            newThread.Abort();
        }
    }

};

/*
mcs main.cs
mono main.exe
*/