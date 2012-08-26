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

///定义Mod,必须写在Py_Initialize前导出
BOOST_PYTHON_MODULE(CppMod) {
  class_<CppClass>("CppClass")
    .def("getNum",&CppClass::getNum);
}


int main(int argc, char ** argv)
{


  try
  {

    Py_Initialize();

    object main_module((handle<> (borrowed(PyImport_AddModule("__main__")))));

    object main_namespace = main_module.attr("__dict__");

    /**initCppMod is a function created by the BOOST_PYTHON_MODULE macro
     * which is used to initialize the Python module CppMod.
     * At this point, your embedded python script may call import CppMod
     * and then access CppClass as a member of the module.
     */
    PyImport_AppendInittab( "CppMod", &initCppMod ); ///导出Mod,可在py中直接import了

    ///取消定义导出到py的类,由前面定义为Mod,避免全部定义都是全局的
//    main_namespace["CppClass"] = class_<CppClass>("CppClass")
//                                   .def("getNum",&CppClass::getNum);

    ///写脚本调用刚导出的py类,此处要加上模块名CppMod,并且py会为我们管理内存问题
    handle<> ignored1(( PyRun_String("import CppMod\n"
                                     "cppa = CppMod.CppClass()\n"
                                     "print cppa.getNum()\n",
                                         Py_file_input,
                                         main_namespace.ptr(),
                                         main_namespace.ptr() ) ));

    ///It would be convenient,
    ///however, if the module was already imported,
    ///so that the programmer does not have to manually load the module at the beginning of each script.
    ///To do this,add the following after the main_namespace object has been initialized
    object cpp_module( (handle<>(PyImport_ImportModule("CppMod"))) );
    main_namespace["CppMod"] = cpp_module;

    ///写脚本调用刚导出的py类,此处要加上模块名CppMod,并且py会为我们管理内存问题
    handle<> ignored(( PyRun_String( "cpp = CppMod.CppClass()\n"
                                     "print cpp.getNum()\n",
                                         Py_file_input,
                                         main_namespace.ptr(),
                                         main_namespace.ptr() ) ));

    ///使py中能调用C++中已定义的实例
    //CppClass cpp;
    //main_namespace["cppa"] = ptr(&cpp);
    //exec("print cppa.getNum()\n", main_namespace, main_namespace); ///exec参数不用ptr

    ///使py中能调用C++中已定义的实例,重新把cpp指向到c++中的类,需要自己注意内存管理
    CppClass cpp;
    scope(cpp_module).attr("cpp")  = ptr(&cpp);
    exec("print cpp.getNum()\n", main_namespace, main_namespace); ///exec参数不用ptr
  }
  catch (error_already_set )
  {
    PyErr_Print();
  }
}
