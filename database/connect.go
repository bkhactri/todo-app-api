package db

import (
	"fmt"

	"github.com/tobias/todo-app/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const addr = "postgresql://kh_c:123456789123@free-tier6.gcp-asia-southeast1.cockroachlabs.cloud:26257/todo?sslmode=verify-full&sslrootcert=F:/cockroachdb/cc-ca.crt&options=--cluster=todo-app-114"

func ConnectTask() *gorm.DB {
	db, err := gorm.Open(postgres.Open(addr), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	db.AutoMigrate(&models.Task{})

	return db
}

func ConnectUser() *gorm.DB {
	db, err := gorm.Open(postgres.Open(addr), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	db.AutoMigrate(&models.User{})

	return db
}
