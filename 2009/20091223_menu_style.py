#!/usr/bin/python
#coding:utf-8
'''
Created on 2011-3-15

@author: bonly
'''
import wx;
import wx.py.images;
import CB;

class ToolbarFrame(wx.Frame):
   def __init__(self,parent,id):
      wx.Frame.__init__(self,parent,id,'Toolbars',size=(300,200))
      panel = wx.Panel(self)
      panel.SetBackgroundColour('White')
      statusBar = self.CreateStatusBar() #创建状态栏
      toolbar = self.CreateToolBar() #创建工具栏
      toolbar.AddSimpleTool(wx.NewId(),wx.py.images.getPyBitmap(),"New","Long help for 'New'") #给工具栏加一个工具
      toolbar.Realize() #准备显示工具栏
      
      menuBar=wx.MenuBar() #创建菜单栏
      
      #创建两个菜单
      menu1 = wx.Menu()
      menuBar.Append(menu1,"&File")
      
      menu2 = wx.Menu()      
      #创建菜单的项目
      menu2.Append(wx.NewId(),"&Copy","Copy in status bar")
      menu2.Append(wx.NewId(),"C&ut","")
      menu2.Append(wx.NewId(),"Past","")
      menu2.AppendSeparator()
      menu2.Append(wx.NewId(),"&Options...","Display Options")
      menuBar.Append(menu2,"&Edit") #在菜单栏上附上菜单
      self.SetMenuBar(menuBar) #在柜架上附上菜单栏
      
      
if __name__ == '__main__':
   app = wx.PySimpleApp()
   frame = ToolbarFrame(parent=None,id=-1)
   frame.Show()
   app.MainLoop()
