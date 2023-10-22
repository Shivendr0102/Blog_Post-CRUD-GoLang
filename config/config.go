package config

import (
	"database/sql"
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type SqlServer struct {
	URL      string
	Port     int
	DBName   string
	Username string
	Password string
}

type Constants struct {
	IsProduction   bool
	Port           string
	AllowHttp      bool
	AllowedOrigins []string
	SqlServer      SqlServer
}

type Config struct {
	Constants
}

// NewConfig is used to generate a configuration instance which will be passed around the codebase
func New() (*Config, error) {
	config := Config{}
	constants, err := initViper()
	config.Constants = constants
	if err != nil {
		return &config, err
	}

	return &config, err
}

func (config *Config) Connect() (*sql.DB, error) {
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s",
		"xx.xx.xx.xx", "xx", "xxxxxxxxx", 0000, "xxxx")
	return sql.Open("mssql", connString)
}

func initViper() (Constants, error) {
	viper.SetConfigName(".local/api.config") // Configuration fileName without the .TOML or .YAML extension
	viper.AddConfigPath(".")                 // Search the root directory for the configuration file
	viper.WatchConfig()                      // Watch for changes to the configuration file and recompile

	viper.SetDefault("Port", "8080")
	viper.SetDefault("AllowHttp", "true")
	if err := viper.ReadInConfig(); err != nil {
		log.Panicf("Error reading config file, %s", err)
	}

	var constants Constants
	err := viper.Unmarshal(&constants)
	return constants, err
}
