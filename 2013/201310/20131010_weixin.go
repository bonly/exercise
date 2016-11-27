// Weixin MP SDK (Golang)
package weixin

import (
	"bytes"
	"crypto/sha1"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"time"
)

const (
	// Event type
	msgEvent         = "event"
	EventSubscribe   = "subscribe"
	EventUnsubscribe = "unsubscribe"
	EventScan        = "SCAN"
	EventView        = "VIEW"
	EventClick       = "CLICK"
	EventLocation    = "LOCATION"

	// Message type
	MsgTypeDefault          = ".*"
	MsgTypeText             = "text"
	MsgTypeImage            = "image"
	MsgTypeVoice            = "voice"
	MsgTypeVideo            = "video"
	MsgTypeShortVideo       = "shortvideo"
	MsgTypeLocation         = "location"
	MsgTypeLink             = "link"
	MsgTypeEvent            = msgEvent + ".*"
	MsgTypeEventSubscribe   = msgEvent + "\\." + EventSubscribe
	MsgTypeEventUnsubscribe = msgEvent + "\\." + EventUnsubscribe
	MsgTypeEventScan        = msgEvent + "\\." + EventScan
	MsgTypeEventView        = msgEvent + "\\." + EventView
	MsgTypeEventClick       = msgEvent + "\\." + EventClick
	MsgTypeEventLocation    = msgEvent + "\\." + EventLocation

	// Media type
	MediaTypeImage = "image"
	MediaTypeVoice = "voice"
	MediaTypeVideo = "video"
	MediaTypeThumb = "thumb"
	// Button type
	MenuButtonTypeKey             = "click"
	MenuButtonTypeUrl             = "view"
	MenuButtonTypeScancodePush    = "scancode_push"
	MenuButtonTypeScancodeWaitmsg = "scancode_waitmsg"
	MenuButtonTypePicSysphoto     = "pic_sysphoto"
	MenuButtonTypePicPhotoOrAlbum = "pic_photo_or_album"
	MenuButtonTypePicWeixin       = "pic_weixin"
	MenuButtonTypeLocationSelect  = "location_select"
	MenuButtonTypeMediaId         = "media_id"
	MenuButtonTypeViewLimited     = "view_limited"
	// Weixin host URL
	weixinHost        = "https://api.weixin.qq.com/cgi-bin"
	weixinQRScene     = "https://api.weixin.qq.com/cgi-bin/qrcode"
	weixinShowQRScene = "https://mp.weixin.qq.com/cgi-bin/showqrcode"
	weixinShortURL    = "https://api.weixin.qq.com/cgi-bin/shorturl"
	weixinFileURL     = "http://file.api.weixin.qq.com/cgi-bin/media"
	// Max retry count
	retryMaxN = 3
	// Reply format
	replyText               = "<xml>%s<MsgType><![CDATA[text]]></MsgType><Content><![CDATA[%s]]></Content></xml>"
	replyImage              = "<xml>%s<MsgType><![CDATA[image]]></MsgType><Image><MediaId><![CDATA[%s]]></MediaId></Image></xml>"
	replyVoice              = "<xml>%s<MsgType><![CDATA[voice]]></MsgType><Voice><MediaId><![CDATA[%s]]></MediaId></Voice></xml>"
	replyVideo              = "<xml>%s<MsgType><![CDATA[video]]></MsgType><Video><MediaId><![CDATA[%s]]></MediaId><Title><![CDATA[%s]]></Title><Description><![CDATA[%s]]></Description></Video></xml>"
	replyMusic              = "<xml>%s<MsgType><![CDATA[music]]></MsgType><Music><Title><![CDATA[%s]]></Title><Description><![CDATA[%s]]></Description><MusicUrl><![CDATA[%s]]></MusicUrl><HQMusicUrl><![CDATA[%s]]></HQMusicUrl><ThumbMediaId><![CDATA[%s]]></ThumbMediaId></Music></xml>"
	replyNews               = "<xml>%s<MsgType><![CDATA[news]]></MsgType><ArticleCount>%d</ArticleCount><Articles>%s</Articles></xml>"
	replyHeader             = "<ToUserName><![CDATA[%s]]></ToUserName><FromUserName><![CDATA[%s]]></FromUserName><CreateTime>%d</CreateTime>"
	replyArticle            = "<item><Title><![CDATA[%s]]></Title> <Description><![CDATA[%s]]></Description><PicUrl><![CDATA[%s]]></PicUrl><Url><![CDATA[%s]]></Url></item>"
	transferCustomerService = "<xml>" + replyHeader + "<MsgType><![CDATA[transfer_customer_service]]></MsgType></xml>"

	// QR scene request
	requestQRScene      = `{"expire_seconds":%d,"action_name":"QR_SCENE","action_info":{"scene":{"scene_id":%d}}}`
	requestQRLimitScene = `{"action_name":"QR_LIMIT_SCENE","action_info":{"scene":{"scene_id":%d}}}`
)

