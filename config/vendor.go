package config

import "time"

type Vendor struct {
	Database Database `mapstructure:"database"`
}

type Database struct {
	Driver          string        `mapstructure:"driver"`
	DataSource      string        `mapstructure:"data_source"`
	MaxOpenConns    int           `mapstructure:"max_open_conns"`
	MaxIdleConns    int           `mapstructure:"max_idle_conns"`
	MaxConnLifeTime time.Duration `mapstructure:"max_conn_life_time"`
	MaxConnIdleTime time.Duration `mapstructure:"max_conn_idle_time"`
}
