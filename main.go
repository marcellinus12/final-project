package main

import (
	"fmt"
	"go-gorm/database"
	"go-gorm/router"
	"os"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	PORT := os.Getenv("PORT")

	fmt.Println("server start at :" + PORT)
	if err := r.Run(":" + PORT); err != nil {
		panic("panic failed to start")
	}

	// r := router.StartApp()
	// r.Run("8080")

}
