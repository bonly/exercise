package main

import (
"animation"
"atext"
"log"
"time"

"golang.org/x/mobile/app"
"golang.org/x/mobile/event"
"golang.org/x/mobile/gl"
//"golang.org/x/mobile/app/debug"
_ "golang.org/x/mobile/f32"
"golang.org/x/mobile/geom"
"golang.org/x/mobile/sprite"

"golang.org/x/mobile/sprite/clock"
"golang.org/x/mobile/sprite/glsprite"

"image/color"
"code.google.com/p/freetype-go/freetype/truetype"

// for load font
	"runtime"
	"fmt"
	"io/ioutil"
	"code.google.com/p/freetype-go/freetype"

// for png
	"image"
	"image/png"
//	"image/jpeg"
_	"image/gif"
	"bytes"

//	"math"
)

var eng sprite.Engine;
var font  *truetype.Font;
var start time.Time;
var game *sprite.Node;
var ar *animation.Arrangement;
var an *animation.Animation;

func main() {
	app.Run(app.Callbacks{
		Start: onStart,
		Draw:  onDraw,
		Touch: onTouch,
		Stop:  onStop,
	});
}

func onStart() {
	log.Println("application start");
	start = time.Now();
	eng = glsprite.Engine();

	var err error;
	font, err = loadFont();
	if err != nil {
		panic(err);
	}

        game = new(sprite.Node);
        eng.Register(game);



    // read pic
    f, err := app.Open("dbback_01.png");
    if err != nil {
		panic(err);
    }
    mb, err := ioutil.ReadAll(f);
    if err != nil {
		panic(err);
    }
    m, err := png.Decode(bytes.NewReader(mb));
    if err != nil {
		panic(err) ;
    }
    Tex, err := eng.LoadTexture(m);
    if err != nil {
		panic(err) ;
    }

    //geom中的基本单位是Pt，我们知道Pt是屏幕的物理尺寸，代表一个点，geom.PixelsPerPt就是一个Pt中有多少个Px，
    //我们需要将我们的Node放大这么多倍才能得到它真正在屏幕上显示时的尺寸
    /*
    dp = 与密度无关的像素,每英寸160点(dpi)的显示器上，1dp = 1px  
    概念：160dpi(中)的手机上每1px(像素) = 1dp
    dip = device independent pixels 等同于dp的别名
    px = 像素 320*480的屏幕在横向有320个象素，在纵向有480个象素

    pt = 表示一个点，是屏幕的物理尺寸。大小为1英寸=72pt,印刷上的点1/72 of an inch (0.3527 mm)
    sp = scaled pixels 放大像素  字体统一用这个
    android上文字的尺寸一律用sp单位，非文字的尺寸一律使用dp单位

    1英寸 = 2.54厘米
    px = dp * (dpi/160)
    sp = pt * (dpi/160)

    ppi (pixels per inch)：图像分辨率 （在图像中，每英寸所包含的像素数目）
    PPI = √（长度像素数² + 宽度像素数²） / 屏幕对角线英寸数
    dpi (dots per inch)： 打印分辨率 （每英寸所能打印的点数，即打印精度）

    假设有一部手机，屏幕的物理尺寸为1.5英寸x2英寸，屏幕分辨率为240×320，
    则我们可以计算出在这部手机的屏幕上，
    每英寸包含的像素点的数量为240/1.5=160dpi（横向）或320/2=160dpi（纵向），160dpi就是这部手机的像素密度，
    像素密度的单位dpi是Dots Per Inch的缩写，即每英寸像素数量。
    横向和纵向的这个值都是相同的，原因是大部分手机屏幕使用正方形的像素点。

    Android系统定义了四种像素密度：低（120dpi）、中（160dpi）、高（240dpi）和超高（320dpi），
    它们对应的dp到px的系数分别为0.75、1、1.5和2，这个系数乘以dp长度就是像素数。

    界面上有一个长度为“100dp”的图片，那么它在240dpi的手机上实际显示为100×(240/160)=150px，
    在320dpi的手机上实际显示为100×(320/160)=200px。
    如果你拿这两部手机放在一起对比，会发现这个图片的物理尺寸“差不多”，这就是使用dp作为单位的效果

    sony xperia v 4.3 inch 1280x720
    dpi = sqrt(pow(1280,2)+pow(720,2))/4.3 = 341.5 属于超高
    72pt = 72pt/inch = 320dip = 320px/inch 
    geom.PixelsPerPt = 320px/72pt = 4.4444 

    */
    ns := geom.PixelsPerPt;
    str := fmt.Sprintf("pixels per pt is: %v", ns);
    width := app.GetConfig().Width;
    height := app.GetConfig().Height;
0
    bg := sprite.SubTex{Tex, image.Rect(0, 0, 600, 600)}; //加载进来时的大小是以px为单位
	gopher := &sprite.Node{
		Arranger: &animation.Arrangement{
			Offset: geom.Point{X: 0, Y: 0},//支点与父点间的偏移
			//Size:   &geom.Point{269, 162}, //可选，拉伸到尺寸
			Size:   &geom.Point{width, height}, //可选，拉伸到尺寸
			// Pivot:  geom.Point{30 / 2, 30 / 2},
			Pivot:  geom.Point{0, 0},  //支点
			SubTex: bg,
			//Rotation: float32(math.Pi/2.0),
		},
	};
	eng.Register(gopher);
	game.AppendChild(gopher);

}

