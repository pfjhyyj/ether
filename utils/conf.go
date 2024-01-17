package utils

import (
	"fmt"

	"github.com/spf13/viper"
)

func InitViper() {
	path, name := GetConfPath()
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)
	viper.SetConfigName(name)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}

func GetConfPath() (string, string) {
	env := GetEnv()
	return "conf/", fmt.Sprintf("config.%s.yml", env)
}
