#define BOOST_PYTHON_SOURCE
#include <boost/python.hpp>

using namespace boost::python;

int main(int argc, char *argv[])
{
  try
  {
    Py_Initialize();
    //object main_mod = import ("__main__"); ///如果在python语句中使用了变量,就须要指定场景参数,则导入这个
    //object main_spa = main_mod.attr("__dict__");
    object subp = import ("subprocess");
    list param;
    param.append("ls");
    param.append("-l");
    subp.attr("call")(param);

    return 0;
  }
  catch(error_already_set)
  {
    PyErr_Print();
  }
}
