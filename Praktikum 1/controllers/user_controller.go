package controllers

import (
	"net/http"
	"praktikum_orm/config"
	"praktikum_orm/models"

	"github.com/labstack/echo"
)

// get all users
func GetUsersController(c echo.Context) error {
	var users []models.User

	err := config.DB.Find(&users).Error
	if err!=nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status": http.StatusInternalServerError,
			"messages": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": http.StatusOK,
		"messages": "success get all users",
		"data":    users,
	})
}

// create new user
func CreateUserController(c echo.Context) error {
  // binding data
	user := models.User{}
	c.Bind(&user)

	err := config.DB.Save(&user).Error
	if err!=nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status": http.StatusInternalServerError,
			"messages": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": http.StatusOK,
		"messages": "success create user",
		"data": user,
	})
}