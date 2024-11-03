package configs

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DB_HOST    		string `mapstructure:"DB_HOST"`
	DB_NAME  		string `mapstructure:"DB_NAME"`
	DB_USER  		string `mapstructure:"DB_USER"`
	DB_PASSWORD 	string `mapstructure:"DB_PASSWORD"`
	DB_PORT 		string `mapstructure:"DB_PORT"`
	JWT_SECRET	string `mapstructure:"JWT_SECRET"`
	// ADMIN_EMAIL 	string `mapstructure:"ADMIN_EMAIL"`
	// ADMIN_PASSWORD 	string `mapstructure:"ADMIN_PASSWORD"`
}

func NewConfig() *Config {
	config := &Config{}
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln("❌ Error reading config file", err)
	}

	if err := viper.Unmarshal(config); err != nil {
		log.Fatalln("❌ Unable to decode into struct", err)
	}

	return config
}