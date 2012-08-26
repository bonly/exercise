/**
 * @file 20100303_strtol.cpp
 * @brief 测试strtol的用法
 *
 * @author bonly
 * @date 2011-10-11 bonly created
 */

#include <cstdlib>
#include <cstring>
#include <iostream>

#define AtoI(X,R) {\
    if(X==NULL) R=0; \
    else R=strtol(X,(char**)NULL,0);}

int main(int argc, char* argv[])
{
    if (argv[1] == NULL
    )
        std::cerr << "argv[1] is null" << std::endl;

    const char* null_val = "";
    std::cerr << "when val=\"\"  is: " << strtol(null_val, (char**) NULL, 0)
                << std::endl;

    long int aval = -1;
    AtoI(argv[1], aval);
    std::cerr << "frist conv val: " << aval << std::endl;

    try ///try机制对strtol无效,当输入为NULL时无法捕捉,函数只设置errno
    {
        long int val = strtol(argv[2], (char**) NULL, 0);
        std::cerr << "second val: " << val << std::endl;
    }
    catch (...)
    {
        std::cerr << "input nothing\n";
    }

    {
        char *null_point=0;//当输入为NULL指针时core
        long int val = strtol(null_point, (char**) NULL, 0);
        std::cerr << "null point: " << val << std::endl;
    }

    return 0;
}

