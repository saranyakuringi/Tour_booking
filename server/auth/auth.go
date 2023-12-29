package auth

import (
	"auth/server/connection"
	"log"
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

func GET_Connection(c *gin.Context) {
	input, err := connection.Query(1)
	if err != nil {
		log.Println("Error in input", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
	}
	c.IndentedJSON(http.StatusOK, input)

}
