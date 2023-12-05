package main

import (
	"github.com/souzagmu/go-admin-web-react/types"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: "root:root@tcp/go_admin",
	}), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.Migrator().DropTable(&types.User{})

}
