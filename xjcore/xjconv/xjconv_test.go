package xjconv

import "testing"

func Test_Int64ArrayToString(t *testing.T) {
	arr := []int64{123, 456}
	str := Int64ArrayToString(arr)
	println(str)
}
