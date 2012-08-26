#include <Python.h>

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

int main(int argc, char* argv[])
{
  txt_test(argc,argv);

  return 0;
}
