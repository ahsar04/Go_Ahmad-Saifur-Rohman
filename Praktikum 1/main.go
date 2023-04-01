package main

import (
	"praktikum_orm/config"
	"praktikum_orm/routes"
)

// -------------------- controller --------------------

// get user by id
// func GetUsersByIdController(c echo.Context) error {
// 	userId, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, map[string]interface{}{
// 			"status": http.StatusBadRequest,
// 			"messages": "invalid user id",
// 		})
// 	}

// 	for _, user := range users {
// 		if user.Id == userId {
// 		return c.JSON(http.StatusOK, map[string]interface{}{
// 			"status": http.StatusOK,
// 			"messages": "success get user",
// 			"data":     user,
// 		})
// 		}
// 	}

// 	return c.JSON(http.StatusNotFound, map[string]interface{}{
// 		"status": http.StatusNotFound,
// 		"messages": "user not found",
// 	})
// }
// // delete user by id
// func DeleteUserController(c echo.Context) error {
// 	userId, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, map[string]interface{}{
// 		"status": http.StatusBadRequest,
// 		"messages": "invalid user id",
// 		})
// 	}

// 	for i, user := range users {
// 		if user.Id == userId {
// 		users = append(users[:i], users[i+1:]...)
// 		return c.JSON(http.StatusOK, map[string]interface{}{
// 			"status": http.StatusOK,
// 			"messages": "success delete user",
// 		})
// 		}
// 	}

// 	return c.JSON(http.StatusNotFound, map[string]interface{}{
// 		"status": http.StatusNotFound,
// 		"messages": "user not found",
// 	})
// }
// // update user by id
// func UpdateUserController(c echo.Context) error {
// 	userId, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, map[string]interface{}{
// 			"status": http.StatusBadRequest,
// 			"messages": "invalid user id",
// 		})
// 	}

// 	for i, user := range users {
// 		if user.Id == userId {
// 		// binding data
// 		updatedUser := User{}
// 		c.Bind(&updatedUser)

// 		users[i].Name = updatedUser.Name
// 		users[i].Email = updatedUser.Email
// 		users[i].Password = updatedUser.Password

// 		return c.JSON(http.StatusOK, map[string]interface{}{
// 			"status": http.StatusOK,
// 			"messages": "success update user",
// 			"data":     users[i],
// 		})
// 		}
// 	}

// 	return c.JSON(http.StatusNotFound, map[string]interface{}{
// 		"status": http.StatusNotFound,
// 		"messages": "user not found",
// 	})
// }



func main() {
	config.InitDB()
	e := routes.New()
	// routing with query parameter
	// e.DELETE("/users/:id", DeleteUserController)
	// e.PUT("/users/:id", UpdateUserController)

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}

