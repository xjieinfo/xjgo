package xjhttp

import (
	"encoding/json"
	"fmt"
	"gitee.com/xjieinfo/xjgo/xjcore/xjtypes"
	"net/http"
	"net/url"
)

type Context struct {
	Request *http.Request
	Writer  http.ResponseWriter
}

func (c *Context) QueryAll() url.Values {
	return c.Request.URL.Query()
}

func (c *Context) QueryStr(key string) string {
	querys := c.Request.URL.Query()
	return querys.Get(key)
}

func (c *Context) String(code int, format string, values ...interface{}) {
	c.Writer.WriteHeader(code)
	fmt.Fprintf(c.Writer, format, values...)
}

func (c *Context) JSON(code int, obj interface{}) {
	c.Writer.WriteHeader(code)
	j, _ := json.Marshal(obj)
	fmt.Fprintf(c.Writer, string(j))
}

func (c *Context) Make(code int, rcode int, data interface{}, msg string) {
	c.Writer.WriteHeader(code)
	r := new(xjtypes.R).Make(rcode, data, msg)
	j, _ := json.Marshal(r)
	fmt.Fprintf(c.Writer, string(j))
}

func (c *Context) Success(code int, data interface{}) {
	c.Writer.WriteHeader(code)
	r := new(xjtypes.R).Success(data)
	j, _ := json.Marshal(r)
	fmt.Fprintf(c.Writer, string(j))
}

func (c *Context) Fail(code int, msg string) {
	c.Writer.WriteHeader(code)
	r := new(xjtypes.R).Fail(msg)
	j, _ := json.Marshal(r)
	fmt.Fprintf(c.Writer, string(j))
}

func (c *Context) Error(code int, err error) {
	c.Writer.WriteHeader(code)
	r := new(xjtypes.R).Fail(err.Error())
	j, _ := json.Marshal(r)
	fmt.Fprintf(c.Writer, string(j))
}
