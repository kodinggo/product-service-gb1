package db

import (
	"fmt"

	"github.com/kodinggo/product-service-gb1/config"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMysql() *gorm.DB {
	// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.MysqlUser(), config.MysqlPass(), config.MysqlHost(), config.MysqlDB())
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Fatalf("Failed to connect to database: %v", err)
	}

	return db
}
