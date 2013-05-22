//============================================================================
// Name        : stream_log.cpp
// Author      : bonly
// Version     :
// Copyright   : bonly's copyright notice
// Description : Hello World in C++, Ansi-style
//============================================================================

#include <iostream>
using namespace std;

#define sw16(x) \
    ((short)( \
        (((short)(x) & (short)0x00ffU) << 8 ) | \
        (((short)(x) & (short)0xff00U) >> 8 ) ))
/*//大小端的转换
1) AAAAAAAABBBBBBBB => AAAAAAAA    BBBBBBBB
2) AAAAAAAA => 00000000AAAAAAAA
3) BBBBBBBB => BBBBBBBB00000000
4) BBBBBBBB00000000 + 00000000AAAAAAAA = BBBBBBBBAAAAAAAA
假设x=0xaabb
(short)(x) & (short)0x00ffU) 与将16位数高8位置0   成了0x00bb 然后<<8 向左移8位后 低8位变成了高8位 低8位补0  结果为 0xbb00
(((short)(x) & (short)0xff00U) >> 8 ) )) 恰好相反 得到的结果为 0x00aa
两个结果再 或一下 0xbb00 | 0x00aa 就成了 0xbbaa
 */
int main()
{
    for(int i=0; i<3; ++i)
    {
        struct timespec tm={0,100000*100000000};
        nanosleep(&tm,NULL);
        cout << ":bonly^_^" << endl; // prints :bonly^_^
    }
    cout << hex << 0xaabb << endl;
    cout << hex << sw16(0xaabb) << endl;
	return 0;
}

