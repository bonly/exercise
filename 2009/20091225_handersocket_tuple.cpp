//============================================================================
// Name        : testcase.cpp
// Author      : bonly
// Version     :
// Copyright   : bonly's copyright notice
// Description : Hello World in C++, Ansi-style
//============================================================================
#define BOOST_PYTHON_SOURCE
#include <boost/python.hpp>
using namespace boost::python;

#include <iostream>
using namespace std;

int main()
{
   try
   {
      Py_Initialize();
      object main_mod = import("__main__");
      object main_spa = main_mod.attr("__dict__");

      //exec("from pyhs.sockets import ReadSocket",main_spa);
      object pyhs = import("pyhs");
      object hs = pyhs.attr("Manager")();

      list val;
      val.append("ind");
      val.append("13");
      tuple tup=make_tuple(val);

      hs.attr("insert")("test","testd",tup);
      ///第三个param是索引名,可以在mysql中用show index from testd查看"key_name",
      ///默认为PRIMARY名字,是建主建时自动建的名字;
      ///python中代码为:hs.insert('test','testd',[('ind','12')],'PRIMARY')

   }
   catch (error_already_set)
   {
      PyErr_Print();
      PyErr_Clear();
   }
   return 0;
}

