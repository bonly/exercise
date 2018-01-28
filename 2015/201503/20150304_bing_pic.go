package main 

import (
"fmt"
"encoding/json"
"net/http"
"io/ioutil"
"image"
"image/jpeg"
"bytes"
"github.com/nfnt/resize"
"os"
"flag"
"math/rand"
"time"
)

var to = flag.String("f", "/usr/share/awesome/themes/default/background.png", "out put file");
var idx = flag.Int("i", 0, "-1 ~ 18 day");
var country = flag.String("r", "zh-cn",
	`zh-cn
	en-US
	zh-CN
	ja-JP
	en-AU
	de-DE
	en-NZ
	en-CA`);
var cnt = flag.Int("c", 100, "count of image");
var screen_y = flag.Uint("h", 800, "screen y");
var screen_x = flag.Uint("w", 1280, "screen_x");

var srv = "http://www.bing.com";
type Bing struct{
	Images []struct{
		Startdate string;
		Fullstartdate string;
		Enddate string;
		Url string;
		Urlbase string;
		Copyright string;
		Copyrightlink string;
		Wp bool;
		Hsh string;
		Drk int;
		Top int;
		Bot int;
		Hs []interface{};
	};
};

func init(){
	flag.Parse();
}

func Net_Get(addr string) (err error,ret []byte){
	resp, err := http.Get(addr);
	if err != nil{
		fmt.Println("http get: ", err);
		return err, ret;
	}
	defer resp.Body.Close();

	ret, err = ioutil.ReadAll(resp.Body);
	if err != nil{
		fmt.Println("body: ", err);
		return err, ret;
	}
	return nil, ret;
}

func Parse_Json(bit []byte) (err error, ret Bing){
	if err = json.Unmarshal(bit, &ret); err != nil{
		fmt.Println("json: ", err);
		return err, ret;
	}
	return nil, ret;
}

func Resize_Image(data []byte, wx uint, wh uint) (err error, ret image.Image){
	// Decoding gives you an Image.
	// If you have an io.Reader already, you can give that to Decode 
	// without reading it into a []byte.
	image, _, err := image.Decode(bytes.NewReader(data));
	if err != nil{
		fmt.Println("decode image: ", err);
		return err, ret;
	}

	ret = resize.Resize(wx, wh, image, resize.Lanczos3);
	return;
}

func main(){
	day := fmt.Sprintf("%d", *idx);
	sum	:= fmt.Sprintf("%d", *cnt);
	str := srv + "/HPImageArchive.aspx?format=js&idx="+ day +"&n=" + sum + "&mkt＝" + *country;
	fmt.Println(str);
	err, bit := Net_Get(str);
	if err != nil{
		return;
	}

	err, js := Parse_Json(bit);
	if err != nil {
		return;
	}

	// fmt.Printf("%v\n", js);

	s1 := rand.NewSource(time.Now().UnixNano());
    r1 := rand.New(s1);
	idx := r1.Intn(len(js.Images));
	
	err, bit = Net_Get(srv + js.Images[idx].Url);
	if err != nil{
		return;
	}

	// fmt.Println("pic:\n %s\n", string(bit));
	// ioutil.WriteFile("/tmp/abc.jpg", bit, 0777);

	err, img := Resize_Image(bit, *screen_x, *screen_y);
	if err != nil{
		return;
	}

	file, _ := os.Create(*to);
	defer file.Close();
	
	jpeg.Encode(file, img, &jpeg.Options{100});
}
