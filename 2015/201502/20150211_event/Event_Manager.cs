/*
auth: bonly
bref: 事件处理
log:  2017.3.10
*/
using System;
using System.Collections.Generic;

namespace techappen{

public interface IEven{
};

public interface IEven_Proc{
	int Process(IEven even);
};

public class Even_Manager{
	public Dictionary<Type, List<IEven_Proc>> Map {get; set;}

	private static Even_Manager instance;
	public  static Even_Manager Instance{
		get { return instance ?? (instance = new Even_Manager()); }
	}

	// 绑定消息类型与处理对象
	public void Bind<T>(IEven_Proc prc){
		if (instance.Map == null){ // 创建存储空间
			instance.Map = new Dictionary<Type, List<IEven_Proc>>();
		}
		if (!Map.ContainsKey(typeof(T))){  // 没有此清单
			List<IEven_Proc> lst = new List<IEven_Proc>();
			lst.Add(prc);
			Map.Add(typeof(T), lst);
		}else{ // 已有此清单，增加
			Map[typeof(T)].Add(prc);
		}
	}

	// 分发消息，返回分发数
	public int Dispatch(IEven even){
		if (instance == null){
			return -1;  // 还没创建实例
		}
		int ret = 0;
		foreach (Type key in Map.Keys){
			 for (ret = 0; ret < Map[key].Count; ret++){ // 分发事件，并处理
			 	// todo: 用协程或线程去调用处理
				(Map[key])[ret].Process(even); 
			 }
		}
		return ret;
	}
};

}
