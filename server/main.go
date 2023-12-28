package main

import (
	"auth/server/connection"
	"auth/server/router"
)

func main() {

	router.Router()

	//Open the SQL database
	connection.Query(1)
	//connection.Query(2)
	//connection.Query(3)
	//query

}
