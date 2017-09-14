package main 

import (
"net/http"
"log"
"io/ioutil"
"encoding/base64"
"encoding/json"
"crypto/des"
"bytes"
"errors"
"fmt"
"strings"
"net/url"
"github.com/lazywei/go-opencv/opencv"
"os"
"flag"
"golang.org/x/text/encoding/simplifiedchinese"
"golang.org/x/text/transform"
"path"
"runtime"
)

var Size = flag.Int("s", 600, "size of image");
var Id = flag.String("i", "440602197809192115", "id card");
var Name = flag.String("n", "何健波", "name");
var Mime = flag.String("m", "image/jpeg", "image mime type");
var Detect = flag.Bool("d", true, "detect face");
var Rate = flag.Int("r", 80, "compass rate for jpg");
var Pic = flag.String("f", "/tmp/test.jpg", "指定的测试图");

var srv_addr = "http://120.25.160.52";
var srv_key = "nXV3Xbhx";

type Data struct {
	Name string `json:"name"`;
	Idnum string `json:"idnum"`;
	Pic string `json:"pic"`;
};

type Ret struct{
	Code int `json:"code"`;
	Msg interface{} `json:"msg"`;
};

type R_residuenum struct{
	Num int `json:"num"`;
};

type R_validate struct{
	Id_num int `json:"id_num"`;
	Name int `json:"name"`;
	Validate_result int `json:"validate_result"`;
	Similarity int `json:"similarity"`;
	Err string `json:"err"`;
	Msg string `json:"msg"`;
	Serial_number string `json:"serial_number"`;
};

func residue_num(){
	resp, err := http.Get(srv_addr + "/joggle/identity/residuenum");
	if err != nil{
		log.Printf("http get: %v\n", err);
		return;
	}
	defer resp.Body.Close();

	body, err := ioutil.ReadAll(resp.Body);
	if err != nil{
		log.Printf("body: %v\n", err);
		return;
	} 

	log.Printf("recv org: %s", string(body));
	
	buf := make([]byte, 512);
	lng, err := base64.StdEncoding.Decode(buf, body); //先base64解码
	if err != nil{
		log.Printf("base64 decode: %v\n", err);
		return;
	}
	// log.Printf("len: %d %d\n", lng, len(string(buf[:lng])));

	ret, err := DesDecrypt([]byte(string(buf[:lng])), []byte(srv_key)); //再des解密
	if err != nil{
		log.Printf("des decode: %v\n", err);
		return;
	}
	log.Printf("recv: %+v", string(ret));

	var ret_data Ret;
	var ret_msg R_residuenum;
	ret_data.Msg = &ret_msg;
	// if err = json.Unmarshal([]byte(`{"code":0,"msg":{"num":5}}`), &ret_data); err != nil{
	if err = json.Unmarshal(ret, &ret_data); err != nil{
		log.Printf("json data decode: %v\n", err);
		return;
	}

	log.Println(fmt.Sprintf("code=%d num=%+v\n", ret_data.Code, ret_msg.Num));
}

