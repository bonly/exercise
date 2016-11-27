package main

import (
	"github.com/google/gxui"
	"github.com/google/gxui/themes/dark"
	"github.com/google/gxui/drivers/gl"
	
	"fmt"
	"log"
	"io/ioutil"
	"math/rand"
	"time"
)

type calc struct{
	X int;
	Y int;
	S int;
};

var dt []calc;
var window gxui.Window;
var theme  gxui.Theme;
var font   gxui.Font;
const size int=60;

func main(){
	rand.Seed(time.Now().UnixNano());
		
	newData();

	gl.StartDriver(appMain);
}

func newData(){	
	dt = make([]calc, 0, size);
	for i:=0; i<size; i++{
		x := rand.Int() % 10 * 10 + 5;
		y := x;
		dt = append(dt, calc{x, y, x*y});
	}	
}
func loadFont(driver gxui.Driver){
	//font comes from windows
    fontData, err := ioutil.ReadFile("/home/bonly/.fonts/YaHei/YaHei.Consolas.1.11b.ttf"); 
    if err != nil {
        log.Fatalf("error reading font: %v", err);
    }
    font, err = driver.CreateFont(fontData, 20);
    if err != nil {
        panic(err);
    }	
}

func appMain(driver gxui.Driver){
	theme = dark.CreateTheme(driver);
	
	window = theme.CreateWindow(900, 700, "十位相同个位为5");
	window.SetScale(1.0);
	
	// load chinese font
	loadFont(driver);
	
	window.OnClose(driver.Terminate);
	
	// create label context
	addElem(dt);
}

func addElem(lst []calc){
	backgroup := theme.CreateLinearLayout();
	backgroup.SetDirection(gxui.LeftToRight);
	window.AddChild(backgroup);	
	
	left := theme.CreateLinearLayout();
	left.SetDirection(gxui.TopToBottom);
	backgroup.AddChild(left);	
	
	mid := theme.CreateLinearLayout();
	mid.SetDirection(gxui.TopToBottom);
	backgroup.AddChild(mid);
		
	right := theme.CreateLinearLayout();
	right.SetDirection(gxui.TopToBottom);
	backgroup.AddChild(right);	
	for i, ca := range lst{
		label_q := theme.CreateLabel();
		txt := fmt.Sprintf("%2d x %2d = ", ca.X, ca.Y);
		label_q.SetText(txt);
		label_q.SetFont(font);
		
		label_v := theme.CreateLabel();
		value := fmt.Sprintf("%d\t", ca.S);
		label_v.SetText(value);
		label_v.SetFont(font);
		label_v.SetVisible(false);
		
		input := theme.CreateTextBox();
		input.SetText("");
		input.SetFont(font);
		input.OnLostFocus(func(){
			label_v.SetVisible(true);
		});
				
	    switch i % 3 {
			case 0:{ 
				lay := theme.CreateLinearLayout();
				lay.SetDirection(gxui.LeftToRight);
				lay.AddChild(label_q);
				lay.AddChild(input);
				lay.AddChild(label_v);
				left.AddChild(lay);
			}
			case 1:{
				lay := theme.CreateLinearLayout();
				lay.SetDirection(gxui.LeftToRight);
				lay.AddChild(label_q);
				lay.AddChild(input);			
				lay.AddChild(label_v);
				mid.AddChild(lay);
			}
			case 2:{
				lay := theme.CreateLinearLayout();
				lay.SetDirection(gxui.LeftToRight);
				lay.AddChild(label_q);
				lay.AddChild(input);			
				lay.AddChild(label_v);
				right.AddChild(lay);			
			}		
		}
	}

}