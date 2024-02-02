package config

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// 加载配置文件
func Load(app *gin.Engine, env string) {
	viper.AddConfigPath("./configs")
	viper.SetConfigType("json")

	// 公共的配置文件，无论哪个环境都加载
	viper.SetConfigName("appsettings")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	// 读取对应环境的配置文件
	viper.SetConfigName("appsettings." + env)
	if err := viper.MergeInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}
