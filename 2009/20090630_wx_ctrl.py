#!/usr/bin/python
#-*-coding:utf-8-*-
#http://www.learningpython.com/2007/01/29/creating-a-gui-using-python-wxwidgets-and-wxpython/
import wx;
class wxHelloFrame(wx.Frame):
  def __init__(self,*args,**kwargs):
      wx.Frame.__init__(self,*args,**kwargs);
      self.create_controls();

  def create_controls(self):
      self.h_sizer=wx.BoxSizer(wx.HORIZONTAL);
      self.text=wx.StaticText(self,label="Enter some thing"); #创建text
      self.edit=wx.TextCtrl(self,size=wx.Size(250,-1));#创建Text编辑区
      self.h_sizer.Add(self.text,0,);
      self.h_sizer.AddSpacer((5,0));
      self.h_sizer.Add(self.edit,1);
      self.SetSizer(self.h_sizer);

class wxHelloApp(wx.App):
  def OnInit(self):
      frame=wxHelloFrame(None,title="wxHello");
      frame.Show();
      self.SetTopWindow(frame);
      return True;

if __name__=="__main__":
    app = wxHelloApp();
    app.MainLoop();