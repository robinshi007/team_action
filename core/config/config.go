package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

// Config - config
type Config struct {
	Mode     string         `envconfig:"APP_MODE"`
	Log      LogConfig      `yaml:"log"`
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
	Contact  ContactConfig  `yaml:"contact"`
}

// ServerConfig -
type ServerConfig struct {
	Name string `default:"team_action"`
	Host string `default:"localhost"`
	Port string `default:"3000"`
}

// LogConfig -
type LogConfig struct {
	Code     string `default:"zap"`
	Level    string `default:"info"`
	FileName string `default:"team_action.log" yaml:"file_name"`
}

// DatabaseConfig -
type DatabaseConfig struct {
	Code       string `default:"sqldb"`
	DriverName string `default:"postgres" yaml:"driver_name"`
	URLAddress string `yaml:"url_address"`
}

// ContactConfig -
type ContactConfig struct {
	Name  string `default:"Robin Shi"`
	Email string `default:"robinshi@outlook.com"`
}

// ReadEnv -
func ReadEnv(env string, valueDefault string) string {
	value := os.Getenv(env)
	if value != "" {
		return value
	}
	return valueDefault
}

// NewConfig  -
func NewConfig() (*Config, error) {
	env := ReadEnv("APP_MODE", "dev")
	return ReadConfig(fmt.Sprintf("config/config.%s.yml", env))
}

// ReadConfig -reads the file of the filename (in the same folder) and put it into the Config
func ReadConfig(filename string) (*Config, error) {
	var c Config
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, errors.Wrap(err, "read error")
	}
	err = yaml.Unmarshal(file, &c)

	if err != nil {
		return nil, errors.Wrap(err, "unmarshal")
	}
	err = validateConfig(c)
	if err != nil {
		return nil, errors.Wrap(err, "validate config")
	}
	// re write mode using env
	c.Mode = ReadEnv("APP_MODE", "dev")

	fmt.Println("Config:", c)
	return &c, nil
}

// ReadEnvConfig -
func ReadEnvConfig(cfg *Config) error {
	// return envconfig.Process("app", cfg)
	return nil
}