// Common message header
type MessageHeader struct {
	ToUserName   string
	FromUserName string
	CreateTime   int
	MsgType      string
}

// Weixin request
type Request struct {
	MessageHeader
	MsgId        int64
	Content      string
	PicUrl       string
	MediaId      string
	Format       string
	ThumbMediaId string
	LocationX    float32 `xml:"Location_X"`
	LocationY    float32 `xml:"Location_Y"`
	Scale        float32
	Label        string
	Title        string
	Description  string
	Url          string
	Event        string
	EventKey     string
	Ticket       string
	Latitude     float32
	Longitude    float32
	Precision    float32
	Recognition  string
}

// Use to reply music message
type Music struct {
	Title        string `json:"title"`
	Description  string `json:"description"`
	MusicUrl     string `json:"musicurl"`
	HQMusicUrl   string `json:"hqmusicurl"`
	ThumbMediaId string `json:"thumb_media_id"`
}

// Use to reply news message
type Article struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	PicUrl      string `json:"picurl"`
	Url         string `json:"url"`
}

// Use to store QR code
type QRScene struct {
	Ticket        string `json:"ticket"`
	ExpireSeconds int    `json:"expire_seconds"`
}

// Custom Menu
type Menu struct {
	Buttons []MenuButton `json:"button,omitempty"`
}

type MenuButton struct {
	Name       string       `json:"name"`
	Type       string       `json:"type,omitempty"`
	Key        string       `json:"key,omitempty"`
	Url        string       `json:"url,omitempty"`
	MediaId    string       `json:"media_id,omitempty"`
	SubButtons []MenuButton `json:"sub_button,omitempty"`
}

// Use to output reply
type ResponseWriter interface {
	// Get weixin
	GetWeixin() *Weixin
	GetUserData() interface{}
	// Reply message
	ReplyOK()
	ReplyText(text string)
	ReplyImage(mediaId string)
	ReplyVoice(mediaId string)
	ReplyVideo(mediaId string, title string, description string)
	ReplyMusic(music *Music)
	ReplyNews(articles []Article)
	TransferCustomerService(serviceId string)
	// Post message
	PostText(text string) error
	PostImage(mediaId string) error
	PostVoice(mediaId string) error
	PostVideo(mediaId string, title string, description string) error
	PostMusic(music *Music) error
	PostNews(articles []Article) error
	// Media operator
	UploadMediaFromFile(mediaType string, filepath string) (string, error)
	DownloadMediaToFile(mediaId string, filepath string) error
	UploadMedia(mediaType string, filename string, reader io.Reader) (string, error)
	DownloadMedia(mediaId string, writer io.Writer) error
}

type responseWriter struct {
	wx           *Weixin
	writer       http.ResponseWriter
	toUserName   string
	fromUserName string
}

type response struct {
	ErrorCode    int    `json:"errcode,omitempty"`
	ErrorMessage string `json:"errmsg,omitempty"`
}

// Callback function
type HandlerFunc func(ResponseWriter, *Request)

