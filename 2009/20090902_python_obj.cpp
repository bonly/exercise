//============================================================================
// Name        : pyobj.cpp
// Author      : bonly
// Version     :
// Copyright   : bonly's copyright notice
// Description : Hello World in C++, Ansi-style
//============================================================================
#include <boost/python.hpp>
#include <iostream>
using namespace std;
using namespace boost;
using namespace boost::python;

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
  Py_Finalize();
}

/*
def f(x, y):
     if (y == 'foo'):
         x[3:7] = 'bar'
     else:
         x.items += y(3, x)
     return x
 */

object f(object x, object y)
{
  if (y == "foo")
    x.slice(3,7)="bar"; //python 中不成功
  else
    x.attr("items")+=y(3,x); //python中不成功
  return x;
}

void getf()
{
  Py_Initialize();
  //PySys_SetArgv(argc, argv);

  object y("foor");
  //str x="test"; 没有重载
  str x;
  string ret = extract<string>(f(x,y)); //结果为空
  //object obj = f(x,y);  string ret = extract<string>(obj);  //结果为空
  cout << "ret is: "<< ret << endl;
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
void pycla(int argc, char** argv)
{
  Py_Initialize();
  PySys_SetArgv(argc, argv); //可不写哦?!
  object wx = import("wx");

  //object fun = wx["App"]; 失败
  object fun = wx.attr("App");
  object app = extract<object>(fun());

  object None;
  /*
  object frame = extract<object>(wx["Frame"](None)); 失败
  object frame = extract<object>(wx.attr("Frame")(None)); 成功
  */
  object frame = wx.attr("Frame")(None);
  //frame["Show"]();失败
  frame.attr("Show")(true);

  //app["MainLoop"]();失败
  app.attr("MainLoop")();

  Py_Finalize();
}

void wxc(int argc,char* argv[])
{
  Py_Initialize();
  PySys_SetArgv(argc, argv);
  PyRun_SimpleString( "import wx" );
  PyRun_SimpleString( "class Frame(wx.Frame):pass");
  /*
  PyRun_SimpleString( "class App(wx.App):" );
  PyRun_SimpleString( "def OnInit(self):" );
  PyRun_SimpleString( " self.frame=Frame(parent=None,title='Spare')" );
  PyRun_SimpleString( "app=App()" );
  PyRun_SimpleString( "app.MainLoop()" ); //失败*/

  PyRun_SimpleString(
        "class App(wx.App):\n"
         "\tdef OnInit(self):\n"
          "\t\tself.frame=Frame(parent=None,title='spare')\n "
          "\t\tself.frame.Show()\n"
          "\t\treturn True\n"
           "\t\tself.SetTopWindow(self.frame)"
  );
  PyRun_SimpleString( "app=App()" );
  PyRun_SimpleString( "app.MainLoop()" );
  Py_Finalize();
}
int main(int argc, char** argv)
{
  //pyobj();
  //wx(argc,argv);
	//pycla(argc,argv);
  //getf();
  wxc(argc,argv);
	return 0;
}

//g++ -I/usr/include/python2.5 -O0 -g3 -Wall -l python -l boost_python -fmessage-length=0 -MMD -MP -MF"src/pyobj.d" -MT"src/pyobj.d" -o"src/pyobj.o" "../src/pyobj.cpp"
