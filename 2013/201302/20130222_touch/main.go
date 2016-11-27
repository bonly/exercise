package main

import (
"animation"
"atext"
"log"
"time"

"golang.org/x/mobile/app"
"golang.org/x/mobile/event"
"golang.org/x/mobile/gl"
"golang.org/x/mobile/app/debug"
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
	"bytes"

// for Abs
"math"
)

var eng sprite.Engine;
var font  *truetype.Font;
var start time.Time;
var game *sprite.Node;
var ar *animation.Arrangement;

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
    f, err := app.Open("png/BG.png");
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

    bg := sprite.SubTex{Tex, image.Rect(0, 0, 472, 954)};
	gopher := &sprite.Node{
		Arranger: &animation.Arrangement{
			Offset: geom.Point{X: 0, Y: 0},
			Size:   &geom.Point{180, 320},
			// Pivot:  geom.Point{30 / 2, 30 / 2},
			Pivot:  geom.Point{0, 0},
			SubTex: bg,
			// Rotation: 0,
		},
	};
	eng.Register(gopher);
	game.AppendChild(gopher);


    //addText(game, "Hello 何!", 7, geom.Point{50, 50});
    addText(game, "Hello world!", 12, geom.Point{20, 50});

    // load texture
    mf, err := app.Open("png/Objects/SignArrow.png"); //打开文件
    mr, err := ioutil.ReadAll(mf);  //读入文件内容
    mc, err := png.Decode(bytes.NewReader(mr)); //以png方式解释文件内容
    mt, err := eng.LoadTexture(mc); //转为纹理
    mv := sprite.SubTex{mt, image.Rect(0,0,84,87)};  //建立纹理精灵
    // add node to pic
    mn := &sprite.Node{    //建立节点精灵
      Arranger:   //建立布局
            &animation.Arrangement{   //动画参数设置
		Offset: geom.Point{X: 20, Y: 20},
		Size: &geom.Point{20,20},
		Pivot: geom.Point{10,10},
		SubTex: mv,
		// Rotation: 0,
      },
    };
    eng.Register(mn);  // 引擎中注册节点
    game.AppendChild(mn);  // 游戏中加入加点

    //把动画中的参数对象给全局指针以方便后面操作
    ar = mn.Arranger.(*animation.Arrangement);
}

func onStop() {
    log.Println("application stop");
}

func onTouch(t event.Touch) {
    log.Println("application touch");
    x := math.Abs(float64(t.Loc.X - ar.Offset.X));
    y := math.Abs(float64(t.Loc.Y - ar.Offset.Y));
    if x<3 && y<3 {
	    //ar.Offset.Y += 10;
	    ar.Offset.X += 10;
    }
}

func onDraw() {
    log.Println("======= application onDraw ===========================");
    // setup opengl
    gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA);
    gl.Enable(gl.BLEND);

    // draw background.
    gl.ClearColor(0, 0, 0, 0);
    gl.Clear(gl.COLOR_BUFFER_BIT);

    debug.DrawFPS();

    t := now();

    updateGame(t);

    eng.Render(game, t);
    //gl.DeleteBuffer(buf);
}

func updateGame(t clock.Time){
	tween := clock.Linear(0, 2, t);
	ar.Offset.Y =  ar.Offset.Y + geom.Pt(tween);
	if ar.Offset.Y > 300 {
		ar.Offset.Y = 0;
	}
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
    log.Println("====== application load Font begin ==================");

	font := ""
	switch runtime.GOOS {
	case "android":
		font = "/system/fonts/DroidSansMono.ttf"
		//font = "/sdcard/NotoSansHans-Regular.ttf"
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
		return nil, err;
	}

        log.Println("====== application load Font end ====================");
	return freetype.ParseFont(b);
}
