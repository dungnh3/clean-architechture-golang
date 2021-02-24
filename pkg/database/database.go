package database

import "fmt"

type Config struct {
	Host     string `json:"host" mapstructure:"host"`
	Port     int    `json:"port" mapstructure:"port"`
	Username string `json:"username" mapstructure:"username"`
	Password string `json:"password" mapstructure:"password"`
	Database string `json:"database" mapstructure:"database"`
	Options  string `json:"options" mapstructure:"options"`
}

func (c *Config) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", c.Username, c.Password, c.Host, c.Port, c.Database)
}
