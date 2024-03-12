package conf

import (
	"go-cc/data"
	"os"

	"github.com/spf13/viper"
)

type Application struct {
	Port uint16 `mapstructure:"port"`
}

type Config struct {
	Application *Application   `mapstructure:"application"`
	Database    *data.Database `mapstructure:"database"`
}

const databaseConfigFile = ".env"

func ReadConfig() (*Config, error) {
	path, err := os.Executable()
	if err != nil {
		return nil, err
	}

	viper.AddConfigPath(path)
	viper.SetConfigFile(databaseConfigFile)
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	config := Config{}
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return &config, err
}
