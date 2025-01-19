package config

import (
	"flag"
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	AppDebug        bool
	DB              DatabaseConfig
	RedisConfig     RedisConfig
	GrpcIp          string
	GrpcPort        string
	ShutdownTimeout int
	RSAKeys         RSAKeys
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string

	Migrate bool
	Seeding bool
}

type RedisConfig struct {
	Url      string
	Password string
	Prefix   string
}

type RSAKeys struct {
	PrivateKey string
	PublicKey  string
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

	readFlags()

	// add value to the config
	config := Config{
		DB:              loadDatabaseConfig(),
		AppDebug:        viper.GetBool("APP_DEBUG"),
		GrpcIp:          viper.GetString("GRPC_IP"),
		GrpcPort:        viper.GetString("GRPC_PORT"),
		ShutdownTimeout: viper.GetInt("SHUTDOWN_TIMEOUT"),
		RedisConfig:     loadRedisConfig(),
		RSAKeys:         loadRSAKeys(),
	}
	return config, nil
}

func loadDatabaseConfig() DatabaseConfig {
	return DatabaseConfig{
		Host:     viper.GetString("DB_HOST"),
		Port:     viper.GetString("DB_PORT"),
		User:     viper.GetString("DB_USER"),
		Password: viper.GetString("DB_PASSWORD"),
		Name:     viper.GetString("DB_NAME"),
		Migrate:  viper.GetBool("DB_MIGRATE"),
		Seeding:  viper.GetBool("DB_SEEDING"),
	}
}

func loadRedisConfig() RedisConfig {
	return RedisConfig{
		Url:      viper.GetString("REDIS_URL"),
		Password: viper.GetString("REDIS_PASSWORD"),
		Prefix:   viper.GetString("REDIS_PREFIX"),
	}
}

func loadRSAKeys() RSAKeys {
	return RSAKeys{
		PrivateKey: viper.GetString("PRIVATE_KEY"),
		PublicKey:  viper.GetString("PUBLIC_KEY"),
	}
}

func setDefaultValues() {
	viper.SetDefault("DB_HOST", "localhost")
	viper.SetDefault("DB_PORT", "5432")
	viper.SetDefault("DB_USER", "postgres")
	viper.SetDefault("DB_PASSWORD", "admin")
	viper.SetDefault("DB_NAME", "postgres")
	viper.SetDefault("APP_DEBUG", true)
	viper.SetDefault("GRPC_IP", "0.0.0.0")
	viper.SetDefault("GRPC_PORT", ":50151")
	viper.SetDefault("SHUTDOWN_TIMEOUT", 5)

	viper.SetDefault("DB_MIGRATE", false)
	viper.SetDefault("DB_SEEDING", false)
}

func readFlags() {
	migrateDb := flag.Bool("m", false, "use this flag to migrate database")
	seedDb := flag.Bool("s", false, "use this flag to seed database")
	flag.Parse()
	if *migrateDb {
		viper.Set("DB_MIGRATE", true)
	}

	if *seedDb {
		viper.Set("DB_SEEDING", true)
	}
}
