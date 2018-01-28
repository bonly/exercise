// using System.Collections;
// using System.Collections.Generic;
// using UnityEngine;
using System;

public enum CFS_KIND {
    HIT,
    BEHIT,
    CAST
};

public  interface IStateTrigger {
    bool check(ref triggerData rule);
    void mod_data(int param);
};

public class BeHit : IStateTrigger {
    private int _beHitTimes;
    public void mod_data(int param) { _beHitTimes += param; }
    public bool check(ref triggerData rule) {
        return (_beHitTimes >= rule.BeHitTimes);
    }
};

public class Hit : IStateTrigger {
    private int _HitTimes;
    public void mod_data(int param) { _HitTimes += param; }
    public bool check(ref triggerData rule) {
        return (_HitTimes >= rule.HitTimes);
    }
};

public class Cast : IStateTrigger {
    private int _CastTimes;
    public void mod_data(int param) { _CastTimes += param; }
    public bool check(ref triggerData rule) {
        return (_CastTimes >= rule.CastTimes);
    }
};

public class Trigger {
    public triggerData rule;
	private IStateTrigger _obj;    
    public bool alive = true;

    // 创建触发器
    public Trigger(CFS_KIND kind, triggerData dt){
        rule = dt;
        switch (kind) {
            case CFS_KIND.HIT: {
                _obj = new Hit();
                break;
            }
            case CFS_KIND.BEHIT: {
                _obj = new BeHit();
                break;
            }
            case CFS_KIND.CAST: {
                _obj = new Cast();
                break;
            }
            default: {
                // Debug.Log("未知的触发器类型");
                Console.WriteLine("未知的触发器类型");
                break;
            }
        }
    }
	/* 
	接受事件
	param 接收事件参数
	*/
	public bool Update_State(int param){
        if (_obj == null){
            // Debug.Log("未正确设置状态类型，不能更新状态");
            Console.WriteLine("未正确设置状态类型，不能更新状态");
            return false;
        }
        _obj.mod_data(param); // 修改值
        if (_obj.check(ref rule) && execute()){  //检查为真时，执行触发事件
            alive = false;  // 执行完后标识为可删除
        }
        return true;
    }

	// 执行触发事件	
	private bool execute(){
        if (_obj == null) {
            // Debug.Log("未正确设置状态类型，不能输出技能");
            Console.WriteLine("未正确设置状态类型，不能输出技能");
            return false;
        }
        // 通知Excecutor执行skillId 
        Console.WriteLine("可施放技能");

        return true; 
    }
}
