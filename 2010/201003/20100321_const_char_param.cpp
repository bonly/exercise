#include <iostream>

int fun(const char* p=0)
{
    if (p!=0)
    std::cerr << p << std::endl;
    return 0;
}

int main()
{
    fun ("100");
    fun ();
    fun ("200");
    return 0;
}

