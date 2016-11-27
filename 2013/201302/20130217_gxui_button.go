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

  window := theme.CreateWindow(800, 600, "test window");

  btn := theme.CreateButton();
  btn.SetText("点 here");
 
  window.AddChild(btn);
  window.OnClose(driver.Terminate);
}

