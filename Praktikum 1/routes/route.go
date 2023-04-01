package routes

import (
	"praktikum_orm/controllers"

	"github.com/labstack/echo"
)

func New() *echo.Echo{
	e := echo.New()
	
	e.GET("/users", controllers.GetUsersController)
	// e.GET("/users/:id", GetUsersByIdController)
	e.POST("/users", controllers.CreateUserController)
	return e
}