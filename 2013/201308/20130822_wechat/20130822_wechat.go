
package main

import (
	"crypto/sha1"
	"encoding/xml"
	"encoding/json"
	"net/url"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strings"
	"time"
	"html/template"
	// "os"
	// "bytes"
)
 
const (
	token = "wechat4xbed";
	tuling = "0fad39aba6e58fb7fdc7fc9d7c52fc27";
	appid = "wx100949d5a719dac2";
	secret = "10b624d9c9faa83165edfc1d4a336935";	
)

type TextRequestBody struct {
	XMLName      xml.Name `xml:"xml"`;
	ToUserName   string;
	FromUserName string;
	CreateTime   time.Duration;
	MsgType      string;
	Content      string;
	MsgId        int;
};

type TextResponseBody struct {
	XMLName      xml.Name `xml:"xml"`;
	ToUserName   CDATAText;
	FromUserName CDATAText;
	CreateTime   time.Duration;
	MsgType      CDATAText;
	Content      CDATAText;
};

/*
type CDATAText struct {
	Text []byte `xml:",innerxml"`
}
*/

type Tuling struct{
	Code int64 `json:"code"`;
	Text string `json:"text"`;
	Url  string `json:"url"`;
};


type News struct{
	Article string `json:"article"`;
	Source string `json:"source"`;
	Detailurl string `json:"detailurl"`;
	Icon	string `json:"icon"`;
};

type TNews struct{
	Tuling,
	List []News;
};

type Train struct{
	Trainnum string `json:"trainnum"`;
	Start string `json:"start"`;
	Terminal string `json:"terminal"`;
	Starttime string `json:"starttime"`;
	Icon string `json:"icon"`;
};

type TTrain struct{
	Tuling,
	List []Train;
};

type Flight struct{
	Flight string `json:"flight"`;
	Route string `json:"route"`;
	
};

/*
200000:链接
302000:新闻
305000:列车
306000:航班
308000:菜谱
*/
type CDATAText struct {
	Text string `xml:",innerxml"`;
}

func askTuling(ask string)(ret string){
	resp, err := http.Get("http://www.tuling123.com/openapi/api?key=" + tuling + "&info=" + ask);
	if err != nil{
		log.Println("http get: ", err);
	}
	defer resp.Body.Close();
	
	body, err := ioutil.ReadAll(resp.Body);
	if err != nil{
		log.Println("body: ", err);
	}
	//log.Println(string(body));
	
	var tuling Tuling; 
	if err := json.Unmarshal(body,&tuling); err != nil{
		log.Println("tuling: ", err);
	}
	//log.Println(tuling);
	
	switch(tuling.Code){
		case 200000:
		   return tuling.Text + "\n" + tuling.Url;
		case 302000:
		   var news TNews;
		   if err := json.Unmarshal(body, &news); err != nil{
			   log.Println(err);
		   }
		   for _, v := range news.List{
		      ret += "\n" + v.Article + "\n" + v.Detailurl;
		   }
		   return tuling.Text + ret;
		default:
		   return tuling.Text;
	}
}

func makeSignature(timestamp, nonce string) string {
	sl := []string{token, timestamp, nonce};
	sort.Strings(sl);
	s := sha1.New();
	io.WriteString(s, strings.Join(sl, ""));
	return fmt.Sprintf("%x", s.Sum(nil));
}

func validateUrl(w http.ResponseWriter, r *http.Request) bool {
	timestamp := strings.Join(r.Form["timestamp"], "");
	nonce := strings.Join(r.Form["nonce"], "");
	signatureGen := makeSignature(timestamp, nonce);

	signatureIn := strings.Join(r.Form["signature"], "");
	if signatureGen != signatureIn {
		return false;
	}
	echostr := strings.Join(r.Form["echostr"], "");
	fmt.Fprintf(w, echostr);
	return true;
}

func parseTextRequestBody(r *http.Request) *TextRequestBody {
	body, err := ioutil.ReadAll(r.Body);
	if err != nil {
		log.Fatal(err);
		return nil;
	}
	fmt.Println(string(body));
	requestBody := &TextRequestBody{};
	xml.Unmarshal(body, requestBody);
	return requestBody;
}

