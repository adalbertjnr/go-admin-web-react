package store

import (
	"github.com/souzagmu/go-admin-web-react/types"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnDB() *gorm.DB {

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: "root:root@tcp/go_admin",
	}), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&types.User{}, &types.Role{})

	return db
}
