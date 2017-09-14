<?php

//DES加密
class Des {
	var $key;
	var $iv; //偏移量
	
    /**
     * 设置加密秘钥
     * @param $key 密鑰（八個字元內）
     * @param $iv 偏移量
     */
	function Des($key, $iv=0)
	{
		$this->key = $key;
		if($iv == 0)
		{
			$this->iv = $key;
		}
		else
		{
			$this->iv = $iv;
		}
	}
    
    /**
    * PHP DES 加密程式
    * @param $encrypt 要加密的明文
    * @return string 密文
    */
    function encrypt ($encrypt)
    {
        // 根據 PKCS#7 RFC 5652 Cryptographic Message Syntax (CMS) 修正 Message 加入 Padding
        $block = mcrypt_get_block_size(MCRYPT_DES, MCRYPT_MODE_ECB);
        $pad = $block - (strlen($encrypt) % $block);
        $encrypt .= str_repeat(chr($pad), $pad);
        
        // 不需要設定 IV 進行加密
        $passcrypt = mcrypt_encrypt(MCRYPT_DES, $this->key, $encrypt, MCRYPT_MODE_ECB);
        return base64_encode($passcrypt);
    }
    
    
    /**
    * PHP DES 解密程式
    * @param $decrypt 要解密的密文
    * @return string 明文
    */
    function decrypt ($decrypt)
    {
        // 不需要設定 IV
        $str = mcrypt_decrypt(MCRYPT_DES, $this->key, base64_decode($decrypt), MCRYPT_MODE_ECB);
        
        // 根據 PKCS#7 RFC 5652 Cryptographic Message Syntax (CMS) 修正 Message 移除 Padding
        $pad = ord($str[strlen($str) - 1]);
        return substr($str, 0, strlen($str) - $pad);
    }
	
    
    /** 以下方法暂时不使用 */
	//加密
	function soureencrypt($str)
	{
		$size = mcrypt_get_block_size ( MCRYPT_DES, MCRYPT_MODE_CBC );
		$str = $this->pkcs5Pad ( $str, $size );
		$data=mcrypt_cbc(MCRYPT_DES, $this->key, $str, MCRYPT_ENCRYPT, $this->iv);
		//$data=strtoupper(bin2hex($data)); //返回大写十六进制字符串
		return base64_encode($data);
	}
	
	//解密
	function souredecrypt($str)
	{
		$str = base64_decode ($str);
		//$strBin = $this->hex2bin( strtolower($str));
		$str = mcrypt_cbc(MCRYPT_DES, $this->key, $str, MCRYPT_DECRYPT, $this->iv );
		$str = $this->pkcs5Unpad( $str );
		return $str;
	}
	
	function hex2bin($hexData)
	{
		$binData = "";
		for($i = 0; $i < strlen ( $hexData ); $i += 2)
		{
			$binData .= chr(hexdec(substr($hexData, $i, 2)));
		}
		return $binData;
	}
	
	function pkcs5Pad($text, $blocksize){
		$pad = $blocksize - (strlen ( $text ) % $blocksize);
		return $text . str_repeat ( chr ( $pad ), $pad );
	}
	
	function pkcs5Unpad($text)
	{
		$pad = ord ( $text {strlen ( $text ) - 1} );
		if ($pad > strlen ( $text ))
			return false;
		if (strspn ( $text, chr ( $pad ), strlen ( $text ) - $pad ) != $pad)
			return false;
		return substr ( $text, 0, - 1 * $pad );
	}
}

/**
 * 发送HTTP请求方法
 * @param  string $url    请求URL
 * @param  array  $params 请求参数
 * @param  string $method 请求方法GET/POST
 * @return array  $data   响应数据
 */
function http($url, $params, $method = 'GET', $header = array(), $multi = false){
    $opts = array(
            CURLOPT_TIMEOUT        => 30,
            CURLOPT_RETURNTRANSFER => 1,
            CURLOPT_SSL_VERIFYPEER => false,
            CURLOPT_SSL_VERIFYHOST => false,
            CURLOPT_HTTPHEADER     => $header
    );

    /* 根据请求类型设置特定参数 */
    switch(strtoupper($method)){
        case 'GET':
            $opts[CURLOPT_URL] = $url . '?' . http_build_query($params);
            break;
        case 'POST':
            //判断是否传输文件
            $params = $multi ? $params : http_build_query($params);
            $opts[CURLOPT_URL] = $url;
            $opts[CURLOPT_POST] = 1;
            $opts[CURLOPT_POSTFIELDS] = $params;
            break;
        default:
            throw new Exception('不支持的请求方式！');
    }

    /* 初始化并执行curl请求 */
    $ch = curl_init();
    curl_setopt_array($ch, $opts);
    $data  = curl_exec($ch);
    $error = curl_error($ch);
    curl_close($ch);
    if($error) throw new Exception('请求发生错误：' . $error);
    return  $data;
}


header("Content-type:text/html;charset=utf-8");
$appsecret = 'nXV3Xbhx'; //机密密钥

//验证图片和身份是否对应
// $name = 'xxx'; //验证人姓名
// $idnum = '421302xxxx'; //身份证号
// $filepath= 'face.jpg'; //头像路径
// $file_info = getimagesize($filepath);
// $fp  = fopen($filepath, 'rb', 0);
// $pic = "data:{$file_info['mime']};base64,".chunk_split(base64_encode(fread($fp,filesize($filepath))));
// fclose($fp);

// $param = array(
//     'name' => $name,
//     'idnum' => $idnum,
//     'pic' => $pic
// );
$Des = new Des($appsecret);
// $data = $Des->encrypt(json_encode($param));

// $url = 'http://120.25.160.52/joggle/identity/validate';
// $postdata = array(
//     'data' => $data
// );
// $result = http($url, $postdata, 'POST');
// $returndata = $Des->decrypt($result);
// echo $returndata;

//获取剩余的验证次数
$url = 'http://120.25.160.52/joggle/identity/residuenum';
$result = http($url, array());
$returndata = $Des->decrypt($result);
echo $returndata;