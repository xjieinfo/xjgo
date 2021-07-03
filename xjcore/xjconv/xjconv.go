package xjconv

import (
	"encoding/json"
	"fmt"
	"gitee.com/xjieinfo/xjgo/xjcore/xjtypes"
	"strconv"
	"strings"
	"time"
)

func JsonToIntArray(str string) ([]int, error) {
	var ints []int
	err := json.Unmarshal([]byte(str), ints)
	return ints, err
}

func StringToIntArray(s string) ([]int, error) {
	var ints []int
	ss := strings.Split(s, ",")
	for _, s1 := range ss {
		i1, err := strconv.Atoi(s1)
		if err == nil {
			ints = append(ints, i1)
		} else {
			return ints, err
		}
	}
	return ints, nil
}

func StringToInt64Array(s string) ([]int64, error) {
	var ints []int64
	ss := strings.Split(s, ",")
	for _, s1 := range ss {
		i1, err := strconv.ParseInt(s1, 10, 64)
		if err == nil {
			ints = append(ints, i1)
		} else {
			return ints, err
		}
	}
	return ints, nil
}

func InterfaceArrayToInt64Array(src []interface{}) ([]int64, error) {
	var ints []int64
	for _, s1 := range src {
		i1 := InterfaceToInt64(s1)
		ints = append(ints, i1)
	}
	return ints, nil
}

func InterfaceArrayToIntArray(src []interface{}) ([]int, error) {
	var ints []int
	for _, s1 := range src {
		i1 := InterfaceToInt(s1)
		ints = append(ints, i1)
	}
	return ints, nil
}

func StringToInterfaceArray(s string) ([]interface{}, error) {
	var ints []interface{}
	ss := strings.Split(s, ",")
	for _, s1 := range ss {
		i1, err := strconv.ParseInt(s1, 10, 64)
		if err == nil {
			ints = append(ints, i1)
		} else {
			return ints, err
		}
	}
	return ints, nil
}
func IntArrayToString(ints []int) string {
	s := ""
	for _, i1 := range ints {
		if s != "" {
			s += ","
		}
		s += strconv.Itoa(i1)
	}
	return s
}

func Int64ArrayToString(ints []int64) string {
	s := ""
	for _, i1 := range ints {
		if s != "" {
			s += ","
		}
		s += strconv.FormatInt(i1, 10)
	}
	return s
}

func InterfaceArrayToString(ints []interface{}) string {
	s := ""
	for _, i1 := range ints {
		if s != "" {
			s += ","
		}
		s += InterfaceToString(i1)
	}
	return s
}

func InterfaceToString(val interface{}) string {
	if val != nil {
		switch val.(type) {
		case bool:
			return strconv.FormatBool(val.(bool))
		case string:
			return val.(string)
		case int8, int, int32, int64:
			strV := fmt.Sprintf("%d", val)
			return strV
		case float32:
			strV := fmt.Sprintf("%f", val)
			return strV
		case float64:
			strV := fmt.Sprintf("%f", val)
			return strV
		default:
			strV := fmt.Sprintf("%s", val)
			return strV
		}
	}
	return ""
}
func InterfaceToDate(val interface{}) xjtypes.XjDate {
	var date xjtypes.XjDate
	if val != nil {
		strV := ""
		switch val.(type) {
		case bool:
			strV = strconv.FormatBool(val.(bool))
		case string:
			strV = val.(string)
		case int8, int, int32, int64:
			strV = fmt.Sprintf("%d", val)
		case float32:
			strV = fmt.Sprintf("%f", val)
		case float64:
			strV = fmt.Sprintf("%f", val)
		default:
			strV = fmt.Sprintf("%s", val)
		}
		if len(strV) == 10 {
			v, _ := time.ParseInLocation("2006-01-02", strV, time.Local)
			return xjtypes.XjDate(v)
		} else if len(strV) >= 19 {
			v, _ := time.ParseInLocation("2006-01-02 15:04:05", strV[:19], time.Local)
			return xjtypes.XjDate(v)
		} else {
			return date
		}
	}
	return date
}

