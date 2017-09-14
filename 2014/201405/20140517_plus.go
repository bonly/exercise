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

	//造表
	tab := make([]calc, 0, 121);
	for i:=10; i<20; i++{
		for j:=i; j<20; j++{ //前半部分已计算过
			tab = append(tab, calc{i, j, i*j});
		}
	}
	// fmt.Println(tab);
	for i:=0; i<size; i++{
		idx := rand.Int() % len(tab);
		ret = append(ret, tab[idx]);
		tab = append(tab[:idx], tab[idx+1:]...);//用过的从tab中去掉
	}		
	return ret;
}

func main() {
	rand.Seed(time.Now().UnixNano());

	dt = NewData();

	// Create and build a window
	win := gwu.NewWindow("plus", "10 至 20 之间的 乘法"); //路径名，窗口名
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

	server := gwu.NewServer("", "0.0.0.0:8081");//服务器网页根地址，侦听地址
	server.AddWin(win); //加入窗口
	server.Start("plus"); //服务起动，等待结束,当窗口路径名为空时，打开的是列表
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
	}, gwu.ETypeChange, gwu.ETypeBlur); //定义事件处理
	hpl.Add(value);
	hpl.Add(check);

	(*pl).Add(hpl);
}