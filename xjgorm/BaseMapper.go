package xjgorm

import (
	"github.com/xjieinfo/xjgo/xjcore/xjconv"
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

func (this BaseMapper) FirstString(sql string, args []interface{}) (item map[string]string, err error) {
	result := make(map[string]interface{})
	err = this.Gorm.Raw(sql, args...).First(&result).Error
	if err != nil {
		return
	}
	item = make(map[string]string)
	for k, v := range result {
		item[k] = xjconv.InterfaceToString(v)
	}
	return
}

func (this BaseMapper) Find(wrapper *xjtypes.GormWrapper, list interface{}) error {
	err := wrapper.SetDb(this.Gorm).Find(list).Error
	return err
}

func (this BaseMapper) FindString(sql string, args []interface{}) (list []map[string]string, err error) {
	var result []map[string]interface{}
	err = this.Gorm.Raw(sql, args...).Find(&result).Error
	if err != nil {
		return
	}
	for _, item := range result {
		it := make(map[string]string)
		for k, v := range item {
			it[k] = xjconv.InterfaceToString(v)
		}
		list = append(list, it)
	}
	return
}

func (this BaseMapper) Count(wrapper *xjtypes.GormWrapper, item interface{}, total *int64) error {
	wrapper.Current = 0
	wrapper.Size = 0
	err := wrapper.SetDb(this.Gorm).Model(item).Count(total).Error
	return err
}

func (this BaseMapper) FindCount(wrapper *xjtypes.GormWrapper, list interface{}, total *int64) error {
	err := this.Find(wrapper, list)
	if err != nil {
		return err
	}
	item := getSliceZeroItem(list)
	err = this.Count(wrapper, item, total)
	return err
}

func (this BaseMapper) FindCount3(wrapper *xjtypes.GormWrapper, list interface{}, total *int64) error {
	err := this.Find(wrapper, list)
	if err != nil {
		return err
	}
	table := getSliceTableName(list)
	wrapper.Current = 0
	wrapper.Size = 0
	err = wrapper.SetDb(this.Gorm).Table(table).Count(total).Error
	return err
}

func (this BaseMapper) FindCount2(wrapper *xjtypes.GormWrapper, list interface{}, total *int64) error {
	err := this.Find(wrapper, list)
	if err != nil {
		return err
	}
	list2 := reflect.ValueOf(list).Elem().Interface()
	len := reflect.ValueOf(list2).Len()
	if len > 0 {
		v := reflect.ValueOf(list2).Index(0)
		ve := v.Interface()
		err = this.Count(wrapper, ve, total)
		return err
	}
	return nil
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
