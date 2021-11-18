package app

import (
	"github.com/gin-gonic/gin"
	"github.com/nubesFilius/bselling-go-users-api/logger"
)

var (
	router = gin.Default()
)

func StartApplication() {
	mapUrls()

	logger.Info("Starting the app...")
	router.Run(":8080")
}
