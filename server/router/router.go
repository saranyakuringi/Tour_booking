package router

import (
	"auth/server/auth"
	"auth/server/connection"

	"github.com/gin-gonic/gin"
)

func Router(config connection.DBConfig) {
	router := gin.Default()
	router.GET("/secure", auth.AuthMiddleware(), auth.GET_DB, func(c *gin.Context) {
		auth.GET_Connection(c, config)
	})
	router.Run("localhost:8092")
}
