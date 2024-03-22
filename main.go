package main

import (
	"fmt"
	"go-gorm/database"
	"go-gorm/router"
)

func main() {
	database.StartDB()
	const PORT = ":8080"

	fmt.Sprintln("server start at", PORT)
	router.StartApp().Run(PORT)

	// r := router.StartApp()
	// r.Run("8080")

}
