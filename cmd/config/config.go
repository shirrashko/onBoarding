package config

import (
	"github.com/spf13/viper" // a popular configuration management library for Go.
)

// DBConfig holds the info needed for connecting to the database
// It's annotated with mapstructure tags to indicate how configuration keys should be mapped to struct fields.
type DBConfig struct {
	Host         string `mapstructure:"HOST"`
	Port         int    `mapstructure:"PORT"`
	User         string `mapstructure:"USER"`
	Password     string `mapstructure:"PASSWORD"`
	DatabaseName string `mapstructure:"DATABASE_NAME"`
}

// ServerConfig Holds information about the server (host and port).
type ServerConfig struct {
	Host string `mapstructure:"HOST"`
	Port int    `mapstructure:"PORT"`
}

// Config The main configuration struct that contains nested fields for database connection information and server information.
// It's also annotated with mapstructure tags.
type Config struct {
	DBConfig   DBConfig     `mapstructure:"DBConfig"`
	ServerInfo ServerConfig `mapstructure:"ServerConfig"`
}

var vp *viper.Viper

// LoadConfig loading the configuration from various sources and returning a populated Config struct:
func LoadConfig() (Config, error) {
	vp = viper.New()  // Initializes a new instance of the viper configuration manager.
	var config Config // Creates an instance of the Config struct to store the loaded configuration.

	vp.SetConfigName("config")     // Specifies the name of the configuration file without the file extension ("config.json").
	vp.SetConfigType("json")       // Sets the expected file type for the configuration file.
	vp.AddConfigPath("cmd/config") // Adds the directory path where the configuration file is located.
	vp.AutomaticEnv()              // Enables automatic binding of environment variables to configuration keys.

	if err := vp.ReadInConfig(); err != nil { //  If successful, it loads the configuration into vp.
		return Config{}, err
	}

	if err := vp.Unmarshal(&config); err != nil { // Unmarshal the configuration data from vp into the config struct.
		return Config{}, err
	}

	return config, nil
}