func InterfaceToTime(val interface{}) xjtypes.XjTime {
	var date xjtypes.XjTime
	if val != nil {
		strV := ""
		switch val.(type) {
		case bool:
			strV = strconv.FormatBool(val.(bool))
		case string:
			strV = val.(string)
		case int8, int, int32, int64:
			strV = fmt.Sprintf("%d", val)
		case float32:
			strV = fmt.Sprintf("%f", val)
		case float64:
			strV = fmt.Sprintf("%f", val)
		default:
			strV = fmt.Sprintf("%s", val)
		}
		if len(strV) == 10 {
			v, _ := time.ParseInLocation("2006-01-02", strV, time.Local)
			return xjtypes.XjTime(v)
		} else if len(strV) >= 19 {
			v, _ := time.ParseInLocation("2006-01-02 15:04:05", strV[:19], time.Local)
			return xjtypes.XjTime(v)
		} else {
			return date
		}
	}
	return date
}

func InterfaceToBool(val interface{}) bool {
	if val != nil {
		strV := ""
		switch val.(type) {
		case bool:
			strV = strconv.FormatBool(val.(bool))
		case string:
			strV = val.(string)
		case int8, int, int32, int64:
			strV = fmt.Sprintf("%d", val)
		case float32:
			strV = fmt.Sprintf("%f", val)
		case float64:
			strV = fmt.Sprintf("%f", val)
		default:
			strV = fmt.Sprintf("%s", val)
		}
		if strV == "1" || strV == "true" || strV == "True" {
			return true
		}
	}
	return false
}

func InterfaceToInt(val interface{}) int {
	var intval int
	var strval string
	if val != nil {
		switch val.(type) {
		case bool:
			strval = strconv.FormatBool(val.(bool))
		case string:
			strval = val.(string)
		case int8, int, int32, int64:
			strval = fmt.Sprintf("%d", val)
		case float32:
			strval = fmt.Sprintf("%.0f", val)
		case float64:
			strval = fmt.Sprintf("%.0f", val)
		default:
			strval = fmt.Sprintf("%s", val)
		}
	}
	var err error
	intval, err = strconv.Atoi(strval)
	if err != nil {
		fmt.Println(err)
	}
	return intval
}

func InterfaceToInt64(val interface{}) int64 {
	var intval int64
	var strval string
	if val != nil {
		switch val.(type) {
		case bool:
			strval = strconv.FormatBool(val.(bool))
		case string:
			strval = val.(string)
		case int8, int, int32, int64:
			strval = fmt.Sprintf("%d", val)
		case float32:
			strval = fmt.Sprintf("%.0f", val)
		case float64:
			strval = fmt.Sprintf("%.0f", val)
		default:
			strval = fmt.Sprintf("%s", val)
		}
	}
	var err error
	intval, err = strconv.ParseInt(strval, 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	return intval
}
func InterfaceToFloat64(val interface{}) float64 {
	var floatval float64
	var strval string
	if val != nil {
		switch val.(type) {
		case bool:
			strval = strconv.FormatBool(val.(bool))
		case string:
			strval = val.(string)
		case int8, int, int32, int64:
			strval = fmt.Sprintf("%d", val)
		case float32:
			strval = fmt.Sprintf("%.0f", val)
		case float64:
			strval = fmt.Sprintf("%.0f", val)
		default:
			strval = fmt.Sprintf("%s", val)
		}
	}
	var err error
	floatval, err = strconv.ParseFloat(strval, 64)
	if err != nil {
		fmt.Println(err)
	}
	return floatval
}

func Int64ArrayToInterfaceArray(ints []int64) []interface{} {
	list := make([]interface{}, 0)
	for _, i := range ints {
		list = append(list, i)
	}
	return list
}

func IntToFloat64(val int) float64 {
	var floatval float64
	strval := fmt.Sprintf("%d", val)
	var err error
	floatval, err = strconv.ParseFloat(strval, 64)
	if err != nil {
		fmt.Println(err)
	}
	return floatval
}

func Float64ToInt(val float64) int {
	var intval int64
	strval := fmt.Sprintf("%f", val)
	var err error
	intval, err = strconv.ParseInt(strval, 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	return int(intval)
}

func Float64ToInt64(val float64) int64 {
	var intval int64
	strval := fmt.Sprintf("%f", val)
	var err error
	intval, err = strconv.ParseInt(strval, 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	return intval
}
