package mysql

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"school-robot-chat/model"
	"school-robot-chat/settings"
)

var db *gorm.DB

// Init 初始化MySQL连接
func Init(cfg *settings.MySQLConfig) (err error) {
	// "user:password@tcp(host:port)/dbname"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DB)
	db, err = gorm.Open("mysql", dsn)
	if err != nil {
		return err
	}

	err = db.AutoMigrate(&model.User{}, &model.Session{}).Error
	if err != nil {
		return errors.New("Unable autoMigrateDB - " + err.Error())
	}

	sqlDB := db.DB()
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	return nil
}

// Close 关闭MySQL连接
func Close() {
	_ = db.Close()
}
