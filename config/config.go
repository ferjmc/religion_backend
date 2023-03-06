package config

import (
	"errors"
	"path/filepath"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	HttpServer HttpServer
	Mysql      MysqlConfig
	Logger     Logger
	Firebase   Firebase
}

type HttpServer struct {
	Mode              string
	Port              string
	Timeout           time.Duration
	ReadTimeout       time.Duration
	WriteTimeout      time.Duration
	CookieLifeTime    int
	SessionCookieName string
}

type MysqlConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Dbname   string
}

type Logger struct {
	Development       bool
	DisableCaller     bool
	DisableStacktrace bool
	Encoding          string
	Level             string
}

type Firebase struct {
	Path string
}

// Load config file from given path
func LoadConfig(filename string) (*viper.Viper, error) {
	v := viper.New()

	v.SetConfigName(filename)
	v.AddConfigPath(".")
	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}

	return v, nil
}

// Parse config file
func ParseConfig(v *viper.Viper) (*Config, error) {
	var c Config

	err := v.Unmarshal(&c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}

// Get config
func GetConfig(configPath string) (*Config, error) {
	cfgFile, err := LoadConfig(configPath)
	if err != nil {
		return nil, err
	}

	cfg, err := ParseConfig(cfgFile)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

func GetConfigPath(configPath string) string {
	if configPath == "docker" {
		return filepath.Join("..", "config", "config-docker")
	}
	return filepath.Join("..", "config", "config-local")
}
