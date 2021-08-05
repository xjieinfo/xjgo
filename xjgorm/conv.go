package xjgorm

import (
	"bytes"
	"log"
	"reflect"
	"strconv"
	"strings"
	"unicode"
)

func getSliceStructName(i interface{}) string {
	// 这里我想获取数组里面的结构体类型名
	s := reflect.TypeOf(i).String()
	return s
}

func getSliceZeroItem(list interface{}) interface{} {
	vType := reflect.TypeOf(list).Elem()
	value := reflect.Zero(vType).Interface()
	return value
}

func getSliceTableName(list interface{}) string {
	vType := reflect.TypeOf(list).Elem()
	value := reflect.Zero(vType).Interface()
	vValue := reflect.ValueOf(value)
	for i := 0; i < vType.NumMethod(); i++ {
		methodName := vType.Method(i).Name
		if methodName == "TableName" {
			values := vValue.Method(i).Call(nil)
			if len(values) > 0 {
				name := values[0].String()
				return name
			}
		}
	}
	name := reflect.TypeOf(list).Elem().Name()
	return Camel2Case(name)
}

func getStructTableName(item interface{}) string {
	vType := reflect.TypeOf(item)
	vValue := reflect.ValueOf(item)
	for i := 0; i < vType.NumMethod(); i++ {
		methodName := vType.Method(i).Name
		if methodName == "TableName" {
			values := vValue.Method(i).Call(nil)
			if len(values) > 0 {
				name := values[0].String()
				return name
			}
		}
	}
	name := reflect.TypeOf(item).Name()
	return Camel2Case(name)
}

// 驼峰式写法转为下划线写法
func Camel2Case(name string) string {
	buffer := NewBuffer()
	for i, r := range name {
		if unicode.IsUpper(r) {
			if i != 0 {
				buffer.Append('_')
			}
			buffer.Append(unicode.ToLower(r))
		} else {
			buffer.Append(r)
		}
	}
	return buffer.String()
}

// 下划线写法转为驼峰写法
func Case2Camel(name string) string {
	name = strings.Replace(name, "_", " ", -1)
	name = strings.Title(name)
	return strings.Replace(name, " ", "", -1)
}

// 首字母大写
func Ucfirst(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return ""
}

// 首字母小写
func Lcfirst(str string) string {
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}
	return ""
}

// 内嵌bytes.Buffer，支持连写
type Buffer struct {
	*bytes.Buffer
}

func NewBuffer() *Buffer {
	return &Buffer{Buffer: new(bytes.Buffer)}
}

func (b *Buffer) Append(i interface{}) *Buffer {
	switch val := i.(type) {
	case int:
		b.append(strconv.Itoa(val))
	case int64:
		b.append(strconv.FormatInt(val, 10))
	case uint:
		b.append(strconv.FormatUint(uint64(val), 10))
	case uint64:
		b.append(strconv.FormatUint(val, 10))
	case string:
		b.append(val)
	case []byte:
		b.Write(val)
	case rune:
		b.WriteRune(val)
	}
	return b
}

func (b *Buffer) append(s string) *Buffer {
	defer func() {
		if err := recover(); err != nil {
			log.Println("*****内存不够了！******")
		}
	}()
	b.WriteString(s)
	return b
}
