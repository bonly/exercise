package main

import (
"log"
"image"
_ "image/jpeg"

"golang.org/x/mobile/app"
"golang.org/x/mobile/event"
"golang.org/x/mobile/gl"
"golang.org/x/mobile/sprite/glsprite"
"golang.org/x/mobile/sprite"
"golang.org/x/mobile/app/debug"
"golang.org/x/mobile/f32"
"golang.org/x/mobile/geom"
)

var eng = glsprite.Engine();
var scene *sprite.Node;
var texture *sprite.SubTex;


func main() {
	app.Run(app.Callbacks{
		Start: onStart,
		Draw:  onDraw,
		Touch: onTouch,
		Stop:  onStop,
	});
}

func onStart() {
    log.Println("application init");

    LoadTexture();
    LoadScene();


    /*
    n := &sprite.Node{};
	eng.Register(n);
	scene.AppendChild(n);
	eng.SetSubTex(n, *texture);

	NODE_SIZE := geom.PixelsPerPt;
	matrix := f32.Affine{
		{NODE_SIZE, 0, float32(15) * NODE_SIZE},
		{0, NODE_SIZE, float32(15) * NODE_SIZE},
	};
	eng.SetTransform(n, matrix);
	*/
}

/*
f32:
封装了大量OpenGL中float32的代数矩阵运算。
*/
func LoadScene(){
	scene = &sprite.Node{};
	eng.Register(scene);
	eng.SetSubTex(scene, *texture);
	NODE_SIZE := geom.PixelsPerPt;
	matrix := f32.Affine{
		{150, 0, float32(15) * NODE_SIZE},
		{0, 150, float32(15) * NODE_SIZE},
	};
	eng.SetTransform(scene, matrix);

	// eng.SetTransform(scene, f32.Affine{
		// {1,0,0},
		// {0,1,0},
		// });
/*
在New出一个Node并贴上图之后，我们会发现显示在屏幕上的图非常非常小，这时可以这么做：
NODE_SIZE := geom.PixelsPerPt
matrix := f32.Affine{
    {NODE_SIZE, 0, 0},
    {0, NODE_SIZE, 0},
}
eng.SetTransform(node.SpriteNode, matrix)

这里的geom中的基本单位是Pt，我们知道Pt是屏幕的物理尺寸，代表一个点，geom.PixelsPerPt就是一个Pt中有多少个Px，
我们需要将我们的Node放大这么多倍才能得到它真正在屏幕上显示时的尺寸。

由于Go的Sprite包目前还过于简陋，所以我们这里需要回顾了一下图形学的知识。。 一个图形的变换是用过一个 3x3 矩阵来做到的：
[ s 0 x ]
[ 0 s y ]
[ 0 0 s ]
如果s＝1, x = y = 0，这就是一个变换的基，s的值就是放大的倍率，x和y分别代表x轴y轴上平移的距离。

而一个f32.Affine则是一个 3x2 矩阵，其z轴固定为[0, 0, 1]，所以表示的是2D的坐标变换。
于是我们构造2D变换矩阵，使用eng.SetTransform(node, matrix)将变换应用在SpriteNode上，就可以移动Node了。
*/
}

/*
app:

包含了一个最基本的使用NativeActivity的Android App框架，目前只提供了2个实用接口

app.Run(cb Callbacks)，app启动的入口，其中
app.Callbacks{
    Start: func(), 
    Stop: func(), 
    Draw: func(), 
    Touch: func(event.Touch) 
}
定义了Start,Stop两个生命周期函数以及Draw渲染和Touch事件回调。

app.Open(name String) (ReadSeekCloser, error)：获取一个资源文件，只能读取asset目录下的文件。
*/
func LoadTexture(){
    //opens a named asset
    //func Open(name string) (ReadSeekCloser, error)
    reader, err := app.Open("head.jpg");
    defer reader.Close();
    if err != nil{
    	log.Println("read err!");
    	log.Fatal(err);
    }

    img, _, err := image.Decode(reader);
    if err != nil{
    	log.Println("Decode err!");
    	log.Fatal(err);
    }

    tex, err := eng.LoadTexture(img);
    if err != nil{
    	log.Println("tex err!");
    	log.Fatal(err);
    }

    texture = &sprite.SubTex{tex, img.Bounds()};
}

func onStop() {
	log.Println("application stop");
}

func onTouch(t event.Touch) {
    log.Println("application touch");

}

func onDraw() {
	// draw background.
	gl.ClearColor(1, 1, 1, 1)
	gl.Clear(gl.COLOR_BUFFER_BIT)
	// render root node (scene). then it will render whole node tree.
	eng.Render(scene, 0)
	debug.DrawFPS()
}
/*
app/debug:
只提供了一个debug.DrawFPS()用来画FPS的，而且固定画在屏幕左下角。
bind:
bind相关有一系列包，用于在Android程序中调用Go编出来的so，目前文档比较匮乏
event:
定义了Touch事件，一个Touch事件包含一个TouchType和geom.Point。TouchType包括TouchStart, TouchMove, TouchEnd三种事件，
目前的Key事件官方还没实现，在back退出程序时打出来的Log中能看到
TODO input event: key。
geom:
定义了一个二维坐标系，坐标原点依旧是在左上角
gl:
封装了大量OpenGL中的渲染、变换、贴图相关的函数。
gl/glutil:
一些工具函数，比如gl.Program的创建，Image的绘制等等。
sprite:
这是一个很有意思的包，提供了一个简单（不能更简单）封装的2D游戏引擎，目前能做到的是构造SpriteNode树，给SpriteNode贴图，渲染Node树。
这个引擎实在简单，以至于一些基本的Sprite的Translate,Scale等等操作都还得通过gl.Affine来做。
sprite/clock:
定义了Sprite世界的时间，以及一个线性时间变换和Bezier时间变换。
*/

