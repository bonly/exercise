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

)

var eng sprite.Engine;
var font  *truetype.Font;
var start time.Time;
var game *sprite.Node;

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
    f, err := app.Open("balloon_sheet.png");
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

    balloon := sprite.SubTex{Tex, image.Rect(0, 0, 194, 42)};
	gopher := &sprite.Node{
		Arranger: &animation.Arrangement{
			Offset: geom.Point{X: 20, Y: 20},
			Size:   &geom.Point{30, 30},
			Pivot:  geom.Point{30 / 2, 30 / 2},
			SubTex: balloon,
		},
	};
	eng.Register(gopher);
	game.AppendChild(gopher);


    // addText(game, "Hello 何睿湜!", 12, geom.Point{50, 50});
    addText(game, "Hello world!", 12, geom.Point{150, 150});

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

	debug.DrawFPS();
	    
    t := now();
	eng.Render(game, t);
	//gl.DeleteBuffer(buf);
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
		font = "/system/fonts/DroidSansMono.ttf"
		//font = "/system/fonts/MTLmr3m.ttf"
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
