#include <boost/python.hpp>

using namespace boost::python;

int add_five(int x) {
  return x + 5;
}

BOOST_PYTHON_MODULE(Pointless)
{
    def("add_five", add_five);
}

int main(int, char **) {

  Py_Initialize(); 
  
  try {
    //BOOST_PYTHON_MODULE宏中已定义了此函数
    initPointless(); // initialize Pointless
  
    PyRun_SimpleString("import Pointless");
    PyRun_SimpleString("print Pointless.add_five(4)");
  } catch (error_already_set) {
    PyErr_Print();
  }
  
  Py_Finalize();
  return 0;
}

// g++ 20090915_pymod.cpp -l python2.6 -l boost_python -I /usr/include/python2.6 -L /usr/lib
