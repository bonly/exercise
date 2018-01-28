using System;
using System.Runtime.InteropServices; // for DllImport
using System.Threading;
using Proto;

public class Net {
    public static void Main(string[] args) {
        Channel channel = new Channel("192.168.1.104:50051", ChannelCredentials.Insecure);

        var client = new Proto.User.UserClient(channel);
        String user = "you";

        var reply = client.Login(new ReqLogin { Name = user });
        Console.WriteLine("Login: " + reply.Msg);
        
        // var secondReply = client.SayHelloAgain(new HelloRequest { Name = user });
        // Console.WriteLine("Greeting: " + secondReply.Message);

        channel.ShutdownAsync().Wait();
        Console.WriteLine("Press any key to exit...");
        Console.ReadKey();
    }

};

