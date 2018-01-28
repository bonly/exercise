using System;
using System.Threading.Tasks;
using System.Threading;

public class Test{
    static void Main(string[] args)
    {
        Console.WriteLine("执行GetReturnResult方法前的时间：" + DateTime.Now.ToString("yyyy-MM-dd HH:mm:ss"));
        var strRes = Task.Run<string>(() => { return GetReturnResult("ok"); });
        Console.WriteLine("执行GetReturnResult方法后的时间：" + DateTime.Now.ToString("yyyy-MM-dd HH:mm:ss"));
        Console.WriteLine("我是主线程，线程ID：" + Thread.CurrentThread.ManagedThreadId);
        Console.WriteLine(strRes.Result);
        Console.WriteLine("得到结果后的时间：" + DateTime.Now.ToString("yyyy-MM-dd HH:mm:ss"));

        Console.ReadLine();
    }

    static string GetReturnResult(string ts)
    {
        Console.WriteLine("我是GetReturnResult里面的线程，线程ID：" + Thread.CurrentThread.ManagedThreadId);
        Thread.Sleep(2000);
        return "我是返回值" + ts;
    }

};
