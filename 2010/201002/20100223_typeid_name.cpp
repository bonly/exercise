#include <stdio.h>
#include <typeinfo>
#include <cxxabi.h>
int main()
{
    char * realname = abi::__cxa_demangle(typeid(int()).name(), 0, 0, 0);
    printf("%10s ==> %15s\n", typeid(int()).name(), realname);

    realname = abi::__cxa_demangle(typeid(int(*)()).name(), 0, 0, 0);
    printf("%10s ==> %15s\n", typeid(int(*)()).name(), realname);

    realname = abi::__cxa_demangle(typeid(int).name(), 0, 0, 0);
    printf("%10s ==> %15s\n", typeid(int).name(), realname);
    return 0;
}
