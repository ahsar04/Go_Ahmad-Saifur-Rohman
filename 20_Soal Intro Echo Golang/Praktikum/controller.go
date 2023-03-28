package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type User struct {
	Id    int    `json:"id" form:"id"`
	Name  string `json:"name" form:"name"`
	Email string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var users []User

// -------------------- controller --------------------

// get all users
func GetUsersController(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": http.StatusOK,
		"messages": "success get all users",
		"data":    users,
	})
}

// get user by id
func GetUsersByIdController(c echo.Context) error {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": http.StatusBadRequest,
			"messages": "invalid user id",
		})
	}

	for _, user := range users {
		if user.Id == userId {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status": http.StatusOK,
			"messages": "success get user",
			"data":     user,
		})
		}
	}

	return c.JSON(http.StatusNotFound, map[string]interface{}{
		"status": http.StatusNotFound,
		"messages": "user not found",
	})
}
// delete user by id
func DeleteUserController(c echo.Context) error {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
		"status": http.StatusBadRequest,
		"messages": "invalid user id",
		})
	}

	for i, user := range users {
		if user.Id == userId {
		users = append(users[:i], users[i+1:]...)
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status": http.StatusOK,
			"messages": "success delete user",
		})
		}
	}

	return c.JSON(http.StatusNotFound, map[string]interface{}{
		"status": http.StatusNotFound,
		"messages": "user not found",
	})
}
// update user by id
func UpdateUserController(c echo.Context) error {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": http.StatusBadRequest,
			"messages": "invalid user id",
		})
	}

	for i, user := range users {
		if user.Id == userId {
		// binding data
		updatedUser := User{}
		c.Bind(&updatedUser)

		users[i].Name = updatedUser.Name
		users[i].Email = updatedUser.Email
		users[i].Password = updatedUser.Password

		return c.JSON(http.StatusOK, map[string]interface{}{
			"status": http.StatusOK,
			"messages": "success update user",
			"data":     users[i],
		})
		}
	}

	return c.JSON(http.StatusNotFound, map[string]interface{}{
		"status": http.StatusNotFound,
		"messages": "user not found",
	})
}

// create new user
func CreateUserController(c echo.Context) error {
  // binding data
	user := User{}
	c.Bind(&user)

	if len(users) == 0 {
		user.Id = 1
	} else {
		newId := users[len(users)-1].Id + 1
		user.Id = newId
	}
	users = append(users, user)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": http.StatusOK,
		"messages": "success create user",
		"data": user,
	})
}