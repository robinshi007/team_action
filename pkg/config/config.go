package config

import (
	"fmt"
	"os"

	"github.com/jinzhu/configor"
)

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
	env := os.Getenv("APP_ENV")
	var conf *configor.Configor
	switch env {
	case "prod":
		conf = &configor.Configor{
			Config: &configor.Config{
				Environment: "prod",
			},
		}
	case "test":
		conf = &configor.Configor{
			Config: &configor.Config{
				Environment: "test",
			},
		}
	default:
		conf = &configor.Configor{
			Config: &configor.Config{
				Environment: "dev",
			},
		}
	}
	fmt.Println("APP_ENV:", env)
	c := &Config{}

	err := configor.Load(
		c,
		fmt.Sprintf("config/config.%s.yml", conf.GetEnvironment()),
		"config/config.yml",
	)
	if err != nil {
		return nil, err
	}
	//fmt.Printf("config %v", c)
	return c, nil
}
