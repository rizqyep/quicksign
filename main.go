package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rizqyep/quicksign/database"
	"github.com/rizqyep/quicksign/routers"
)

func main() {
	err := godotenv.Load("config/.env")
	if err != nil {
		fmt.Println("failed load file environment")
	} else {
		fmt.Println("successfully read file environment")
	}

	DB := database.GetDBConnection()
	database.DBMigrate()
	defer DB.Close()

	r := gin.Default()
	routers.RouteHandlers(r)

	if err := r.Run(":" + os.Getenv("PORT")); err != nil {
		log.Fatalf("Error in Starting the HTTP Server, Err: %s", err.Error())
	}

}
