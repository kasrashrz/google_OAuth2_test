package configs

import (
	"fmt"
	"github.com/spf13/viper"
)

func InitializeViper() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetConfigType("yml")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("error reading config file, %s", err)
	}
}
