/**
 * Protocal.cpp
 * @file
 * @brief 传协议的封装及解释实现
 * @author bonly
 *
 * @date 2011-5-4 bonly created
 * @date 2011-8-16 bonly 增加路由组状态查询及路由状态变更接口协议
 */

#include "20100603_Protocal.h"
#include "20100603_chat.c"
#include "20100605_chat.c"
#include <string.h>
#include <iostream>
using namespace std;

namespace bas {
/*
 int str2l(const string &buf,int &res)
 {
 errno = 0;
 res = strtol(buf.c_str(),0,10);\
        if((errno==ERANGE && (res == LONG_MAX||res == LONG_MIN))
 ||(errno!=0 && res==0))
 {
 perror("数据不正确");
 return 1;
 }
 return 0;
 }
 */

}

//*//test main
using namespace bas;
using namespace std;
int main()
{
    {   /// 测试一个文件中的协议
        char buf[1024] = "";

        Protocol p(buf, 1024);
        CmdTalk tk;
        tk._from_id._key = 1001;
        tk._from_id._type = _PEOPLE_ID;
        tk._to_id._type = _PEOPLE_ID;
        tk._to_id._key = 1010;
        tk._say._msg_len = 6;
        strcpy(tk._say._msg, "hello");

        p << 3 << tk;
        clog << "len is: " << p.Encode() << endl;

        p.init(buf, 1024);
        CmdTalk rk;
        clog << "len is: " << p.Decode() << endl;
        p >> &rk;
        clog << "msg: " << rk._say._msg << endl;
    }

    { /// 测试第二个文件中的协议
        char buf[1024]="";
        Protocol p(buf, 1024);
        CmdAct ac;
        ac._from_id._type = _PEOPLE_ID;
        ac._from_id._key = 1023;
        ac._to_id._type = _GROUP_ID;
        ac._to_id._key = 122;
        ac._action = 1223;
        p << 4 << ac;
        clog << "len is: " << p.Encode() << endl;

        Pack *hd;
        hd = (Pack*)buf;
        clog << "cmd is: " << hd->cmd << endl;

        CmdAct ot;
        p.init(buf, 1024);
        clog << "len is: " << p.Decode() << endl;
        p >> ot;
        clog << "action: " << ot._action << endl;
    }

    return 0;
}
// */

