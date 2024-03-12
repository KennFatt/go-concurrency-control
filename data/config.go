package data

import (
	"time"
)

type Database struct {
	// DSN
	Driver   string `mapstructure:"driver"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Hostname string `mapstructure:"hostname"`
	Port     int    `mapstructure:"port"`
	Name     string `mapstructure:"name"`

	// SQL Client config
	MaxLifeTime *time.Duration `mapstructure:"max_life_time,omitempty"`
	MaxIdleConn *int           `mapstructure:"max_idle_conn,omitempty"`
	MaxOpenConn *int           `mapstructure:"max_open_conn,omitempty"`
}
