using System;
using System.Runtime.InteropServices; // for DllImport
using System.Threading;
using Grpc.Core;
using He;

public class Net {
    public static void Main(string[] args) {
        Channel channel = new Channel("192.168.1.104:50051", ChannelCredentials.Insecure);

        var client = new He.Greeter.GreeterClient(channel);
        String user = "you";

        var reply = client.SayHello(new HelloRequest { Name = user });
        Console.WriteLine("Greeting: " + reply.Message);
        
        // var secondReply = client.SayHelloAgain(new HelloRequest { Name = user });
        // Console.WriteLine("Greeting: " + secondReply.Message);

        channel.ShutdownAsync().Wait();
        Console.WriteLine("Press any key to exit...");
        Console.ReadKey();
    }

};

