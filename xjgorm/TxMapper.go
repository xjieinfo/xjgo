package xjgorm

import (
	"github.com/xjieinfo/xjgo/xjcore/xjconv"
	"github.com/xjieinfo/xjgo/xjcore/xjtypes"
	"gorm.io/gorm"
)

type TxMapper struct{}

func (this TxMapper) Create(tx *gorm.DB, item interface{}) (bool, error) {
	db := tx.Create(item)
	return db.RowsAffected > 0, db.Error
}

func (this TxMapper) Save(tx *gorm.DB, item interface{}) (bool, error) {
	db := tx.Save(item)
	return db.RowsAffected > 0, db.Error
}

func (this TxMapper) Delete(tx *gorm.DB, wrapper *xjtypes.GormWrapper, item interface{}) (bool, error) {
	db := wrapper.SetDb(tx).Delete(item)
	return db.RowsAffected > 0, db.Error
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

//func (this TxMapper) FindCount(tx *gorm.DB, wrapper *xjtypes.GormWrapper, list interface{}, total *int64) error {
//	err := wrapper.SetDb(tx).Find(list).Offset(0).Limit(-1).Count(total).Error
//	return err
//}

func (this TxMapper) Update(tx *gorm.DB, wrapper *xjtypes.GormWrapper, item interface{}, column string, value interface{}) (bool, error) {
	db := wrapper.SetDb(tx).Model(item).Update(column, value)
	return db.RowsAffected > 0, db.Error
}

func (this TxMapper) Updates(tx *gorm.DB, wrapper *xjtypes.GormWrapper, item interface{}, values interface{}) (bool, error) {
	db := wrapper.SetDb(tx).Model(item).Updates(values)
	return db.RowsAffected > 0, db.Error
}

func (this TxMapper) UpdatesItem(tx *gorm.DB, wrapper *xjtypes.GormWrapper, item interface{}) (bool, error) {
	db := wrapper.SetDb(tx).Model(item).Updates(wrapper.Sets)
	return db.RowsAffected > 0, db.Error
}

func (this TxMapper) UpdatesTable(tx *gorm.DB, wrapper *xjtypes.GormWrapper, table string) (bool, error) {
	db := wrapper.SetDb(tx).Table(table).Updates(wrapper.Sets)
	return db.RowsAffected > 0, db.Error
}