func value2CDATA(v string) CDATAText {
	//return CDATAText{[]byte("<![CDATA[" + v + "]]>")}
	return CDATAText{"<![CDATA[" + v + "]]>"};
}

func makeTextResponseBody(fromUserName, toUserName, content string) ([]byte, error) {
	textResponseBody := &TextResponseBody{};
	textResponseBody.FromUserName = value2CDATA(fromUserName);
	textResponseBody.ToUserName = value2CDATA(toUserName);
	textResponseBody.MsgType = value2CDATA("text");
	
	textResponseBody.Content = value2CDATA(content);
	textResponseBody.CreateTime = time.Duration(time.Now().Unix());
	return xml.MarshalIndent(textResponseBody, " ", "  ");
}

func procRequest(w http.ResponseWriter, r *http.Request) {
	r.ParseForm();
	if !validateUrl(w, r) {
		log.Println("Wechat Service: this http request is not from Wechat platform!");
		return;
	}

	if r.Method == "POST" {
		textRequestBody := parseTextRequestBody(r);
		if textRequestBody != nil {
			fmt.Printf("Wechat Service: Recv text msg [%s] from user [%s]!\n",
				textRequestBody.Content,
				textRequestBody.FromUserName);
			var str_ret string;
			if textRequestBody.Content==""{
				str_ret = "Hi，终于等到你了。XBED是国内创新精品互联网酒店公寓。没有前台的自助入住办理服务，充分保护你的隐私。不同房间的个性化的设计，值得细细品味和感受。高品质的配套和服务，为你创造最佳的商旅睡眠空间。XBED，给你点不同的！";
			}else{
				str_ret = askTuling(textRequestBody.Content);
			}
			responseTextBody, err := makeTextResponseBody(textRequestBody.ToUserName,
				textRequestBody.FromUserName,
				str_ret);
			if err != nil {
				log.Println("Wechat Service: makeTextResponseBody error: ", err);
				return;
			}
	
			w.Header().Set("Content-Type", "text/xml");
			fmt.Println(str_ret);
			fmt.Fprintf(w, string(responseTextBody));
		}
	}
}


type Mail struct {
    UserName string
    SiteName string
    ActiveLink string
}

type Data struct{
	Mtoker string;
};

func srvRequest(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL);
    uri, _ := url.Parse(r.URL.String());
	vl := uri.Query();

	code := vl.Get("code");

	tk := get_openid(code);
	// fmt.Fprintf(w, get_openid(code));

	ok := Data{tk};
    // tmp := template.New("need_clean"); //创建一个模板
    tmp, err := template.ParseFiles("html/need_clean.html");  //解析模板文件
    if err != nil{
    	fmt.Println(err);
    }
    // tmp.ExecuteTemplate(w, "header", nil);
    // user := GetUser() //获取当前用户信息
    tmp.Execute(w , ok)  //执行模板的merger操作	
    // fmt.Fprintf(w, tmp);
}

func get_openid(cd string)(ret string){
	val := url.Values{};
    val.Add("grant_type", "authorization_code");
	val.Add("code", cd);
	val.Add("appid", appid);
	val.Add("secret", secret);
	addr := "https://api.weixin.qq.com/sns/oauth2/access_token?"+val.Encode();

    resp, err := http.Get(addr);
    if err != nil {
        log.Println("http get: ",err);
    }
    defer resp.Body.Close();
    
	body, err := ioutil.ReadAll(resp.Body);
    if err != nil {
        log.Println("body: ", err);
    }
 	fmt.Println("send: ", addr);
    fmt.Println("recv: ", string(body));	
    return string(body);
}

func main() {
	log.Println("Wechat Service: Start!");
	http.Handle("/test",http.StripPrefix("/test",http.FileServer(http.Dir("."))));
	http.Handle("/img", http.FileServer(http.Dir("/html/img")));	
	http.HandleFunc("/", procRequest);
	http.HandleFunc("/wechat", srvRequest);
	err := http.ListenAndServe(":8090", nil);
	if err != nil {
		log.Fatal("Wechat Service: ListenAndServe failed, ", err);
	}
	log.Println("Wechat Service: Stop!");
}
