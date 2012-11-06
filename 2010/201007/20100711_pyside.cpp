#include <boost/python.hpp>
#include <iostream>
using namespace boost::python;

void hello()
{
    std::cout << "Hello from c" << std::endl;
}

BOOST_PYTHON_MODULE(emb)  //宏自动生成initemb()函数
{
    def("hello", hello);
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
