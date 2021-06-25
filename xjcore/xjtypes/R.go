package xjtypes

import "time"

type R struct {
	Code      int         `json:"code"`
	Timestamp int64       `json:"timestamp"`
	Msg       string      `json:"msg"`
	Data      interface{} `json:"data"`
}

func (r *R) Success(data interface{}) R {
	return R{
		Code:      0,
		Timestamp: time.Now().UnixNano() / 1e6,
		Msg:       "success",
		Data:      data,
	}
}

func (r *R) Fail(msg string) R {
	return R{
		Code:      1,
		Timestamp: time.Now().UnixNano() / 1e6,
		Msg:       msg,
	}
}

func (r *R) Make(code int, data interface{}, msg string) R {
	return R{
		Code:      code,
		Timestamp: time.Now().UnixNano() / 1e6,
		Msg:       msg,
		Data:      data,
	}
}
