using System;
using System.Runtime.InteropServices; // for DllImport
using System.Threading;

public class Net {
	[UnmanagedFunctionPointer(CallingConvention.StdCall)]
	delegate void MsgCallback(uint sn, string str);  // 定义供外部调用的回调函数

	[DllImport ("libtechappen.so", EntryPoint = "SetCallBack", CallingConvention = CallingConvention.StdCall)]
	extern static System.IntPtr set_callback([MarshalAs(UnmanagedType.FunctionPtr)] MsgCallback cb_fn);

    [DllImport ("libtechappen.so", EntryPoint = "Run", CallingConvention = CallingConvention.StdCall, CharSet = CharSet.Ansi)]
    private static extern void Srv();

    [DllImport ("libtechappen.so", EntryPoint = "Exit", CallingConvention = CallingConvention.StdCall, CharSet = CharSet.Ansi)]
    private static extern void End();  

    [DllImport ("libtechappen.so", EntryPoint = "Proc", CallingConvention = CallingConvention.StdCall, CharSet = CharSet.Ansi)]
    private static extern int Proc([MarshalAs(UnmanagedType.LPStr)] string json);

    public static void Main(string[] args) {
        // String user = "you";
        Srv();
        int ret = 0;

        // Net net = new Net();
        set_callback(MsgCallback_impl);

        for (int idx=0; idx<900000; ++idx){
        // for (int idx=0; idx<100000; ++idx){
        // while (true){
            ret = //登录
                //数据
                Proc("{\"Key\":{\"Scope\":\"Q_User\",\"Argv\":\"Q_User\",\"Func\":\"Login\"},\"Data\":{\"Sn\":0,\"Name\":\"1\",\"Passwd\":\"1\"}}");
            Console.WriteLine("数据 recv login " + ret);

            // ret = ( //登录
            //     //发出错误的数据包
            //     Proc("{\"Key\":{\"Argv\":\"Q_User\",\"Func\":\"Login\"},\"Data\":{\"Sn\":1,\"Name\":\"boy\",\"Passwd\":\"pwd\"},\"Rep\":\"abc\"}"));
            // Console.WriteLine("发送错误的数据包 recv login: " + ret);

            // ret = ( //登录
            //     //发出正确的数据包
            //     Proc("{\"Key\":{\"Scope\":\"Q_User\",\"Argv\":\"Q_User\",\"Func\":\"Login\"},\"Data\":{\"Sn\":2, \"Name\":\"boy\",\"Passwd\":\"pwd\"},\"Rep\":\"abc\"}"));
            // Console.WriteLine("发出正确的数据包 recv login: " + ret);

            // ret = ( //登录
            //     //发出正确的数据包
            //     Proc("{\"Key\":{\"Scope\":\"Q_User\",\"Argv\":\"Q_User\",\"Func\":\"Login\"},\"Data\":{\"Sn\":3, \"Name\":\"girl\",\"Passwd\":\"pwd\"},\"Rep\":\"abc\"}"));
            // Console.WriteLine("不存在的用户 recv login: " + ret);

            // // str = Marshal.PtrToStringAnsi( //注册
            // //     Proc("{\"Key\":{\"Scope\":\"Q_User\",\"Argv\":\"Q_User\",\"Func\":\"Register\"},\"Data\":{\"Sn\":4, \"Name\":\"boy\",\"Passwd\":\"pwd\"},\"Rep\":\"abc\"}"));     
            // // Console.WriteLine("recv register: " + str);   

            // ret = ( //取场景数据
            //     Proc("{\"Key\":{\"Scope\":\"Scene\",\"Argv\":\"Q_KV\",\"Func\":\"Get\"},\"Data\":{\"Sn\":5, \"Key\":\"0\"}}"));     
            // Console.WriteLine("取不存在的场景数据 scene: " + ret);    

            // ret = ( //存场景数据
            //     Proc("{\"Key\":{\"Scope\":\"Scene\",\"Argv\":\"Q_KV\",\"Func\":\"Set\"},\"Data\":{\"Sn\":6, \"Key\":\"1\",\"Data\":\"{}\"}}"));     
            // Console.WriteLine("存场景数据 scene: " + ret);   

            // ret = ( //取场景数据
            //     Proc("{\"Key\":{\"Scope\":\"Scene\",\"Argv\":\"Q_KV\",\"Func\":\"Get\"},\"Data\":{\"Sn\":7, \"Key\":\"1\"}}"));     
            // Console.WriteLine("取存在的场景数据 scene: " + ret);  

            // ret = ( //存卡牌数据
            //     Proc("{\"Key\":{\"Scope\":\"Cards\",\"Argv\":\"Q_Card\",\"Func\":\"Set\"},\"Data\":{\"Sn\":8, \"Key\":\"1\",\"Data\":\"{}\"}}"));     
            // Console.WriteLine("存卡牌数据 cards: " + ret);   

            // ret = ( //取卡牌数据
            //     Proc("{\"Key\":{\"Scope\":\"Cards\",\"Argv\":\"Q_Card\",\"Func\":\"Get\"},\"Data\":{\"Sn\":9, \"Key\":\"1\"}}"));     
            // Console.WriteLine("取卡牌数据 cards: " + ret);    

            // ret = ( //匹配数据
            //     Proc("{\"Key\":{\"Scope\":\"PK\",\"Argv\":\"Q_PK\",\"Func\":\"Get\"},\"Data\":{\"Sn\":10, \"Id\":\"1\"}}"));     
            // Console.WriteLine("匹配数据 match: " + ret);                                                                   
        // }
        // Console.WriteLine("Press any key to exit...");
        Console.ReadKey();
        End();

    }

	public static void MsgCallback_impl(uint sn, string str){
		Console.WriteLine("ok, recv... {0}", str);
	}

};

