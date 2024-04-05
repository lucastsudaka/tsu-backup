package pkg

import (
	"fmt"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
	"os"
	"path"
	"path/filepath"
	"runtime"
)

type IAppConfig interface {
	GetConfigValueByKey(key string) string
	GetDatabaseDSN() string
	GetSentryDSN() string
	GetAppRootDir() string
}
type AppConfig struct {
	client *viper.Viper
}

func (cfg *AppConfig) GetAppRootDir() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return filepath.Dir(d)
}

func NewAppConfig() (*AppConfig, error) {
	m := &AppConfig{}
	m.Init()
	return m, nil
}

func (cfg *AppConfig) Init() {
	cfg.client = viper.GetViper()
	cfg.client.AddConfigPath(".")
	cfg.client.SetConfigName("")
	cfg.client.SetConfigFile(".env")

	cfg.client.SetConfigType("env")
	if err := cfg.client.ReadInConfig(); err != nil {
		fmt.Println("Error reading env file", err)
	}
	fmt.Println("--- INIT APP CONFIG ---")

}

func (cfg *AppConfig) GetConfigValueByKey(key string) string {
	if key == "APP_BACKEND_MASTER_DB_HOST" {
		if len(cast.ToString(os.Getenv("IS_CONTAINER"))) > 0 {
			return "localhost"
		}
	}

	return cast.ToString(viper.GetViper().GetString(key))
}

func (cfg *AppConfig) GetDatabaseDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		cfg.GetConfigValueByKey("APP_BACKEND_MASTER_DB_USER"),
		cfg.GetConfigValueByKey("APP_BACKEND_MASTER_DB_PASSWORD"),
		cfg.GetConfigValueByKey("APP_BACKEND_MASTER_DB_HOST"),
		cfg.GetConfigValueByKey("APP_BACKEND_MASTER_DB_PORT"),
		cfg.GetConfigValueByKey("APP_BACKEND_MASTER_DB_NAME"),
	)
}

func (cfg *AppConfig) GetSentryDSN() string {
	return cfg.GetConfigValueByKey("APP_BACKEND_JWT_SECRET")
}
