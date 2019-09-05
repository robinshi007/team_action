package config

import "github.com/jinzhu/configor"

// Config - config
type Config struct {
	AppName string `default:"team_action"`
	Port    string `default:"3000"`
	Logger  struct {
		Use         string `default:"zapLogger"`
		Environment string `default:"prod"`
		LogLevel    string `default:"info"`
		FileName    string `default:"team_action.log"`
	}
	DB struct {
		Use        string `default:"postgres"`
		Connection struct {
			Enabled  bool   `default:"true"`
			Host     string `default:"postgres"`
			Port     string `default:"5432"`
			UserName string `default:"postgres"`
			Password string `default:"postgres"`
			Database string `default:"denti"`
		}
	}
	Contacts struct {
		Name  string `default:"Robin Shi"`
		Email string `default:"robinshi@outlook.com"`
	}
}

// NewConfig - NewConfig
func NewConfig() (*Config, error) {
	c := &Config{}
	err := configor.Load(c, "config/config.yml")
	if err != nil {
		return nil, err
	}
	return c, nil
}
