package bootstrap

import (
	"github.com/spf13/viper"
	"iddd.com/iddd-go/app"
	"log"
)

var (
	appConfig *app.Config
)

func initConfig() {
	log.Println("init configuration...")
	appConfig = new(app.Config)
	appConfig.App.Port = 8888

	viper.SetConfigName("app")
	viper.AddConfigPath("config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("error reading config file, %s", err)
	}

	if err := viper.Unmarshal(appConfig); err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
}
