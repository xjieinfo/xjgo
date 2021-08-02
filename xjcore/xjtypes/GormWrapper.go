package xjtypes

import (
	"gorm.io/gorm"
)

type GormWrapper struct {
	Where   string
	Args    []interface{}
	Cols    string
	Orderby string
	Current int
	Size    int
	Sets    map[string]interface{}
}

func (this *GormWrapper) Set(field string, val interface{}) *GormWrapper {
	if this.Sets == nil {
		this.Sets = make(map[string]interface{})
	}
	this.Sets[field] = val
	return this
}

func (this *GormWrapper) Eq(field string, val interface{}) *GormWrapper {
	if this.Where != "" {
		this.Where += " and "
	}
	this.Where += field + "=?"
	if this.Args == nil || len(this.Args) == 0 {
		this.Args = make([]interface{}, 0)
	}
	this.Args = append(this.Args, val)
	return this
}

func (this *GormWrapper) Ge(field string, val interface{}) *GormWrapper {
	if this.Where != "" {
		this.Where += " and "
	}
	this.Where += field + ">=?"
	if this.Args == nil || len(this.Args) == 0 {
		this.Args = make([]interface{}, 0)
	}
	this.Args = append(this.Args, val)
	return this
}

func (this *GormWrapper) Le(field string, val interface{}) *GormWrapper {
	if this.Where != "" {
		this.Where += " and "
	}
	this.Where += field + "<=?"
	if this.Args == nil || len(this.Args) == 0 {
		this.Args = make([]interface{}, 0)
	}
	this.Args = append(this.Args, val)
	return this
}

func (this *GormWrapper) Like(field string, val interface{}) *GormWrapper {
	if this.Where != "" {
		this.Where += " and "
	}
	this.Where += field + "like('%',?,'%')"
	if this.Args == nil || len(this.Args) == 0 {
		this.Args = make([]interface{}, 0)
	}
	this.Args = append(this.Args, val)
	return this
}

func (this *GormWrapper) In(field string, val interface{}) *GormWrapper {
	if this.Where != "" {
		this.Where += " and "
	}
	this.Where += field + " in(?)"
	if this.Args == nil || len(this.Args) == 0 {
		this.Args = make([]interface{}, 0)
	}
	this.Args = append(this.Args, val)
	return this
}

func (this *GormWrapper) InSql(field string, val string) *GormWrapper {
	if this.Where != "" {
		this.Where += " and "
	}
	this.Where += field + " in(" + val + ")"
	return this
}

func (this *GormWrapper) OrderByDesc(field string) {
	if this.Orderby != "" {
		this.Orderby += ","
	}
	this.Orderby += field + " desc"
}

func (this *GormWrapper) SetDb(db *gorm.DB) *gorm.DB {
	if this.Where != "" {
		if len(this.Args) > 0 {
			db = db.Where(this.Where, this.Args...)
		} else {
			db = db.Where(this.Where)
		}
	}
	if this.Cols != "" {
		db = db.Select(this.Cols)
	}
	if this.Orderby != "" {
		db = db.Order(this.Orderby)
	}
	start := (this.Current - 1) * this.Size
	if start > 0 {
		db = db.Offset(start)
	}
	if this.Size > 0 {
		db = db.Limit(this.Size)
	}
	return db
}
