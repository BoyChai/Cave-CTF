package config

import (
	"github.com/spf13/viper"
	"os"
)

func (c *config) Read() {
	viper.SetConfigType("yaml")
	cave, _ := os.Open("conf/cave.yaml")
	_ = viper.ReadConfig(cave)
	_ = viper.Unmarshal(&c)
}
