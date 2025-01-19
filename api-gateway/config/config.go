package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	AppDebug           bool
	Email              EmailConfig
	RedisConfig        RedisConfig
	MicroserviceConfig MicroserviceConfig
	ServerIp           string
	ServerPort         string
	ShutdownTimeout    int
	AuthServiceIp      string
	UserServiceIp      string
	ChatServiceIp      string
	AuthServicePort    string
	UserServicePort    string
	ChatServicePort    string
}

type RedisConfig struct {
	Url      string
	Password string
	Prefix   string
}

type EmailConfig struct {
	ApiKey    string
	FromName  string
	FromEmail string
}

type MicroserviceConfig struct {
	Auth string
	User string
	Chat string
}

func LoadConfig() (Config, error) {
	// Set default values
	setDefaultValues()

	viper.AddConfigPath(".")
	viper.AddConfigPath("..")
	viper.AddConfigPath("../..")
	viper.SetConfigType("dotenv")
	viper.SetConfigName(".env")

	// Allow Viper to read environment variables
	viper.AutomaticEnv()

	// Read the configuration file
	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("Error reading config file: %s, using default values or environment variables", err)
	}

	// add value to the config
	config := Config{
		AppDebug:        viper.GetBool("APP_DEBUG"),
		ServerIp:        viper.GetString("SERVER_IP"),
		ServerPort:      viper.GetString("SERVER_PORT"),
		ShutdownTimeout: viper.GetInt("SHUTDOWN_TIMEOUT"),

		Email:              loadEmailConfig(),
		RedisConfig:        loadRedisConfig(),
		MicroserviceConfig: loadMicroserviceConfig(),
		AuthServiceIp:      viper.GetString("AUTH_SERVICE_IP"),
		UserServiceIp:      viper.GetString("USER_SERVICE_IP"),
		ChatServiceIp:      viper.GetString("CHAT_SERVICE_IP"),
		AuthServicePort:    viper.GetString("AUTH_SERVICE_PORT"),
		UserServicePort:    viper.GetString("USER_SERVICE_PORT"),
		ChatServicePort:    viper.GetString("USER_SERVICE_PORT"),
	}
	return config, nil
}

func loadRedisConfig() RedisConfig {
	return RedisConfig{
		Url:      viper.GetString("REDIS_URL"),
		Password: viper.GetString("REDIS_PASSWORD"),
		Prefix:   viper.GetString("REDIS_PREFIX"),
	}
}

func loadEmailConfig() EmailConfig {
	return EmailConfig{
		ApiKey:    viper.GetString("MAILERSEND_API_KEY"),
		FromName:  viper.GetString("MAILERSEND_FROM_NAME"),
		FromEmail: viper.GetString("MAILERSEND_FROM_EMAIL"),
	}
}

func loadMicroserviceConfig() MicroserviceConfig {
	return MicroserviceConfig{
		Auth: viper.GetString("AUTH_SERVICE_IP") + ":" + viper.GetString("AUTH_SERVICE_PORT"),
		User: viper.GetString("USER_SERVICE_IP") + ":" + viper.GetString("USER_SERVICE_PORT"),
		Chat: viper.GetString("CHAT_SERVICE_IP") + ":" + viper.GetString("CHAT_SERVICE_PORT"),
	}
}

func setDefaultValues() {
	viper.SetDefault("APP_DEBUG", true)
	viper.SetDefault("SERVER_PORT", "8181")
	viper.SetDefault("SHUTDOWN_TIMEOUT", 5)
}
