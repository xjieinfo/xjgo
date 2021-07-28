package xjgorm

import (
	"fmt"
	"github.com/xjieinfo/xjgo/xjcore/xjgorm"
	"github.com/xjieinfo/xjgo/xjcore/xjtypes"
	"testing"
)

var baseMapper BaseMapper

func init() {
	dsn := "root:root@tcp(localhost:3306)/test?charset=utf8mb4&parseTime=true"
	datasource := xjtypes.Datasource{
		Drivername: "mysql",
		Dsn:        dsn,
	}
	baseMapper.Gorm = xjgorm.GormInit(datasource)
}

type Student struct {
	ID       int
	Username string
	Address  string
}

func (this Student) TableName() string {
	return "student"
}

func Test_Create(t *testing.T) {
	item := Student{
		Username: "张三",
		Address:  "长沙市",
	}
	flag, _ := baseMapper.Create(&item)
	fmt.Println(flag)
}

func Test_Save(t *testing.T) {
	var item Student
	baseMapper.First(new(xjtypes.GormWrapper).Eq("id", 2), &item)
	item.Address = "星沙"
	flag, _ := baseMapper.Save(&item)
	fmt.Println(flag)
}

func Test_Delete(t *testing.T) {
	var item Student
	flag, _ := baseMapper.Delete(new(xjtypes.GormWrapper).Eq("id", 2), &item)
	fmt.Println(flag)
}

func Test_Find(t *testing.T) {
	var list []Student
	baseMapper.Find(new(xjtypes.GormWrapper).Eq("username", "张三"), &list)
	fmt.Println(list)
}

func Test_Count(t *testing.T) {
	var item Student
	var total int64
	baseMapper.Count(new(xjtypes.GormWrapper).Eq("username", "张三"), &item, &total)
	fmt.Println(total)
}

func Test_ListCount(t *testing.T) {
	var list []Student
	var total int64
	baseMapper.FindCount(new(xjtypes.GormWrapper).Eq("username", "张三"), &list, &total)
	fmt.Println(list)
	fmt.Println(total)
}
