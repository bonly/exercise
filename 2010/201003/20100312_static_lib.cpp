#include <iostream>
int foo()
{
    std::cerr << "this is in foo\n";
    return 0;
}


extern "C"{
class AC
{
    AC()
    {
        std::cerr << "this is in AC\n";
    }

};
}
