package main

import (
    "fmt"
    "os"
    "io/ioutil"

    "github.com/Luxurioust/excelize"
    "encoding/json"
    "reflect"
    "strconv"
    "strings"
    "flag"
)

var excel     = flag.String("f", "skill.xlsx", "input file (excel)");
var json_file = flag.String("o", "", "output file (json)");
var sheet     = flag.String("s", "Sheet1", "Sheet name");
var sheet_idx = flag.Int("i", 1, "Sheet Index");

func init(){
    flag.Parse();
}

func main() {
    excel_to_json(*excel, *json_file);
}

type Lanch_Data struct {
    Id string;
    Name string;
    LanchPos int64;
    LanchAngle float64;
    ScanRange float64;
    TargetCamp int64;
    TargetType int64;
    DelayTime float64;
    LanchDelay float64;
    LanchNum int64;
    LanchTime int64;
    ProjectileID string;
};



type Data struct{
    Key []string;
    Value []Lanch_Data;
};

func excel_to_json(file_name string, json_file string){
    xlf, err := excelize.OpenFile(file_name);
    if err != nil {
        fmt.Println("OpenFile: ", err);
        os.Exit(1);
    }

    fmt.Println("All sheet: ");
    for k, v := range xlf.GetSheetMap(){
        fmt.Println(k, v);
    }
    xlf.SetActiveSheet(*sheet_idx);
    fmt.Println("now in : ", xlf.GetActiveSheetIndex());

    rows := xlf.GetRows(*sheet);

    prc_key := false;
    prc_data := false;

    var key int; // key所在行
    keys := make(map[string]string); // key行数据

    var datas Data;
    var lst []Lanch_Data;
    var ID  []string;
    for x, row := range rows{
        fmt.Printf("处理[%d]行\n", x);
        has_data := false;

        var data Lanch_Data;        
        for y, col := range row{
            fmt.Printf("[%d][%d]%s\t", x, y, col);
            if (col == "k"){
                prc_key = true;
                prc_data = false;
                key = x;
            }
            if (col == "v"){
                prc_key = false;
                prc_data = true;
            }            
            if len(col) > 0 {
                if prc_key{ // 处理key
                    keys[fmt.Sprintf("%d,%d", x, y)] = col;
                }else if prc_data{ // 处理data
                    val := reflect.ValueOf(&data).Elem().FieldByName(keys[fmt.Sprintf("%d,%d", key, y)]);
                    if val.IsValid(){
                        ref := reflect.ValueOf(&data).Elem().FieldByName(keys[fmt.Sprintf("%d,%d", key, y)]).Interface();
                        switch ref.(type){
                            case string:{
                                val.SetString(col);
                                break;
                            }
                            case int64:{
                                inum, err := strconv.ParseInt(col, 10, 64);
                                if err != nil{
                                    fmt.Printf("错误![%d,%d]应该是[int], 现配置为=%s: %v", x, y, col, err);
                                    return;
                                }
                                val.SetInt(inum);
                                break;
                            }
                            case float64:{
                                ifloat, err := strconv.ParseFloat(col, 64);
                                if err != nil{
                                    fmt.Printf("错误![%d,%d]应该是[float], 现配置为=%s: %v", x, y, col, err);
                                    return;
                                }
                                val.SetFloat(ifloat);
                                break;
                            }
                            case []string:{
                                subs := strings.Split(col, ";");
                                vl := reflect.ValueOf(subs);
                                val.Set(vl);
                                break;
                            }
                        }


                        if has_data == false{ // 首次出现，必定是ID
                            has_data = true;
                            ID = append(ID, col);
                        }                            
                    }
                }
            }
        }
        if prc_data && has_data{
            lst = append(lst, data);
            prc_key = false;
            has_data = false;
   
            // fmt.Printf("data: %#v\n", data);    
        }
    } // end for

    datas.Key = ID;
    datas.Value = lst;

    // fmt.Printf("key in %d\n", key);
    // fmt.Printf("keys: %#v\n", keys);
    // fmt.Printf("data: %#v\n", lst);
    js, _ := json.MarshalIndent(datas, " ", " ");

    if len(json_file) > 0{
        write_file(json_file, js);
    }else{
        fmt.Printf("%s\n", js);   
    }
}

func write_file(file_name string, data []byte){
    err := ioutil.WriteFile(file_name, data, 0644);
    if err != nil{
        panic(err);
    }
}