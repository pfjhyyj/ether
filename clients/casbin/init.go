package casbin

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/pfjhyyj/ether/clients/gorm"
	"sync"
)

var (
	enforcer *casbin.Enforcer

	once sync.Once
)

func GetModel() string {
	return `
		[request_definition]
		r = sub, dom, obj, act
		
		[policy_definition]
		p = sub, dom, obj, act
		
		[role_definition]
		g = _, _, _
		
		[policy_effect]
		e = some(where (p.eft == allow))
		
		[matchers]
		m = g(r.sub, p.sub, r.dom) && r.dom == p.dom && r.obj == p.obj && r.act == p.act || r.sub == "root"
	`
}

func Init() {
	once.Do(func() {
		db := gorm.GetDB()
		a, err := gormadapter.NewAdapterByDB(db)
		if err != nil {
			panic(err)
		}

		m, err := model.NewModelFromString(GetModel())
		if err != nil {
			panic(err)
		}

		enforcer, err = casbin.NewEnforcer(m, a)
	})
}

func GetEnforcer() *casbin.Enforcer {
	Init()
	return enforcer
}
