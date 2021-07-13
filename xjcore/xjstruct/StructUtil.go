package xjstruct

import (
	"encoding/json"
	"github.com/xjieinfo/xjgo/xjcore/xjtypes"
	"log"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func StructCopy(DstStructPtr interface{}, SrcStructPtr interface{}) {
	srcv := reflect.ValueOf(SrcStructPtr)
	dstv := reflect.ValueOf(DstStructPtr)
	srct := reflect.TypeOf(SrcStructPtr)
	dstt := reflect.TypeOf(DstStructPtr)
	if srct.Kind() != reflect.Ptr || dstt.Kind() != reflect.Ptr ||
		srct.Elem().Kind() == reflect.Ptr || dstt.Elem().Kind() == reflect.Ptr {
		panic("Fatal error:type of parameters must be Ptr of value")
	}
	if srcv.IsNil() || dstv.IsNil() {
		panic("Fatal error:value of parameters should not be nil")
	}
	srcV := srcv.Elem()
	dstV := dstv.Elem()
	srcfields := DeepFields(reflect.ValueOf(SrcStructPtr).Elem().Type())
	for _, v := range srcfields {
		if v.Anonymous {
			continue
		}
		dst := dstV.FieldByName(v.Name)
		src := srcV.FieldByName(v.Name)
		if !dst.IsValid() {
			continue
		}
		if src.Type() == dst.Type() && dst.CanSet() {
			dst.Set(src)
			continue
		}
		if src.Kind() == reflect.Ptr && !src.IsNil() && src.Type().Elem() == dst.Type() {
			dst.Set(src.Elem())
			continue
		}
		if dst.Kind() == reflect.Ptr && dst.Type().Elem() == src.Type() {
			dst.Set(reflect.New(src.Type()))
			dst.Elem().Set(src)
			continue
		}
	}
	return
}

func DeepFields(ifaceType reflect.Type) []reflect.StructField {
	var fields []reflect.StructField

	for i := 0; i < ifaceType.NumField(); i++ {
		v := ifaceType.Field(i)
		if v.Anonymous && v.Type.Kind() == reflect.Struct {
			fields = append(fields, DeepFields(v.Type)...)
		} else {
			fields = append(fields, v)
		}
	}

	return fields
}

func CopyStruct(src, dst interface{}) {
	sval := reflect.ValueOf(src).Elem()
	dval := reflect.ValueOf(dst).Elem()

	for i := 0; i < sval.NumField(); i++ {
		value := sval.Field(i)
		name := sval.Type().Field(i).Name

		dvalue := dval.FieldByName(name)
		if dvalue.IsValid() == false {
			continue
		}
		if dvalue.Type() == value.Type() {
			dvalue.Set(value) //这里默认共同成员的类型一样，否则这个地方可能导致 panic，需要简单修改一下。
		} else {
			log.Printf("name is:%s,src type is:%s, dst type is:%s\n", name, value.Type().String(), dvalue.Type().String())
		}
	}
}

func CopyStruct2(src, dst interface{}) {
	aj, err := json.Marshal(src)
	if err != nil {
		log.Println(err)
	}
	err = json.Unmarshal(aj, dst)
	if err != nil {
		log.Println(err)
	}
}

/**
 * Map转struct
 * @param map 需要初始化的数据，key字段必须与实体类的成员名字一样，否则赋值为空
 * @param entity  需要转化成的实体类
 * @return
 */
func MapToStructWithOutType(src map[string]interface{}, dst interface{}) {
	dval := reflect.ValueOf(dst).Elem()
	for k, v := range src {
		key := Capitalize(k)
		//println(key)
		dvalue := dval.FieldByName(key)
		if !dvalue.IsValid() {
			//log.Printf("%s is not found. \n", key)
			continue
		}
		valueStr := ""
		srcType := ""
		switch v.(type) {
		case string:
			srcType = "string"
			valueStr = v.(string)
		case int64:
			srcType = "int64"
			valueStr = strconv.FormatInt(v.(int64), 10)
		case float64:
			srcType = "float64"
			valueStr = strconv.FormatFloat(v.(float64), 'f', -1, 64)
		case []interface{}:
			srcType = "[]interface{}"
		default:
			srcType = "other"
			log.Printf("not do,field is %s,src type is %s.\n", key, srcType)
		}
		tgtType := dvalue.Type().Name()
		if srcType == tgtType {
			//fmt.Printf("same,field is %s,src type is %s,tgt type is %s.\n",key,srcType,tgtType)
			vv := reflect.ValueOf(v)
			dvalue.Set(vv)
		} else {
			//fmt.Printf("diff,field is %s,src type is %s,tgt type is %s.\n",key,srcType,tgtType)
			switch tgtType {
			case "string":
				vv := reflect.ValueOf(valueStr)
				dvalue.Set(vv)
			case "int":
				v, _ := strconv.Atoi(valueStr)
				vv := reflect.ValueOf(v)
				dvalue.Set(vv)
			case "int8":
				v, _ := strconv.Atoi(valueStr)
				v2 := int8(v)
				vv := reflect.ValueOf(v2)
				dvalue.Set(vv)
			case "int64":
				v, _ := strconv.ParseInt(valueStr, 10, 64)
				vv := reflect.ValueOf(v)
				dvalue.Set(vv)
			case "float64":
				v, _ := strconv.ParseFloat(valueStr, 10)
				vv := reflect.ValueOf(v)
				dvalue.Set(vv)
			case "Time":
				if len(valueStr) == 10 {
					v, _ := time.ParseInLocation("2006-01-02", valueStr, time.Local)
					vv := reflect.ValueOf(v)
					dvalue.Set(vv)
				} else if len(valueStr) == 19 {
					v, _ := time.ParseInLocation("2006-01-02 15:04:05", valueStr, time.Local)
					vv := reflect.ValueOf(v)
					dvalue.Set(vv)
				} else {
					//v, _ := time.ParseInLocation("2006-01-02", "0000-01-01", time.Local)
					//vv := reflect.ValueOf(v)
					//dvalue.Set(vv)
				}
			default:
				log.Printf("error:field is %s,src type is %s,tgt type is %s.\n", key, srcType, tgtType)
			}
		}
	}
}

// Capitalize 字符首字母大写
func Capitalize(str string) string {
	var upperStr string
	vv := []rune(str) // 后文有介绍
	for i := 0; i < len(vv); i++ {
		if i == 0 {
			if vv[i] >= 97 && vv[i] <= 122 { // 后文有介绍
				vv[i] -= 32 // string的码表相差32位
				upperStr += string(vv[i])
			} else {
				log.Println("Not begins with lowercase letter,")
				return str
			}
		} else {
			upperStr += string(vv[i])
		}
	}
	return upperStr
}

////Map转struct(包含子struct)
func MapToStructWithOutTypeDeep(src map[string]interface{}, dst interface{}) interface{} {
	dval := reflect.ValueOf(dst).Elem()
	for k, v := range src {
		key := Capitalize(k)
		//println(key)
		dvalue := dval.FieldByName(key)
		if !dvalue.IsValid() {
			key = UnderlineLargeLetters(key)
			dvalue = dval.FieldByName(key)
			if !dvalue.IsValid() {
				continue
			}
		}
		valueStr := ""
		srcType := ""
		switch v.(type) {
		case string:
			srcType = "string"
			valueStr = v.(string)
		case int64:
			srcType = "int64"
			valueStr = strconv.FormatInt(v.(int64), 10)
		case float64:
			srcType = "float64"
			valueStr = strconv.FormatFloat(v.(float64), 'f', -1, 64)
		case []interface{}:
			srcType = "[]interface{}"
			if dvalue.Type().Kind() == reflect.Slice {
				SliceReflect := reflect.MakeSlice(dvalue.Type(), 0, 0)
				for _, item := range v.([]interface{}) {
					typ := dvalue.Type().Elem()
					ditem := reflect.New(typ)
					dst := MapToStructWithOutTypeDeep(item.(map[string]interface{}), ditem.Interface())
					SliceReflect = reflect.Append(SliceReflect, reflect.ValueOf(dst).Elem())
				}
				dvalue.Set(SliceReflect)
				continue
			}
		case map[string]interface{}:
			srcType = "map[string]interface{}"
			if dvalue.Type().Kind() == reflect.Struct {
				typ := dvalue.Type()
				ditem := reflect.New(typ)
				MapToStructWithOutTypeDeep(v.(map[string]interface{}), ditem.Interface())
				dvalue.Set(ditem.Elem())
				continue
			}
		default:
			srcType = "other"
			log.Printf("not do,field is %s,src type is %s.\n", key, srcType)
		}
		tgtType := dvalue.Type().Name()
		ptr := false
		if dvalue.Type().Kind().String() == "ptr" {
			elem := dvalue.Type().Elem()
			tgtType = elem.Name()
			ptr = true
		}
		if srcType == tgtType {
			//fmt.Printf("same,field is %s,src type is %s,tgt type is %s.\n",key,srcType,tgtType)
			vv := reflect.ValueOf(v)
			dvalue.Set(vv)
		} else {
			//fmt.Printf("diff,field is %s,src type is %s,tgt type is %s.\n",key,srcType,tgtType)
			switch tgtType {
			case "string":
				vv := reflect.ValueOf(valueStr)
				dvalue.Set(vv)
			case "int":
				v, _ := strconv.Atoi(valueStr)
				vv := reflect.ValueOf(v)
				dvalue.Set(vv)
			case "int8":
				v, _ := strconv.Atoi(valueStr)
				v2 := int8(v)
				vv := reflect.ValueOf(v2)
				dvalue.Set(vv)
			case "int64":
				v, _ := strconv.ParseInt(valueStr, 10, 64)
				vv := reflect.ValueOf(v)
				dvalue.Set(vv)
			case "float64":
				v, _ := strconv.ParseFloat(valueStr, 10)
				vv := reflect.ValueOf(v)
				dvalue.Set(vv)
			case "Time":
				if len(valueStr) == 10 {
					v, _ := time.ParseInLocation("2006-01-02", valueStr, time.Local)
					if ptr {
						vv := reflect.ValueOf(&v)
						dvalue.Set(vv)
					} else {
						vv := reflect.ValueOf(v)
						dvalue.Set(vv)
					}
				} else if len(valueStr) == 19 {
					v, _ := time.ParseInLocation("2006-01-02 15:04:05", valueStr, time.Local)
					if ptr {
						vv := reflect.ValueOf(&v)
						dvalue.Set(vv)
					} else {
						vv := reflect.ValueOf(v)
						dvalue.Set(vv)
					}
				}
			case "XjTime":
				if len(valueStr) == 10 {
					v1, _ := time.ParseInLocation("2006-01-02", valueStr, time.Local)
					v := xjtypes.XjTime(v1)
					if ptr {
						vv := reflect.ValueOf(&v)
						dvalue.Set(vv)
					} else {
						vv := reflect.ValueOf(v)
						dvalue.Set(vv)
					}
				} else if len(valueStr) == 19 {
					v1, _ := time.ParseInLocation("2006-01-02 15:04:05", valueStr, time.Local)
					v := xjtypes.XjTime(v1)
					if ptr {
						vv := reflect.ValueOf(&v)
						dvalue.Set(vv)
					} else {
						vv := reflect.ValueOf(v)
						dvalue.Set(vv)
					}
				}
			case "XjDate":
				if len(valueStr) >= 10 {
					v1, _ := time.ParseInLocation("2006-01-02", valueStr, time.Local)
					v := xjtypes.XjDate(v1)
					if ptr {
						vv := reflect.ValueOf(&v)
						dvalue.Set(vv)
					} else {
						vv := reflect.ValueOf(v)
						dvalue.Set(vv)
					}
				}
			default:
				log.Printf("error:field is %s,src type is %s,tgt type is %s.\n", key, srcType, tgtType)
			}
		}
	}
	return dst
}

func UnderlineLargeLetters(str string) string {
	index := strings.Index(str, "_")
	for index > -1 {
		str = str[:index] + Capitalize(str[index+1:])
		index = strings.Index(str, "_")
	}
	return str
}
