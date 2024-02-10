package app

import (
	"github.com/pfjhyyj/ether/app/notice"
	"github.com/pfjhyyj/ether/app/tenant"
	"github.com/pfjhyyj/ether/app/user"
)

func Init() {
	user.AutoMigrate()
	tenant.AutoMigrate()
	notice.AutoMigrate()
}
