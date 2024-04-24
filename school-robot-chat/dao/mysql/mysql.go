package mysql

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
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
	//dao.SetMaxOpenConns(cfg.MaxOpenConns)
	//dao.SetMaxIdleConns(cfg.MaxIdleConns)
	return err
}

// Close 关闭MySQL连接
func Close() {
	_ = db.Close()
}
