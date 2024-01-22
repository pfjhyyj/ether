package utils

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/pfjhyyj/ether/clients/gorm"
)

func InitPermission() {
	db := gorm.GetDB()
	a, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		panic(err)
	}
	e, _ := casbin.NewEnforcer("examples/rbac_model.conf", a)

	// Or you can use an existing DB "abc" like this:
	// The adapter will use the table named "casbin_rule".
	// If it doesn't exist, the adapter will create it automatically.
	// a := gormadapter.NewAdapter("mysql", "mysql_username:mysql_password@tcp(127.0.0.1:3306)/abc", true)

	// Load the policy from DB.
	e.LoadPolicy()

	// Check the permission.
	e.Enforce("alice", "data1", "read")

	// Modify the policy.
	// e.AddPolicy(...)
	// e.RemovePolicy(...)

	// Save the policy back to DB.
	e.SavePolicy()
}
