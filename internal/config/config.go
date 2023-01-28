package config

import (
	"github.com/spf13/viper"
)

// Configurations wraps all the config variables required by the catchall service
type Configurations struct {
	DBHost   string
	DBName   string
	DBUser   string
	DBPass   string
	DBPort   string
	DBSchema string
}

// NewConfigurations returns a new Configuration object
func NewConfigurations() *Configurations {

	viper.AutomaticEnv()

	viper.SetDefault("DB_HOST", "postgres-db") // 'postgres' If running docker-compose '0.0.0.0' localhost
	viper.SetDefault("DB_NAME", "catchall-db")
	viper.SetDefault("DB_USER", "postgres")
	viper.SetDefault("DB_PASSWORD", "root")
	viper.SetDefault("DB_PORT", "5433")
	viper.SetDefault("DB_SCHEMA", "domain_names")

	configs := &Configurations{
		DBHost:   viper.GetString("DB_HOST"),
		DBName:   viper.GetString("DB_NAME"),
		DBUser:   viper.GetString("DB_USER"),
		DBPass:   viper.GetString("DB_PASSWORD"),
		DBPort:   viper.GetString("DB_PORT"),
		DBSchema: viper.GetString("DB_SCHEMA"),
	}

	return configs
}
