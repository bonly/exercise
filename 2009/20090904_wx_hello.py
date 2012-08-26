#-*-coding:utf-8-*-
#!/usr/bin/env python   
"""Hello, wxPython! program."""
import wx    
 
class Frame(wx.Frame):   #2 wx.Frame子类
    """Frame class that displays an image.""" 
    def __init__(self, image, parent=None, id=-1,   
                 pos=wx.DefaultPosition, 
                 title='Hello, wxPython!'):     #3图像参数
        """Create a Frame instance and display image.""" 
    #4 显示图像
        temp = image.ConvertToBitmap()                           
        size = temp.GetWidth(), temp.GetHeight()                
        wx.Frame.__init__(self, parent, id, title, pos, size)   
        self.bmp = wx.StaticBitmap(parent=self, bitmap=temp)     
 
#class App(wx.App):  #5 wx.App子类 
#    """Application class.""" 
#    def OnInit(self): 
    #6 图像处理 
#        image = wx.Image('/home/bonly/Downloads/hart.jpg', wx.BITMAP_TYPE_JPEG)   
#        self.frame = Frame(image)
#        self.frame.Show() 
#        self.SetTopWindow(self.frame) 
#        return True 
	
#def main():  #7        
#    app = App()      
#    app.MainLoop()   

#if __name__ == '__main__':    
#     main() 
