package config

type Schema struct {
	App    App    `mapstructure:"app"`
	Vendor Vendor `mapstructure:"vendor"`
}
