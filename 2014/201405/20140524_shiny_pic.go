package main

import (
	"fmt"
	"image"
	"log"
	"os"

	"golang.org/x/exp/shiny/driver"
	"golang.org/x/exp/shiny/screen"
	"golang.org/x/exp/shiny/widget"
	"golang.org/x/exp/shiny/widget/theme"
	"golang.org/x/exp/shiny/unit"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"image/color"

	_ "golang.org/x/image/bmp"
	_ "golang.org/x/image/tiff"
	_ "golang.org/x/image/webp"
)

var px = unit.Pixels
func decode(filename string) (image.Image, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	m, _, err := image.Decode(f)
	if err != nil {
		return nil, fmt.Errorf("could not decode %s: %v", filename, err)
	}
	return m, nil
}

func main() {
	log.SetFlags(0)
	driver.Main(func(s screen.Screen) {
		
		src, err := decode("a_png_file.png")
		if err != nil {
			log.Fatal(err)
		}
		m := widget.NewImage(src, src.Bounds());

		flow := widget.NewFlow(widget.AxisHorizontal);
		flow.Insert(m, nil); //加图片
		
		//widget.go中 t theme.Theme修改其数据,才能白色字
		/*
		t.Palette = &theme.Palette{
			Light:      &image.Uniform{C: color.RGBA{0xf5, 0xf5, 0xf5, 0xff}}, // Material Design "Grey 100".
			Neutral:    &image.Uniform{C: color.RGBA{0xee, 0xee, 0xee, 0xff}}, // Material Design "Grey 200".
			Dark:       &image.Uniform{C: color.RGBA{0xe0, 0xe0, 0xe0, 0xff}}, // Material Design "Grey 300".
			Accent:     &image.Uniform{C: color.RGBA{0x21, 0x96, 0xf3, 0xff}}, // Material Design "Blue 500".
			Foreground: &image.Uniform{C: color.RGBA{0xff, 0xff, 0xff, 0xff}}, // Material Design "Black".
			Background: &image.Uniform{C: color.RGBA{0xff, 0xff, 0xff, 0xff}}, // Material Design "White".
		};
		*/

		thm := theme.Default;
		// thm.GetPalette().Background = //没作用
		//       &image.Uniform{C: color.RGBA{0x21, 0x96, 0xf3, 0xff}};
		// flow.Layout(thm); //没作用

		// Make the widget node tree.
		hf := widget.NewFlow(widget.AxisHorizontal); //横向数据
		hf.Insert(widget.NewLabel("Cyan:"), nil);
		hf.Insert(widget.NewUniform(color.RGBA{0x00, 0x7f, 0x7f, 0xff}, px(0), px(20)),nil); //上色
		hf.LastChild.LayoutData = widget.FlowLayoutData{ExpandAlongWeight: 1}; //横向长度比重
		hf.Insert(widget.NewLabel("Magenta:"),nil);
		hf.Insert(widget.NewUniform(color.RGBA{0x7f, 0x00, 0x7f, 0xff}, px(0), px(30)), nil); //上色
		hf.LastChild.LayoutData = widget.FlowLayoutData{ExpandAlongWeight: 2}; //横向长度比重
		hf.Insert(widget.NewLabel("Yellow:"), nil);
		hf.Insert(widget.NewUniform(color.RGBA{0x7f, 0x7f, 0x00, 0xff}, px(0), px(40)), nil); //上色
		hf.LastChild.LayoutData = widget.FlowLayoutData{ExpandAlongWeight: 3}; //横向长度比重

		vf := widget.NewFlow(widget.AxisVertical); //竖向数据
		vf.Insert(widget.NewUniform(color.RGBA{0xff, 0x00, 0x00, 0xff}, px(80), px(40)), nil);
		vf.Insert(widget.NewUniform(color.RGBA{0x00, 0xff, 0x00, 0xff}, px(50), px(50)), nil);
		vf.Insert(widget.NewUniform(color.RGBA{0x00, 0x00, 0xff, 0xff}, px(20), px(60)), nil);
		vf.Insert(hf, nil); //加了前一个布局 一个子只能挂一个父
		vf.LastChild.LayoutData = widget.FlowLayoutData{ExpandAcross: true}; //扩展加入后面的数据
		vf.Insert(widget.NewLabel(fmt.Sprintf(
			"The black rectangle is 1.5 inches x 1 inch when viewed at %v DPI.", thm.GetDPI())), nil);
		vf.Insert(widget.NewUniform(color.White, unit.Inches(1.5), unit.Inches(1)), nil);
		
		flow.Insert(vf, nil);

		if err := widget.RunWindow(s, flow); err != nil {
			log.Fatal(err);
		}
	})
}
