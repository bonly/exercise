using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using System.Text;
using System.Net;
using System.Net.Sockets;
using System;

using System.Threading;
using System.Runtime.InteropServices; // for DllImport
using System.ComponentModel;


public class sc : MonoBehaviour {
#if UNITY_EDITOR && UNITY_EDITOR_WIN
    public const string lib = "Windows/libtechappen.so";
#elif UNITY_EDITOR && UNITY_EDITOR_OSX
    public const string lib = "osx/libtechappen.so";
#elif UNITY_ANDROID 
    public const string lib = "techappen";
#elif UNITY_IOS || UNITY_IPHONE
    public const string lib = "__Internal";
#elif UNITY_EDITOR
    public const string lib = "Linux/libtechappen.so";
#else
    public const string lib = "Windows/libtechappen.so";
#endif

    // [DllImport(lib, EntryPoint = "Run", CallingConvention = CallingConvention.StdCall, CharSet = CharSet.Ansi)]
    // public static extern int Run();

	Socket conn;
	EndPoint local;
	// EndPoint remote;

	// AsyncCallback thr = null;
	Thread thr = null;

	// Use this for initialization
	void Start () {
		// IPEndPoint raddr = new IPEndPoint(IPAddress.Parse("127.0.0.1"), 9999);
		// remote = (EndPoint) (raddr);

		// IPEndPoint laddr = new IPEndPoint(IPAddress.Parse("127.0.0.1"), 9998);
		// local = (EndPoint)(laddr);

		// conn = new Socket(AddressFamily.InterNetwork, SocketType.Dgram, ProtocolType.Udp);
		// conn.Bind(local);

/* thread
	try{
		thr = new Thread(new ThreadStart(srv));
		thr.IsBackground = true;
		thr.Start();

	}catch(Exception ex){
		Debug.LogFormat("err: {0}\n", ex);
	}
//*/

/* APM
	// var func = new Func<string, string>(fc =>{
	// 	while(true){
	// 		Debug.LogFormat("in fn");
	// 		Run();
	// 		return fc;
	// 	}
	// });

	// var asyncResult = func.BeginInvoke("func", t =>
	// {
	// 	string str = func.EndInvoke(t);
	// 	Debug.Log(str);
	// }, null);		
*/

/*EAP
	BackgroundWorker worker = new BackgroundWorker();
	worker.DoWork += new DoWorkEventHandler((s1, s2) =>{
		Debug.LogFormat("in fn");
		Run();
	});

	worker.RunWorkerAsync(worker);
*/

	// Debug.Log("srv ok");

	}
	
	// Update is called once per frame
	void Update () {
		
	}

	~sc(){
		Debug.Log("~sc");
	}

	void OnGUI(){
		// if (GUI.Button(new Rect(150, 50, 100, 100), "send")){
		// 	string str="i send to u ";
		// 	conn.SendTo(Encoding.ASCII.GetBytes(str), str.Length, SocketFlags.None, remote);	
		// 	Debug.LogFormat("send: {0}", str);
		// }

		// if (GUI.Button(new Rect(150, 300, 100, 100), "recv")){
		// 	byte[] data = new byte[1024];
		// 	int recv = conn.ReceiveFrom(data, ref remote);
		// 	Debug.LogFormat("recv: {0}", Encoding.ASCII.GetString(data, 0, recv));			
		// }

		if(GUI.Button(new Rect(300, 50, 100, 100), "as send")){
			snd = new UdpClient();

			snd.Connect("127.0.0.1", 9999);
			byte[] sendByte = 	Encoding.ASCII.GetBytes("send from cs");

			snd.BeginSend(sendByte, sendByte.Length, new AsyncCallback(asSend), snd);		
		}

		if (GUI.Button(new Rect(300, 300, 100, 100), "as recv")){
			cli = new UdpClient(9998);
			remote = new IPEndPoint(IPAddress.Parse("127.0.0.1"), 9999);

			cli.BeginReceive(new AsyncCallback(asRecv), null);
		}
	}

	static UdpClient cli;
	static IPEndPoint remote;
	static void asRecv(IAsyncResult res){
		byte[] data = cli.EndReceive(res, ref remote);

		Debug.LogFormat("recv {0}: {1}", remote.Address.ToString(), Encoding.ASCII.GetString(data));

		cli.BeginReceive(new AsyncCallback(asRecv), null); //null 可以是自定义的对象数据

		//if res is end
		//需要cli.EndReceive(res, ref cli);  ???
	}

	static UdpClient snd;
	private void asSend(IAsyncResult res){
		Debug.LogFormat("send end");
		cli.EndSend(res);
	}
	// public static void srv(){
	// 	while(true){
	// 		Debug.Log("begin srv");
	// 		Run();
	// 		Debug.Log("end srv");
	// 	}
	// }
}
