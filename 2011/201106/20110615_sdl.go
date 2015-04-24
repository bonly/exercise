package main
 
import (
"github.com/neagix/Go-SDL/sdl"
"log"
"unsafe"
"math/rand"
)
 
func draw_point(x int32,y int32,value uint32,screen* sdl.Surface) {
  var pix = uintptr(screen.Pixels);
  pix += (uintptr)((y*screen.W)+x)*unsafe.Sizeof(value);
  var pu = unsafe.Pointer(pix);
  var pp *uint32;
  pp = (*uint32)(pu);
  *pp = value;
}
 
func main() {
 
  var screen = sdl.SetVideoMode(640, 480, 32, sdl.RESIZABLE)
 
  if screen == nil {
    log.Fatal(sdl.GetError())
  }
 
  var n int32;
  for n=0;n<1000000;n++ {
 
    var y int32 =rand.Int31()%480;
    var x int32 =rand.Int31()%640;
    var value uint32 = rand.Uint32();
    draw_point(x,y,value,screen);
 
    screen.Flip();
  }
}
