package config

import (
	"log"
	"os"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type Config struct {
	AppConfig
	LoggerConfig
	ServerConfig
	GrpcServerConfig
	StorageConfig
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

type GrpcServerConfig struct {
	Address string
}

type StorageConfig struct {
	Dsn string
}

func NewConfig() *Config {
	parseFlags()

	configFolder := viper.GetString("configFolder")
	viper.SetConfigType("yml")
	viper.AddConfigPath(configFolder)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln(err)
	}

	var config Config
	err := viper.Unmarshal(&config)
	if err != nil {
		log.Fatalln(err)
	}

	config.StorageConfig.Dsn = os.Getenv("DSN")

	return &config
}

func parseFlags() {
	pflag.String("configFolder", "./configs", "path to configs folder")
	pflag.Parse()
	err := viper.BindPFlags(pflag.CommandLine)
	if err != nil {
		log.Fatalln(err)
	}
}
