#!/usr/bin/python
#-*-coding:utf-8-*-
import wx

global wxf
def OnButtonClick(event):
        print "on  out of button click\n"
        wxf.ok.SetLabel("clicked Me!")
        wxf.panel.Refresh()
        event.Skip()

def OnEnterWindow(event,wxf):
        print "on enter window\n"
        wxf.button.SetLabel("Over Me!")
        event.Skip()

def OnLeaveWindow(event):
        button.SetLabel("Not Over")
        event.Skip()

class MouseEventFrame(wx.Frame):
 def __init__(self, parent, id):
        wx.Frame.__init__(self, parent, id, 'Frame With Button', size=(300, 100))
        self.panel = wx.Panel(self)
        global wxf
        wxf = self
        self.button = wx.Button(self.panel,label="Not Over", pos=(100, 15))
        self.ok = wx.Button(self.panel,label="Not Click", pos=(200, 15))
        self.Bind(wx.EVT_BUTTON, self.OnButtonClick, self.button)    #1 绑定按钮事件
        self.Bind(wx.EVT_BUTTON, OnButtonClick, self.ok)    #1 绑定按钮事件

        #self.button.Bind(wx.EVT_ENTER_WINDOW, OnEnterWindow, self)     #2 绑定鼠标位于其上事件
        #self.button.Bind(wx.EVT_LEAVE_WINDOW,
        #    self.OnLeaveWindow)     #3 绑定鼠标离开事件

 def OnButtonClick(self, event):
        print "inner button click\n"
        OnButtonClick(event)
        self.panel.SetBackgroundColour('Green')
        self.panel.Refresh()



if __name__ == '__main__':
    app = wx.PySimpleApp()
    frame = MouseEventFrame(parent=None, id=-1)
    frame.Show()
    app.MainLoop()
