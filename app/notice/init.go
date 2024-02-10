package notice

import (
	"github.com/pfjhyyj/ether/app/notice/model"
	"github.com/pfjhyyj/ether/clients/gorm"
	"github.com/spf13/viper"
)

func AutoMigrate() {
	if viper.GetBool("postgres.auto_migrate") {
		db := gorm.GetDB()
		if err := db.AutoMigrate(
			&model.Message{},
			&model.MessageText{},
		); err != nil {
			panic(err)
		}
	}
}
