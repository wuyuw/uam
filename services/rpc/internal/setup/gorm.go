package setup

import (
	"log"
	"uam/services/rpc/internal/config"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupGormMysql(c config.Config) *gorm.DB {
	var (
		err error
		db  *gorm.DB
	)
	dsn := c.Mysql.Dsn()
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	if db, err = gorm.Open(mysql.New(mysqlConfig), gormConfig()); err != nil {
		log.Fatal(errors.Wrap(err, "MySQL初始化失败"))
	} else {
		// 连接池配置
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(c.Mysql.MaxIdleConns)
		sqlDB.SetMaxOpenConns(c.Mysql.MaxOpenConns)
		// sqlDB.SetConnMaxLifetime(time.Minute)
		logx.Info("MySQL链接初始化完成!")
	}
	return db

}

// gorm相关配置
func gormConfig() *gorm.Config {
	return &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}
}
