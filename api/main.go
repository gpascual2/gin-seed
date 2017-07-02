package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"

	"bitbucket.org/gpascual2/gin-seed/api/config"
)

var log = logrus.New()

func init() {
	// Config logger
	log.Formatter = new(logrus.TextFormatter)

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
	config.Init(env, log)

	// Read config
	c := config.GetConfig()

	// Set configed log level
	cLogLvl, err := logrus.ParseLevel(c.GetString("env.log_level"))
	if err != nil {
		cLogLvl = logrus.WarnLevel // Default to Warning
	}
	log.Level = cLogLvl

	// Log initial config
	log.WithFields(logrus.Fields{"env": env}).Info("Environment set")
	log.WithFields(logrus.Fields{"log_level": c.GetString("env.log_level")}).Info("Log level set")

	// HTTP Port to use
	log.WithFields(logrus.Fields{
		"port": c.Get("server.port"),
	}).Info("HTTP Port configed")

}
