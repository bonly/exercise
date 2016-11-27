package main

import (
	"github.com/google/gxui"
	"github.com/google/gxui/themes/dark"
	"github.com/google/gxui/drivers/gl"
	
	"log"
	"math/rand"
	"time"
	"math"
	"io/ioutil"
	"fmt"
)

type calc struct {
	X int;
	Y int;
	S int;
};

var dt []calc;
var font gxui.Font;
var window gxui.Window;
var theme gxui.Theme;
const size int = 40;


func addElem(lst []calc){
	backgroup := theme.CreateLinearLayout();
	backgroup.SetDirection(gxui.LeftToRight);
	window.AddChild(backgroup); //底层,左到右
	
	left := theme.CreateLinearLayout();
	left.SetDirection(gxui.TopToBottom);
	backgroup.AddChild(left); //左边，上到下
	
	// mid := theme.CreateLinearLayout();
	// mid.SetDirection(gxui.TopToBottom);
	// backgroup.AddChild(mid); //中间，上到下
	
	right := theme.CreateLinearLayout();
	right.SetDirection(gxui.TopToBottom);
	backgroup.AddChild(right); //右边，上到下
	
	for i, ca := range lst{
		label_q := theme.CreateLabel();
		var txt string;
		if ca.Y >=0 {
			txt = fmt.Sprintf("%6d + %6d = ", ca.X,
				int(math.Abs(float64(ca.Y))));
		}else{
			txt = fmt.Sprintf("%6d - %6d = ", ca.X,
				int(math.Abs(float64(ca.Y))));
		}
		label_q.SetText(txt);
		label_q.SetFont(font);
		
		label_v := theme.CreateLabel();
		value := fmt.Sprintf("%6d\t", ca.S); //结果标签
		label_v.SetText(value);
		label_v.SetFont(font);
		label_v.SetVisible(false);
		
		input := theme.CreateTextBox(); //结果输入框
		input.SetText("");
		input.SetFont(font);
		input.OnLostFocus(func(){
			label_v.SetVisible(true);
		});
		
		switch  i % 2{
			case 0:{
				lay := theme.CreateLinearLayout();
				lay.SetDirection(gxui.LeftToRight);
				lay.AddChild(label_q);
				lay.AddChild(input);
				lay.AddChild(label_v);
				left.AddChild(lay);
				break;
			}
			case 1:{
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

func newData(){
	dt = make([]calc, 0, size); 
	for i:=0; i<size; i++{
		x := rand.Int() % 1000000;
		y := rand.Int() % 1000000;
		if x > y{
			y = -y;
		}else{
			y = +y;
		}
		dt = append(dt, calc{x, y, x+y});
	}
}

func loadFont(driver gxui.Driver){
	fontData, err := ioutil.ReadFile("/home/bonly/.fonts/YaHei/YaHei.Consolas.1.11b.ttf");
	if err != nil{
		log.Fatalf("error reading font: %v", err);
	}
	font, err = driver.CreateFont(fontData, 20);
	if err != nil{
		panic(err);
	}
}

func appMain(driver gxui.Driver){
	theme = dark.CreateTheme(driver);
	window = theme.CreateWindow(850, 700, "多位数加减法");
	window.SetScale(1.0); //放大比率
	
	loadFont(driver);
	
	window.OnClose(driver.Terminate);
	
	//放入数据排版
	addElem(dt);
}

func main(){
	rand.Seed(time.Now().UnixNano());
	
	//生成数据
	newData();
	
	//显示
	gl.StartDriver(appMain);
}