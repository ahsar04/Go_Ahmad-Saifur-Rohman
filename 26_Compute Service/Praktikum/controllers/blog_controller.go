package controllers

import (
	"code_structure/config"
	"code_structure/models"
	"net/http"

	"github.com/labstack/echo"
)

// get all blogs
func GetBlogsController(c echo.Context) error {
	var blogs []models.Blog


	if err := config.DB.Preload("User").Find(&blogs).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": http.StatusOK,
		"message": "success get all blogs",
		"Data":   blogs,
	})
}
// get blog by id
func GetBlogController(c echo.Context) error {
	blog := models.Blog{}
	blogID := c.Param("id")

	if err := config.DB.Preload("User").First(&blog, blogID).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": http.StatusOK,
		"message": "success get blog by id",
		"data":    blog,
	})
}
// create new blog
func CreateBlogController(c echo.Context) error {
	blog := models.Blog{}
	c.Bind(&blog)


	if err := config.DB.Preload("User").Save(&blog).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": http.StatusOK,
		"message": "success create new blog",
		"data":    blog,
	})
}
// delete blog by id
func DeleteBlogController(c echo.Context) error {
	blog := models.Blog{}
	blogID := c.Param("id")

	if err := config.DB.Preload("User").First(&blog, blogID).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := config.DB.Delete(&blog).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": http.StatusOK,
		"message": "success delete blog by id",
	})
}
// update blog by id
func UpdateBlogController(c echo.Context) error {
	blog := models.Blog{}
	blogID := c.Param("id")

	if err := config.DB.First(&blog, blogID).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Bind(&blog); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := config.DB.Save(&blog).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": http.StatusOK,
		"message": "success update blog by id",
		"data":    blog,
	})
}