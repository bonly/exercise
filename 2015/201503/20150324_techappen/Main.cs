using System;
using System.Runtime.InteropServices; // for DllImport
using System.Threading;

public class Net {
    [DllImport ("libtechappen.so", EntryPoint = "Run", CallingConvention = CallingConvention.StdCall, CharSet = CharSet.Ansi)]
    private static extern void Srv();
    
    [DllImport ("libtechappen.so", EntryPoint = "Proc", CallingConvention = CallingConvention.StdCall, CharSet = CharSet.Ansi)]
    private static extern int Proc([MarshalAs(UnmanagedType.LPStr)] string cmd, [MarshalAs(UnmanagedType.LPStr)] string data);

    public static void Main(string[] args) {
        String user = "you";

        Srv();
        // Send(user);
        Proc("Q_User", "{\"Req\":{\"Func\":\"Login\", \"Name\":\"boy\",\"Passwd\":\"pwd\"},\"Rep\":\"abc\"}");
        // Proc("Q_User", "{\"Func\":\"Login\", \"Name\":\"boy\",\"Passwd\":\"pwd\"}");

        Console.WriteLine("User: " + user);
        // Console.WriteLine("Press any key to exit...");
        // Console.ReadKey();
    }

};