type route struct {
	regex   *regexp.Regexp
	handler HandlerFunc
}

type accessToken struct {
	token   string
	expires time.Time
}

type Weixin struct {
	token     string
	routes    []*route
	tokenChan chan accessToken
	userData  interface{}
}

// Convert qr scene to url
func (qr *QRScene) ToURL() string {
	return (weixinShowQRScene + "?ticket=" + qr.Ticket)
}

// Create a Weixin instance
func New(token string, appid string, secret string) *Weixin {
	wx := &Weixin{}
	wx.token = token
	if len(appid) > 0 && len(secret) > 0 {
		wx.tokenChan = make(chan accessToken)
		go createAccessToken(wx.tokenChan, appid, secret)
	}
	return wx
}

func NewWithUserData(token string, appid string, secret string, userData interface{}) *Weixin {
	wx := New(token, appid, secret)
	wx.userData = userData
	return wx
}

// Register request callback.
func (wx *Weixin) HandleFunc(pattern string, handler HandlerFunc) {
	regex, err := regexp.Compile(pattern)
	if err != nil {
		panic(err)
		return
	}
	route := &route{regex, handler}
	wx.routes = append(wx.routes, route)
}

// Post text message
func (wx *Weixin) PostText(touser string, text string) error {
	var msg struct {
		ToUser  string `json:"touser"`
		MsgType string `json:"msgtype"`
		Text    struct {
			Content string `json:"content"`
		} `json:"text"`
	}
	msg.ToUser = touser
	msg.MsgType = "text"
	msg.Text.Content = text
	return postMessage(wx.tokenChan, &msg)
}

// Post image message
func (wx *Weixin) PostImage(touser string, mediaId string) error {
	var msg struct {
		ToUser  string `json:"touser"`
		MsgType string `json:"msgtype"`
		Image   struct {
			MediaId string `json:"media_id"`
		} `json:"image"`
	}
	msg.ToUser = touser
	msg.MsgType = "image"
	msg.Image.MediaId = mediaId
	return postMessage(wx.tokenChan, &msg)
}

// Post voice message
func (wx *Weixin) PostVoice(touser string, mediaId string) error {
	var msg struct {
		ToUser  string `json:"touser"`
		MsgType string `json:"msgtype"`
		Voice   struct {
			MediaId string `json:"media_id"`
		} `json:"voice"`
	}
	msg.ToUser = touser
	msg.MsgType = "voice"
	msg.Voice.MediaId = mediaId
	return postMessage(wx.tokenChan, &msg)
}

// Post video message
func (wx *Weixin) PostVideo(touser string, m string, t string, d string) error {
	var msg struct {
		ToUser  string `json:"touser"`
		MsgType string `json:"msgtype"`
		Video   struct {
			MediaId     string `json:"media_id"`
			Title       string `json:"title"`
			Description string `json:"description"`
		} `json:"video"`
	}
	msg.ToUser = touser
	msg.MsgType = "video"
	msg.Video.MediaId = m
	msg.Video.Title = t
	msg.Video.Description = d
	return postMessage(wx.tokenChan, &msg)
}

// Post music message
func (wx *Weixin) PostMusic(touser string, music *Music) error {
	var msg struct {
		ToUser  string `json:"touser"`
		MsgType string `json:"msgtype"`
		Music   *Music `json:"music"`
	}
	msg.ToUser = touser
	msg.MsgType = "video"
	msg.Music = music
	return postMessage(wx.tokenChan, &msg)
}

// Post news message
func (wx *Weixin) PostNews(touser string, articles []Article) error {
	var msg struct {
		ToUser  string `json:"touser"`
		MsgType string `json:"msgtype"`
		News    struct {
			Articles []Article `json:"articles"`
		} `json:"news"`
	}
	msg.ToUser = touser
	msg.MsgType = "news"
	msg.News.Articles = articles
	return postMessage(wx.tokenChan, &msg)
}

