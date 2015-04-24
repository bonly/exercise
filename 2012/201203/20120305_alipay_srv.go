/*
支付平台的处理后，模拟其返回结果给请求方
*/
package main 

import (
  "net/http"
  //"net/url"
  "log"
  "strings"
  "io/ioutil"
  //"crypto/md5"
  //"encoding/hex"
  //"crypto/des"
  //"crypto/cipher"  
  //"os/exec"
  //"bytes"
)

/*
模拟支付宝直接发送
*/
func main(){
	cli := &http.Client{};

  //www := "http://192.168.211.234:80/PHP-UTF-8/notify_url.php";
  www := "http://183.61.112.5:8091/zhifubao";
  
  req, err := http.NewRequest("POST",www, 
    //支付完成
    strings.NewReader("discount=0.00&payment_type=1&subject=%E6%94%AF%E4%BB%98%E5%AE%9D%E5%BF%AB%E6%8D%B7%E5%85%85%E5%80%BC&trade_no=2014062531083193&buyer_email=18673111967&gmt_create=2014-06-25+11%3A41%3A51&notify_type=trade_status_sync&quantity=1&out_trade_no=192&seller_id=2088411388671371&notify_time=2014-06-25+11%3A41%3A51&body=%E6%94%AF%E4%BB%98%E5%AE%9D%E5%BF%AB%E6%8D%B7&trade_status=TRADE_FINISHED&is_total_fee_adjust=N&total_fee=0.01&gmt_payment=2014-06-25+11%3A41%3A51&seller_email=mdsw%40game83.com&gmt_close=2014-06-25+11%3A41%3A51&price=0.01&buyer_id=2088902627757935&notify_id=344f97a346ff49a07b56185390787ca176&use_coupon=N&sign_type=RSA&sign=fzpTvexpzPNt8iTOhxDd8Iqtvz6YdefSNkBhPUSvEImgr70hCE54x5Z0qltdZouPnSnVl0uNol7SFsJ%2BT3GUT7fqbWq58Sj3L%2F1VEnGKVrmxeWt97y1HXJYOVnBvGeaBmtu4nlvk1jLZiVWhq7eyJT2urM9WTSlv6pNU%2F8rwM94%3D"));
    //等待支付
    //strings.NewReader(`discount=0.00&payment_type=1&subject=%E6%94%AF%E4%BB%98%E5%AE%9D%E5%BF%AB%E6%8D%B7%E5%85%85%E5%80%BC&trade_no=2014062531083193&buyer_email=18673111967&gmt_create=2014-06-25+11%3A41%3A51&notify_type=trade_status_sync&quantity=1&out_trade_no=192&seller_id=2088411388671371&notify_time=2014-06-25+11%3A41%3A51&body=%E6%94%AF%E4%BB%98%E5%AE%9D%E5%BF%AB%E6%8D%B7&trade_status=WAIT_BUYER_PAY&is_total_fee_adjust=Y&total_fee=0.01&seller_email=mdsw%40game83.com&price=0.01&buyer_id=2088902627757935&notify_id=fb31507b3e0fb032b1c31a251f6cfff376&use_coupon=N&sign_type=RSA&sign=ht5JxaRyABUVNgnBmPsz%2FkTQQIJSVpFoDGQnTGvZ2DtXOXiNWAPl6qyxpXJ1O1QZ%2FnH3yZ4v9z4KJ0LiUC1GBLBoYfwpJONuWPHHTexd0mbk6f2skBMIpBTYxIjUZA1A4%2FgeHKQJoJd9kdz9vVtxz8XsCcFtUCOH6UDT1AH%2BAy8%3D`));
  
  resp, err := cli.Do(req);

  if err != nil {
     log.Println("handle error");
  }

  body, err := ioutil.ReadAll(resp.Body);

	log.Println("res: ", string(body));
}
