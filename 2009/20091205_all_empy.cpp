#define BOOST_PYTHON_SOURCE
#include <boost/python.hpp>
#include <iostream>
using namespace boost::python;
using namespace std;

object main_spa;

class CppClass
{
  public:
    void ClickButton(object ob)
    {
      clog << "u click the button\n";
    }
};

void cb(object obj)  //event 参数用object 代替
{
  clog << "u cb function\n";
  exec("wxf.ok.SetLabel('clicked Me!')",main_spa);
  exec("wxf.panel.Refresh()",main_spa);
  //exec("event.Skip()",main_spa);
}

BOOST_PYTHON_MODULE(CB)
{
  def("cb", cb);
}

BOOST_PYTHON_MODULE(CppMod)
{
  class_<CppClass>("CppClass")
        .def("ClickButton", &CppClass::ClickButton);
}

int main(int argc, char *argv[])
{
  try
  {
    Py_Initialize();
    object main_mod = import ("__main__"); ///如果在python语句中使用了变量,就须要指定场景参数,则导入这个
    main_spa = main_mod.attr("__dict__");
    exec("import wx",main_spa);

    PyImport_AppendInittab((char*)"CB", &initCB);///初始化后可以用import导入,注意initXXX是boost生动成生的
    object cb = import("CB");   ///导入C,脚本中不能直接使用CB.fun(),需要在场景中再转化
    main_spa["CB"]=cb;///C中定义的对象转到py中
    //main_spa["CB"]=ptr(&cb);///如果py中已有CB对象(必须是已有),可以重新指到C中新生成对象
    //exec("import CB",main_spa); ///此方法导入,脚本中可以直接CB.fun()调用

    exec("global wxf",main_spa);

    PyImport_AppendInittab((char*)"CppMod", &initCppMod);
    exec("import CppMod",main_spa);  ///注意执行所在的场景
    exec("class MouseEventFrame(wx.Frame):\n"
         "  def __init__(self, parent, id):\n"
         "    wx.Frame.__init__(self, parent, id, 'Frame With Button', size=(300, 100))\n"
         "    self.panel = wx.Panel(self)\n"
         "    global wxf\n"
         "    wxf = self\n"
         "    self.button = wx.Button(self.panel,label='Not Over', pos=(100, 15))\n"
         "    self.ok = wx.Button(self.panel,label='Not Click', pos=(200, 15))\n"
         "    cpm = CppMod.CppClass()\n"
         "    self.Bind(wx.EVT_BUTTON, cpm.ClickButton, self.ok)\n"
         "    self.Bind(wx.EVT_BUTTON, CB.cb, self.button)\n"
         "app = wx.PySimpleApp()\n"
         "frame = MouseEventFrame(parent=None, id=-1)\n"
         "frame.Show()\n"
         "app.MainLoop()\n", main_spa
    );



    return 0;
  }
  catch(error_already_set)
  {
    PyErr_Print();
  }
}
