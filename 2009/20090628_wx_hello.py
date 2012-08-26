import wx;
class wxHelloApp(wx.App):
    def OnInit(self):
        frame=wx.Frame(None,title="wxHello");
        frame.Show();
        self.SetTopWindow(frame);
        return True;
if __name__ == "__main__":
    app=wxHelloApp();
    app.MainLoop();