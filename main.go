package main

import (
	"fmt"
	"go-gorm/database"
	"go-gorm/router"
	"os"
)

func main() {
	database.StartDB()
	PORT := os.Getenv("PORT")

	fmt.Sprintln("server start at :", PORT)
	router.StartApp().Run(PORT)

	// r := router.StartApp()
	// r.Run("8080")

}
