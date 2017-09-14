package main 

import (
"golang.org/x/mobile/app"
"golang.org/x/mobile/event/key"
"golang.org/x/mobile/event/paint"
"golang.org/x/mobile/event/size"
"golang.org/x/mobile/event/touch"
"golang.org/x/mobile/event/lifecycle"
"fmt"
)

func main(){
	app.Main(func(ap app.App){
		for ev := range ap.Events(){
			switch ev := ap.Filter(ev).(type){
				case lifecycle.Event:{
					switch ev.Crosses(lifecycle.StageVisible){
						case lifecycle.CrossOn:{
							fmt.Println("crossOn");
							break;
						}
						case lifecycle.CrossOff:{
							fmt.Println("crossOff");
							break;
						}
					}
					break;
				} //end lifecycle
				case size.Event:{
					fmt.Println("size");
					break;
				}
				case paint.Event:{
					fmt.Println("paint");
					break;
				}
				case touch.Event:{
					fmt.Println("touch");
					break;
				}
				case key.Event:{
					fmt.Println("key");
					break;
				}
			}
		}
	});
}