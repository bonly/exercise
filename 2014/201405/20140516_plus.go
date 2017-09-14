package main

import (
	"fmt"
	"github.com/icza/gowut/gwu"
	"math/rand"
	"time"
)

type calc struct{
	X int;
	Y int;
	V int;
};

var dt []calc;
const size int = 30;


func NewData()(ret []calc){
	ret = make([]calc, 0, size);
	for i:=0; i<size; i++{
		x := rand.Int() % 10 + 10;
		y := rand.Int() % 10 + 10;

		ret = append(ret, calc{x, y, x*y});
	}		
	return ret;
}

func main() {
	rand.Seed(time.Now().UnixNano());

	dt = NewData();

	// Create and build a window
	win := gwu.NewWindow("bonly", "bonly Window"); //路径名，窗口名
	// win.Style().SetFullWidth(); //全屏
	// win.SetHAlign(gwu.HACenter); //中对齐
	// win.SetCellPadding(20);

	pl := gwu.NewHorizontalPanel();
	// pl := gwu.NewPanel(); //创建panel,默认为LayoutVertical
	pl.SetCellPadding(20); //总panel

	vpl := gwu.NewVerticalPanel(); //竖向左边
	vpl.SetCellPadding(5); 

	vpr := gwu.NewVerticalPanel(); //竖向右边
	vpr.SetCellPadding(5); 

	pl.Add(vpl);
	pl.Add(vpr);

	for i, cl := range dt{
		if i % 2 == 0{
			add_obj(&vpl, cl);
		}else{
			add_obj(&vpr, cl);
		}
	}

	win.Add(pl);

	server := gwu.NewServer("", "localhost:8081");//服务器网页根地址，侦听地址
	server.AddWin(win); //加入窗口
	server.Start("bonly"); //服务起动，等待结束,当窗口路径名为空时，打开的是列表
}


func add_obj(pl *gwu.Panel, ti calc){
	hpl := gwu.NewHorizontalPanel(); //横向
	hpl.SetCellPadding(5); //每行之间的间隔

	str_title := fmt.Sprintf("%d x %d = ", ti.X, ti.Y);
	title := gwu.NewLabel(str_title);  //标签
	title.Style().SetWidth("200px"); //大小
	title.Style().SetColor(gwu.ClrRed); //颜色
	hpl.Add(title);

	cvalue := fmt.Sprintf("%d", ti.V);
	check := gwu.NewCheckBox(cvalue);
	check.Style().SetColor(gwu.ClrWhite);

	value := gwu.NewTextBox(""); //输入框
	value.AddSyncOnETypes(gwu.ETypeKeyUp); //增加响应事件类型
	value.AddEHandlerFunc(func(e gwu.Event) {
		if value.Text() == check.Text() {
			check.SetState(true);
			e.MarkDirty(check);
		}
	}, gwu.ETypeChange, gwu.ETypeKeyUp); //定义事件处理
	hpl.Add(value);
	hpl.Add(check);

	(*pl).Add(hpl);
}