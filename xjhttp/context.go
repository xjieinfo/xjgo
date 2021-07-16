package xjhttp

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/xjieinfo/xjgo/xjcache"
	"github.com/xjieinfo/xjgo/xjcore/xjstruct"
	"github.com/xjieinfo/xjgo/xjcore/xjtypes"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type Context struct {
	Request  *http.Request
	Writer   http.ResponseWriter
	CacheMgr *xjcache.CacheMgr
	index    int
	handlers []HandlerFunc
}

func (c *Context) GetHeader(key string) string {
	return c.Request.Header.Get(key)
}

func (c *Context) Form() url.Values {
	c.Request.ParseForm()
	return c.Request.Form
}

func (c *Context) Bind(obj interface{}) error {
	return c.BodyJson(obj)
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

func (c *Context) DefaultQuery(key, def string) string {
	val := c.QueryStr(key)
	if val == "" {
		return def
	} else {
		return val
	}
}

func (c *Context) Query(key string) string {
	return c.QueryStr(key)
}

func (c *Context) GetString(key string) string {
	return c.QueryStr(key)
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

func (c *Context) Param(key string) string {
	pattern := c.Request.Header.Get("xjgo-path-pattern")
	paths := strings.Split(pattern, "/")
	for i, item := range paths {
		if item == ":"+key {
			return c.PathParam(i)
		}
	}
	return ""
}

func (c *Context) ParamInt(key string) (int, error) {
	pattern := c.Request.Header.Get("xjgo-path-pattern")
	paths := strings.Split(pattern, "/")
	for i, item := range paths {
		if item == ":"+key {
			str := c.PathParam(i)
			val, err := strconv.Atoi(str)
			return val, err
		}
	}
	return 0, errors.New("not found")
}

func (c *Context) ParamInt64(key string) (int64, error) {
	pattern := c.Request.Header.Get("xjgo-path-pattern")
	paths := strings.Split(pattern, "/")
	for i, item := range paths {
		if item == ":"+key {
			str := c.PathParam(i)
			val, err := strconv.ParseInt(str, 10, 64)
			return val, err
		}
	}
	return 0, errors.New("not found")
}

func (c *Context) ParamFloat64(key string) (float64, error) {
	pattern := c.Request.Header.Get("xjgo-path-pattern")
	paths := strings.Split(pattern, "/")
	for i, item := range paths {
		if item == ":"+key {
			str := c.PathParam(i)
			val, err := strconv.ParseFloat(str, 64)
			return val, err
		}
	}
	return 0, errors.New("not found")
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

func (c *Context) PathParamInt(index int) (int, error) {
	uri := c.Request.RequestURI
	strs := strings.Split(uri, "/")
	if len(strs) > index {
		i, err := strconv.Atoi(strs[index])
		return i, err
	} else {
		return 0, errors.New("index error.")
	}
}

func (c *Context) PathParamInt64(index int) (int64, error) {
	uri := c.Request.RequestURI
	strs := strings.Split(uri, "/")
	if len(strs) > index {
		i, err := strconv.ParseInt(strs[index], 10, 64)
		return i, err
	} else {
		return 0, errors.New("index error.")
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

func (c *Context) Next() {
	c.index++
	for c.index < len(c.handlers) {
		c.handlers[c.index](c)
		c.index++
	}
}

func (c *Context) Abort() {
	c.index = 100
}

func (c *Context) MultipartForm() (*multipart.Form, error) {
	err := c.Request.ParseMultipartForm(10 * 1024 * 1024)
	return c.Request.MultipartForm, err
}
