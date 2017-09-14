package main

import (
	"fmt"
	"image"
	"log"
	"os"

	"golang.org/x/exp/shiny/driver"
	"golang.org/x/exp/shiny/screen"
	"golang.org/x/exp/shiny/widget"
	// "golang.org/x/exp/shiny/widget/theme"
	"golang.org/x/exp/shiny/unit"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	// "image/color"

	_ "golang.org/x/image/bmp"
	_ "golang.org/x/image/tiff"
	_ "golang.org/x/image/webp"
)

// TODO: scrolling, such as when images are larger than the window.
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
		
		flow := widget.NewFlow(widget.AxisHorizontal);
		src, err := decode("a_png_file.png")
		if err != nil {
			log.Fatal(err)
		}
		m := widget.NewImage(src, src.Bounds())
		flow.AppendChild(m); //加图片
		
		// thm := theme.Default;

		// Make the widget node tree.
		// hf := widget.NewFlow(widget.AxisHorizontal); //横向数据
		// hf.AppendChild(widget.NewLabel("Cyan:"))
		// hf.AppendChild(widget.NewUniform(color.RGBA{0x00, 0x7f, 0x7f, 0xff}, px(0), px(20))); //上色
		// hf.LastChild.LayoutData = widget.FlowLayoutData{ExpandAlongWeight: 1}; //横向长度比重
		// hf.AppendChild(widget.NewLabel("Magenta:"))
		// hf.AppendChild(widget.NewUniform(color.RGBA{0x7f, 0x00, 0x7f, 0xff}, px(0), px(30))) //上色
		// hf.LastChild.LayoutData = widget.FlowLayoutData{ExpandAlongWeight: 2} //横向长度比重
		// hf.AppendChild(widget.NewLabel("Yellow:"))
		// hf.AppendChild(widget.NewUniform(color.RGBA{0x7f, 0x7f, 0x00, 0xff}, px(0), px(40))) //上色
		// hf.LastChild.LayoutData = widget.FlowLayoutData{ExpandAlongWeight: 3} //横向长度比重

		// vf := widget.NewFlow(widget.AxisVertical) //竖向数据
		// vf.AppendChild(widget.NewUniform(color.RGBA{0xff, 0x00, 0x00, 0xff}, px(80), px(40)))
		// vf.AppendChild(widget.NewUniform(color.RGBA{0x00, 0xff, 0x00, 0xff}, px(50), px(50)))
		// vf.AppendChild(widget.NewUniform(color.RGBA{0x00, 0x00, 0xff, 0xff}, px(20), px(60)))
		// vf.AppendChild(hf) //加了前一个布局
		// vf.LastChild.LayoutData = widget.FlowLayoutData{ExpandAcross: true} //扩展加入后面的数据
		// vf.AppendChild(widget.NewLabel(fmt.Sprintf(
		// 	"The black rectangle is 1.5 inches x 1 inch when viewed at %v DPI.", thm.GetDPI())))
		// vf.AppendChild(widget.NewUniform(color.Black, unit.Inches(1.5), unit.Inches(1)))
		


		if err := widget.RunWindow(s, flow); err != nil {
			log.Fatal(err)
		}
	})
}
