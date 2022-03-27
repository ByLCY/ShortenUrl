package db

import (
	"log"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	once sync.Once
	db   *gorm.DB
)

func NewDb() *gorm.DB {
	once.Do(func() {
		var err error
		dsn := "host=10.10.10.12 user=isaacl password=123abc dbname=isaacl port=5432 sslmode=disable TimeZone=Asia/Shanghai"
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("无法连接数据库: %v", err)
		}
	})

	return db
}
