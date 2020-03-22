package bootstrap

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"strings"
)

var App *Application

type Config viper.Viper

type Application struct {
	Name      string `json:"name"`
	Version   string `json:"version"`
	ENV       string `json:"env"`
	AppConfig Config `json:"application_config"`
	DBConfig  Config `json:"database_config"`
}

func init() {
	App = &Application{}
	App.Name = "APP_NAME"
	App.Version = "APP_VERSION"
	App.loadAppConfig()
	App.loadENV()
	App.loadDBConfig()
}

// loadAppConfig: read application config and build viper object
func (app *Application) loadAppConfig() {
	var (
		appConfig *viper.Viper
		err       error
	)
	appConfig = viper.New()
	appConfig.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	appConfig.SetEnvPrefix("APP_")
	appConfig.AutomaticEnv()
	appConfig.SetConfigName("config")
	appConfig.AddConfigPath(".")
	appConfig.SetConfigType("yaml")
	if err = appConfig.ReadInConfig(); err != nil {
		panic(err)
	}
	appConfig.WatchConfig()
	appConfig.OnConfigChange(func(e fsnotify.Event) {
		//	glog.Info("App Config file changed %s:", e.Name)
	})
	app.AppConfig = Config(*appConfig)
}

// loadDBConfig: read application config and build viper object
func (app *Application) loadDBConfig() {
	var (
		dbConfig *viper.Viper
		err      error
	)
	dbConfig = viper.New()
	dbConfig.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	dbConfig.SetEnvPrefix("DB_")
	dbConfig.AutomaticEnv()
	dbConfig.SetConfigName("config")
	dbConfig.AddConfigPath(".")
	dbConfig.SetConfigType("yaml")
	if err = dbConfig.ReadInConfig(); err != nil {
		panic(err)
	}
	dbConfig.WatchConfig()
	dbConfig.OnConfigChange(func(e fsnotify.Event) {
		//	glog.Info("App Config file changed %s:", e.Name)
	})
	app.DBConfig = Config(*dbConfig)
}

// loadENV
func (app *Application) loadENV() {
	var APPENV string
	var appConfig viper.Viper
	appConfig = viper.Viper(app.AppConfig)
	APPENV = appConfig.GetString("app_env")
	switch APPENV {
	case "dev":
		app.ENV = "dev"
		break
	case "staging":
		app.ENV = "staging"
		break
	case "production":
		app.ENV = "production"
		break
	default:
		app.ENV = "dev"
		break
	}
}

// String: read string value from viper.Viper
func (config *Config) String(key string) string {
	var viperConfig viper.Viper
	viperConfig = viper.Viper(*config)
	return viperConfig.GetString(fmt.Sprintf("%s.%s", App.ENV, key))
}

// Int: read int value from viper.Viper
func (config *Config) Int(key string) int {
	var viperConfig viper.Viper
	viperConfig = viper.Viper(*config)
	return viperConfig.GetInt(fmt.Sprintf("%s.%s", App.ENV, key))
}

// Boolean: read boolean value from viper.Viper
func (config *Config) Boolean(key string) bool {
	var viperConfig viper.Viper
	viperConfig = viper.Viper(*config)
	return viperConfig.GetBool(fmt.Sprintf("%s.%s", App.ENV, key))
}
