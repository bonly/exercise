
package main

import (
	"crypto/sha1"
	"encoding/xml"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strings"
	"time"
)
 
const (
	token = "wechat4xbed";
	tuling = "0fad39aba6e58fb7fdc7fc9d7c52fc27";
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
			resp := askTuling(textRequestBody.Content);
			responseTextBody, err := makeTextResponseBody(textRequestBody.ToUserName,
				textRequestBody.FromUserName,
				resp);
			if err != nil {
				log.Println("Wechat Service: makeTextResponseBody error: ", err);
				return;
			}
			w.Header().Set("Content-Type", "text/xml");
			fmt.Println(string(responseTextBody));
			fmt.Fprintf(w, string(responseTextBody));
		}
	}
}

func main() {
	log.Println("Wechat Service: Start!");
	http.HandleFunc("/", procRequest);
	err := http.ListenAndServe(":80", nil);
	if err != nil {
		log.Fatal("Wechat Service: ListenAndServe failed, ", err);
	}
	log.Println("Wechat Service: Stop!");
}