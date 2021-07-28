# 一个简单封装的go框架xjgo

go语言标准库提供了http库，功能比较完善，在此基础上稍加封装，就能成为一个称手的工具，请看如下示例：
```go
package main

import (
	"github.com/xjieinfo/xjgo/xjhttp"
)

func main() {
	xjhttp := xjhttp.Default()
	xjhttp.HandleFunc("GET", "/str", Str)
	xjhttp.GET("/success", Success)
	xjhttp.GET("/fail", Fail)
	xjhttp.GET("/make", Make)
	xjhttp.POST("/:order", order)
	xjhttp.PUT("/user/saler/:userName", saler)
	xjhttp.DELETE("/user/:id/xyz", xyz)
	xjhttp.GET("/ctx", okctx)
	xjhttp.Run(":6001")
}

func Str(ctx *xjhttp.Context) {
	name := "xjgo"
	ctx.String(200, "hello %s .", name)
}

func Success(ctx *xjhttp.Context) {
	data := struct {
		Name string
		Age  int
	}{
		Name: "ZhangSan",
		Age:  18,
	}
	ctx.Success(200, data)
}

func Fail(ctx *xjhttp.Context) {
	ctx.Fail(200, "fail")
}

func Make(ctx *xjhttp.Context) {
	ctx.Make(200, 101806, "data is error.", "access error")
}
func okctx(ctx *xjhttp.Context) {
	ctx.JSON(200, "ok")
}
func order(ctx *xjhttp.Context) {
	ctx.String(200, "order query.")
}

func saler(ctx *xjhttp.Context) {
	ctx.String(200, "saler.")
}

func xyz(ctx *xjhttp.Context) {
	ctx.String(200, "xyz")
}

```
项目开源地址：[https://github.com/xjieinfo/xjgo](https://github.com/xjieinfo/xjgo)
如果你有好的想法，请告诉我，谢谢！
如果此项目对你有所帮助或启发，请给个star支持一下，谢谢！
