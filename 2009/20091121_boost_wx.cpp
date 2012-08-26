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

  object wx = import("wx");
  object app = wx.attr("App")();
  object frame = wx.attr("Frame")(object());
  object sh = frame.attr("Show")();
  object lop = app.attr("MainLoop")();
  return 0;
}



