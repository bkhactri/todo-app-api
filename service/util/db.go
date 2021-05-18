package util

import (
	Type "todo-app/service/api"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectCockRoachDB() {
	const addr = "postgresql://kh_c:buikhacTri123@free-tier6.gcp-asia-southeast1.cockroachlabs.cloud:26257/todo?sslmode=verify-full&sslrootcert=F:/cockroachdb/cc-ca.crt&options=--cluster=todo-app-114"
	db, err := gorm.Open(postgres.Open(addr), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Type.Task{})
}
