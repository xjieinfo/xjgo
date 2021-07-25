package main

import "github.com/xjieinfo/xjgo/xjcore/xjexcel"

//通过excelize将列表转换成excel文件
func main() {
	list := make([]User, 0)
	user1 := User{
		Name:    "张三",
		Age:     18,
		Address: "北京东三环",
	}
	user2 := User{
		Name:    "李四",
		Age:     21,
		Address: "上海人民路",
	}
	user3 := User{
		Name:    "王五",
		Age:     22,
		Address: "长沙开福区",
	}
	list = append(list, user1, user2, user3)
	f := xjexcel.ListToExcel(list, "员工信息表", "员工表")
	f.SaveAs("/员工表.xls")
}

type User struct {
	Name    string `excel:"column:B;desc:姓名;width:30"`
	Age     int    `excel:"column:C;desc:年龄;width:10"`
	Address string `excel:"column:D;desc:地址;width:50"`
}
