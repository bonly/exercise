#include <iostream>
#include <string.h>
bool check_asic(const char* str, const size_t len)
{
    const char *p = str;
    for (size_t i=0; i<len; ++i)
    {
        if(p[i]>127 || p[i]<0) return false;
    }
    return true;
}

int main(int argc, char* argv[])
{
    if (check_asic(argv[1], strlen(argv[1])))
        std::clog << "true" << std::endl;
    else
        std::clog << "false" << std::endl;
    return 0;
}
