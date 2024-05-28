package config

import (
	"fmt"
	"log"
	"stp/internal/pkg/util"

	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var Config = &Configuration{}

func Init() {
	// Init viper
	v := viper.New()

	// 根据环境，加载不同的配置文件
	if gin.Mode() == gin.ReleaseMode {
		v.SetConfigFile(util.GetCwd() + "/config.yaml")
	} else {
		v.SetConfigFile(util.GetCwd() + "/config/config.yaml")
	}

	v.SetConfigType("yaml")

	// Read config file
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("failed to read config file: %s", err))
	}

	// Watch changes for config file and enable hot reload
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		// Reload config
		if err := readConfig(v, Config); err != nil {
			log.Printf("ERR failed to hot reload config file: %s\n", err.Error())
		}
	})

	// Initialize global var for config
	if err := readConfig(v, Config); err != nil {
		log.Printf("ERR failed to load config file: %s\n", err.Error())
	}
}

func readConfig(v *viper.Viper, config *Configuration) error {
	if err := v.Unmarshal(config); err != nil {
		return err
	}

	return nil
}
