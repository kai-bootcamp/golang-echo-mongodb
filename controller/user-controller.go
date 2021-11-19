package controller

import (
	"context"
	"fmt"
	"go-echo-mongodb/db"
	"go-echo-mongodb/types/dto"
	"go-echo-mongodb/types/model"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func CreateUser(c echo.Context) error {
	var userDto dto.User
	var userModel model.User
	if err := c.Bind(&userDto); err != nil {
		return err
	}
	userModel = model.CreateUserRequest(&userDto)

	ctx := context.Background()
	result, err := initDBConnection().CreateUser(ctx, &userModel)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, result)
}

func GetUser(c echo.Context) error {
	name := c.Param("username")
	result, err := initDBConnection().ReadUser(name)
	if err != nil {
		log.Panic(err)
	}
	return c.JSON(http.StatusOK, result)
}

func RemoveUser(c echo.Context) error {
	var userDto dto.User
	if err := c.Bind(&userDto); err != nil {
		return err
	}
	username := userDto.Username
	fmt.Printf("user name is %s \n", username)
	result, err := initDBConnection().RemoveUser(username)
	if err != nil {
		log.Panic(err)
	}
	return c.JSON(http.StatusOK, result)
}

func UpdateUser(c echo.Context) error {
	// var userDto dto.User
	var userModel model.User
	if err := c.Bind(&userModel); err != nil {
		return err
	}
	// userModel = model.CreateUserRequest(&userDto)
	result, err := initDBConnection().UpdateUser(&userModel)
	if err != nil {
		log.Panic(err)
	}
	return c.JSON(http.StatusOK, result)
}

func initDBConnection() *db.MongoDB {
	database, err := db.NewMongoDB(os.Getenv("MONGODB_URI"))
	if err != nil {
		log.Panic("Unable to connect DB")
	}
	return database
}
