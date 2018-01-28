using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using System.Runtime.InteropServices; //for DllImport
using System.Runtime.Serialization.Formatters.Binary;
using System.IO;

public class Main : MonoBehaviour {
	// [UnmanagedFunctionPointer(CallingConvention.Cdecl)]
	[UnmanagedFunctionPointer(CallingConvention.StdCall)]
	delegate void MsgCallback(string str);  // 定义供外部调用的回调函数
	[DllImport ("libf.so", EntryPoint = "SetCallBack", CallingConvention = CallingConvention.StdCall)]
	extern static System.IntPtr set_callback([MarshalAs(UnmanagedType.FunctionPtr)] MsgCallback cb_fn);

	// [DllImport ("libfc", EntryPoint = "gidToUid")] extern static System.IntPtr gidToUid(ulong ullGID);
	[DllImport ("libfc.so", EntryPoint = "PrintHello")] extern static System.IntPtr gidToUid();
	[DllImport ("libf.so", EntryPoint = "Foo")] extern static System.IntPtr Foo();
	[DllImport ("libf.so", EntryPoint = "GetString")] extern static System.IntPtr HfromG();
	[DllImport ("libf.so", EntryPoint = "MergeString", CharSet = CharSet.Ansi)] 
	extern static System.IntPtr merge([MarshalAs(UnmanagedType.LPStr)] string left, [MarshalAs(UnmanagedType.LPStr)] string right);

	MsgCallback callback;
	// Use this for initialization
	void Start () {
		Debug.Log("in start");
		callback = MsgCallback_impl;	
	}
	
	// Update is called once per frame
	void Update () {
		
	}

	/// <summary>
	/// OnGUI is called for rendering and handling GUI events.
	/// This function can be called multiple times per frame (one call per event).
	/// </summary>
	void OnGUI()
	{
		if (GUILayout.Button("Press me")){
			System.IntPtr intPtr = gidToUid();
			string uid = Marshal.PtrToStringAnsi(intPtr);
			Debug.Log(uid);
		}
		if (GUILayout.Button("other dll")){

			// System.IntPtr intPtr = Foo();
			// string uid = Marshal.PtrToStringAnsi(intPtr);
			// Debug.Log(uid);

			System.IntPtr intPtr = HfromG();
			string uid = Marshal.PtrToStringAnsi(intPtr);
			Debug.Log(uid);
		}
		if (GUILayout.Button("two")){
			Debug.Log("click");
			System.IntPtr intPtr = merge("one", "two");
			string uid = Marshal.PtrToStringAnsi(intPtr);
			Debug.Log(uid);
		}	
		if (GUILayout.Button("load json from file")){
			var fn = Application.dataPath + "/Resources/Test.json";
			// BinaryFormatter bf = new BinaryFormatter();
			if (!File.Exists(fn)){
				Debug.LogFormat("file not exists {0}", fn);
			}

			StreamReader str = new StreamReader(fn);

			if (str == null){
				Debug.Log("file context is null");
				return;
			}

			string json = str.ReadToEnd();
			if (json.Length > 0){
				var js = JsonUtility.FromJson<GameStatus>(json);
				Debug.LogFormat("json: {0}", js.statusList[0].name);
			}
		}	
		if (GUILayout.Button("run")){
			Debug.Log("run");
			// var human = GameObject.Find("WomanWarrior");
			// human.CrossFade("walk");
			// Debug.LogFormat("com: {0}", human);
		}	
		if (GUILayout.Button("set callback")){
			Debug.LogFormat("begin set callback: {0}", callback);
			string cb_str = Marshal.PtrToStringAnsi(set_callback(MsgCallback_impl));	
			Debug.LogFormat("end set callback {0}", cb_str);
			// callback("aaaa");
		}
	}

	public void MsgCallback_impl(string str){
		Debug.LogFormat("ok, recv... {0}", str);
	}
};

[System.Serializable]
public class GameStatus
{
    public string gameName;
    public string version;
    public bool isStereo;
    public bool isUseHardWare;
    public refencenes[] statusList;
}

[System.Serializable]
public class refencenes
{
    public string name;
    public int id;
}

/*
放到任意目录下，在unity中设置windows平台，
运行时寻找Asset\Mono目录的dll,实际上unity打包时所有dll汇总到Asset\Plugins目录，
因此需要DllImport时需要写../Plugins/abc.dll
*/