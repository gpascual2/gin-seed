package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"

	"bitbucket.org/gpascual2/gin-seed/api/config"
	"bitbucket.org/gpascual2/gin-seed/api/server"
)

var logger = logrus.New()

func init() {
	// Config logger
	logger.Formatter = new(logrus.TextFormatter)

}

func main() {
	var err error
	var env string
	flag.StringVar(&env, "e", "dev", "Working environment. Default: dev")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	fmt.Println("Environment: ", env)
	config.Init(env, logger)

	// Read config
	cfg := config.GetConfig()

	// Set configed log level
	cLogLvl, err := logrus.ParseLevel(cfg.GetString("env.log_level"))
	if err != nil {
		cLogLvl = logrus.WarnLevel // Default to Warning
	}
	logger.Level = cLogLvl

	// Log initial config
	logger.WithFields(logrus.Fields{"env": env}).Info("Environment set")
	logger.WithFields(logrus.Fields{"log_level": cfg.GetString("env.log_level")}).Info("Log level set")

	// HTTP Port to use
	logger.WithFields(logrus.Fields{
		"port": cfg.Get("server.port"),
	}).Info("HTTP Port configed")

	// Initialize and run server
	server.Init(cfg, logger)

}
