//============================================================================
// Name        : wxhe.cpp
// Author      : bonly
// Version     :
// Copyright   : bonly's copyright notice
// Description : Hello World in C++, Ansi-style
//============================================================================

#include <iostream>
#include <boost/python.hpp>
using namespace std;
using namespace boost;

void show(int argc, char* argv[])
{
  Py_Initialize();
  PySys_SetArgv(argc, argv);
  PyRun_SimpleString( "import Tkinter" );
  PyRun_SimpleString( "root = Tkinter.Tk()" );
  PyRun_SimpleString( "root.mainloop()" );
  Py_Finalize();

}

void txt_test(int argc, char* argv[])
{
  Py_Initialize();
  PySys_SetArgv(argc, argv);
  PyRun_SimpleString( "from Tkinter import *" );
  PyRun_SimpleString( "root=Tk()" );
  PyRun_SimpleString( "w = Label(root,text = 'hello world')");
  PyRun_SimpleString( "w.pack()" );
  PyRun_SimpleString( "root.mainloop()" );

  Py_Finalize();
}

void wx(int argc,char* argv[])
{
  Py_Initialize();
  PySys_SetArgv(argc, argv);
  PyRun_SimpleString( "import wx" );
  PyRun_SimpleString( "app=wx.App()" );
  PyRun_SimpleString( "frame=wx.Frame(None)");
  PyRun_SimpleString( "frame.Show()" );
  PyRun_SimpleString( "app.MainLoop()" );

  Py_Finalize();
}
void script_file(int argc, char* argv[])
{
  Py_Initialize();
  PySys_SetArgv(argc, argv);
  FILE *fp = fopen("/home/bonly/src/20090721_wx_point.py", "r");
  PyRun_SimpleFile( fp, "/home/bonly/src/20090721_wx_point.py");

  Py_Finalize();
}

void botobj()
{
  Py_Initialize();
  boost::python::object main_module = boost::python::import("__main__");
  boost::python::object global = main_module.attr("__dict__");
  //boost::python::dict global = main_module["__dict__"];
  boost::python::object res = boost::python::exec_file(
      "/home/bonly/src/20090721_wx_point.py",
      global
      );
  Py_Finalize();
}

void execpy()
{
  Py_Initialize();
  //引入__main__作用域
  using namespace boost::python;
  object main_module = import("__main__");
  object main_namespace = main_module.attr("__dict__");
  exec ("print('Hello world')",main_namespace);
}

void pyobj()
{
  Py_Initialize();
  using namespace boost::python;
  object mm = import ("__main__");
  object mn = mm.attr("__dict__");
  object sm = exec_file("/home/bonly/src/20090722_simple.py",mn);
  object fo = mn["foo"];
  int val = extract<int>(fo(4));
  clog << "python foo return: " << val << endl;

}
int main(int argc, char* argv[])
{
  //show(argc,argv);
  //txt_test(argc,argv);
  //wx(argc,argv);
  //botobj();
  //script_file(argc,argv);
  //execpy();
  pyobj();
  return 0;
}
