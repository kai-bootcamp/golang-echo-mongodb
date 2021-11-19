package controller

import (
	"context"
	"github.com/labstack/echo/v4"
	"go-echo-mongodb/db"
	"go-echo-mongodb/types/dto"
	"go-echo-mongodb/types/model"
	"net/http"
)

func CreateUser(c echo.Context) error {
	var userDto *dto.User
	var userModel model.User
	if err := c.Bind(&userDto); err != nil {
		return err
	}
	userModel= model.CreateUserRequest(userDto)
	database, err := db.NewMongoDB()
	if err != nil {
		return err
	}
	ctx := context.Background()
	result, err := database.CreateUser(ctx, &userModel)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, result)
}
