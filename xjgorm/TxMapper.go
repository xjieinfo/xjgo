package xjgorm

import (
	"github.com/xjieinfo/xjgo/xjcore/xjconv"
	"github.com/xjieinfo/xjgo/xjcore/xjtypes"
	"gorm.io/gorm"
	"reflect"
)

type TxMapper struct{}

func (this TxMapper) Create(tx *gorm.DB, item interface{}) (int64, error) {
	db := tx.Create(item)
	return db.RowsAffected, db.Error
}

func (this TxMapper) Save(tx *gorm.DB, item interface{}) (int64, error) {
	db := tx.Save(item)
	return db.RowsAffected, db.Error
}

func (this TxMapper) Delete(tx *gorm.DB, wrapper *xjtypes.GormWrapper, item interface{}) (int64, error) {
	db := wrapper.SetDb(tx).Delete(item)
	return db.RowsAffected, db.Error
}

func (this TxMapper) First(tx *gorm.DB, wrapper *xjtypes.GormWrapper, item interface{}) error {
	err := wrapper.SetDb(tx).First(item).Error
	return err
}

func (this TxMapper) FirstString(tx *gorm.DB, sql string, args []interface{}) (item map[string]string, err error) {
	var result = make(map[string]interface{})
	err = tx.Raw(sql, args...).First(&result).Error
	if err != nil {
		return
	}
	item = make(map[string]string)
	for k, v := range item {
		item[k] = xjconv.InterfaceToString(v)
	}
	return
}

func (this TxMapper) Find(tx *gorm.DB, wrapper *xjtypes.GormWrapper, list interface{}) error {
	err := wrapper.SetDb(tx).Find(list).Error
	return err
}

func (this TxMapper) FindString(tx *gorm.DB, sql string, args []interface{}) (list []map[string]string, err error) {
	var result []map[string]interface{}
	err = tx.Raw(sql, args...).Find(&result).Error
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

func (this TxMapper) Count(tx *gorm.DB, wrapper *xjtypes.GormWrapper, item interface{}, total *int64) error {
	wrapper.Current = 0
	wrapper.Size = 0
	err := wrapper.SetDb(tx).Model(item).Count(total).Error
	return err
}

func (this TxMapper) FindCount(tx *gorm.DB, wrapper *xjtypes.GormWrapper, list interface{}, total *int64) error {
	err := this.Find(tx, wrapper, list)
	if err != nil {
		return err
	}
	item := getSliceZeroItem(list)
	err = this.Count(tx, wrapper, item, total)
	return err
}

func (this TxMapper) FindCount2(tx *gorm.DB, wrapper *xjtypes.GormWrapper, list interface{}, total *int64) error {
	err := this.Find(tx, wrapper, list)
	if err != nil {
		return err
	}
	list2 := reflect.ValueOf(list).Elem().Interface()
	len := reflect.ValueOf(list2).Len()
	if len > 0 {
		v := reflect.ValueOf(list2).Index(0)
		ve := v.Interface()
		err = this.Count(tx, wrapper, ve, total)
		return err
	}
	return nil
}

func (this TxMapper) Update(tx *gorm.DB, wrapper *xjtypes.GormWrapper, item interface{}, column string, value interface{}) (int64, error) {
	db := wrapper.SetDb(tx).Model(item).Update(column, value)
	return db.RowsAffected, db.Error
}

func (this TxMapper) Updates(tx *gorm.DB, wrapper *xjtypes.GormWrapper, item interface{}, values interface{}) (int64, error) {
	db := wrapper.SetDb(tx).Model(item).Updates(values)
	return db.RowsAffected, db.Error
}

func (this TxMapper) UpdatesItem(tx *gorm.DB, wrapper *xjtypes.GormWrapper, item interface{}) (int64, error) {
	db := wrapper.SetDb(tx).Model(item).Updates(wrapper.Sets)
	return db.RowsAffected, db.Error
}

func (this TxMapper) UpdatesTable(tx *gorm.DB, wrapper *xjtypes.GormWrapper, table string) (int64, error) {
	db := wrapper.SetDb(tx).Table(table).Updates(wrapper.Sets)
	return db.RowsAffected, db.Error
}
