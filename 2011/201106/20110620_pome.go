//skip  //不知此标识作用

package pomelib

/*
#cgo CFLAGS: -I/home/bonly/libpomelo/include -I/home/bonly/libpomelo -I/home/bonly/libpomelo/deps/uv/include -I/home/bonly/libpomelo/deps/jansson/src
#cgo linux CFLAGS: -DLINUX=1
#cgo LDFLAGS: -L/home/bonly/libpomelo -lpomelo -L/home/bonly/libpomelo/deps/jansson/src/.libs/ -ljansson -L/home/bonly/libpomelo/deps/uv -luv
#include "20110621_pome.h"
*/
import "C"

import "fmt"

func Gfun(){
     fmt.Println("Gfun()");
     C.init();
}

//export Gfun4C
func Gfun4C(){
     fmt.Println("Gfun4C()");
     var mys C.struct_Tmy;
     mys.abc = 10;
     fmt.Println("abc is: ", mys.abc);
     mys.kstc.def = 5;
     fmt.Println("def is: ", mys.kstc.def);
     

}
