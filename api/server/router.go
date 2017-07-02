package server

import (
	"net/http"
	"time"

	"bitbucket.org/gpascual2/gin-seed/api/middleware"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// NewRouter creates the Gin router with all routes
func NewRouter(config *viper.Viper, logger *logrus.Logger) *gin.Engine {
	gin.SetMode(config.GetString("server.mode"))

	// Validate CORS config
	var corsOrigins []string
	corsOrigins = config.GetStringSlice("cors.allow_origin")
	corsEnabled := config.GetBool("cors.enable")
	if len(corsOrigins) == 0 {
		corsEnabled = false
	}

	// Setup router
	router := gin.New()
	router.Use(middleware.Ginrus(logger, time.RFC3339, false))
	if corsEnabled {
		router.Use(middleware.CORSMiddleware(config))
	}
	router.Use(gin.Recovery())

	// Setup routes
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	router.LoadHTMLGlob("./templates/*")
	router.Static("/templates", "./templates")
	// root route
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
	// 404 handler
	router.NoRoute(func(c *gin.Context) {
		c.HTML(404, "404.html", gin.H{})
	})
	return router
}
