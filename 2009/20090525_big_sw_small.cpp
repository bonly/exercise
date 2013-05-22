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
/*//��С�˵�ת��
1) AAAAAAAABBBBBBBB => AAAAAAAA    BBBBBBBB
2) AAAAAAAA => 00000000AAAAAAAA
3) BBBBBBBB => BBBBBBBB00000000
4) BBBBBBBB00000000 + 00000000AAAAAAAA = BBBBBBBBAAAAAAAA
����x=0xaabb
(short)(x) & (short)0x00ffU) �뽫16λ����8λ��0   ����0x00bb Ȼ��<<8 ������8λ�� ��8λ����˸�8λ ��8λ��0  ���Ϊ 0xbb00
(((short)(x) & (short)0xff00U) >> 8 ) )) ǡ���෴ �õ��Ľ��Ϊ 0x00aa
��������� ��һ�� 0xbb00 | 0x00aa �ͳ��� 0xbbaa
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

