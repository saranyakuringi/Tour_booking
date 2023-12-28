package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	authorizedUsername = "admin"
	authorizedpassword = "admin123"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		username, password, ok := c.Request.BasicAuth()
		if ok && username == authorizedUsername && password == authorizedpassword {
			c.Next()
		} else {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Unauthorised"})
			c.Abort()
		}
	}

}

func GET_DB(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Access Granted!"})
	//c.JSON(200, gin.H{"message": connection.Query(1)})

}
