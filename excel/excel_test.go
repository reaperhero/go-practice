package excel

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"testing"
)


// write
func Test_excel_01(t *testing.T) {
	categories := map[string]string{"A2": "Small", "A3": "Normal", "A4": "Large", "B1": "Apple", "C1": "Orange", "D1": "Pear"}
	values := map[string]int{"B2": 2, "C2": 3, "D2": 3, "B3": 5, "C3": 2, "D3": 4, "B4": 6, "C4": 7, "D4": 8}
	f := excelize.NewFile()
	for k, v := range categories {
		f.SetCellValue("Sheet1", k, v)
	}
	for k, v := range values {
		f.SetCellValue("Sheet1", k, v)
	}
	if err := f.SaveAs("Book1.xlsx"); err != nil {
		fmt.Println(err)
	}
}
//	     Apple Orange Pear
//Small	 2	   3	3
//Normal 5	   2	4
//Large	 6     7	8



// read
func Test_excel_02(t *testing.T) {
	f, err := excelize.OpenFile("Book1.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	f.GetCellValue("Sheet1", "B2") // 2

	// Get all the rows in the Sheet1.
	rows := f.GetRows("Sheet1")
	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println()
	}
}

