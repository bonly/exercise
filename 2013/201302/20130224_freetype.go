package main
import (
	"code.google.com/p/freetype-go/freetype"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"log"
	"os"
)
const (
	dx       = 100         // 图片的大小 宽度
	dy       = 40          // 图片的大小 高度
	//fontFile = "/home/bonly/.fonts/simsun.ttc" // 需要使用的字体文件
	//fontFile = "/home/bonly/src/20130222_touch/NotoSansHans-Regular.ttf" // 需要使用的字体文件
	fontFile = "/home/bonly/src/20130222_touch/NotoSansHans-Regular.otf" // 需要使用的字体文件
	fontSize = 20          // 字体尺寸
	fontDPI  = 72          // 屏幕每英寸的分辨率
)
func main() {
	// 需要保存的文件
	imgcounter := 123
	imgfile, _ := os.Create(fmt.Sprintf("%03d.png", imgcounter))
	defer imgfile.Close()
	// 新建一个 指定大小的 RGBA位图
	img := image.NewNRGBA(image.Rect(0, 0, dx, dy))
	// 画背景
	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			// 设置某个点的颜色，依次是 RGBA
			img.Set(x, y, color.RGBA{uint8(x), uint8(y), 0, 255})
		}
	}
	// 读字体数据
	fontBytes, err := ioutil.ReadFile(fontFile)
	if err != nil {
		log.Println(err)
		return
	}
	font, err := freetype.ParseFont(fontBytes)
	if err != nil {
		log.Println(err)
		return
	}
	c := freetype.NewContext()
	c.SetDPI(fontDPI)
	c.SetFont(font)
	c.SetFontSize(fontSize)
	c.SetClip(img.Bounds())
	c.SetDst(img)
	c.SetSrc(image.White)
	pt := freetype.Pt(10, 10+int(c.PointToFix32(fontSize)>>8)) // 字出现的位置
	//_, err = c.DrawString("ABCDE", pt)
	_, err = c.DrawString("何睿湜", pt)
	if err != nil {
		log.Println(err)
		return
	}
	// 以PNG格式保存文件
	err = png.Encode(imgfile, img)
	if err != nil {
		log.Fatal(err)
	}
}

