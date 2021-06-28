package xjtypes

import (
	"database/sql/driver"
	"encoding/binary"
	"errors"
	"strconv"
	"time"
)

type XjTime time.Time

const (
	timeFormart = "2006-01-02 15:04:05"
	dateFormart = "2006-01-02"
)

// GobEncode implements the gob.GobEncoder interface.
func (t XjTime) GobEncode() ([]byte, error) {
	return time.Time(t).MarshalBinary()
}

//GobDecode implements the gob.GobDecoder interface.
func (t *XjTime) GobDecode(data []byte) error {
	t1 := time.Time(*t)
	return t1.UnmarshalBinary(data)
}

func (t *XjTime) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+timeFormart+`"`, string(data), time.Local)
	*t = XjTime(now)
	return
}

func (t XjTime) MarshalJSON() ([]byte, error) {
	var tt time.Time
	tt = time.Time(t)
	if tt.IsZero() {
		b := make([]byte, 0, 2)
		b = append(b, '"')
		b = append(b, '"')
		return b, nil
	} else {
		b := make([]byte, 0, len(timeFormart)+2)
		b = append(b, '"')
		b = time.Time(t).AppendFormat(b, timeFormart)
		b = append(b, '"')
		return b, nil
	}
}

func (t XjTime) String() string {
	return time.Time(t).Format(timeFormart)
}

func (t XjTime) StringDiy(sFormat string) string {
	return time.Time(t).Format(sFormat)
}

func (t XjTime) Value() (driver.Value, error) {
	// MyTime 转换成 time.Time 类型
	tTime := time.Time(t)
	return tTime.Format("2006/01/02 15:04:05"), nil
}

func (t *XjTime) Scan(v interface{}) error {
	switch vt := v.(type) {
	case string:
		// 字符串转成 time.Time 类型
		tTime, _ := time.Parse("2006/01/02 15:04:05", vt)
		*t = XjTime(tTime)
	case time.Time:
		*t = XjTime(vt)
	default:
		return errors.New("类型处理错误")
	}
	return nil
}

type EsTime time.Time

func (t *EsTime) UnmarshalJSON(data []byte) (err error) {
	now := time.Unix(int64(binary.BigEndian.Uint32(data)/1000), 0)
	*t = EsTime(now)
	return
}

func (t EsTime) MarshalJSON() ([]byte, error) {
	if time.Time(t).IsZero() {
		return []byte("0"), nil
	} else {
		i := time.Time(t).Unix() * 1000
		str := strconv.FormatInt(i, 10)
		return []byte(str), nil
	}
}
func Int64ToBytes(i int64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(i))
	return buf
}

func BytesToInt64(buf []byte) int64 {
	return int64(binary.BigEndian.Uint64(buf))
}

//func (t EsTime) MarshalJSON() ([]byte, error) {
//var buf = make([]byte, len(timeFormart)+2)
//if time.Time(t).IsZero(){
//	buf = []byte(`"1970-01-01 00:00:00"`)
//}else{
//	i := time.Time(t).Unix()*1000
//	binary.BigEndian.PutUint64(buf, uint64(i))
//}
//return buf, nil
//
//	b := make([]byte, 0, len(timeFormart)+2)
//	b = append(b, '"')
//	b = time.Time(t).AppendFormat(b, timeFormart)
//	b = append(b, '"')
//	if string(b)==`"0001-01-01 00:00:00"`{
//		b = []byte(`"1970-01-01 00:00:00"`)
//	}
//	return b, nil
//}

type EsDate time.Time

func (t *EsDate) UnmarshalJSON(data []byte) (err error) {
	now := time.Unix(int64(binary.BigEndian.Uint32(data)/1000), 0)
	*t = EsDate(now)
	return
}

func (t EsDate) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(dateFormart)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, dateFormart)
	b = append(b, '"')
	if string(b) == `"0001-01-01"` {
		b = []byte(`""`)
	}
	return b, nil
}

type XjDate time.Time

// GobEncode implements the gob.GobEncoder interface.
func (t XjDate) GobEncode() ([]byte, error) {
	return time.Time(t).MarshalBinary()
}

//GobDecode implements the gob.GobDecoder interface.
func (t *XjDate) GobDecode(data []byte) error {
	t1 := time.Time(*t)
	return t1.UnmarshalBinary(data)
}

func (t *XjDate) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+dateFormart+`"`, string(data), time.Local)
	*t = XjDate(now)
	return
}

func (t XjDate) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(dateFormart)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, dateFormart)
	b = append(b, '"')
	return b, nil
}

func (t XjDate) String() string {
	return time.Time(t).Format(dateFormart)
}

func (t XjDate) StringDiy(sFormat string) string {
	return time.Time(t).Format(sFormat)
}

func (t XjDate) Value() (driver.Value, error) {
	// MyTime 转换成 time.Time 类型
	tTime := time.Time(t)
	return tTime.Format("2006/01/02"), nil
}

func (t *XjDate) Scan(v interface{}) error {
	switch vt := v.(type) {
	case string:
		// 字符串转成 time.Time 类型
		tTime, _ := time.Parse("2006/01/02", vt)
		*t = XjDate(tTime)
	case time.Time:
		*t = XjDate(vt)
	default:
		return errors.New("类型处理错误")
	}
	return nil
}

type Json string

func (j *Json) UnmarshalJSON(data []byte) (err error) {
	*j = Json(data)
	return
}

func (j Json) MarshalJSON() ([]byte, error) {
	b := []byte(j)
	return b, nil
}
