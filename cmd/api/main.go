package main

import (
	"go-echo-mongodb/controller"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	server := echo.New()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	server.POST(os.Getenv("CREATE_PATH"), controller.CreateUser)
	server.GET(os.Getenv("READ_PATH"), controller.GetUser)
	server.POST(os.Getenv("DELETE_PATH"), controller.RemoveUser)
	server.POST(os.Getenv("UPDATE_PATH"), controller.UpdateUser)

	server.Logger.Fatal(server.Start(":8080"))

}
