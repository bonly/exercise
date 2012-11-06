/*
 * 20100701_wxpy.cpp
 *
 *  Created on: 2012-10-10
 *      Author: bonly
 */
#include <boost/python.hpp>
#include <iostream>
using namespace boost::python;

char const* greet()
{
    return "hello, world";
}

BOOST_PYTHON_MODULE(hello_ext)
{
    def("greet", greet);
}

int main()
{
    try
    {
        Py_Initialize();
        PyImport_AppendInittab((char*)"hello_ext", &inithello_ext); ///< 导出到python中
        object cb = import("hello_ext"); ///< 从python实例到c中
    }
    catch (error_already_set)
    {
        PyErr_Print();
    }
    return 0;

}
