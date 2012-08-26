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
  try
  {
    Py_Initialize();
    object main_module = import("__main__");
    object main_namespace = main_module.attr("__dict__");
    //dict global;
    //exec_file("simple.py",global,global);
    exec_file("simple.py", main_namespace, main_namespace);

    object wx = import("wx");
    object app = wx.attr("PySimpleApp")();
    object frame = main_module.attr("MouseEventFrame")(object(),object(-1));
    frame.attr("Show")();
    app.attr("MainLoop")();
  }
  catch (error_already_set)
  {
    PyErr_Print();
  }
  return 0;
}

/**
 #!/usr/bin/python
 #-*-coding:utf-8-*-
 import wx
 class MouseEventFrame(wx.Frame):
 def __init__(self, parent, id):
 wx.Frame.__init__(self, parent, id, 'Frame With Button',
 size=(300, 100))
 self.panel = wx.Panel(self)
 self.button = wx.Button(self.panel,
 label="Not Over", pos=(100, 15))
 self.Bind(wx.EVT_BUTTON, self.OnButtonClick,
 self.button)    #1 绑定按钮事件
 self.button.Bind(wx.EVT_ENTER_WINDOW,
 self.OnEnterWindow)     #2 绑定鼠标位于其上事件
 self.button.Bind(wx.EVT_LEAVE_WINDOW,
 self.OnLeaveWindow)     #3 绑定鼠标离开事件

 def OnButtonClick(self, event):
 self.panel.SetBackgroundColour('Green')
 self.panel.Refresh()

 def OnEnterWindow(self, event):
 self.button.SetLabel("Over Me!")
 event.Skip()

 def OnLeaveWindow(self, event):
 self.button.SetLabel("Not Over")
 event.Skip()

 /*下面这部分没有写,在程序中以代码实现
 if __name__ == '__main__':
 app = wx.PySimpleApp()
 frame = MouseEventFrame(parent=None, id=-1)
 frame.Show()
 app.MainLoop()
 */

 */
