package permission

import (
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/pfjhyyj/ether/app/permission/model"
	"github.com/pfjhyyj/ether/clients/gorm"
	"github.com/spf13/viper"
)

func AutoMigrate() {
	db := gorm.GetDB()
	if viper.GetBool("postgres.auto_migrate") {
		if err := db.AutoMigrate(
			&model.Role{}, &model.UserRole{}, &model.Permission{},
		); err != nil {
			panic(err)
		}
	} else {
		gormadapter.TurnOffAutoMigrate(db)
	}
}
