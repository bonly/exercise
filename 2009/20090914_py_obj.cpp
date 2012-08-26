#include <boost/python.hpp>
#include <iostream>

int main(int, char **) {
  using namespace boost::python;

  Py_Initialize();
  
  try {
    PyRun_SimpleString("result = 5 ** 2");
    
    object module(handle<>(borrowed(PyImport_AddModule("__main__"))));
    object dictionary = module.attr("__dict__");
    object result = dictionary["result"];
    int result_value = extract<int>(result);
    
    std::cout << result_value << std::endl;
    
    dictionary["result"] = 20;

    PyRun_SimpleString("print result");
  } catch (error_already_set) {
    PyErr_Print();
  }
  
  Py_Finalize();
  return 0;
}

//g++ 20090914_py_obj.cpp -l python2.6 -l boost_python -I /usr/include/python2.6 -L /usr/lib
