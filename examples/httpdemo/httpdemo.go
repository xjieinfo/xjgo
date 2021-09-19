package main

import (
	"github.com/xjieinfo/xjgo/xjhttp"
)

func main() {
	xjhttp := xjhttp.Default()
	xjhttp.HandleFunc("GET", "/str", "", Str)
	xjhttp.GET("/", "", Hello)
	xjhttp.GET("/success", "", Success)
	xjhttp.GET("/fail", "", Fail)
	xjhttp.GET("/make", "", Make)
	xjhttp.POST("/:order", "", order)
	xjhttp.PUT("/user/saler/:userName", "", saler)
	xjhttp.DELETE("/user/:id/xyz", "", xyz)
	xjhttp.GET("/ctx", "", okctx)
	xjhttp.Run(":6001")
}

func Hello(ctx *xjhttp.Context) {
	name := "xjgo"
	ctx.String(200, "hello %s .", name)
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
