package controllers

import (
	"code_structure/config"
	"code_structure/models"
	"net/http"

	"github.com/labstack/echo"
)

// get all users
func GetBooksController(c echo.Context) error {
	var books []models.Book


	if err := config.DB.Find(&books).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": http.StatusOK,
		"message": "success get all books",
		"Data":   books,
	})
}
// get book by id
func GetBookController(c echo.Context) error {
	book := models.Book{}
	bookID := c.Param("id")

	if err := config.DB.First(&book, bookID).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": http.StatusOK,
		"message": "success get book by id",
		"data":    book,
	})
}
// create new book
func CreateBookController(c echo.Context) error {
	book := models.Book{}
	c.Bind(&book)


	if err := config.DB.Save(&book).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": http.StatusOK,
		"message": "success create new book",
		"data":    book,
	})
}
// delete book by id
func DeleteBookController(c echo.Context) error {
	book := models.Book{}
	bookID := c.Param("id")

	if err := config.DB.First(&book, bookID).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := config.DB.Delete(&book).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": http.StatusOK,
		"message": "success delete book by id",
	})
}
// update book by id
func UpdateBookController(c echo.Context) error {
	book := models.Book{}
	bookID := c.Param("id")

	if err := config.DB.First(&book, bookID).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Bind(&book); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := config.DB.Save(&book).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": http.StatusOK,
		"message": "success update book by id",
		"data":    book,
	})
}