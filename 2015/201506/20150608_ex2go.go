package main

import (
	"fmt"
	_ "game/gameStruct"
	"os"

	"github.com/xuri/excelize"
)

func main() {
	fmt.Println("begin")
	defer fmt.Println("end")

	excel_to_struct("abc")
}

func excel_to_struct(file_name string) {
	xlsx, err := excelize.OpenFile("场景单位表.xlsx")
	if err != nil {
		fmt.Printf("Open File: %v\n", err)
		os.Exit(1)
	}

	// Get value from cell by given worksheet name and axis.
	cell := xlsx.GetCellValue("Sheet1", "B2")
	fmt.Println(cell)
	// Get all the rows in the Sheet1.
	rows := xlsx.GetRows("Sheet1")
	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println()
	}
}
