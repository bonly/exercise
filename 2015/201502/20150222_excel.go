package main

import (
    "fmt"
    "os"

    "github.com/Luxurioust/excelize"
    "encoding/json"
    "reflect"
    "strconv"
    "strings"
)

func main() {
    xlsx, err := excelize.OpenFile("/home/bonly/techappen/skill.xlsx");
    if err != nil {
        fmt.Println(err);
        os.Exit(1);
    }

    rows := xlsx.GetRows("Sheet1");

    prc_key := false;
    prc_data := false;

    var key int; // key所在行
    keys := make(map[string]string); // key行数据

    var datas Data;
    var lst []Skill_Data;
    var ID  []string;
    for x, row := range rows{
        // fmt.Printf("处理[%d]行\n", x);
        has_data := false;

        var data Skill_Data;        
        for y, col := range row{
            // fmt.Printf("[%d][%d]%s\t", x, y, col);
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
                            case int:{
                                inum, err := strconv.ParseInt(col, 10, 64);
                                if err != nil{
                                    fmt.Println(err);
                                    return;
                                }
                                val.SetInt(inum);
                                break;
                            }
                            case float64:{
                                ifloat, err := strconv.ParseFloat(col, 64);
                                if err != nil{
                                    fmt.Println(err);
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
    fmt.Printf("%s\n", js);   

}

type Skill_Data struct {
    Id string;
    Name string;
    Icon string;
    CoolDown float64;
    LockType int64;
    LockDistance float64;
    TargetCamp int64;
    LanchIDs []string;
    TouchUiID string;
    TouchMeshID string;
    Desc string;
    StartTime float64;
    StopTime float64;
    StartMoveTime float64;
    StopMoveTime float64;
    StartRotateTime float64;
    StopRotateTime float64;
    SkillControlTime float64;
};

type Data struct{
    Key []string;
    Value []Skill_Data;
};