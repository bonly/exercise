package main

import (
"golang.org/x/mobile/audio"
"animation"
"atext"
"log"
"time"
"io"


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
	"image/jpeg"
_	"image/gif"
	"bytes"

	"math"
)

var eng sprite.Engine;
var font  *truetype.Font;
var start time.Time;
var game *sprite.Node;
var ar *animation.Arrangement;
var an *animation.Animation;
const numtr = 8;
var fi [numtr]io.Closer;
var playlist[numtr]*audio.Player;

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
    f, err := app.Open("tbg.jpg");
    //f, err := app.Open("bg.png");
    if err != nil {
		panic(err);
    }
    mb, err := ioutil.ReadAll(f);
    if err != nil {
		panic(err);
    }
    m, err := jpeg.Decode(bytes.NewReader(mb));
    //m, err := png.Decode(bytes.NewReader(mb));
    if err != nil {
		panic(err) ;
    }
    Tex, err := eng.LoadTexture(m);
    if err != nil {
		panic(err) ;
    }

    bg := sprite.SubTex{Tex, image.Rect(0, 0, 600, 2000)};
	gopher := &sprite.Node{
		Arranger: &animation.Arrangement{
			Offset: geom.Point{X: 0, Y: 0},//支点与父点间的偏移
			Size:   &geom.Point{155,520}, //可选，拉伸到尺寸
			// Pivot:  geom.Point{30 / 2, 30 / 2},
			Pivot:  geom.Point{0, 0},  //支点
			SubTex: bg,
			//Rotation: float32(math.Pi/2.0),
		},
	};
	eng.Register(gopher);
	game.AppendChild(gopher);


    // addText(game, "Hello 何睿湜!", 12, geom.Point{50, 50});
    //addText(game, "Hello world!", 12, geom.Point{20, 50});
//*
    // load texture
    mf, err := app.Open("skomas.png"); //打开文件
    mr, err := ioutil.ReadAll(mf);  //读入文件内容
    mc, err := png.Decode(bytes.NewReader(mr)); //以png方式解释文件内容
    mt, err := eng.LoadTexture(mc); //转为纹理
    mv := sprite.SubTex{mt, image.Rect(450,450,540,543)};  //建立纹理精灵
    // add node to pic
    //*  直接布局
    mn := &sprite.Node{    //建立节点精灵
      Arranger:   //建立布局
            &animation.Arrangement{   //动画参数设置
		Offset: geom.Point{X: 35, Y: 0},
		Size: &geom.Point{27,40},
		Pivot: geom.Point{13,20},
		SubTex: mv,
		Rotation: float32(math.Pi/-2.0),
      },
    };
    eng.Register(mn);  // 引擎中注册节点
    game.AppendChild(mn);  // 游戏中加入加点

    //把动画中的参数对象给全局指针以方便后面操作
    ar = mn.Arranger.(*animation.Arrangement);
//*/
    //*
    //加入使用动画布局
    ma := &sprite.Node{
	Arranger:
		&animation.Animation{
		    Current: "init",
		    States: map[string]animation.State{
			    "init": animation.State{},
			    "offscreen": animation.State{
				    Next: "offscreen",
				    Transforms:map[*sprite.Node]animation.Transform{
					    mn: animation.Transform{
						    Tween: clock.EaseInOut,
						    Transformer: animation.Rotate(0.32),
					    },
				    },
			    },
		    },
	    },
    };
    ma.Arranger.(*animation.Animation).Transition(0, "init");

    //eng.Register(ma);  // 引擎中注册节点, 动画节点可以不注册！
    //game.AppendChild(ma);  // 游戏中加入加点,动画节点可以不加入!

    //把动画对象给全局指针以方便后面操作
    an = ma.Arranger.(*animation.Animation);

    //*/
    addText(game, "Hello, 欢迎进入游戏", 12, geom.Point{ 50, 50}); 

    //加载声音
    for i := 0; i<8; i++{
	    rc, err := app.Open(fmt.Sprintf("track%d.wav",i));
	    if err != nil{
		    //log.Fatal(err);
	    }
	    fi[i]=rc;
	    p, err := audio.NewPlayer(rc, audio.Stereo16, 44100);
	    if err != nil{
		    //log.Fatal(err);
	    }
	    playlist[i] = p;
    }
}

func onStop() {
    log.Println("application stop");
}

func onTouch(t event.Touch) {
    log.Println("application touch");
    if an.Current== "init"{
	    an.Current = "offscreen";
    }else{
	    an.Current = "init";
    }
    playlist[0].Play();
}

func onDraw() {
    log.Println("============= application onDraw ===========================");
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
        an.Transition(t, an.Current);
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
