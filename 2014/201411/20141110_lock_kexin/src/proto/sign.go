/*
auth: bonly
create: 2015.10.27
*/
package proto

import(
// "reflect"
"fmt"
// "sort"
// "strings"
"crypto/md5"
"encoding/hex"
"encoding/json"
)

func Gen_Sign(obj interface{}, key string)(ret string, err error){
	org, err := json.Marshal(obj);
	if err != nil {
		return ret, err;
	}

	all_data := fmt.Sprintf("%s+%s", (string)(org), key);
	
	
	fmt.Printf("待签名数据: %s\n", all_data);

	md := md5.New();
	md.Write([]byte(all_data));
	
	// return strings.ToUpper(hex.EncodeToString(md.Sum(nil)));
	return hex.EncodeToString(md.Sum(nil)), nil;	
}
