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

"time"
"golang.org/x/mobile/exp/sprite/clock"

"golang.org/x/mobile/asset"
"image"
_ "image/png"
//_ "image/jpeg"

// "animation"
// "golang.org/x/mobile/geom"
"golang.org/x/mobile/exp/f32"
"os"
)

func main(){
	app.Main(Loop);
}

var glx 		gl.Context; //定义opengl内容
var eng 		sprite.Engine;  //引擎
var	images    	*glutil.Images; //引擎操作的画板
var root     	*sprite.Node; //场景的根结点

var startTime   = time.Now();

func Loop(ap app.App){
	for ev := range ap.Events(){
		var sz size.Event;  //定义窗口大小变化事件
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
				sz = ev; //窗口大小变化了
				break;
			}
			case paint.Event:{
				// fmt.Println("paint");
				if glx == nil || ev.External{
					continue;
				}
				onPaint(glx, sz);  //给出窗口的大小
				ap.Publish(); //flush draw buffer
				ap.Send(paint.Event{}); //保持重绘动画
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
	
	root = &sprite.Node{}; //生成根精灵结点
	eng.Register(root); //注册根到引擎中
	// eng.SetTransform(root, f32.Affine{
	// 	{1, 0, 0},
	// 	{0, 1, 0},
	// });

	// loadTexture();//加载纹理对象

	fl, err := asset.Open("shift.png"); //打开图片文件 
	if err != nil{
		fmt.Println(err);
		os.Exit(1);
	}
	defer fl.Close();

	mb, _, err := image.Decode(fl); //解释文件格式
	if err != nil{
		fmt.Println(err);
		os.Exit(1);
	}
	Tex, err := eng.LoadTexture(mb); //读入纹理
	if err != nil{
		fmt.Println(err);
		os.Exit(1);
	}

	bg := sprite.SubTex{Tex, image.Rect(0, 0, 521, 521)};

	sp := &sprite.Node{};
	eng.Register(sp); //要先注册挂到树上再给数据
	root.AppendChild(sp);	

	eng.SetSubTex(sp, bg);
	eng.SetTransform(sp, f32.Affine{   //这个地方各种可能都不行???
		{0,500,0},
		{500,0,0}});
}

func onPaint(glx gl.Context, sz size.Event){
	log.Println("App paint");

	glx.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA);
	glx.Enable(gl.BLEND);  //设置gl特性

	glx.ClearColor(0, 0, 0, 0);
	glx.Clear(gl.COLOR_BUFFER_BIT); //清屏

	now := clock.Time(time.Since(startTime) * 60 / time.Second);
	if root != nil{
		eng.Render(root, now, sz);  //引擎绘画
	}
}

func loadTexture(){
	fl, err := asset.Open("tbg.jpg"); //打开图片文件 
	if err != nil{
		fmt.Println(err);
	}
	defer fl.Close();

	mb, _, err := image.Decode(fl); //解释文件格式
	if err != nil{
		fmt.Println(err);
	}
	Tex, err := eng.LoadTexture(mb); //读入纹理
	if err != nil{
		fmt.Println(err);
	}

	bg := sprite.SubTex{Tex, image.Rect(128+1, 0, 128*2-1, 128)};
	// sp := &sprite.Node{
	// 	Arranger: &animation.Arrangement{
	// 		Offset: geom.Point{X: 0, Y: 0},//支点与父点间的偏移
	// 		Size:   &geom.Point{155,520}, //可选，拉伸到尺寸
	// 		// Pivot:  geom.Point{30 / 2, 30 / 2},
	// 		Pivot:  geom.Point{0, 0},  //支点
	// 		SubTex: bg,
	// 		//Rotation: float32(math.Pi/2.0),
	// 	},
	// };
	root = &sprite.Node{};
	eng.SetSubTex(root, bg);
	eng.Register(root);
	// root.AppendChild(sp);	
}
