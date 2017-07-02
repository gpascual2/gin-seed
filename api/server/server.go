package server

import (
	"runtime"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Init the server config for Gin
func Init(config *viper.Viper, logger *logrus.Logger) {
	// Use all cpu cores
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Create router and listen on the configed port
	r := NewRouter(config, logger)
	r.Run(":" + config.GetString("server.port"))
}
