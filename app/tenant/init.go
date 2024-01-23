package tenant

import (
	"github.com/pfjhyyj/ether/app/tenant/model"
	"github.com/pfjhyyj/ether/clients/gorm"
	"github.com/spf13/viper"
)

func AutoMigrate() {
	if viper.GetBool("postgres.auto_migrate") {
		db := gorm.GetDB()
		if err := db.AutoMigrate(
			&model.Tenant{},
		); err != nil {
			panic(err)
		}
	}
}
