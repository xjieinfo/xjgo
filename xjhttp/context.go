package xjhttp

import (
	"encoding/json"
	"fmt"
	"gitee.com/xjieinfo/xjgo/xjcore/xjstruct"
	"gitee.com/xjieinfo/xjgo/xjcore/xjtypes"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type Context struct {
	Request *http.Request
	Writer  http.ResponseWriter
}

func (c *Context) Form() url.Values {
	c.Request.ParseForm()
	return c.Request.Form
}

func (c *Context) BodyJson(obj interface{}) error {
	str, err := c.Body()
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(str), obj)
	return err
}

func (c *Context) Body() (string, error) {
	body, err := ioutil.ReadAll(c.Request.Body)
	//c.Request.Body.Close()
	return string(body), err
}

func (c *Context) QueryMap() (m map[string]interface{}) {
	m = make(map[string]interface{})
	values := c.Request.URL.Query()
	for k, v := range values {
		if len(v) == 1 {
			m[k] = v[0]
		} else {
			m[k] = v
		}
	}
	return
}

func (c *Context) QueryStruct(dst interface{}) {
	m := c.QueryMap()
	xjstruct.MapToStructWithOutTypeDeep(m, dst)
	return
}

func (c *Context) QueryAll() url.Values {
	return c.Request.URL.Query()
}

func (c *Context) QueryStr(key string) string {
	querys := c.Request.URL.Query()
	return querys.Get(key)
}

func (c *Context) QueryInt(key string) (int, error) {
	querys := c.Request.URL.Query()
	str := querys.Get(key)
	val, err := strconv.Atoi(str)
	return val, err
}

func (c *Context) QueryIntDefault(key string, def int) (int, error) {
	querys := c.Request.URL.Query()
	str := querys.Get(key)
	val, err := strconv.Atoi(str)
	if err != nil {
		val = def
	}
	return val, err
}

func (c *Context) QueryInt64(key string) (int64, error) {
	querys := c.Request.URL.Query()
	str := querys.Get(key)
	val, err := strconv.ParseInt(str, 10, 64)
	return val, err
}

func (c *Context) PathParam(index int) string {
	uri := c.Request.RequestURI
	strs := strings.Split(uri, "/")
	if len(strs) > index {
		return strs[index]
	} else {
		return ""
	}
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
