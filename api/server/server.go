package server

import (
	"runtime"

	"github.com/sirupsen/logrus"

	"bitbucket.org/gpascual2/gin-seed/api/config"
)

// Init the server config for Gin
func Init(logger *logrus.Logger) {
	config := config.GetConfig()

	// Use all cpu cores
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Create router and listen on the configed port
	r := NewRouter(config.GetString("server.mode"), logger)
	r.Run(":" + config.GetString("server.port"))
}
