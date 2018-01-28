using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using System;
using System.Reflection;

public class inject : MonoBehaviour {

	// Use this for initialization
	void Start () {
		
	}
	
	// Update is called once per frame
	void Update () {

	}

	/// <summary>
	/// OnGUI is called for rendering and handling GUI events.
	/// This function can be called multiple times per frame (one call per event).
	/// </summary>
	void OnGUI() {
		if (GUILayout.Button("inject")){
			//注入类类型
            InjectorFactory.Instance.Bind<IMyClass, MyClass2>();
            //进入框架生命周期
            var test = InjectorFactory.Instance.CreateInstance<Test>();
            //调用此类的方法GetName
            Debug.LogFormat("{0}", test.cls.GetName());
		}		
	}
};

public interface IMyClass{ // 接口，与实现类绑定
	string GetName();
};

[AttributeUsage(AttributeTargets.All, Inherited = false, AllowMultiple = true)]
sealed class InjectAttribute : Attribute{
	readonly string positionalString;
	public InjectAttribute(string injectType) {
		positionalString = injectType;
	}
	public InjectAttribute() { }
};

public class MyClass1 : IMyClass{  // 实现类
	public string GetName() {
		return "Keyle Is Inject MyClass1 ...";
	}
};
public class MyClass2 : IMyClass{  // 实现类
	public string GetName() {
		return "Keyle Is Inject MyClass2 ...";
	}
};

public class Test{  // 需要被注入的测试类
	[Inject]
	public IMyClass cls { get; set; }
};

public class InjectorFactory{
	private static InjectorFactory instance;
	public  static InjectorFactory Instance{
		get { return instance ?? (instance = new InjectorFactory()); }
	}
	private InjectorFactory() { BindCache = new Dictionary<Type, Type>(); }
	public Dictionary<Type, Type> BindCache { get; set; }
	public void Bind<T, V>(){
		if (!BindCache.ContainsKey(typeof(T))){
			BindCache.Add(typeof(T), typeof(V));
		}
		else{
			BindCache[typeof(T)] = typeof(V);
		}
	}

	public T CreateInstance<T>() where T : new(){
		var a = new T();

		//注入此类内部属性
		KeyValuePair<Type, PropertyInfo>[] pairs = new KeyValuePair<Type, PropertyInfo>[0];
		object[] names = new object[0];
		MemberInfo[] privateMembers = a.GetType().FindMembers(MemberTypes.Property,
												BindingFlags.FlattenHierarchy |
												BindingFlags.SetProperty |
												BindingFlags.Public |
												BindingFlags.NonPublic |
												BindingFlags.Instance,
												null, null);

		foreach (MemberInfo member in privateMembers){
			object[] injections = member.GetCustomAttributes(typeof(InjectAttribute), true);
			if (injections.Length > 0){
				PropertyInfo point = member as PropertyInfo;
				Type pointType = point.PropertyType;
				point.SetValue(a, Activator.CreateInstance(BindCache[pointType]),null);
			}
		}
		return a;
	}
};


