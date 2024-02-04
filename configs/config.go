package config

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
)

var v *viper.Viper

// 加载配置文件
func Load(app *gin.Engine, env string) {
	v = viper.New()
	addJson(env)
	addConsul(env)
}

func GetString(key string) (value string) {
	return v.GetString(key)
}

func addJson(env string) {
	v.AddConfigPath("./configs")
	v.SetConfigType("json")

	// 读取对应环境的配置文件
	v.SetConfigName("appsettings." + env)
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
	v.WatchConfig()
}

func addConsul(env string) {
	consul_url := v.GetString("Consul_Url")
	v.AddRemoteProvider("consul", consul_url, "appsettings."+env+".json")
	v.SetConfigType("json")
	if err := v.ReadRemoteConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	v.WatchRemoteConfigOnChannel()
}
