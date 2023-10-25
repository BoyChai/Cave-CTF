package dao

import (
	"Cave-CTF/config"
	"Cave-CTF/log"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// 数据库链接
func Client() {
	cfg := config.Config.DATA
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User,
		cfg.Pass,
		cfg.Host,
		cfg.Port,
		cfg.Database,
	)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: false,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false,
		},
		DisableForeignKeyConstraintWhenMigrating: true, // 逻辑外键（代码里自动体验外键关系）
	})
	if err != nil {
		log.NewCriticalLog(000, err, "数据库链接错误")
	}
}
