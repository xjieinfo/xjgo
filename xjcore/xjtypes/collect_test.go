package xjtypes

import (
	"fmt"
	"testing"
)

type Student struct {
	ID       int
	Username string
	Address  string
	Amount   float64
}

func Test_CollectString(t *testing.T) {
	list := []Student{
		{1, "aaa", "aa1", 110.5},
		{2, "bbb", "bb1", 200.8},
		{1, "aaa", "aa1", 110.5},
	}
	names := CollectString(list, "Username")
	fmt.Println(names)
}

func Test_CollectInt(t *testing.T) {
	list := []Student{
		{1, "aaa", "aa1", 110.5},
		{2, "bbb", "bb1", 200.8},
		{1, "aaa", "aa1", 110.5},
	}
	ids := CollectInt(list, "ID")
	fmt.Println(ids)
}

func Test_CollectInt64(t *testing.T) {
	list := []Student{
		{1, "aaa", "aa1", 110.5},
		{2, "bbb", "bb1", 200.8},
		{1, "aaa", "aa1", 110.5},
	}
	ids := CollectInt64(list, "ID")
	fmt.Println(ids)
}

func Test_CollectFloat64(t *testing.T) {
	list := []Student{
		{1, "aaa", "aa1", 110.5},
		{2, "bbb", "bb1", 200.8},
		{1, "aaa", "aa1", 110.5},
	}
	vals := CollectFloat64(list, "Amount")
	fmt.Println(vals)
}

func Test_CollectSetString(t *testing.T) {
	list := []Student{
		{1, "aaa", "aa1", 110.5},
		{2, "bbb", "bb1", 200.8},
		{1, "aaa", "aa1", 110.5},
	}
	names := CollectSetString(list, "Username")
	fmt.Println(names)
}

func Test_CollectSetInt(t *testing.T) {
	list := []Student{
		{1, "aaa", "aa1", 110.5},
		{2, "bbb", "bb1", 200.8},
		{1, "aaa", "aa1", 110.5},
	}
	ids := CollectSetInt(list, "ID")
	fmt.Println(ids)
}

func Test_CollectSetInt64(t *testing.T) {
	list := []Student{
		{1, "aaa", "aa1", 110.5},
		{2, "bbb", "bb1", 200.8},
		{1, "aaa", "aa1", 110.5},
	}
	ids := CollectSetInt64(list, "ID")
	fmt.Println(ids)
}

func Test_CollectSetFloat64(t *testing.T) {
	list := []Student{
		{1, "aaa", "aa1", 110.5},
		{2, "bbb", "bb1", 200.8},
		{1, "aaa", "aa1", 110.5},
	}
	vals := CollectSetFloat64(list, "Amount")
	fmt.Println(vals)
}
