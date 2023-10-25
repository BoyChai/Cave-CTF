package config

import (
	"fmt"
	"github.com/BoyChai/Ant"
	"github.com/spf13/viper"
	"os"
)

func (c *config) Read() {
	var err error
	viper.SetConfigType("yaml")
	cave, _ := os.Open("conf/cave.yaml")
	err = viper.ReadConfig(cave)
	if err != nil {
		fmt.Println("Log:viper读取配置出现错误:", err)
		return
	}
	err = viper.Unmarshal(&c)
	if err != nil {
		fmt.Println("Log:viper加载配置:", err)
		return
	}
	validator := Ant.New(Ant.Validator{})
	e := validator.Struct(Config)
	if e.Is {
		fmt.Println("Log:数据校验错误 ", err.Error())
		return
	}
}
