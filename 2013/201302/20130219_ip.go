package main

import (
"github.com/google/gxui"
"github.com/google/gxui/themes/dark"
"github.com/google/gxui/drivers/gl"
_ "code.google.com/p/freetype-go/freetype"
)


func main(){
  gl.StartDriver(appMain);
}


func appMain(driver gxui.Driver){
  theme := dark.CreateTheme(driver);


  lay := theme.CreateLinearLayout();
  lay.SetDirection(gxui.LeftToRight);

  edt := theme.CreateTextBox();
  edt.SetText("192.168.1.1");
  lay.AddChild(edt);


  btn := theme.CreateButton();
  btn.SetText("show as hex");
  lay.AddChild(btn);

  lab := theme.CreateLabel();
  lay.AddChild(lab);

  btn.OnClick(func(gxui.MouseEvent){
	  lab.SetText("clicked");
  });

  window := theme.CreateWindow(800, 600, "ip conver");
  window.AddChild(lay);
  window.OnClose(driver.Terminate); //关闭时退出程序
}

/*
chinese input:
https://github.com/mattn/glfw/commit/d48fc47dc96d7486539c3cf22617f747f36f5c62
916
-                if (XFilterEvent(event, None))
-                {
-                    // Discard intermediary (dead key) events for character input
-                    break;
-                }
-

1743
     {
	              XEvent event;
		       XNextEvent(_glfw.x11.display, &event);
	       +        if (XFilterEvent(event, None))
	       +            continue;
			processEvent(&event);
     }

     */
