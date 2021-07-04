package xjexcel

import (
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"log"
	"reflect"
	"strconv"
	"strings"
)

//将数据列表转成excel表格
func ListToExcel(list interface{}, title, sheetName string) *excelize.File {
	f := excelize.NewFile()
	lines := 0
	titleLines := 0
	if title != "" {
		titleLines = 1
	}

	fieldList := make([]ExcelField, 0)
	switch reflect.TypeOf(list).Kind() {
	case reflect.Slice, reflect.Array:
		s := reflect.ValueOf(list)
		lines = s.Len()
		for i := 0; i < lines; i++ {
			val := s.Index(i)
			//fmt.Println(val)
			if i == 0 {
				st := reflect.TypeOf(val.Interface())
				for j := 0; j < st.NumField(); j++ {
					field := st.Field(j)
					tag := field.Tag.Get("excel")
					if tag != "" {
						column := TagField(tag, "column")
						desc := TagField(tag, "desc")
						width := TagField(tag, "width")
						w, _ := strconv.ParseFloat(width, 10)
						fielditem := ExcelField{
							Name:   field.Name,
							Column: column,
							Desc:   desc,
							Width:  w,
						}
						fieldList = append(fieldList, fielditem)
					}
				}
			}
			if !FindExcelCol(fieldList, "A") {
				//如果没有指定A列,则A列为序号列
				f.SetCellValue("Sheet1", "A"+strconv.Itoa(1+titleLines), "序号")
				f.SetCellValue("Sheet1", "A"+strconv.Itoa(i+2+titleLines), i+1)
			}
			for _, item := range fieldList {
				//设置单元格值
				f.SetCellValue("Sheet1", item.Column+strconv.Itoa(i+2+titleLines), val.FieldByName(item.Name))
				//设置单元格宽度
				f.SetColWidth("Sheet1", item.Column, item.Column, item.Width)
			}
		}
	}

	//标题
	maxCol := FindMaxExcelCol(fieldList)
	if title != "" {
		f.SetCellValue("Sheet1", "A1", title)
		//合并
		f.MergeCell("Sheet1", "A1", maxCol+"1")
		//格式:居中
		style1, _ := f.NewStyle(`{"alignment":{"horizontal":"center"}}`)
		f.SetCellStyle("Sheet1", "A1", maxCol+"1", style1)
	}
	//列名
	for _, item := range fieldList {
		f.SetCellValue("Sheet1", item.Column+strconv.Itoa(1+titleLines), item.Desc)
	}
	//格式,有边框,单元格居中
	sty := `{"border":[{"type":"left","color":"000000","style":1},{"type":"top","color":"000000","style":1},{"type":"bottom","color":"000000","style":1},{"type":"right","color":"000000","style":1}],
		"alignment":{"horizontal":"center"}}`
	style, err := f.NewStyle(sty)
	if err != nil {
		log.Println(err)
	}
	f.SetCellStyle("Sheet1", "A"+strconv.Itoa(1+titleLines), maxCol+strconv.Itoa(lines+1+titleLines), style)
	//冻结0列2行
	f.SetPanes("Sheet1", `{
		"freeze": true,
		"x_split": 0,
		"y_split": `+strconv.Itoa(1+titleLines)+"}",
	)
	// 修改表名
	if sheetName != "" {
		f.SetSheetName("Sheet1", sheetName)
	}
	return f

}

func TagField(tag, field string) string {
	i1 := strings.Index(tag, field)
	if i1 > -1 {
		i2 := i1 + len(field) + 1
		i3 := strings.Index(tag[i2:], ";")
		if i3 > -1 {
			return tag[i2 : i2+i3]
		} else {
			return tag[i2:]
		}
	}
	return ""
}

func FindExcelCol(list []ExcelField, col string) bool {
	for _, item := range list {
		if item.Column == col {
			return true
		}
	}
	return false
}

func FindMaxExcelCol(list []ExcelField) string {
	col := "A"
	for _, item := range list {
		if strings.Compare(strings.ToLower(col), strings.ToLower(item.Column)) < 0 {
			col = item.Column
		}
	}
	return col
}

type ExcelField struct {
	Name   string
	Column string
	Desc   string
	Width  float64
}