// Upload media from local file
func (wx *Weixin) UploadMediaFromFile(mediaType string, fp string) (string, error) {
	file, err := os.Open(fp)
	if err != nil {
		return "", err
	}
	defer file.Close()
	return wx.UploadMedia(mediaType, filepath.Base(fp), file)
}

// Download media and save to local file
func (wx *Weixin) DownloadMediaToFile(mediaId string, fp string) error {
	file, err := os.Create(fp)
	if err != nil {
		return err
	}
	defer file.Close()
	return wx.DownloadMedia(mediaId, file)
}

// Upload media with media
func (wx *Weixin) UploadMedia(mediaType string, filename string, reader io.Reader) (string, error) {
	return uploadMedia(wx.tokenChan, mediaType, filename, reader)
}

// Download media with media
func (wx *Weixin) DownloadMedia(mediaId string, writer io.Writer) error {
	return downloadMedia(wx.tokenChan, mediaId, writer)
}

// Get ip list
func (wx *Weixin) GetIpList() ([]string, error) {
	reply, err := sendGetRequest(weixinHost+"/getcallbackip?access_token=", wx.tokenChan)
	if err != nil {
		return nil, err
	}
	var result struct {
		IpList []string `json:"ip_list"`
	}
	if err := json.Unmarshal(reply, &result); err != nil {
		return nil, err
	}
	return result.IpList, nil
}

// Create QR scene
func (wx *Weixin) CreateQRScene(sceneId int, expires int) (*QRScene, error) {
	reply, err := postRequest(weixinQRScene+"/create?access_token=", wx.tokenChan, []byte(fmt.Sprintf(requestQRScene, expires, sceneId)))
	if err != nil {
		return nil, err
	}
	var qr QRScene
	if err := json.Unmarshal(reply, &qr); err != nil {
		return nil, err
	}
	return &qr, nil
}

// Create  QR limit scene
func (wx *Weixin) CreateQRLimitScene(sceneId int) (*QRScene, error) {
	reply, err := postRequest(weixinQRScene+"/create?access_token=", wx.tokenChan, []byte(fmt.Sprintf(requestQRLimitScene, sceneId)))
	if err != nil {
		return nil, err
	}
	var qr QRScene
	if err := json.Unmarshal(reply, &qr); err != nil {
		return nil, err
	}
	return &qr, nil
}

// Long url to short url
func (wx *Weixin) ShortURL(url string) (string, error) {
	var request struct {
		Action  string `json:"action"`
		LongUrl string `json:"long_url"`
	}
	request.Action = "long2short"
	request.LongUrl = url
	data, err := marshal(request)
	if err != nil {
		return "", err
	}
	reply, err := postRequest(weixinShortURL+"?access_token=", wx.tokenChan, data)
	if err != nil {
		return "", err
	}
	var shortUrl struct {
		Url string `json:"short_url"`
	}
	if err := json.Unmarshal(reply, &shortUrl); err != nil {
		return "", err
	}
	return shortUrl.Url, nil
}

// Custom menu
func (wx *Weixin) CreateMenu(menu *Menu) error {
	data, err := marshal(menu)
	if err != nil {
		return err
	}
	_, err = postRequest(weixinHost+"/menu/create?access_token=", wx.tokenChan, data)
	return err
}

func (wx *Weixin) GetMenu() (*Menu, error) {
	reply, err := sendGetRequest(weixinHost+"/menu/get?access_token=", wx.tokenChan)
	if err != nil {
		return nil, err
	}
	var result struct {
		MenuCtx *Menu `json:"menu"`
	}
	if err := json.Unmarshal(reply, &result); err != nil {
		return nil, err
	}
	return result.MenuCtx, nil
}

func (wx *Weixin) DeleteMenu() error {
	_, err := sendGetRequest(weixinHost+"/menu/delete?access_token=", wx.tokenChan)
	return err
}

// Create handler func
func (wx *Weixin) CreateHandlerFunc(w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		wx.ServeHTTP(w, r)
	}
}

