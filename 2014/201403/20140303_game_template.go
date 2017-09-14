package main 

import (
"golang.org/x/mobile/app"
"golang.org/x/mobile/event/key"
"golang.org/x/mobile/event/paint"
"golang.org/x/mobile/event/size"
"golang.org/x/mobile/event/touch"
"golang.org/x/mobile/event/lifecycle"

"golang.org/x/mobile/gl"

"fmt"
)

func main(){
	app.Main(Loop);
}

var glx gl.Context; //定义opengl内容

func Loop(ap app.App){
	for ev := range ap.Events(){
		switch ev := ap.Filter(ev).(type){
			case lifecycle.Event:{
				switch ev.Crosses(lifecycle.StageVisible){
					case lifecycle.CrossOn:{ //开始
						fmt.Println("crossOn");
						glx, _ = ev.DrawContext.(gl.Context); //创建opengl
						break;
					}
					case lifecycle.CrossOff:{ //结束
						fmt.Println("crossOff");
						glx = nil; //删除opengl
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
				if glx == nil || ev.External{
					continue;
				}
				onPaint(glx);
				ap.Publish(); //flush draw buffer
				ap.Send(paint.Event{}); //保持重绘
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
}	

func onPaint(glx gl.Context){
	glx.ClearColor(1, 1, 1, 1);
	glx.Clear(gl.COLOR_BUFFER_BIT);
}
