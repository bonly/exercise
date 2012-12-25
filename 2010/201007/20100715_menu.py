#!/usr/bin/python
# -*- coding: utf-8 -*-
import wx;
import wx.py.images as images;

class ToolbarFrame(wx.Frame):
    def __init__(self, parent, id):
        wx.Frame.__init__(self, parent, id, 'Toolbars', size= (300,200));
        panel = wx.Panel(self);
        panel.SetBackgroundColour('White');
        statusBar = self.CreateStatusBar(); #1创建状态栏
        toolbar = self.CreateToolBar(); #2创建工具栏
        toolbar.AddSimpleTool(wx.NewId(), images.getPyBitmap(), "New", "Long help for 'New'"); #3给工具栏加一个工具
        toolbar.Realize(); #4准备显示工具栏
        menuBar = wx.MenuBar() #创建菜单栏
        #创建两个菜单
        menu1 = wx.Menu();
        menuBar.Append(menu1, "&File");
        menu2 = wx.Menu();
        #6创建菜单项目 
        menu2.Append(wx.NewId(), "&Copy", "Copy in status bar");
        menu2.Append(wx.NewId(), "C&ut", "");
        menu2.Append(wx.NewId(), "Paste", "");
        menu2.AppendSeparator();
        menu2.Append(wx.NewId(), "&Options...", "Display Options");
        menuBar.Append(menu2, "&Edit"); #在菜单栏上附上菜单
        self.SetMenuBar(menuBar) #在框架中附上菜单栏
        
#if __name__ == '__main__':
#    app = wx.PySimpleApp();
#    frame = ToolbarFrame(parent=None, id=-1);
#    frame.Show();
#    app.MainLoop();
  
# wxpython in action 44/565        
