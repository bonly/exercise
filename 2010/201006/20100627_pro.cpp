/**
 * Protocal.cpp
 * @file
 * @brief 传协议的封装及解释实现
 * @author bonly
 *
 * @date 2011-5-4 bonly created
 * @date 2011-8-16 bonly 增加路由组状态查询及路由状态变更接口协议
 */

#include "20100627_pro.h"
#include "20100627_enum.hpp"

extern unsigned long Message_get_delimited_size(void *_buffer, int offset);
extern int Message_can_read_delimited_from(void *_buffer, int offset, int length);

unsigned long (*cmd_size)(void*, int)= Message_get_delimited_size;
int (*cmd_can_read)(void *, int, int) = Message_can_read_delimited_from;

//*//test main
#include <string.h>
#include <iostream>
using namespace std;
int main()
{
    {   /// 测试一个文件中的协议
        char buf[1024] = "";

        Protocol p(buf, 1024);
        CmdHeartbeat tk;
        tk._stat = _good;

        p << 0 << tk;
        clog << "len is: " << p.Encode() << endl;

        p.init(buf, 1024);
        CmdHeartbeat rk;
        clog << "len is: " << p.Decode() << endl;
        p >> &rk;
        clog << "msg: " << hex << (int)rk._stat << endl;
    }


    return 0;
}
// */
