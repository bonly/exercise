using System;
using System.Runtime.InteropServices; // for DllImport
using System.Threading;

public class Net {
    static Net nt;

    [DllImport("libtechappen.so", EntryPoint = "Srv", CallingConvention = CallingConvention.StdCall)]
    extern static System.IntPtr srv();

    [DllImport("libtechappen.so", EntryPoint = "Send", CallingConvention = CallingConvention.StdCall)]
    extern static System.IntPtr send();

    static void Main(string[] args){
        Console.WriteLine("begin...");
        
        srv();
        nt = new Net();

        Thread.Sleep(1000);
        nt.Loop();

        Console.WriteLine("end.");
    }


    void Loop(){
        for (;;){
            // Console.WriteLine("running...");
            send();
        }
    }
};