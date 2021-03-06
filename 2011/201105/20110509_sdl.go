package main 

import (
  "github.com/DeedleFake/sdl"
  //"github.com/DeedleFake/sdl/img"
  "time"
)

func main() {
  err := sdl.Init(sdl.INIT_VIDEO);
  if err != nil {
    panic(err);
  }
  defer sdl.Quit();
  
  win, err := sdl.CreateWindow(
    "test",
    sdl.WINDOWPOS_UNDEFINED,
    sdl.WINDOWPOS_UNDEFINED,
    640,
    480,
    sdl.WINDOW_SHOWN,
  );
  if err != nil {
    panic(err);
  }
  defer win.Destroy();
  
  ren, err := win.CreateRenderer(-1, sdl.RENDERER_ACCELERATED);
  if err != nil {
    panic(err);
  }
  defer ren.Destroy();
  ren.SetDrawColor(100, 100, 255, sdl.ALPHA_OPAQUE);
  
  
  input := make(map[sdl.Keycode]bool);
  
  fps := time.Tick(time.Second / 60);
  for fps != nil{
    var ev sdl.Event;
    for sdl.PollEvent(&ev){
      switch ev := ev.(type){
        case *sdl.KeyboardEvent:
          if ev.Type == sdl.KEYDOWN {
            input[ev.Keysym.Sym] = true;
          }else{
            input[ev.Keysym.Sym] = false;
          }
        case *sdl.QuitEvent:
          fps = nil
      }
    }

  ren.Clear();
  ren.Present();
  }
}

