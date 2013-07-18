package main
import (
    "fmt"
    "github.com/tealeg/xlsx" //需要引入的包
)

func main() {
    var xlFile *xlsx.File
    var sheetLen int

    xlFile, _ = xlsx.OpenFile("haha2.xlsx")
    //defer xlFile.CloseFile

    // 按 sheet 的顺序编号从 0 开始处理到最后 sheet
    sheetLen = len(xlFile.Sheets)
    for i := 0; i < sheetLen; i++ { // 第一种方式
        sheet := xlFile.Sheets[i]
        fmt.Println("现在是 sheet:", i)
        for rowIndex, row := range sheet.Rows {
            for cellIndex, cell := range row.Cells {
                fmt.Printf("第%d行，第%d列：%v\n", rowIndex, cellIndex, cell)
            }
        }
    }

    for i, _ := range xlFile.Sheets { // 第二种方式, 比较好
        sheet := xlFile.Sheets[i]
        fmt.Println("现在是 sheet:", i)
        for rowIndex, row := range sheet.Rows {
            for cellIndex, cell := range row.Cells {
                fmt.Printf("第%d行，第%d列：%v\n", rowIndex, cellIndex, cell)
            }
        }
    }

    // 按 sheet 的名字进行处理, 顺序是随机的不确定
    for shname, _ := range xlFile.Sheet {
        sheet := xlFile.Sheet[shname]
        fmt.Println("现在是 sheet:", shname)
        for rowIndex, row := range sheet.Rows {
            for cellIndex, cell := range row.Cells {
                fmt.Printf("第%d行，第%d列：%v\n", rowIndex, cellIndex, cell)
            }
        }
    }

    // 打出各种结构看看
    fmt.Printf("文件结构:%#v\n\n", xlFile)
    fmt.Printf("Sheet 结构:%#v\n\n", xlFile.Sheet)   // sheet 名称放在 map 中
    fmt.Printf("Sheets 结构:%#v\n\n", xlFile.Sheets) // 按sheet 的顺序号从零开始
}
