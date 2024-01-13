package auth

import (
	"auth/server/connection"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Defining user access
/*var (
	authorizedUsername = "admin"
	authorizedpassword = "admin123"
)*/

// Configuration struct represents the structure of your authorization configuration
type Configuration struct {
	AuthorizedUsername string `json:"authorizedUsername"`
	AuthorizedPassword string `json:"authorizedpassword"`
}

var config Configuration

// SetAuthorizationConfig sets the authorization configuration
func SetAuthorizationConfig(username, password string) {
	config.AuthorizedUsername = username
	config.AuthorizedPassword = password
}

// Function checks the autorization calls GET_DB function and grant/deny the access
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		username, password, ok := c.Request.BasicAuth()
		if ok && username == config.AuthorizedUsername && password == config.AuthorizedPassword {
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

// v.2.1 Calls the Query function from connection package and perform the query
func GET_Connection(c *gin.Context, config connection.DBConfig) {
	connection.Connection(config)
	input, err := connection.Query(config.Host, config.User, config.Password, config.Dbname, config.Port, 1)
	if err != nil {
		log.Println("Error in input", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
	}
	c.IndentedJSON(http.StatusOK, input)
}
