/*
auth: bonly
create: 2015.11.3
*/

package main

/*
#include <jni.h>
#include <stdlib.h>
#include "jstring.h"
#cgo linux LDFLAGS: -L. -ljstring
*/
import "C"

import (
"net/url"
"strings"
"fmt"
"encoding/base64"
"crypto/sha1"
"crypto/hmac"
// "unsafe"
)


//export Java_Qunar_Encode
func Java_Qunar_Encode(env *C.JNIEnv, clazz C.jobject, data C.jstring, key C.jstring)(ret C.jstring){
	dat := C.jstringToChar(env, data);
	ky  := C.jstringToChar(env, key);

	// buf := C.malloc(50);
	// defer C.free(unsafe.Pointer(buf));

	val, err := Parse(C.GoString(dat));
	if err != nil{
		return C.charTojstring(env, C.CString(""));
	}
	rs := Sign(val, ([]byte)(C.GoString(ky)));
	return C.charTojstring(env, C.CString(rs));
}

func main(){
}

func Parse(str string)(val url.Values, err error){
	strs := strings.Split(str, "&");

	val = url.Values{};
	for _, str := range strs{
		st := strings.Split(str, "=");
		if len(st) == 2{
			val.Add(st[0], st[1]);
		}
	}
	return val, nil;
}

func Sign(val url.Values, key []byte)(string){
	// tn := time.Now();
	// loc, _ := time.LoadLocation("Local");
	// val.Set("timeStamp", tn.Format("2006-01-02 15:04:05"));

	org := val.Encode();
	// fmt.Printf("原数据: %s\n", org);
	sig := url.QueryEscape(
		base64.StdEncoding.EncodeToString(
			HmacSHA1(key, strings.ToLower(org))));
	return fmt.Sprintf("%s&signature=%s", org, sig);
}

func HmacSHA1(key []byte, data string) (ret []byte){
	//hmac ,use sha1
	mac := hmac.New(sha1.New, key);
	mac.Write([]byte(data));
	ret = mac.Sum(nil);
	return ret;
}

