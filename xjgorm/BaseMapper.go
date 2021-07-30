package xjgorm

import (
	"github.com/xjieinfo/xjgo/xjcore/xjtypes"
	"gorm.io/gorm"
	"reflect"
)

type BaseMapper struct {
	Gorm *gorm.DB
}

func (this BaseMapper) Create(item interface{}) (bool, error) {
	db := this.Gorm.Create(item)
	return db.RowsAffected > 0, db.Error
}

func (this BaseMapper) Save(item interface{}) (bool, error) {
	db := this.Gorm.Save(item)
	return db.RowsAffected > 0, db.Error
}

func (this BaseMapper) Delete(wrapper *xjtypes.GormWrapper, item interface{}) (bool, error) {
	db := wrapper.SetDb(this.Gorm).Delete(item)
	return db.RowsAffected > 0, db.Error
}

func (this BaseMapper) First(wrapper *xjtypes.GormWrapper, item interface{}) error {
	err := wrapper.SetDb(this.Gorm).First(item).Error
	return err
}

func (this BaseMapper) Find(wrapper *xjtypes.GormWrapper, list interface{}) error {
	err := wrapper.SetDb(this.Gorm).Find(list).Error
	return err
}

func (this BaseMapper) Count(wrapper *xjtypes.GormWrapper, item interface{}, total *int64) error {
	err := wrapper.SetDb(this.Gorm).Model(item).Count(total).Error
	return err
}

func (this BaseMapper) FindCount(wrapper *xjtypes.GormWrapper, list interface{}, total *int64) error {
	err := wrapper.SetDb(this.Gorm).Find(list).Offset(0).Limit(-1).Count(total).Error
	return err
}

func (this BaseMapper) Update(wrapper *xjtypes.GormWrapper, item interface{}, column string, value interface{}) (bool, error) {
	db := wrapper.SetDb(this.Gorm).Model(item).Update(column, value)
	return db.RowsAffected > 0, db.Error
}

func (this BaseMapper) Updates(wrapper *xjtypes.GormWrapper, item interface{}, values interface{}) (bool, error) {
	db := wrapper.SetDb(this.Gorm).Model(item).Updates(values)
	return db.RowsAffected > 0, db.Error
}

func (this BaseMapper) UpdatesItem(wrapper *xjtypes.GormWrapper, item interface{}) (bool, error) {
	db := wrapper.SetDb(this.Gorm).Model(item).Updates(wrapper.Sets)
	return db.RowsAffected > 0, db.Error
}

func (this BaseMapper) UpdatesTable(wrapper *xjtypes.GormWrapper, table string) (bool, error) {
	db := wrapper.SetDb(this.Gorm).Table(table).Updates(wrapper.Sets)
	return db.RowsAffected > 0, db.Error
}

func getTableName(item interface{}) string {
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