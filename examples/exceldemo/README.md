# golang如何快速生成excel表格

# golang如何快速生成excel表格
      操作excel表格，将数据保存为excel表格，便于数据的交换，展示和下载。在我们平常的工作中是一项基本需求。
      我比较喜欢的库是360开源的excelize，目前最新版本为v2.4.0，这个库非常的方便和好用。
      但对于新手来说，掌握也需要摸索一下，本人最近为了将list表格转换成excel表格，特意研究了一下：
      要求如下：
1.  将list表格转换成excel表格；
2.  给每一列设定一个宽度，并设定列的标题，居中对齐；
3.  给表格加上边框线；
4.  给表格加上一个title；
5.  给表格设定一个名字；
6.  固定标题行，免得数据多时滚动看不到标题行。
    做完之后，我想这种需求以后应该会经常用，何不封装一下，变成一个小工具，以后随时调用，岂不方便很多，最后我将这些功能封装到一个函数中：

```go
func ListToExcel(list interface{}, title, sheetName string) *excelize.File { }
```
以后使用的时候如何调用呢？请看以下示例：
```go

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
```

通过struct的tag将每一列的位置，名称，宽度进行设置，如果不设置A列，则A列为序号列，自动生成。

具体的项目源码，请访问以下网址：[https://github.com/xjieinfo/xjgo](https://github.com/xjieinfo/xjgo)

如果此项目对你有所帮助或启发，请给个star支持一下，谢谢！
