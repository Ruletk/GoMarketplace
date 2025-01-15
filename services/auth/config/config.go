package config

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

// Config is the configuration for the service
type Config struct {
	// Port is the port the service will listen on
	Port int
	// Database is the database configuration
	Database DatabaseConfig
	// Jwt is the jwt configuration
	Jwt JwtConfig
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

// JwtConfig is the configuration for the jwt
type JwtConfig struct {
	// Algo is the jwt algorithm
	Algo jwt.SigningMethod
	// Secret is the jwt secret
	Secret string
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
		Jwt: JwtConfig{
			Algo: jwt.SigningMethodHS256,
			// This is a randomly generated secret, do not use it in production
			// Change it to a secure secret
			// Preferably use a secret from an environment variable
			Secret: "NwNK8PrfxCwcPqSnTrWrFDf4EIe89N4RklvpXyAIdDXeEdOj/Nwc8Y5Iuu5+TbVBka9JNmZ53FKlCce7j654T21nltLoehmfbN7PTQvVYcAZJqZbeHTQu6bxdN93eYgr57nXAxF8DDNkXFRtxg0IfpuVSfUkyyflda/FuLN0YIXZsTN0cccvfDhsEp1Q/osQvqzxxF7hUGu1Z/WstDMCXK3dlrfLAuY6WExuEj5yXPAujkvNUzmfB5Stp2okMc/c+yMHCpmxuW6Pv8MC6RPo1utjtF//zLhedZr9DVH2PUg1eT9cqKZLthbFjcJck2wN5xsPBOErbh16oML92EHJaB8VZjLO3HkQgVjSd63tIfSD1svZdlhUdd+1HrvbTcZOXbSfQOYL4xtz11R0CvGfaqsshOrCW8u+i8467JgzEc56mVbvg/fiYVFg2YcIzuntJf8wx5myNueOJgjcUv/mi0VzcnuhtHgv7wW/1QckgKxPE2DBjQ+Z14ro0/6fhrM3j9ZoJQ8dARKB//zK9rTOdNuXr45cXFRQu3N6nmiiMvgpRgKwAWO7l27byTFJeU9u3/3wK6LMuudjtrd5xj5YjdFyGmUsChnP6WspV+o/QJTQN6DlAKg2y375zdkyH4XgR8Y7IvNBKibdD1D8KCwudybWF8KsgGSzJ9PadkNZS8c=",
		},
	}
}
