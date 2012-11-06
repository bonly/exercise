#include "20100713_pyside.hpp"
#include <iostream>

void hello()
{
    std::cout << "Hello from c" << std::endl;
}

PY_MODULE_IMPLEMENT(emb)
{
    PY_DEF("hello", hello);
}

int main(int argc, char* argv[])
{
    using namespace boost::python;

    Py_Initialize();
    try
    {
        initemb();
        object main_module = import ("__main__");
        object main_namespace = main_module.attr("__dict__");


        exec_file ("20100710_pyside.py", main_namespace);
    }
    catch(...)
    {
        PyErr_Print();
    }
    return 0;
}
