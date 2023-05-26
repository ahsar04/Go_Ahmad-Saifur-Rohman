package handler

import (
	"net/http"
	"strconv"

	"remidial/internal/application/usecase"
	"remidial/internal/entity"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type ItemHandler struct {
	ItemUsecase usecase.ItemUseCase
}

func (handler ItemHandler) GetAllItems() echo.HandlerFunc {
	return func(e echo.Context) error {
		var items []entity.Item

		items, err := handler.ItemUsecase.GetAllItems()
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status code": http.StatusInternalServerError,
				"message": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"status code": http.StatusOK,
			"message": "success get all items",
			"data":   items,
		})
	}
}

func (handler ItemHandler) GetItem() echo.HandlerFunc {
	return func(e echo.Context) error {
		var item entity.Item
		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"status code": http.StatusBadRequest,
				"message": err.Error(),
			})
		}

		item, err = handler.ItemUsecase.GetItem(id)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status code": http.StatusInternalServerError,
				"message": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"status code": http.StatusOK,
			"message": "success get items by id",
			"data":   item,
		})
	}
}

func (handler ItemHandler) GetItemByCategory() echo.HandlerFunc {
	return func(e echo.Context) error {
		var item entity.Item
		category_id, err := strconv.Atoi(e.Param("category_id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"status code": http.StatusBadRequest,
				"message": err.Error(),
			})
		}

		item, err = handler.ItemUsecase.GetItemByCategory(category_id)
		if err != nil {
			return e.JSON(http.StatusNotFound, map[string]interface{}{
				"status code": http.StatusNotFound,
				"message": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"status code": http.StatusOK,
			"message": "success get items by category",
			"data":   item,
		})
	}
}

func (handler ItemHandler) GetItemByName() echo.HandlerFunc {
	return func(e echo.Context) error {
		var item entity.Item
		name:=e.QueryParam("keyword")
		item, err := handler.ItemUsecase.GetItemByName(name)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status code": http.StatusInternalServerError,
				"message": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"status code": http.StatusOK,
			"message": "success get items by name",
			"data":   item,
		})
	}
}

func (handler ItemHandler) CreateItem() echo.HandlerFunc {
	return func(e echo.Context) error {
		var item entity.Item
		if err := e.Bind(&item); err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"status code": http.StatusBadRequest,
				"message": err.Error(),
			})
		}

		// Validasi input menggunakan package validator
		validate := validator.New()
		if err := validate.Struct(item); err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"status code": http.StatusBadRequest,
				"message": err.Error(),
			})
		}

		err := handler.ItemUsecase.CreateItem(item)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status code": http.StatusInternalServerError,
				"message": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"status code": http.StatusOK,
			"message": "success create item",
			"data":   item,
		})
	}
}
func (handler ItemHandler) UpdateItem() echo.HandlerFunc {
	var item entity.Item

	return func(e echo.Context) error {
		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"status code": http.StatusBadRequest,
				"message": err.Error(),
			})
		}

		err = handler.ItemUsecase.FindItem(id)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status code": http.StatusInternalServerError,
				"message": err.Error(),
			})
		}

		if err := e.Bind(&item); err != nil {
			return e.JSON(http.StatusNotFound, map[string]interface{}{
				"status code": http.StatusNotFound,
				"message": err.Error(),
			})
		}

		err = handler.ItemUsecase.UpdateItem(id, item)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status code": http.StatusInternalServerError,
				"message": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
				"status code": http.StatusOK,
				"message": "success update category",
				"data":item,

			})
	}
}
func (handler ItemHandler) DeleteItem() echo.HandlerFunc {
	return func(e echo.Context) error {
		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"status code": http.StatusBadRequest,
				"message": err.Error(),
			})
		}

		err = handler.ItemUsecase.DeleteItem(id)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status code": http.StatusInternalServerError,
				"message": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"status code": http.StatusOK,
			"message": "Success Delete Item`",
		})
	}
}
