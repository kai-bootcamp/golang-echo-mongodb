package main

import (
	"github.com/labstack/echo/v4"
	"go-echo-mongodb/controller"
)

func main() {
	server := echo.New()

	server.POST("/create", controller.CreateUser)

	server.Logger.Fatal(server.Start(":8080"))
}