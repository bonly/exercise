/**
 * @file 20100230_ulimit.cpp
 * @brief 与环境命令ulimit -c unlimited效果一样
 *
 * @author bonly
 * @date 2011-10-9 bonly created
 */
#include <sys/time.h>
#include <sys/resource.h>

int main()
{
    //允许core文件
     struct rlimit fdlim;
     fdlim.rlim_cur = RLIM_INFINITY;
     fdlim.rlim_max = RLIM_INFINITY;
     setrlimit(RLIMIT_CORE, &fdlim);

     int a = 10/0;
     return 0;
}


