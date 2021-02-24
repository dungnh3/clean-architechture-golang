package config

import (
	"github.com/dungnh3/clean-architechture/pkg/database"
	"github.com/dungnh3/clean-architechture/pkg/server"
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Server   *server.Config   `json:"server" mapstructure:"server"`
	Database *database.Config `json:"database" mapstructure:"database"`
}

func Load() (*Config, error) {
	cfg := &Config{}

	viper.SetConfigType("yaml")
	viper.AddConfigPath("./app/config")
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		log.Println("Read config file failed", err.Error())
		return nil, err
	}

	viper.AutomaticEnv()
	return cfg, viper.Unmarshal(cfg)
}
