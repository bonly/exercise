package mydes;

import java.io.UnsupportedEncodingException;
import java.net.URLDecoder;
import java.net.URLEncoder;

import javax.crypto.Cipher;
import javax.crypto.spec.SecretKeySpec;

import org.apache.commons.codec.binary.Base64;


public class Aes {

	private static String key = "A1f@1l*g6EGjfnbf";

	public static String encrypt(String input){
		return encrypt(input, key);
	}

	public static String encrypt(String input, String key){
		byte[] crypted = null;
		try{
			SecretKeySpec skey = new SecretKeySpec(key.getBytes(), "AES");
			Cipher cipher = Cipher.getInstance("AES/ECB/PKCS5Padding");
			cipher.init(Cipher.ENCRYPT_MODE, skey);
			crypted = cipher.doFinal(input.getBytes());
		}catch(Exception e){
			System.out.println(e.toString());
		}
		return new String(Base64.encodeBase64(crypted));
	}


	public static String decrypt(String input){
		return decrypt(input, key);
	}

	public static String decrypt(String input, String key){
		byte[] output = null;
		try{
			SecretKeySpec skey = new SecretKeySpec(key.getBytes(), "AES");
			Cipher cipher = Cipher.getInstance("AES/ECB/PKCS5Padding");
			cipher.init(Cipher.DECRYPT_MODE, skey);
			output = cipher.doFinal(Base64.decodeBase64(input));
		}catch(Exception e){
			System.out.println(e.toString());
		}
		return new String(output);
	}

	public static void main(String[] args) throws UnsupportedEncodingException {
		if (args.length > 0) {
			String input = args[0];
			//String key = args[1];
			//System.out.println(Security.decrypt(URLDecoder.decode("T1WxvcRZuNozdc8%2FPWiC4KgLaVp%2F4K5jda3vvZtRGZc%3D", "utf-8")));
			//明文-》des->utf-8
			//明文-》des->utf-8
			//String strTest= "1234567890qwertyuiopasdfghjklzxcvbnm+=-|?><呵呵";
			//System.out.println(strTest);
			//String des = URLEncoder.encode(encrypt(strTest), "utf-8");
			//System.out.println(des);
			System.out.println(URLEncoder.encode(encrypt(input), "utf-8"));
		} else {
			System.out.println("wrong input!!!!!");
		}
	}	
}

