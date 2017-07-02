package config

import (
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var config *viper.Viper

// Init is an exported method that takes the environment, starts the viper
// (external lib) and returns the configuration struct.
func Init(env string, log *logrus.Logger) {
	var err error
	v := viper.New()
	v.SetConfigType("yaml")
	v.SetConfigName(env)
	v.AddConfigPath("config/")
	err = v.ReadInConfig()
	if err != nil {
		log.WithFields(logrus.Fields{
			"time": time.Now().Format(time.RFC3339),
			"env":  env,
			"err":  err,
		}).Fatal("Error on parsing configuration file")
	}
	config = v
}

// GetConfig returns a config instance after loading from the proper configuration file
func GetConfig() *viper.Viper {
	return config
}
