import wx;
class wxHelloFrame(wx.Frame):
  def __init__(self,*args,**kwargs):
      wx.Frame.__init__(self,*args,**kwargs);
      self.create_controls();

  def create_controls(self):
      pass;

class wxHelloApp(wx.App):
  def OnInit(self):
      frame=wxHelloFrame(None,title="wxHello");
      frame.Show();
      self.SetTopWindow(frame);
      return True;

if __name__=="__main__":
    app = wxHelloApp();
    app.MainLoop();