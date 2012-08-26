#include <typeinfo>
#include <iostream>

template <typename T>
void fun()
{
        T res;
        typeid(T).name();
}

int main(int , char *[])
{
        fun<int>();
        return 0;
}
