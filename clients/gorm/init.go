package gorm

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"sync"
)

var (
	db *gorm.DB

	once sync.Once
)

func Init() {
	once.Do(func() {
		host := viper.GetString("postgres.host")
		port := viper.GetString("postgres.port")
		dbName := viper.GetString("postgres.db")
		user := viper.GetString("postgres.user")
		password := viper.GetString("postgres.password")
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
			host, user, password, dbName, port)

		var err error
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		if viper.GetBool("postgres.debug") {
			db = db.Debug()
		}
	})
}

func GetDB() *gorm.DB {
	Init()
	return db
}
