package main

import (
	"github.com/pfjhyyj/ether/clients/gorm"
	"github.com/pfjhyyj/ether/clients/redis"
	"golang.org/x/exp/rand"
	"time"
)

func Init() {
	rand.Seed(uint64(time.Now().Unix()))
	{
		gorm.Init()
		redis.Init()
	}
}

func main() {
	Init()

}
