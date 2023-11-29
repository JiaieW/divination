package database

import (
	"divination/config"
	"fmt"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	once   sync.Once
	DBConn *gorm.DB
)

func init() {
	once.Do(func() {
		var err error
		config := config.C
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			config.Database.User,
			config.Database.Password,
			config.Database.Host,
			config.Database.Port,
			config.Database.Name,
		)
		DBConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("failed to connect to database")
		}
	})
}
