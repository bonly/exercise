using System;
using System.Threading;
using System.Net.Sockets;
using System.Net;

static class Program
{
    private static Socket icmpSocket;
    private static byte[] receiveBuffer = new byte[256];
    private static EndPoint remoteEndPoint = new IPEndPoint(IPAddress.Any, 0);
 
    /// <summary>
    /// The main entry point for the application.
    /// </summary>
    static void Main()
    {
        CreateIcmpSocket();
        while (true) { Thread.Sleep(10); }
    }
 
    private static void CreateIcmpSocket()
    {
        icmpSocket = new Socket(AddressFamily.InterNetwork, SocketType.Raw, ProtocolType.Icmp);
        //icmpSocket.Bind(new IPEndPoint(IPAddress.Any, 0));
        //icmpSocket.Bind(new IPEndPoint(IPAddress.Loopback, 9999));
        icmpSocket.Bind(new IPEndPoint(IPAddress.Parse("127.0.0.1"), 9999));
        // Uncomment to receive all ICMP message (including destination unreachable).
        // Requires that the socket is bound to a particular interface. With mono,
        // fails on any OS but Windows.
        //if (Environment.OSVersion.Platform == PlatformID.Win32NT)
        //{
        //    icmpSocket.IOControl(IOControlCode.ReceiveAll, new byte[] { 1, 0, 0, 0 }, new byte[] { 1, 0, 0, 0 });
        //}
        BeginReceiveFrom();
    }
 
    private static void BeginReceiveFrom()
    {
        icmpSocket.BeginReceiveFrom(receiveBuffer, 0, receiveBuffer.Length, SocketFlags.None,
            ref remoteEndPoint, ReceiveCallback, null);
    }
 
    private static void ReceiveCallback(IAsyncResult ar)
    {
        int len = icmpSocket.EndReceiveFrom(ar, ref remoteEndPoint);
        Console.WriteLine(string.Format("{0} Received {1} bytes from {2}",
            DateTime.Now, len, remoteEndPoint));
        LogIcmp(receiveBuffer, len);
        BeginReceiveFrom();
    }
 
    private static void LogIcmp(byte[] buffer, int length)
    {
        for (int i = 0; i < length; i++)
        {
            Console.Write(String.Format("{0:X2} ", buffer[i]));
        }
        Console.WriteLine("");
    }
}
