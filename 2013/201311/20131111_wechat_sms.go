package main
 
import (
    "crypto/md5"
    "encoding/hex"
    "encoding/json"
    "fmt"
    "io"
    "io/ioutil"
    "log"
    "net/http"
    "net/url"
    "regexp"
    "strings"
)
 
type WebWeChat struct {
    token   string
    cookies []*http.Cookie
}
 
type MsgItem struct {
    FakeId   string `json:"fakeId"`
    NickName string `json:"nickName"`
    DateTime string `json:"dateTime"`
    Content  string `json:"content"`
}
 
func NewWebWeChat() *WebWeChat {
    w := new(WebWeChat)
    w.login()
    return w
}
 
func (w *WebWeChat) login() {
    login_url := "http://mp.weixin.qq.com/cgi-bin/login?lang=en_US"
    email := "xxxxx@gmail.com"
    password := "xxxxx"
    h := md5.New()
    h.Write([]byte(password))
    password = hex.EncodeToString(h.Sum(nil))
    post_arg := url.Values{"username": {email}, "pwd": {password}, "imgcode": {""}, "f": {"json"}}
 
    req, err := http.NewRequest("POST", login_url, strings.NewReader(post_arg.Encode()))
 
    if err != nil {
        log.Fatal(err)
    }
 
    client := new(http.Client)
    resp, _ := client.Do(req)
    data, _ := ioutil.ReadAll(resp.Body)
 
    s := string(data)
    fmt.Printf("%s", s)
 
    doc := json.NewDecoder(strings.NewReader(s))
 
    type Msg struct {
        Ret                     int
        ErrMsg                  string
        ShowVerifyCode, ErrCode int
    }
 
    var m Msg
    if err := doc.Decode(&m); err == io.EOF {
        fmt.Println(err)
    } else if err != nil {
        log.Fatal(err)
    }
 
    w.token = strings.Split(m.ErrMsg, "=")[3]
 
    fmt.Printf("%v\n", w.token)
 
    w.cookies = resp.Cookies()
}
 
func (w *WebWeChat) SendTextMsg(fakeid string, content string) bool {
    send_url := "http://mp.weixin.qq.com/cgi-bin/singlesend?t=ajax-response"
    referer_url := "https://mp.weixin.qq.com/cgi-bin/singlemsgpage?token=%s&fromfakeid=%s&msgid=&source=&count=20&t=wxm-singlechat&lang=zh_CN"
 
    post_arg := url.Values{
        "tofakeid": {fakeid},
        "type":     {"1"},
        "content":  {content},
        "ajax":     {"1"},
        "token":    {w.token},
    }
 
    req, _ := http.NewRequest("POST", send_url, strings.NewReader(post_arg.Encode()))
 
    req.Header.Set("Referer", fmt.Sprintf(referer_url, w.token, fakeid))
 
    for i := range w.cookies {
        req.AddCookie(w.cookies[i])
    }
 
    client := new(http.Client)
    resp, _ := client.Do(req)
    data, _ := ioutil.ReadAll(resp.Body)
 
    doc := json.NewDecoder(strings.NewReader(string(data)))
 
    type Msg struct {
        Ret string
        Msg string
    }
 
    var m Msg
    if err := doc.Decode(&m); err == io.EOF {
        fmt.Println(err)
    } else if err != nil {
        log.Fatal(err)
    }
    fmt.Println(m.Msg)
 
    if m.Msg == "ok" {
        return true
    } else {
        return false
    }
}
 
func (w *WebWeChat) GetFakeId() []MsgItem {
    msg_url := "https://mp.weixin.qq.com/cgi-bin/getmessage?t=wxm-message&token=%s&lang=zh_CN&count=50"
    referer_url := "https://mp.weixin.qq.com/cgi-bin/indexpage?t=wxm-index&token=%s&lang=zh_CN"
 
    req, _ := http.NewRequest("GET", fmt.Sprintf(msg_url, w.token), nil)
 
    req.Header.Set("Referer", fmt.Sprintf(referer_url, w.token))
 
    for i := range w.cookies {
        req.AddCookie(w.cookies[i])
    }
 
    client := new(http.Client)
    resp, _ := client.Do(req)
    data, _ := ioutil.ReadAll(resp.Body)
 
    re := regexp.MustCompile(`(?s)(?U)<script type="json" id="json-msgList">.+</script>`)
    list := re.FindString(string(data))
    list = strings.Replace(list, `<script type="json" id="json-msgList">`, "", -1)
    list = strings.Replace(list, `</script>`, "", -1)
    list = strings.Replace(list, `&nbsp;`, " ", -1)
 
    list = `{"MsgList":` + list + "}"
 
    type Msgs struct {
        MsgList []MsgItem
    }
 
    var m Msgs
 
    err := json.Unmarshal([]byte(list), &m)
 
    if err != nil {
        log.Println(err)
    }
 
    return m.MsgList
}
 
// // type MSG int8
 
func main() {
    // fakeid := "125638495"
 
    wechat := NewWebWeChat()
 
    log.Println(wechat.GetFakeId())
 
    // wechat.SendTextMsg(fakeid, "Hello.")
}