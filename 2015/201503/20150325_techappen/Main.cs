using System;
using System.Runtime.InteropServices; // for DllImport
using System.Threading;

public class Net {
    // [DllImport ("libtechappen.so", EntryPoint = "Run", CallingConvention = CallingConvention.StdCall, CharSet = CharSet.Ansi)]
    [DllImport ("client.dll", EntryPoint = "Run", CallingConvention = CallingConvention.StdCall, CharSet = CharSet.Ansi)]
    private static extern void Srv();
    
    // [DllImport ("libtechappen.so", EntryPoint = "Proc", CallingConvention = CallingConvention.StdCall, CharSet = CharSet.Ansi)]
    [DllImport ("client.dll", EntryPoint = "Proc", CallingConvention = CallingConvention.StdCall, CharSet = CharSet.Ansi)]
    private static extern int Proc([MarshalAs(UnmanagedType.LPStr)] string json);

    public static void Main(string[] args) {
        // String user = "you";

        Srv();
        for (int idx=0; idx<10; ++idx){
            Proc("{\"Key\":{\"Sc\":\"Q_User\",\"Func\":\"Login\"},\"Data\":{\"Name\":\"boy\",\"Passwd\":\"pwd\"},\"Rep\":\"abc\"}");
        }
        // Console.WriteLine("Login for User: " + user);
        // Console.WriteLine("Press any key to exit...");
        // Console.ReadKey();
    }

};

