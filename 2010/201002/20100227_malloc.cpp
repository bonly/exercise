/**
 * @file 20100227_malloc.cpp
 * @brief 测试数据指针的删除方法
 *
 * @author bonly
 * @date 2011-9-14 bonly created
 */
#include <cstdio>
#include <cstdlib>

int main()
{
    char *mc[3]={0};
    for (int i=0; i<3; ++i)
    {
        mc[i] = (char*)malloc(32);
        sprintf(mc[i],"this is %d\n",i);
    }
    for (int i=0; i<3; ++i)
    {
        printf("%d is: %s", i, mc[i]);
    }
    for (int i=0; i<3; ++i)
    {
        free(mc[i]);
    }
    return 0;
}



