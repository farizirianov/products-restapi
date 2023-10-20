package configs

import (
	"github.com/farizirianov/products-restapi/helper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	dsn := "root:your_password@tcp(127.0.0.1:3306)/store_table?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	helper.ErrorPanic(err)
}
