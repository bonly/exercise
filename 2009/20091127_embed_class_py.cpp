#include <boost/python.hpp>

using namespace boost::python;

class CppClass
{
  public:
    int getNum()
    {
      return 7;
    }
};

int main(int argc, char ** argv)
{
  try
  {
    Py_Initialize();

    object main_module((handle<> (borrowed(PyImport_AddModule("__main__")))));

    object main_namespace = main_module.attr("__dict__");

    ///定义导出到py的类
    main_namespace["CppClass"] = class_<CppClass>("CppClass")
                                   .def("getNum",&CppClass::getNum);

    ///写脚本调用刚导出的py类
    handle<> ignored(( PyRun_String( "cpp = CppClass()\n"
                                     "print cpp.getNum()\n",
                                         Py_file_input,
                                         main_namespace.ptr(),
                                         main_namespace.ptr() ) ));

    ///使py中能调用C++中已定义的实例
    CppClass cpp;
    main_namespace["cppa"] = ptr(&cpp);///使用同一个实例,用ptr以防产生多个实例
    exec("print cppa.getNum()\n", main_namespace, main_namespace); ///exec参数不用ptr

  }
  catch (error_already_set )
  {
    PyErr_Print();
  }
}
