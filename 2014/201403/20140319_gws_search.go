package main
import (
    "fmt"
    "strconv"
    "strings"
    "net/http"
    "time"
    "os"
    "crypto/tls"
)
var AllIP = ""
func ping_ip(host string, port string) bool {
    fmt.Println("正在测试："+host)
    var (
        remote = "https://"+host
    )
    tr := &http.Transport{
        TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
    }
    client := &http.Client{Transport: tr}
    resp,err := client.Get(remote)
    if err != nil {
        fmt.Println(err)
        return false
    }
    if resp.StatusCode == 200 {
        AllIP+=host+"|"
        fmt.Println("确定可用："+host)
    }
    resp.Body.Close()
    return true
}
func ip_to_num(ip string) int {
    canSplit := func(c rune) bool { return c == '.' }
    lisit := strings.FieldsFunc(ip, canSplit)
    //fmt.Println(lisit)
    ip1_str_int, _ := strconv.Atoi(lisit[0])
    ip2_str_int, _ := strconv.Atoi(lisit[1])
    ip3_str_int, _ := strconv.Atoi(lisit[2])
    ip4_str_int, _ := strconv.Atoi(lisit[3])
    return ip1_str_int<<24 | ip2_str_int<<16 | ip3_str_int<<8 | ip4_str_int
}
func num_to_ip(num int64) string {
    ip1_int := (num & 0xff000000) >> 24
    ip2_int := (num & 0x00ff0000) >> 16
    ip3_int := (num & 0x0000ff00) >> 8
    ip4_int := num & 0x000000ff
    //fmt.Println(ip1_int)
    data := fmt.Sprintf("%d.%d.%d.%d", ip1_int, ip2_int, ip3_int, ip4_int)
    return data
}
func main() {
    args := len(os.Args)
    x := "210.242.125.0"
    y := "210.242.125.255"
    if args > 2 {
        x = os.Args[1]
        y = os.Args[2]
        fmt.Println("开始测试自定义IP段")
    }else{
        fmt.Println("开始测试默认IP段")
    }
    ip1 := ip_to_num(x)
    ip2 := ip_to_num(y)
    for ip1 <= ip2 {
        ipint64 := int64(ip1)
        ip_data := num_to_ip(ipint64)
        go ping_ip(ip_data, "80")
        time.Sleep(50 * time.Millisecond)
        ip1++
    }
    fmt.Println("测试结束...等待自动退出")
    time.Sleep(5*time.Second)
    content := "[iplist]\r\n"
    content += "google_cn = "+AllIP+"\r\n"+"google_hk = "+AllIP
    file_name := "proxy.user.ini"
    fd,err := os.Create(file_name)
    defer fd.Close()
    if err != nil {
        fmt.Println(file_name,err)
    }
    fd.WriteString(content)
}
