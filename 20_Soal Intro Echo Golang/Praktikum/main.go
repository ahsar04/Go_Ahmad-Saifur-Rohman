package main

import (
	"github.com/labstack/echo/v4"
)

// run menggunakan perintah "go run ."
func main() {
	e := echo.New()
	// routing with query parameter
	e.GET("/users", GetUsersController)
	e.GET("/users/:id", GetUsersByIdController)
	e.POST("/users", CreateUserController)
	e.DELETE("/users/:id", DeleteUserController)
	e.PUT("/users/:id", UpdateUserController)

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}