package config

import "github.com/spf13/viper"

// Config is the configuration for the service
type Config struct {
	// Port is the port the service will listen on
	Port int
	// Database is the database configuration
	Database DatabaseConfig
}

// DatabaseConfig is the configuration for the database
type DatabaseConfig struct {
	// Host is the database host
	Host string
	// Port is the database port
	Port int
	// User is the database user
	User string
	// Password is the database password
	Password string
	// Name is the database name
	Name string
}

// LoadConfig loads the configuration from the given file
func LoadConfig(file string) (*Config, error) {
	viper.SetConfigFile(file)
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

// LoadDefaultConfig loads the default configuration
func LoadDefaultConfig() *Config {
	return &Config{
		Port: 8080,
		Database: DatabaseConfig{
			Host:     "db",
			Port:     5432,
			User:     "postgres",
			Password: "postgres",
			Name:     "db",
		},
	}
}
