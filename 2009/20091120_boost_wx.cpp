//============================================================================
// Name        : py.cpp
// Author      : bonly
// Version     :
// Copyright   : bonly's copyright notice
// Description : Hello World in C++, Ansi-style
//============================================================================

#include <iostream>
using namespace std;
#include <boost/python.hpp>
using namespace boost::python;

int main(int argc, char *argv[])
{
  Py_Initialize();

  object main_module = import("__main__");
  object main_namespace = main_module.attr("__dict__");
  exec("hello = file('hello.txt', 'w')\n"
       "hello.write('Hello world!')\n"
       "hello.close()",
       main_namespace);

  exec("result = 5 ** 2", main_namespace);
  int five_squared = extract<int>(main_namespace["result"]);
  cout << "The five_squeared caculated by python is " << five_squared << endl;

   // Load the sys module.
  object sys = import("sys");

  // Extract the python version.
  std::string version = extract<std::string>(sys.attr("version"));
  std::cout << version << std::endl;


 //要求simple.py与可执行文件在相同路径下! 运行ok
  object simple = exec_file("simple.py",main_namespace, main_namespace);
  //dict global;
  //object result = exec_file("simple.py", global, global);
  object foo = main_namespace["foo"];
  int val = extract<int>(foo(5));
  cout << "Python has caculated foo as " << val << endl;

  object wx = import("wx");
  object app = wx.attr("App")();
  object frame = wx.attr("Frame")(object());
  object sh = frame.attr("Show")();
  object lop = app.attr("MainLoop")();
  return 0;
}

/*simple.py:
 def foo(i = 4):
  return i + 2008
  */


