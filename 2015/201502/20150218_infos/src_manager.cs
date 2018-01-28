using System.Collections;
using System.Collections.Generic;
using System;
// using UnityEngine;

// public enum INFO_TYPE
// {
//     SCRIPTS,
//     SKILLMESH,
//     UIINFO,
// }

//初始化代码捆绑_begin
[System.Serializable]
public struct scriptInfo : Info{
    public string itemName;
    public string script;
    public string tag;
    public string space;
}

// public class scriptInfos
// {
//     public List<scriptInfo> info;
// }
// //初始化代码捆绑_end

// //初始化技能指示器_begin
// [System.Serializable]
// public class meshPars
// {
//     public float radius;
//     public float angle;
//     public int segments;
//     public int angleDegreePrecision;
//     public int radiusPrecision;

//     public meshPars(float _r = 0.0f, float _a = 0.0f, int _s = 0, int _angleDegreePrecision = 0, int _radiusPrecision = 0)
//     {
//         radius = _r;
//         angle = _a;
//         segments = _s;
//         angleDegreePrecision = _angleDegreePrecision;
//         radiusPrecision = _radiusPrecision;
//     }
//     /// <summary>
//     /// 创建一个扇形Mesh
//     /// </summary>
//     /// <param name = "radius" > 扇形半价 </ param >
//     /// < param name="angleDegree">扇形角度</param>
//     /// <param name = "segments" > 扇形弧线分段数 </ param >
//     /// < param name="angleDegreePrecision">扇形角度精度（在满足精度范围内，认为是同个角度）</param>
//     /// <param name = "radiusPrecision" >
//     /// < pre >
//     /// 扇形半价精度（在满足半价精度范围内，被认为是同个半价）。
//     /// 比如：半价精度为1000，则：1.001和1.002不被认为是同个半径。因为放大1000倍之后不相等。
//     /// 如果半价精度设置为100，则1.001和1.002可认为是相等的。
//     /// </pre>
//     /// </param>
//     /// <returns></returns>
// }
// //初始化技能指示器_end

//UIinfos_begin
[System.Serializable]
public class uiInfos : Info {
    public string prefabName;
    public float posx;
    public float posy;

    public uiInfos(string prefab, float _posx, float _posy)
    {
        prefabName = prefab;
        posx = _posx;
        posy = _posy;
    }
}
// UIinfos_end

public interface Info {
}

public class loadInfos {
    public Dictionary<Type, Dictionary<string, Info>> allInfos;

    private static loadInfos instance;
    public  static loadInfos Instance{
        get {return instance ?? (instance = new loadInfos());}
    }

    public Info getInfo<DT>(string key){
        return allInfos[typeof(DT)][key];
    }

    public void putInfo<DT>(string key, Info data) {
        if (allInfos == null){
            allInfos = new Dictionary<Type, Dictionary<string, Info>>();
        }
        Dictionary<string, Info> dt = new Dictionary<string, Info>();
        dt[key] = data;

        allInfos[typeof(DT)] = dt;
    }

    // public static void loadDicJson<T>(INFO_TYPE key,ref T dic)
    // {
    //     var infos = getInfosByKey(key);
    //     if(infos != null)
    //     {
    //         dic = JsonUtility.FromJson<T>(getJsonString(infos.path));
    //     }
    // }

    // public static void loadDicJson<TKey, TVaule>(INFO_TYPE key,ref Dictionary<TKey, TVaule> dic)
    // {
    //     var infos = getInfosByKey(key);
    //     if(infos != null)
    //     {
    //         dic = JsonUtility.FromJson<Serialization<TKey, TVaule>>(getJsonString(infos.path)).ToDictionary();
    //     }
    // }

    // public static void getJsonInfoByKey<T>(INFO_TYPE t,ref T item,string key = null)
    // {
    //     if(key == null)
    //     {
    //         loadDicJson(t,ref item);
    //     }else
    //     {
    //         var tp = new Dictionary<string, T>();
    //         loadDicJson(t,ref tp);
    //         if(tp != null)
    //         {
    //             tp.TryGetValue(key, out item);
    //         }
    //     }
    // }

    //public static meshPars getSkillMeshByKey(string key)//获取对应的技能指示器
    //{
    //    var meshList = loadDicJson(INFO_TYPE.SKILLMESH, new Dictionary<string, meshPars>()) as Dictionary<string, meshPars>;
    //    meshPars tp = null;
    //    if (meshList != null)
    //    {
    //        meshList.TryGetValue(key, out tp);
    //    }
    //    return tp;
    //}

    //public static scriptInfos getEntryInfo()//获取入口捆绑函数集
    //{
    //    return loadDicJson(INFO_TYPE.SCRIPTS, new scriptInfos()) as scriptInfos;
    //}
}


public class TMain {
    public static void Main(string[] argc){
        // 脚本资源
        scriptInfo inf1 = new scriptInfo();
        inf1.itemName = "abc";

        loadInfos.Instance.putInfo<scriptInfo>("abc", inf1);
        loadInfos.Instance.getInfo<scriptInfo>("abc");


        // UI资源
        uiInfos uf1 = new uiInfos("pre1", 13.0f, 12.0f);
        uf1.prefabName = "pref1";

        loadInfos.Instance.putInfo<uiInfos>("pre1", uf1);
        uiInfos adb = loadInfos.Instance.getInfo<uiInfos>("pre1");        
    }
};
