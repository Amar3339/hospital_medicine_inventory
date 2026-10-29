package main

import (
	"fmt"
	"medical_inventory/db"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("error loadin env")

	}
	// connect with pg
	err = db.Connect()
	if err != nil {
		panic("connection failed with database")

	}
	fmt.Println("sucessfully connected with databse ")
	router := gin.Default()
	router.Use(cors.New(cors.Config{

		AllowOrigins: []string{
			"http://127.0.0.1:5500",
			"http://localhost:8080",
		},
		AllowMethods: []string{
			"POST",
			"GET",
			"PUT",
			"DELETE",
			"OPTION"},
		AllowHeaders: []string{
			"Origin", "Content-Type", "Accept",
		},
	},
	))

	SetupRouters(router)

	router.Run(":8080")

}
