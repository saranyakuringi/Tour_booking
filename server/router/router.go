package router

import (
	"auth/server/auth"

	"github.com/gin-gonic/gin"
)

func Router() {
	router := gin.Default()
	router.GET("/secure", auth.AuthMiddleware(), auth.GET_DB)
	router.Run("localhost:8092")
}
