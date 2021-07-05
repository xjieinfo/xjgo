package xjgorm

import (
	"github.com/xjieinfo/xjgo/xjcore/xjtypes"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"os"
)

func GormInit(conf xjtypes.Datasource) *gorm.DB {
	if conf.Drivername == "mysql" {
		return GormMySqlInit(conf.Dsn)
	} else if conf.Drivername == "sqlserver" {
		return GormSqlServerInit(conf.Dsn)
	}
	return nil
}

func GormMySqlInit(dsn string) *gorm.DB {
	//log.Println("开始连接数据库...")
	Gorm, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
		},
	})
	if err != nil {
		log.Println("数据库连接失败")
		log.Println(err)
		os.Exit(1)
	}
	log.Println("连接数据库...OK")
	return Gorm
}

func GormSqlServerInit(dsn string) *gorm.DB {
	//log.Println("开始连接数据库...")
	Gorm, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
		},
	})
	if err != nil {
		log.Println("数据库连接失败")
		log.Println(err)
		os.Exit(1)
	}
	log.Println("连接数据库...OK")
	return Gorm
}
