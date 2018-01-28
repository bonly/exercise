using System.Collections;
using System.Collections.Generic;
// using UnityEngine;

public enum LOCK_TYPE//锁定类型
{
    LOCK,
    UNLOCK,
}

public enum CAMP//阵营类型
{
    ARMY,//敌方
    UNION,//友军
    ALL,//全部
    SELF,//自己
}

public enum SKILL_TYPE
{
    DOWNLANCH,//按下立刻释放
    RELEASE,//松手释放
    HANDLE,//蓄力
    ACTIVE,//激活
}

public enum TARGET_TYPE//目标类型
{
    HERO,
    BUILDING,
    TRAP,
    BATMAN,
}
[System.Serializable]
public class skillData{//主要技能结构
    public string Id;
    public string Name;
    public string Icon;
    public float CoolDown;
    public LOCK_TYPE LockType;
    public float LockDistance;
    public CAMP TargetCamp;
    public TARGET_TYPE LockTargetType;
    public List<string> LanchIDs;
    public string TouchUiID;
    public string TouchSoundID;
    public string TouchMeshID;
    public string Desc;
    public SKILL_TYPE SkillType;

    public float StartTime;
    public float StopTime;
    public float StartMoveTime;
    public float StopMoveTime;
    public float StartRotateTime;
    public float StopRotateTime;
    public float SkillcontrolTime;
}

public enum LANCH_POS_TYPE//发射器位置类型
{
    SELF,
    JOYSTICK,
}

public class lanchData//发射器主结构数据
{
    public string Id;
    public string Name;
    public LANCH_POS_TYPE LanchPos;
    // public Vector3 MixPos; // @todo 确认此字段的转换
    public float LanchAngle;
    public float DelayTime;
    public float LanchDelay;
    public int LanchNum;
    public int LanchTimes;
    public string ProjectileID;
}

public enum RANGE_TYPE//碰撞区域类型
{
    CIRCLE,
    RECT,
    MESH,
}

public class skillEvent//技能事件
{
    public float damage;
    public List<string> states;
}

public class projectileData//弹道主结构数据
{
    public string Id;
    public string Name;
    public List<string> ItemIDs;
    public float XSpeed;
    public float XAddSpeed;
    public float YSpeed;
    public float YAddSpeed;
    public float ZSpeed;
    public float ZAddSpeed;
    public bool IfCollider;
    public float MaxDistance;
    public float FlyTime;
    public bool IfDestroyAfterCollider;
    public RANGE_TYPE RangeType;
    public float MeshAngle;
    public float MaxRadiu;
    public skillEvent ColliderEvent;
    public skillEvent EndEvent;
    public string Desc;
}

public enum STATE_TYPE
{

}

public enum NESTIFICATION_TYPE//叠加类型
{
    REFRESH,
    DISREFRESH,
}

public enum STATE_LOGIC
{
    CHANGEPAR,
    CHANGESTATE,
    CHANGEPOS,
    CHANGESIZE,
    CHANGERES,
    MOVE,
    ALIVE,
    SKILL,
    ATTACK,
    IDLE,
}

public class stateData//状态主数据结构
{
    public string id;
    public string name;
    public STATE_TYPE stateType;
    public int priority;
    public float duration;
    public bool ifNestfication;
    public int nestNums;
    public STATE_LOGIC logicType;
    public string logicPar;
    public bool ifBack;
    public string triggerID;

    public List<string> lanchIDs;
    public List<string> states;
    public string item;
    public List<string> effects;
    public List<string> audios;
}

public enum SEEN_TYPE//可见类
{

}

public class itemData//单位主数据结构
{
    public string id;
    public string name;
    public string res;
    public string scale;
    public SEEN_TYPE seenType;
    public List<string> skills;
    public List<string> states;
    public string item;
    public string AI;
    // public Vector3 mixRoot;
    public float existTime;
}

public enum ActionType{
    Add,
    Del
}

public class triggerData//触发器主数据结构
{
    public string Id;
    public string Name;
    public int    StateId;    // 监听状态
    public int    BeHitTimes; // 受击次数
    public int    HitTimes;   //攻击次数
    public int    CastTimes;  //施法次数
    public int    OpType;     //动作（增加/删除）
    public int    Executor;   //执行者
    public int    ExecTimes;  //执行次数
    public float  ExecInterval; //执行间隔
    public string OutJudge;    //跳出条件
}