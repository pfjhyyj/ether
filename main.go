package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pfjhyyj/ether/app/auth"
	"github.com/pfjhyyj/ether/app/permission"
	"github.com/pfjhyyj/ether/app/user"
	"github.com/pfjhyyj/ether/clients/casbin"
	"github.com/pfjhyyj/ether/clients/gorm"
	"github.com/pfjhyyj/ether/clients/redis"
	"github.com/pfjhyyj/ether/middleware"
	"github.com/pfjhyyj/ether/utils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"golang.org/x/exp/rand"
	"sync"
	"time"
)

var (
	serviceMapping = map[string]func(){
		"api": runApiServer,
	}
)

func Init() {
	rand.Seed(uint64(time.Now().Unix()))
	utils.Init()
	{
		gorm.Init()
		redis.Init()
		casbin.Init()
	}
	{
		user.AutoMigrate()
	}
}

func runApiServer() {
	r := gin.New()
	r.Use(gin.Logger(), middleware.ErrorMiddleware(), middleware.TenantMiddleware(false))
	apiRouter := r.Group("/api")
	{
		auth.SetRouter(apiRouter)
		user.SetRouter(apiRouter)
		permission.SetRouter(apiRouter)
	}

	port := viper.GetString("service.api.port")
	_ = r.Run(":" + port)
}

func main() {
	Init()

	services := viper.GetStringSlice("service.enabled")
	wg := sync.WaitGroup{}
	for _, srv := range services {
		method, ok := serviceMapping[srv]
		if !ok {
			logrus.Errorf("fail to get srv %s", srv)
			continue
		}
		wg.Add(1)
		go func() {
			defer func() {
				if r := recover(); r != nil {
					logrus.Errorf("panic: %v", r)
				}
			}()
			method()
			wg.Done()
		}()
	}
	wg.Wait()
}
