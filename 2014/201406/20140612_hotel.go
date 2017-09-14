package main

import (
	_ "github.com/alexbrainman/odbc"
	"database/sql"
    "log"
    "fmt"
    "os"
    "strconv"
    "strings"
    "golang.org/x/text/encoding/simplifiedchinese"
    "golang.org/x/text/transform"
    "io/ioutil"
    "bytes"
)

func main() {
	db, err := sql.Open("odbc","DSN=FromAccess");
	if err != nil {
		log.Fatal("open: ",err);
	}
	defer db.Close();

	var (
		id string;
		name string;
		addr string;
		idcode string;
		idphoto sql.RawBytes;
	);

	rows, err := db.Query("select id,name,idcode,address,idphoto from t_guest");

	if err != nil {
		log.Fatal("qry: ", err);
	}
	defer rows.Close();


	for rows.Next() {
		err := rows.Scan(&id, &name, &idcode, &addr, &idphoto);
		if err != nil {
			log.Fatal("scan: ", err);
		}
		fmt.Printf("%s %s %s %s %d\n", id, decrypt(name), decrypt(idcode), decrypt(addr), len(idphoto));
		if id != "1186"{
			continue;
		}
		fmt.Printf("photo: %+v \n", idphoto);			
		filename := fmt.Sprintf("%s.png", id);
		file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0666);
		if err != nil{
			fmt.Println("create file: ", err);
		}
		_, err = file.Write(idphoto);
		if err != nil{
			fmt.Println("write: ", err);
		}
	}
	err = rows.Err();
	if err != nil {
		log.Fatal("row: ", err);
	}
}

func decrypt(src string) (ret string){
	arr_str := strings.Split(src, "@");
	for idx:=0; idx<len(arr_str); idx++{
		chr, _ := strconv.ParseUint(arr_str[idx], 16, 32);
		// fmt.Printf("len is: %d\n", len(arr_str[idx]));
		var achr uint64;
		if len(arr_str[idx]) == 4{
			achr = chr ^ 0xB209;
		}else{
			achr = chr ^ 0xB2;
		}
		// fmt.Printf("%x\n", achr);
		high := achr >> 8;
		low := achr & 0xFF;
		word, _ := GBK2UTF8([]byte{byte(high), byte(low)});
		ret += string(word);
	}
	
	return ret;
}

func GBK2UTF8(sb []byte)([]byte, error){
	reader := transform.NewReader(bytes.NewReader(sb), simplifiedchinese.GBK.NewDecoder());
	dst, err := ioutil.ReadAll(reader);
	if err != nil{
		return nil, err;
	}
	return dst, nil;
}

func UTF82GBK(sb []byte)([]byte, error){
	reader := transform.NewReader(bytes.NewReader(sb), simplifiedchinese.GBK.NewEncoder());
	dst, err := ioutil.ReadAll(reader);
	if err != nil{
		return nil, err;
	}
	return dst, nil;
}