func onStop() {
    log.Println("application stop");
}

func onTouch(t event.Touch) {
    log.Println("application touch");
}

func onDraw() {
    log.Println("============= application onDraw ===========================");
    // setup opengl
    gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA);
    gl.Enable(gl.BLEND);

    // draw background.
    gl.ClearColor(0, 0, 0, 0);
    gl.Clear(gl.COLOR_BUFFER_BIT);

    //debug.DrawFPS();

    t := now();

    updateGame(t);

    eng.Render(game, t);
    //gl.DeleteBuffer(buf);
}

func updateGame(t clock.Time){
}


func addText(parent *sprite.Node, str string, size geom.Pt, pos geom.Point) {
	p := &sprite.Node{
		Arranger: &animation.Arrangement{
			Offset: pos,
		},
	}
	eng.Register(p)
	parent.AppendChild(p)
	pText := &sprite.Node{
		Arranger: &atext.String{
			Size:  size,
			Color: color.Black,
			Font:  font,
			Text:  str,
		},
	}
	eng.Register(pText)
	p.AppendChild(pText)
}

func now() clock.Time {
	d := time.Since(start)
	return clock.Time(60 * d / time.Second)
}

func loadFont() (*truetype.Font, error) {
        log.Println("============= application load Font begin ===========================");

	font := ""
	switch runtime.GOOS {
	case "android":
		//font = "/system/fonts/DroidSansMono.ttf"
		//font = "/system/fonts/NotoSansHans-Regular.otf"
		//font = "/system/fonts/MTLmr3m.ttf"
		font = "/sdcard/cn.ttf"
		// font = "/system/fonts/DroidSansFallback.ttf"
	case "darwin":
		//font = "/Library/Fonts/Andale Mono.ttf"
		font = "/Library/Fonts/Arial.ttf"
		//font = "/Library/Fonts/儷宋 Pro.ttf"
	case "linux":
		font = "/usr/share/fonts/truetype/droid/DroidSansMono.ttf"
	default:
		return nil, fmt.Errorf("go.mobile/app/debug: unsupported runtime.GOOS %q", runtime.GOOS)
	}
	b, err := ioutil.ReadFile(font)
	if err != nil {
		return nil, err
	}

    log.Println("============= application load Font end ===========================");
	return freetype.ParseFont(b);

}