// Process weixin request and send response.
func (wx *Weixin) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !checkSignature(wx.token, w, r) {
		http.Error(w, "", http.StatusUnauthorized)
		return
	}
	// Verify request
	if r.Method == "GET" {
		fmt.Fprintf(w, r.FormValue("echostr"))
		return
	}
	// Process message
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Weixin receive message failed:", err)
		http.Error(w, "", http.StatusBadRequest)
	} else {
		var msg Request
		if err := xml.Unmarshal(data, &msg); err != nil {
			log.Println("Weixin parse message failed:", err)
			http.Error(w, "", http.StatusBadRequest)
		} else {
			wx.routeRequest(w, &msg)
		}
	}
	return
}

func (wx *Weixin) routeRequest(w http.ResponseWriter, r *Request) {
	requestPath := r.MsgType
	if requestPath == msgEvent {
		requestPath += "." + r.Event
	}
	for _, route := range wx.routes {
		if !route.regex.MatchString(requestPath) {
			continue
		}
		writer := responseWriter{}
		writer.wx = wx
		writer.writer = w
		writer.toUserName = r.FromUserName
		writer.fromUserName = r.ToUserName
		route.handler(writer, r)
		return
	}
	http.Error(w, "", http.StatusNotFound)
	return
}

func marshal(v interface{}) ([]byte, error) {
	data, err := json.Marshal(v)
	if err == nil {
		data = bytes.Replace(data, []byte("\\u003c"), []byte("<"), -1)
		data = bytes.Replace(data, []byte("\\u003e"), []byte(">"), -1)
		data = bytes.Replace(data, []byte("\\u0026"), []byte("&"), -1)
	}
	return data, err
}

func checkSignature(t string, w http.ResponseWriter, r *http.Request) bool {
	r.ParseForm()
	var signature string = r.FormValue("signature")
	var timestamp string = r.FormValue("timestamp")
	var nonce string = r.FormValue("nonce")
	strs := sort.StringSlice{t, timestamp, nonce}
	sort.Strings(strs)
	var str string
	for _, s := range strs {
		str += s
	}
	h := sha1.New()
	h.Write([]byte(str))
	return fmt.Sprintf("%x", h.Sum(nil)) == signature
}

func authAccessToken(appid string, secret string) (string, time.Duration) {
	resp, err := http.Get(weixinHost + "/token?grant_type=client_credential&appid=" + appid + "&secret=" + secret)
	if err != nil {
		log.Println("Get access token failed: ", err)
	} else {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println("Read access token failed: ", err)
		} else {
			var res struct {
				AccessToken string `json:"access_token"`
				ExpiresIn   int64  `json:"expires_in"`
			}
			if err := json.Unmarshal(body, &res); err != nil {
				log.Println("Parse access token failed: ", err)
			} else {
				//log.Printf("AuthAccessToken token=%s expires_in=%d", res.AccessToken, res.ExpiresIn)
				return res.AccessToken, time.Duration(res.ExpiresIn * 1000 * 1000 * 1000)
			}
		}
	}
	return "", 0
}

func createAccessToken(c chan accessToken, appid string, secret string) {
	token := accessToken{"", time.Now()}
	c <- token
	for {
		if time.Since(token.expires).Seconds() >= 0 {
			var expires time.Duration
			token.token, expires = authAccessToken(appid, secret)
			token.expires = time.Now().Add(expires)
		}
		c <- token
	}
}

func sendGetRequest(reqURL string, c chan accessToken) ([]byte, error) {
	for i := 0; i < retryMaxN; i++ {
		token := <-c
		if time.Since(token.expires).Seconds() < 0 {
			r, err := http.Get(reqURL + token.token)
			if err != nil {
				return nil, err
			}
			defer r.Body.Close()
			reply, err := ioutil.ReadAll(r.Body)
			if err != nil {
				return nil, err
			}
			var result response
			if err := json.Unmarshal(reply, &result); err != nil {
				return nil, err
			}
			switch result.ErrorCode {
			case 0:
				return reply, nil
			case 42001: // access_token timeout and retry
				continue
			default:
				return nil, errors.New(fmt.Sprintf("WeiXin send get request reply[%d]: %s", result.ErrorCode, result.ErrorMessage))
			}
		}
	}
	return nil, errors.New("WeiXin post request too many times:" + reqURL)
}

