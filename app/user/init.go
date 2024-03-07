package user

import (
	"github.com/pfjhyyj/ether/app/user/model"
	"github.com/pfjhyyj/ether/client/gorm"
	"github.com/spf13/viper"
)

func AutoMigrate() {
	if viper.GetBool("postgres.auto_migrate") {
		db := gorm.GetDB()
		if err := db.AutoMigrate(
			&model.User{},
			&model.Role{},
			&model.UserRole{},
			&model.Permission{},
			&model.RolePermission{},
			&model.Menu{},
			&model.RoleMenu{},
		); err != nil {
			panic(err)
		}
	}
}
