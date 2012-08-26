/**
 * @file 20100226_switch.cpp
 * @brief
 *
 * @author bonly
 * @date 2011-9-11 bonly created
 */
#include <cstdio>

class A
{
    public:
        operator int()
        {
            return kint;
        }

        int kint;


};

int main()
{
    A a;
    a.kint = 2;
    switch(a)
    {
        case 1:
            printf ("values is 1: %d\n", a.kint);
            break;
        case 2:
            printf("values is 2: %d\n", a.kint);
            break;
        default:
            printf("values unknow\n");
            break;
    }
    return 0;
}


