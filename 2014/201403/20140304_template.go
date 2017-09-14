package main

import (
"golang.org/x/mobile/app"
"golang.org/x/mobile/event/key"
"golang.org/x/mobile/event/paint"
"golang.org/x/mobile/event/size"
"golang.org/x/mobile/event/touch"
"golang.org/x/mobile/event/lifecycle"

"golang.org/x/mobile/gl"

"golang.org/x/mobile/exp/sprite"
"golang.org/x/mobile/exp/sprite/glsprite"
"golang.org/x/mobile/exp/gl/glutil"

"fmt"
"log"
)

func main(){
	app.Main(Loop);
}

var glx 		gl.Context; //定义opengl内容
var eng 		sprite.Engine;  //引擎
var	images    	*glutil.Images; //引擎操作的画板
var root     	*sprite.Node; //场景的根结点

func Loop(ap app.App){
	for ev := range ap.Events(){
		switch ev := ap.Filter(ev).(type){
			case lifecycle.Event:{
				switch ev.Crosses(lifecycle.StageVisible){
					case lifecycle.CrossOn:{ //开始
						glx, _ = ev.DrawContext.(gl.Context); //创建opengl
						onStart(glx);
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

func onStart(glx gl.Context){
	log.Println("App start");
	images = glutil.NewImages(glx);  //生成基础图片作为画板
	eng = glsprite.Engine(images);  //设置引擎
	
	root := &sprite.Node{}; //生成根精灵结点
	eng.Register(root); //注册根到引擎中
}

func onPaint(glx gl.Context){
	glx.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA);
	glx.Enable(gl.BLEND);

	glx.ClearColor(1, 1, 1, 1);
	glx.Clear(gl.COLOR_BUFFER_BIT);

	// eng.Render()
}


