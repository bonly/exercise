<?

  		///////生成通过des加密后的cardinfo，并进行base64加密
     	 function GetDesCardInfo($cardmoney,$cardnum,$cardpwd,$deskey)
    	 {
         	 $str=$cardmoney."@".$cardnum."@".$cardpwd;	
         	 $size = mcrypt_get_block_size('des', 'ecb'); 
         	 $input = pkcs5_pad($str, $size);
         	
         	 $td = mcrypt_module_open(MCRYPT_DES,'','ecb',''); //使用MCRYPT_DES算法,ecb模式   
             $iv = mcrypt_create_iv(mcrypt_enc_get_iv_size($td), MCRYPT_RAND);   
             $ks = mcrypt_enc_get_key_size($td);   
             $key=base64_decode($deskey);
             mcrypt_generic_init($td, $key, $iv); //初始处理   
             //加密   
             $encrypted_data = mcrypt_generic($td, $input);   
       
             //结束处理   
             mcrypt_generic_deinit($td);   
             mcrypt_module_close($td); 
             /////作base64的编码     
             $encode = base64_encode($encrypted_data); 
             return $encode; 
         }
         

         function pkcs5_pad ($text, $blocksize)
         {    	
          	 $pad = $blocksize - (strlen($text) % $blocksize);
          	 return $text . str_repeat(chr($pad), $pad);	
         }
         //测试
         //echo GetDesCardInfo("50","1111","2222","123456");
         if ($argc > 5) { 
            print_r($argv);
            echo GetDesCardInfo($argv[1], $argv[2], $argv[3], $argv[4]);
            die("\nUsage: ". $argv[0] ." <cardmoney> <cardnum> <cardpwd> <deskey>\n");
         }
         echo GetDesCardInfo($argv[1], $argv[2], $argv[3], $argv[4]);

?>