package main

import (
	"fmt"
	"gitee.com/xjieinfo/xjgo/xjhttp"
	"net/http"
)

func main() {
	xjhttp := xjhttp.Default()
	xjhttp.HandleFunc("GET", "/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello xjgo.")
	})
	xjhttp.HandleFunc("GET", "/ok", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "ok.")
	})
	xjhttp.HandleFunc("GET", "/:order", order)
	xjhttp.HandleFunc("GET", "/user/saler/:userName", saler)
	xjhttp.HandleFunc("GET", "/user/:id/xyz", xyz)
	xjhttp.Run(":7001")
}

func order(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "order query.")
}

func saler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "saler.")
}

func xyz(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "xyz")
}
