using System;
using System.Runtime.InteropServices; // for DllImport
using System.Threading;

public class Net {
    [DllImport ("libtechappen.so", EntryPoint = "Run", CallingConvention = CallingConvention.StdCall, CharSet = CharSet.Ansi)]
    private static extern void Srv();
    
    [DllImport ("libtechappen.so", EntryPoint = "Proc", CallingConvention = CallingConvention.StdCall, CharSet = CharSet.Ansi)]
    private static extern System.IntPtr Proc([MarshalAs(UnmanagedType.LPStr)] string json);

    public static void Main(string[] args) {
        // String user = "you";
        string str;
        Srv();
        // for (int idx=0; idx<10; ++idx){
            str = Marshal.PtrToStringAnsi( //登录
                Proc("{\"Key\":{\"Scope\":\"Q_User\",\"Argv\":\"Q_User\",\"Func\":\"Login\"},\"Data\":{\"Name\":\"boy\",\"Passwd\":\"pwd\"},\"Rep\":\"abc\"}")
            );
            Console.WriteLine("recv login: " + str);

            str = Marshal.PtrToStringAnsi( //注册
                Proc("{\"Key\":{\"Scope\":\"Q_User\",\"Argv\":\"Q_User\",\"Func\":\"Register\"},\"Data\":{\"Name\":\"boy\",\"Passwd\":\"pwd\"},\"Rep\":\"abc\"}")
            );     
            Console.WriteLine("recv register: " + str);       
        // }
        // Console.WriteLine("Press any key to exit...");
        // Console.ReadKey();
    }

};

