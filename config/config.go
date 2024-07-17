package config

import (
	"errors"
	"log"
	"sync"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var (
	instance *Config
	once     sync.Once
)

func GetConfig() *Config {

	once.Do(func() {
		instance = &Config{}

		if err := initConfig("config"); err != nil {
			log.Fatalf("error initializing config: %v", err)
		}
	})

	return instance
}

// Environment represents the environment of the application
type Environment string

const (
	Conf Environment = "config"
)

func initConfig(env Environment) error {
	if env == "" {
		return errors.New("environment not set")
	}

	v := viper.New()
	v.AddConfigPath("./config/")
	v.SetConfigType("yaml")
	v.SetConfigName(string(env))

	if err := v.ReadInConfig(); err != nil {
		var configFileNotFoundError *viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			return errors.New("configuration file not found")
		}
		return errors.New("error on parsing configuration file")
	}

	// Unmarshal the parsed config into the Config struct
	if err := v.Unmarshal(instance); err != nil {
		log.Println(err)
		return errors.New("error on parsing configuration file to struct")
	}

	// Watch for changes to the config file and update the Config instance
	// accordingly
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("Config file changed: %s", e.Name)
		if err := v.Unmarshal(instance); err != nil {
			log.Printf("Failed to reload config: %v", err)
		}
	})

	return nil
}

// Config represents the application's configuration
type Config struct {
	App       AppConfig       `mapstructure:"app"`
	Database  DatabaseConfig  `mapstructure:"database"`
	WebSocket WebSocketConfig `mapstructure:"websocket"`
}

type AppConfig struct {
	Port int    `mapstructure:"port"`
	Env  string `mapstructure:"env"`
}

type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
	SSLMode  string `mapstructure:"sslmode"`
}

type WebSocketConfig struct {
	Endpoint        string `mapstructure:"endpoint"`
	ReadBufferSize  int    `mapstructure:"readBufferSize"`
	WriteBufferSize int    `mapstructure:"writeBufferSize"`
}
