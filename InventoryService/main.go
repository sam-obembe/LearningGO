package main

import (
	"fmt"
	"inventoryService/database"
	"inventoryService/handlers"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

const apiBasePath = "/api"

func main() {

	database.SetupDatabase()
	fmt.Printf("Listening on port %v", 5000)

	handlers.SetupRoutes(apiBasePath)

	http.ListenAndServe(":5000", nil)

}
