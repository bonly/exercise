using System;

public class Test_Trigger{
	static void Main(string[] args){
		Console.WriteLine("begin test");

		triggerData dt = new triggerData();
		dt.Id = "测试触发器";
		dt.Name = "测试打击";
		dt.StateId = 12;
		dt.OpType = 1;
		dt.HitTimes = 3;
		// dt.SkillId = 34;
		dt.Executor = 11;

		Trigger tr = new Trigger(CFS_KIND.HIT, dt);

		for (int idx=0; idx < 4; ++idx){
			tr.Update_State(1);
		}
	}
}