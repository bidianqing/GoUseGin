package config

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var c *viper.Viper

// 加载配置文件
func Load(app *gin.Engine, env string) {
	c = viper.New()
	c.AddConfigPath("./configs")
	c.SetConfigType("json")

	// 读取对应环境的配置文件
	c.SetConfigName("appsettings." + env)
	if err := c.MergeInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	c.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
	c.WatchConfig()
}

func GetString(key string) (value string) {
	return c.GetString(key)
}