func validate(filename string){
	data := Data{
		Name : *Name,
		Idnum : *Id,
	};

	jpg, err := ioutil.ReadFile(filename); //读图片
	if err != nil{
		log.Println("read jpg file: ", err);
		return;
	}
	data.Pic = fmt.Sprintf("data:%s;base64,%s", *Mime, base64.StdEncoding.EncodeToString(jpg)); //base64图片

	str_dat, err := json.Marshal(data); //json化
	if err != nil{
		log.Println("json encode: ", err);
		return;
	}
	log.Println("json: ", string(str_dat));

	encode_data, err := DesEncrypt([]byte(str_dat), []byte(srv_key));//加密
	if err != nil{
		log.Println("des Encrypt: ", err);
		return;
	}

	//数据需要base64再转换掉url规则的字符
	post_data := "data=" + url.QueryEscape(base64.StdEncoding.EncodeToString(encode_data));
	request, err := http.NewRequest("POST", srv_addr + "/joggle/identity/validate", strings.NewReader(post_data));
  	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:12.0) Gecko/20100101 Firefox/12.0");
  	request.Header.Set("Content-Type", "application/x-www-form-urlencoded");

  	client := &http.Client{};

  	resp, err := client.Do(request);
  	if err != nil{
  		log.Printf("POST: %v\n", err);
  		return;
  	}
  	defer resp.Body.Close();

  	body, err := ioutil.ReadAll(resp.Body);
  	if err != nil{
  		log.Printf("body: %v\n", err);
  		return;
  	}
  	log.Printf("recv org: %v\n", string(body));

	buf := make([]byte, 512);
	lng, err := base64.StdEncoding.Decode(buf, body); //先base64解码
	if err != nil{
		log.Printf("base64 decode: %v\n", err);
		return;
	}
	// log.Printf("len: %d %d\n", lng, len(string(buf[:lng])));

	ret, err := DesDecrypt([]byte(string(buf[:lng])), []byte(srv_key)); //再des解密
	if err != nil{
		log.Printf("des decode: %v\n", err);
		return;
	}
	log.Printf("recv: %+v", string(ret));  	

	var ret_data Ret;
	var ret_msg []R_validate;
	ret_data.Msg = &ret_msg;
	if err = json.Unmarshal(ret, &ret_data); err != nil{
		var msg string;
		ret_data.Msg = &msg;
		if err = json.Unmarshal(ret, &ret_data); err == nil{
			log.Println(fmt.Sprintf("ret=%d msg=%+v\n", ret_data.Code, msg));
		}else{
			log.Printf("json data decode: %v\n", err);
			return;
		}
	}

	if ret_data.Code == 0{
		log.Println(fmt.Sprintf("ret=%d result=%d msg=%+v\n", ret_data.Code, 
			ret_msg[0].Validate_result, ret_msg[0].Msg));
	}

	// Utf8ToGbk(ret.)
}

func main(){
	flag.Parse();
	
	win := opencv.NewWindow("xbed");
	defer win.Destroy();

	cam := opencv.NewCameraCapture(0);
	if cam == nil{
		fmt.Println("无法打开摄像头");
		return;
	}
	defer cam.Release();
	// cam.SetProperty(opencv.CV_CAP_PROP_FRAME_WIDTH, 1280);
	// cam.SetProperty(opencv.CV_CAP_PROP_FRAME_HEIGHT, 960);

	_, currentfile, _, _ := runtime.Caller(0);//取当前程序所在目录

	var img *opencv.IplImage = nil; 
	for {
		if cam.GrabFrame(){
			img = cam.RetrieveFrame(1);
			if img != nil{
				if *Detect==true{
					cascade := opencv.LoadHaarClassifierCascade( 
						path.Join(path.Dir(currentfile), 
							"20131023_haarcascade_frontalface_alt.xml"));

					faces := cascade.DetectObjects(img);

					for _, value := range faces{
						opencv.Rectangle(img,
							opencv.Point{value.X() + value.Width(), value.Y()},
							opencv.Point{value.X(), value.Y() + value.Height()},
							opencv.ScalarAll(255.0), 1, 1, 0)
					}
				}
				win.ShowImage(img);
			}
		}

		key := opencv.WaitKey(1); 
		switch{
			case key == 27: { //Esc
				os.Exit(0);
				break;
			}
			case key == 32: {//空格
				if img != nil{
					resize := opencv.Resize(img, *Size, 0, 0);
					filename := fmt.Sprintf("/tmp/test_%d.jpg", *Size);
					// opencv.SaveImage(filename, resize, opencv.CV_IMWRITE_JPEG_QUALITY);
					// opencv.SaveImage(filename, resize, opencv.CV_IMWRITE_PNG_COMPRESSION);
					params :=[]int32{int32(opencv.CV_IMWRITE_JPEG_QUALITY), int32(*Rate)};
					opencv.ExportImage(filename, resize, params);
					validate(filename);
					fmt.Printf("Size[%d] Rate[%d%%] file[%s]\n", *Size, *Rate, filename);
				}
				break;
			}
			case key == 'n': {
				residue_num();
				break;
			}
			case key == 's':{
				fmt.Printf("input new size [%d]: ", *Size);
				fmt.Scanf("%d", Size);
				fmt.Printf("size[%d] rate[%d%%]\n", *Size, *Rate);
				break;
			}
			case key == 'r':{
				fmt.Printf("input new compass rate [%d]: ", *Rate);
				fmt.Scanf("%d", Rate);
				fmt.Printf("size[%d] rate[%d%%]\n", *Size, *Rate);
			}
			case key == 'p':{
				filename := fmt.Sprintf("%s", *Pic);
				validate(filename);
				fmt.Printf("使用指定的测试文件[%s]\n", *Pic);
			}
		}
	}
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize;
	padtext := bytes.Repeat([]byte{byte(padding)}, padding);
	return append(ciphertext, padtext...);
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData);
	unpadding := int(origData[length-1]);
	return origData[:(length - unpadding)];
}

func ZeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize;
	padtext := bytes.Repeat([]byte{0}, padding);
	return append(ciphertext, padtext...);
}

func ZeroUnPadding(origData []byte) []byte {
	return bytes.TrimFunc(origData,
		func(r rune) bool {
			return r == rune(0);
		})
}

func DesEncrypt(src, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key);
	if err != nil {
		return nil, err;
	}
	bs := block.BlockSize();
	// src = ZeroPadding(src, bs);
	src = PKCS5Padding(src, bs);
	if len(src)%bs != 0 {
		return nil, errors.New("Need a multiple of the blocksize");
	}
	out := make([]byte, len(src));
	dst := out;
	for len(src) > 0 {
		block.Encrypt(dst, src[:bs]);
		src = src[bs:];
		dst = dst[bs:];
	}
	return out, nil;
}
func DesDecrypt(src, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key);
	if err != nil {
		return nil, err;
	}
	out := make([]byte, len(src));
	dst := out;
	bs := block.BlockSize();
	if len(src)%bs != 0 {
		return nil, errors.New("crypto/cipher: input not full blocks");
	}
	for len(src) > 0 {
		block.Decrypt(dst, src[:bs]);
		src = src[bs:];
		dst = dst[bs:];
	}
	// out = ZeroUnPadding(out);
	out = PKCS5UnPadding(out);
	return out, nil;
}

func GbkToUtf8(s []byte) ([]byte, error) {
    reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
    d, e := ioutil.ReadAll(reader)
    if e != nil {
        return nil, e
    }
    return d, nil
}

func Utf8ToGbk(s []byte) ([]byte, error) {
    reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewEncoder())
    d, e := ioutil.ReadAll(reader)
    if e != nil {
        return nil, e
    }
    return d, nil
}

/*
保存图像 imwrite
 bool imwrite(const string& filename, InputArray img, const vector& params=vector() )
params中的每个参数成对出现，即paramId_1, paramValue_1, paramId_2, paramValue_2, … ，当前支持如下参数：

JPEG：压缩质量 ( CV_IMWRITE_JPEG_QUALITY )，从0到100（数值越高质量越好），默认值为95。
PNG：compression level ( CV_IMWRITE_PNG_COMPRESSION ) 从0到9。 数值越高，文件大小越小，压缩时间越长。默认值为3。
PPM, PGM, or PBM：二进制标志 ( CV_IMWRITE_PXM_BINARY )，0 或 1。默认值为1。

还有一组函数，用于从内存读取数据和向内存写入数据。
*/

// func SaveImage(filename string, image *IplImage, params []int32) int {
// 	name_c := C.CString(filename);
// 	defer C.free(unsafe.Pointer(name_c));
// 	// params_c := C.int(params);
// 	// rv := C.cvSaveImage(name_c, unsafe.Pointer(image), &params_c)
// 	rv := C.cvSaveImage(name_c, unsafe.Pointer(image), (*C.int)(unsafe.Pointer(&params[0])));
// 	return int(rv);
// }

//http://www.imys.net/20150916/webapp-input-use-camera.html