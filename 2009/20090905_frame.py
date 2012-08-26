#-*-coding:utf-8-*-
#!/usr/bin/env python   
import wx
import sys

class frame(wx.Frame):
  def __init__(self, parent, id, title):
    print "frame __init__"
    wx.Frame.__init__(self, parent, id, title)
    
class app(wx.App):
  def __init__(self, redirect=True, filename=None):
    print "app _init_"
    wx.App.__init__(self, redirect, filename)
    
  def OnInit(self):
    print "OnInit"
    self.myframe=frame(parent=None,id=-1,title='startup')
    self.myframe.Show()
    self.SetTopWindow(self.myframe)
    print>>sys.stderr,"A pretend error message"
    return True
    
  def OnExit(self):
    print "OnExit"
    
if __name__=='__main__':
  myapp=app(redirect=True)
  print "befor mainloop"
  myapp.MainLoop()
  print "after MainLoop"    
    