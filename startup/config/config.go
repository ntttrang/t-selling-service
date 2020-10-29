package config

import "os"

// Init init config
func Init() {
	loadConfig()
}

var config GeneralConfig

type (
	// MariaDBConfig mariadb configuration
	MariaDBConfig struct {
		Username string
		Password string
		Port     string
		Host     string
		Name     string
	}

	// ServerConfig server config
	ServerConfig struct {
		Host string
		Port string
	}

	// GeneralConfig all configuration in project
	GeneralConfig struct {
		ServerConfig ServerConfig
		MariaDb      MariaDBConfig
	}
)

func loadConfig() *GeneralConfig {
	config = GeneralConfig{
		ServerConfig: ServerConfig{
			Host: "0.0.0.0",
			Port: "8089",
		},
		MariaDb: MariaDBConfig{
			Name:     GetEnv("DB_NAME", "t_selling"),
			Username: GetEnv("DB_USER_NAME", "root"),
			Password: GetEnv("DB_PASSWORD", "P@ssw0rd"),
			Port:     GetEnv("DB_PORT", "3306"),
			Host:     GetEnv("DB_HOST", "localhost"),
		},
	}
	return &config
}

// GetEnv Get environment for project
func GetEnv(key string, defaultvalue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultvalue
}

// GetMariaDbConf Get mariadb config
func GetMariaDbConf() MariaDBConfig {
	return config.MariaDb
}

// GetServerConf Get server config
func GetServerConf() ServerConfig {
	return config.ServerConfig
}
