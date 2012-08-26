#include <boost/python.hpp>
#include <string>
using namespace boost::python;
using namespace std;

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

    //object main_module((handle<> (borrowed(PyImport_AddModule("__main__")))));
    object main_module = import ("__main__");

    object main_namespace = main_module.attr("__dict__");

    //PyImport_AppendInittab( (char*)"CppMod", &initCppMod ); ///导出Mod,可在py中直接import了

    exec("import sys \nname = raw_input('Enter login name: ')\n",main_namespace); ///要写上执行的空间
    string name = extract<string>(main_namespace["name"]);
    cout << "\nwelcome " << name << endl;

  }
  catch (error_already_set )
  {
    PyErr_Print();
  }
}
