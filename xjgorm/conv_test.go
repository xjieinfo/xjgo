package xjgorm

import "testing"

func Test_StructTableName(t *testing.T) {
	var stu Student
	name := getStructTableName(stu)
	println(name)
}

func Test_SliceTableName(t *testing.T) {
	var list []Student
	name := getSliceTableName(list)
	println(name)
}
