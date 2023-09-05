package util

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"mime/multipart"
	"sort"
)

// ExcelDataFormat excel文件处理
type ExcelDataFormat struct {
	A interface{}
	B interface{}
	C interface{}
	D interface{}
	E interface{}
	F interface{}
	G interface{}
	H interface{}
	I interface{}
	J interface{}
	K interface{}
	L interface{}
	M interface{}
	N interface{}
	O interface{}
}

//ReadExcel .读取excel 转成切片
//func ReadExcel(xlsx *excelize.File) ([]LxProduct, error) {
func ReadExcel(fileInfo multipart.File) ([]ExcelDataFormat, error) {
	res := make([]ExcelDataFormat, 0)
	xlsxFile, err := excelize.OpenReader(fileInfo)
	if err != nil {
		return res, err
	}

	//根据名字获取cells的内容，返回的是一个[][]string
	rows, err := xlsxFile.GetRows(xlsxFile.GetSheetName(xlsxFile.GetActiveSheetIndex()))
	if err != nil {
		return res, err
	}
	//声明一个数组
	var lxProducts []ExcelDataFormat
	for i, row := range rows {
		// 去掉第一行是excel表头部分
		if i == 0 {
			continue
		}
		var data ExcelDataFormat
		for k, v := range row {
			if k == 0 {
				data.A = v
			}
			if k == 1 {
				data.B = v
			}
			if k == 2 {
				data.C = v
			}
			if k == 3 {
				data.D = v
			}
			if k == 4 {
				data.E = v
			}
			if k == 5 {
				data.F = v
			}
			if k == 6 {
				data.G = v
			}
			if k == 7 {
				data.H = v
			}
			if k == 8 {
				data.I = v
			}
			if k == 9 {
				data.J = v
			}
			if k == 10 {
				data.K = v
			}
			if k == 11 {
				data.L = v
			}
			if k == 12 {
				data.M = v
			}
			if k == 13 {
				data.N = v
			}
			if k == 14 {
				data.O = v
			}
		}

		//将数据追加到集合中
		lxProducts = append(lxProducts, data)
	}
	return lxProducts, nil
}

func ExportExcel(titleList []string, data []map[string]interface{}, sheetName, filepath string) error {
	// 创建一个工作表
	file := excelize.NewFile()
	index, _ := file.NewSheet(sheetName)
	file.SetActiveSheet(index)
	_ = file.SetSheetRow(sheetName, "A1", &titleList)

	rowNum := 1
	for _, value := range data {
		row := make([]interface{}, 0)
		var dataSlice []string
		for key := range value {
			dataSlice = append(dataSlice, key)
		}
		sort.Strings(dataSlice)
		for _, v := range dataSlice {
			if val, ok := value[v]; ok {
				row = append(row, val)
			}
		}
		rowNum++
		if err := file.SetSheetRow(sheetName, fmt.Sprintf("A%d", rowNum), &row); err != nil {
			return err
		}

	}
	// 根据指定路径保存文件
	return file.SaveAs(filepath)
}
