using System;
using System.Collections.Generic;
using System.Text;
using System.Net;
using System.Net.Sockets;

class Program{
    static void Main(string[] args){
        byte[] data = new byte[1024];
        
        IPEndPoint laddr = new IPEndPoint(IPAddress.Parse("127.0.0.1"), 9998);
        Socket sk = new Socket(AddressFamily.InterNetwork, SocketType.Dgram, ProtocolType.Udp);

        sk.Bind(laddr);

        // IPEndPoint raddr = new IPEndPoint(IPAddress.Any, 0);
        IPEndPoint raddr = new IPEndPoint(IPAddress.Parse("127.0.0.1"), 9999);
        EndPoint remote = (EndPoint) (raddr);

        int recv = sk.ReceiveFrom(data, ref remote);

        Console.WriteLine("recv {0}: {1}",  remote.ToString(), Encoding.ASCII.GetString(data, 0, recv));

        string back = "i am ok";
        sk.SendTo(Encoding.ASCII.GetBytes(back), back.Length, SocketFlags.None, remote);
    }

};