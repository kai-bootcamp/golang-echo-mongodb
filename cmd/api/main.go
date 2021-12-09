package main

import (
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"go-echo-mongodb/controller"
)

func main() {
	server := echo.New()

	//http.ListenAndServe(":5005", nil)
	server.POST("/create", controller.CreateUser)

	//Enable metrics middleware
	p := prometheus.NewPrometheus("echo", nil)
	p.Use(server)

	server.Logger.Fatal(server.Start(":8080"))
}