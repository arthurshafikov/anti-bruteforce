package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	AppConfig
	LoggerConfig
	ServerConfig
	GrpcServerConfig ServerConfig
	DatabaseConfig
	LeakyBucketConfig
}

type AppConfig struct {
	NumberOfAttemptsForLogin    int64
	NumberOfAttemptsForPassword int64
	NumberOfAttemptsForIP       int64
}

type LoggerConfig struct {
	Level string
	File  string
}

type ServerConfig struct {
	Address string
}

type DatabaseConfig struct {
	DSN string
}

type LeakyBucketConfig struct {
	ResetIntervalInSeconds int
}

func NewConfig(configFolder string) *Config {
	viper.SetConfigType("yml")
	viper.AddConfigPath(configFolder)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln(err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalln(err)
	}

	return &config
}
