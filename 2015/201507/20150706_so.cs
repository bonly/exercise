using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using System.Runtime.InteropServices; // for DllImport
using AOT;

using System;
using System.Threading;
using System.Net.Sockets;
using System.Net;
using System.Text;


public class so : MonoBehaviour {
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

    [DllImport(lib, EntryPoint = "Run", CallingConvention = CallingConvention.StdCall, CharSet = CharSet.Ansi)]
    public static extern void Run();

    // [DllImport(lib, EntryPoint = "SetSoCallCs", CallingConvention = CallingConvention.StdCall)]
    // extern static System.IntPtr SetSoCallCs([MarshalAs(UnmanagedType.FunctionPtr)] SoCallCsFun cb_fn);

    // [UnmanagedFunctionPointer(CallingConvention.StdCall)]
    // delegate void SoCallCsFun(int sn, string str);  // 定义供外部调用的回调函数

    [DllImport(lib, EntryPoint = "Stop", CallingConvention = CallingConvention.StdCall, CharSet = CharSet.Ansi)]
    public static extern int Stop();

	Thread thr = null;

	// Use this for initialization
	void Start () {
		// Debug.LogFormat("load library: {0}", lib);
		// SetSoCallCs(SoCallCs);
	}
	
	// Update is called once per frame
	void Update () {
	}

    // [MonoPInvokeCallback(typeof (SoCallCsFun))]
    // static void SoCallCs(int Sn, string str) {
    //     Debug.LogFormat("接收到So直接调用函数{0}  {1}", Sn, str);
    // }
	
	// ~so(){
	// 	Debug.LogFormat("~so");
	// }

	void OnGUI(){
		if (GUI.Button(new Rect(10,10, 100, 100), "init")){
			try{
				cli = new UdpClient(new IPEndPoint(IPAddress.Parse("127.0.0.1"), 9998));
				remote = new IPEndPoint(IPAddress.Parse("127.0.0.1"), 9999);

				// cli.BeginReceive(new AsyncCallback(asRecv), null);
			}catch(Exception e){
				Debug.LogFormat(e.ToString());
			}
		}

		if (GUI.Button(new Rect(10, 150, 100, 100), "start")){
			thr = new Thread(new ThreadStart(Run));
			thr.Start();
			// Run();
	        Debug.LogFormat("起动接收服务");
		}

		if (GUI.Button(new Rect(10, 300, 100, 100), "stop")){
			Stop();	
		}
	}	

	static UdpClient cli;
	static IPEndPoint remote;
	static byte[] data;
	static void asRecv(IAsyncResult res){
		Debug.LogFormat("in end recv");
		data = cli.EndReceive(res, ref remote);
		Debug.LogFormat("recv {0}: {1}", remote.Address.ToString(), Encoding.ASCII.GetString(data));
		cli.BeginReceive(new AsyncCallback(asRecv), null);
	}
}