func postRequest(reqURL string, c chan accessToken, data []byte) ([]byte, error) {
	for i := 0; i < retryMaxN; i++ {
		token := <-c
		if time.Since(token.expires).Seconds() < 0 {
			r, err := http.Post(reqURL+token.token, "application/json; charset=utf-8", bytes.NewReader(data))
			if err != nil {
				return nil, err
			}
			defer r.Body.Close()
			reply, err := ioutil.ReadAll(r.Body)
			if err != nil {
				return nil, err
			}
			var result response
			if err := json.Unmarshal(reply, &result); err != nil {
				return nil, err
			}
			switch result.ErrorCode {
			case 0:
				return reply, nil
			case 42001: // access_token timeout and retry
				continue
			default:
				return nil, errors.New(fmt.Sprintf("WeiXin send post request reply[%d]: %s", result.ErrorCode, result.ErrorMessage))
			}
		}
	}
	return nil, errors.New("WeiXin post request too many times:" + reqURL)
}

func postMessage(c chan accessToken, msg interface{}) error {
	data, err := marshal(msg)
	if err != nil {
		return err
	}
	_, err = postRequest(weixinHost+"/message/custom/send?access_token=", c, data)
	return err
}

func uploadMedia(c chan accessToken, mediaType string, filename string, reader io.Reader) (string, error) {
	reqURL := weixinFileURL + "/upload?type=" + mediaType + "&access_token="
	for i := 0; i < retryMaxN; i++ {
		token := <-c
		if time.Since(token.expires).Seconds() < 0 {
			bodyBuf := &bytes.Buffer{}
			bodyWriter := multipart.NewWriter(bodyBuf)
			fileWriter, err := bodyWriter.CreateFormFile("filename", filename)
			if err != nil {
				return "", err
			}
			if _, err = io.Copy(fileWriter, reader); err != nil {
				return "", err
			}
			contentType := bodyWriter.FormDataContentType()
			bodyWriter.Close()
			r, err := http.Post(reqURL+token.token, contentType, bodyBuf)
			if err != nil {
				return "", err
			}
			defer r.Body.Close()
			reply, err := ioutil.ReadAll(r.Body)
			if err != nil {
				return "", err
			}
			var result struct {
				response
				Type      string `json:"type"`
				MediaId   string `json:"media_id"`
				CreatedAt int64  `json:"created_at"`
			}
			err = json.Unmarshal(reply, &result)
			if err != nil {
				return "", err
			}
			switch result.ErrorCode {
			case 0:
				return result.MediaId, nil
			case 42001: // access_token timeout and retry
				continue
			default:
				return "", errors.New(fmt.Sprintf("WeiXin upload[%d]: %s", result.ErrorCode, result.ErrorMessage))
			}
		}
	}
	return "", errors.New("WeiXin upload media too many times")
}

func downloadMedia(c chan accessToken, mediaId string, writer io.Writer) error {
	reqURL := weixinFileURL + "/get?media_id=" + mediaId + "&access_token="
	for i := 0; i < retryMaxN; i++ {
		token := <-c
		if time.Since(token.expires).Seconds() < 0 {
			r, err := http.Get(reqURL + token.token)
			if err != nil {
				return err
			}
			defer r.Body.Close()
			if r.Header.Get("Content-Type") != "text/plain" {
				_, err := io.Copy(writer, r.Body)
				return err
			}
			reply, err := ioutil.ReadAll(r.Body)
			if err != nil {
				return err
			}
			var result response
			if err := json.Unmarshal(reply, &result); err != nil {
				return err
			}
			switch result.ErrorCode {
			case 0:
				return nil
			case 42001: // access_token timeout and retry
				continue
			default:
				return errors.New(fmt.Sprintf("WeiXin download[%d]: %s", result.ErrorCode, result.ErrorMessage))
			}
		}
	}
	return errors.New("WeiXin download media too many times")
}

