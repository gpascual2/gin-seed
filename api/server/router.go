package server

import (
	"net/http"
	"time"

	"bitbucket.org/gpascual2/gin-seed/api/middleware"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// NewRouter creates the Gin router with all routes
func NewRouter(mode string, logger *logrus.Logger) *gin.Engine {
	gin.SetMode(mode)

	router := gin.New()
	router.Use(middleware.Ginrus(logger, time.RFC3339, false))
	// router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	return router
}
