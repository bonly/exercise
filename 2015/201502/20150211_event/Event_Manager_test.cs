/*
auth: bonly
bref: 测试Event_Manager.cs
log:  2017.3.10
*/
using System;

namespace techappen{

public class My_Event : IEven{
};

public class My_Event_Process :IEven_Proc {
    public int Process(IEven even){
       Console.WriteLine("in my event process");
       return 0; 
    }
};

public class Test_Even {
    static void Main(string[] args){
        Console.WriteLine("in main");

        
        My_Event_Process prc = new My_Event_Process();
        Even_Manager.Instance.Bind<My_Event>(prc);

        My_Event even = new My_Event();
        Even_Manager.Instance.Dispatch(even);
    }
}

}