// Format reply message header
func (w responseWriter) replyHeader() string {
	return fmt.Sprintf(replyHeader, w.toUserName, w.fromUserName, time.Now().Unix())
}

// Return weixin instance
func (w responseWriter) GetWeixin() *Weixin {
	return w.wx
}

// Return user data
func (w responseWriter) GetUserData() interface{} {
	return w.wx.userData
}

// Reply empty message
func (w responseWriter) ReplyOK() {
	w.writer.Write([]byte("success"))
}

// Reply text message
func (w responseWriter) ReplyText(text string) {
	msg := fmt.Sprintf(replyText, w.replyHeader(), text)
	w.writer.Write([]byte(msg))
}

// Reply image message
func (w responseWriter) ReplyImage(mediaId string) {
	msg := fmt.Sprintf(replyImage, w.replyHeader(), mediaId)
	w.writer.Write([]byte(msg))
}

// Reply voice message
func (w responseWriter) ReplyVoice(mediaId string) {
	msg := fmt.Sprintf(replyVoice, w.replyHeader(), mediaId)
	w.writer.Write([]byte(msg))
}

// Reply video message
func (w responseWriter) ReplyVideo(mediaId string, title string, description string) {
	msg := fmt.Sprintf(replyVideo, w.replyHeader(), mediaId, title, description)
	w.writer.Write([]byte(msg))
}

// Reply music message
func (w responseWriter) ReplyMusic(m *Music) {
	msg := fmt.Sprintf(replyMusic, w.replyHeader(), m.Title, m.Description, m.MusicUrl, m.HQMusicUrl, m.ThumbMediaId)
	w.writer.Write([]byte(msg))
}

// Reply news message (max 10 news)
func (w responseWriter) ReplyNews(articles []Article) {
	var ctx string
	for _, article := range articles {
		ctx += fmt.Sprintf(replyArticle, article.Title, article.Description, article.PicUrl, article.Url)
	}
	msg := fmt.Sprintf(replyNews, w.replyHeader(), len(articles), ctx)
	w.writer.Write([]byte(msg))
}

// Transfer customer service
func (w responseWriter) TransferCustomerService(serviceId string) {
	msg := fmt.Sprintf(transferCustomerService, serviceId, w.fromUserName, time.Now().Unix())
	w.writer.Write([]byte(msg))
}

// Post text message
func (w responseWriter) PostText(text string) error {
	return w.wx.PostText(w.toUserName, text)
}

// Post image message
func (w responseWriter) PostImage(mediaId string) error {
	return w.wx.PostImage(w.toUserName, mediaId)
}

// Post voice message
func (w responseWriter) PostVoice(mediaId string) error {
	return w.wx.PostVoice(w.toUserName, mediaId)
}

// Post video message
func (w responseWriter) PostVideo(mediaId string, title string, desc string) error {
	return w.wx.PostVideo(w.toUserName, mediaId, title, desc)
}

// Post music message
func (w responseWriter) PostMusic(music *Music) error {
	return w.wx.PostMusic(w.toUserName, music)
}

// Post news message
func (w responseWriter) PostNews(articles []Article) error {
	return w.wx.PostNews(w.toUserName, articles)
}

// Upload media from local file
func (w responseWriter) UploadMediaFromFile(mediaType string, filepath string) (string, error) {
	return w.wx.UploadMediaFromFile(mediaType, filepath)
}

// Download media and save to local file
func (w responseWriter) DownloadMediaToFile(mediaId string, filepath string) error {
	return w.wx.DownloadMediaToFile(mediaId, filepath)
}

// Upload media with reader
func (w responseWriter) UploadMedia(mediaType string, filename string, reader io.Reader) (string, error) {
	return w.wx.UploadMedia(mediaType, filename, reader)
}

// Download media with writer
func (w responseWriter) DownloadMedia(mediaId string, writer io.Writer) error {
	return w.wx.DownloadMedia(mediaId, writer)
}