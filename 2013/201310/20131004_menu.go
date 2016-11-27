/*
auth: bonly
create: 2015.9.15
*/

package main 

import (
	"github.com/google/gxui"
	"github.com/google/gxui/themes/dark"
	"github.com/google/gxui/drivers/gl"

	// "net/http"
	// "net/url"
	"io/ioutil"
	"fmt"
	"log"
	// "strings"
  	// "encoding/json"	
)

func main(){
	gl.StartDriver(appMain);
}

var font gxui.Font;
var window gxui.Window;
var theme gxui.Theme;


var ed gxui.CodeEditor;
var pk_file gxui.TextBox;

func appMain(driver gxui.Driver){
	theme = dark.CreateTheme(driver);
	window = theme.CreateWindow(800, 600, "微信公众号菜单修改");
	window.SetScale(1.0);

	loadFont(driver);

	window.OnClose(driver.Terminate);

	ui();
}

func loadFont(driver gxui.Driver){
	fontData, err := ioutil.ReadFile("./myfont.ttf");
	if err != nil{
		log.Fatalf("error reading font: %v", err);
	}
	font, err = driver.CreateFont(fontData, 20);
	if err != nil{
		panic(err);
	}
}

func ui(){
	backgroup := theme.CreateLinearLayout();
	backgroup.SetDirection(gxui.LeftToRight);
	window.AddChild(backgroup); //底层,左到右

	up := theme.CreateLinearLayout();
	up.SetDirection(gxui.TopToBottom);
	backgroup.AddChild(up); //左边，上到下

	down := theme.CreateLinearLayout();
	down.SetDirection(gxui.LeftToRight);
	backgroup.AddChild(down); //左边，上到下

	pk_file = theme.CreateTextBox();
	pk_file.SetText("wechat_dev.json");
	pk_file.OnLostFocus(load_file);

	ed = theme.CreateCodeEditor();
	ed.SetFont(font);
	
	up.AddChild(pk_file);
	down.AddChild(ed);

}

func load_file(){
	dat, err := ioutil.ReadFile(pk_file.Text());
	if err != nil{
		fmt.Println(err);
		return;
	}

	ed.SetDesiredWidth(800);
	ed.SetMultiline(true);
	ed.SetText(string(dat));	
}

