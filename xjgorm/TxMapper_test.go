package xjgorm

import (
	"errors"
	"github.com/xjieinfo/xjgo/xjcore/xjgorm"
	"github.com/xjieinfo/xjgo/xjcore/xjtypes"
	"gorm.io/gorm"
	"log"
	"testing"
)

//var baseMapper BaseMapper

func init() {
	dsn := "root:root@tcp(localhost:3306)/test?charset=utf8mb4&parseTime=true"
	datasource := xjtypes.Datasource{
		Drivername: "mysql",
		Dsn:        dsn,
	}
	baseMapper.Gorm = xjgorm.GormInit(datasource)
}

//type Student struct {
//	ID       int
//	Username string
//	Address  string
//}
//
//func (this Student) TableName() string {
//	return "student"
//}

func Test_Transaction(t *testing.T) {
	err := baseMapper.Gorm.Transaction(func(tx *gorm.DB) error {
		var item Student
		TxMapper{}.First(tx, new(xjtypes.GormWrapper).Eq("id", 1), &item)
		item.Address = "星沙10"
		_, er := TxMapper{}.Save(tx, &item)
		if er != nil {
			return er
		}
		var item2 Student
		TxMapper{}.First(tx, new(xjtypes.GormWrapper).Eq("id", 3), &item2)
		item2.Address = "法新10"
		_, er = TxMapper{}.Save(tx, &item2)
		if er != nil {
			return er
		}
		var item3 Student
		TxMapper{}.First(tx, new(xjtypes.GormWrapper).Eq("id", 7), &item3)
		item3.Address = "小庄9"
		ok, er := TxMapper{}.Save(tx, &item3)
		if er != nil {
			return er
		}
		if ok > 0 {
			return errors.New("update error.")
		}
		return nil
	})
	if err != nil {
		log.Println(err)
	} else {
		log.Println("ok")
	}
}
