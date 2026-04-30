package main

import (
	"api/config"
	"api/internals/products"
	"api/internals/user/handler"
	"api/route"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	Port := loadenv()
	DB,err := config.Database()
	if err != nil {
		log.Fatal(err)
	}
	defer DB.Close()
	products := products.NewModule(DB)
	user := handler.NewModule(DB)
	router := route.RouterAPI(user.Controller, products.Controller)

	router.Run(":" + Port)
	
}

func loadenv() string{
	err := godotenv.Load(".env")
	if err != nil{
		log.Println("Error Port")
	}
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	return PORT
}