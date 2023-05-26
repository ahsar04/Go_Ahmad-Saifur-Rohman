package http

import (
	db "remidial/internal/adapters/db/mysql"
	handler "remidial/internal/adapters/http/handler"
	middlewares "remidial/internal/adapters/http/middleware"
	repository "remidial/internal/adapters/repository"
	usecase "remidial/internal/application/usecase"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	// user management
	userRepo    repository.UserRepository
	userHandler handler.UserHandler
	userUsecase usecase.UserUseCase
	// auth
	AuthHandler handler.AuthHandler
	// category management
	categoryRepo    repository.CategoryRepository
	categoryHandler handler.CategoryHandler
	categoryUsecase usecase.CategoryUseCase
	// item management
	itemRepo    repository.ItemRepository
	itemHandler handler.ItemHandler
	itemUsecase usecase.ItemUseCase
)

func declare() {
	userRepo = repository.UserRepository{DB: db.DbMysql}
	userUsecase = usecase.UserUseCase{Repo: userRepo}
	userHandler = handler.UserHandler{UserUsecase: userUsecase}
	// auth
	AuthHandler = handler.AuthHandler{Usecase: userUsecase}
	// category
	categoryRepo = repository.CategoryRepository{DB: db.DbMysql}
	categoryUsecase = usecase.CategoryUseCase{Repo: categoryRepo}
	categoryHandler = handler.CategoryHandler{CategoryUsecase: categoryUsecase}
	// item
	itemRepo = repository.ItemRepository{DB: db.DbMysql}
	itemUsecase = usecase.ItemUseCase{Repo: itemRepo}
	itemHandler = handler.ItemHandler{ItemUsecase: itemUsecase}

}

func InitRoutes() *echo.Echo {
	db.Init()
	declare()

	e := echo.New()
	e.POST("/login", AuthHandler.Login())
	e.POST("/register", AuthHandler.Register())

	// admin group
	admin := e.Group("/")
	admin.Use(middleware.Logger())
	admin.Use(middlewares.AuthMiddleware())

	admin.GET("users", userHandler.GetAllUsers())
	admin.GET("users/:id", userHandler.GetUser())
	// admin.POST("users", userHandler.CreateUser())
	// admin.DELETE("users/:id", userHandler.DeleteUser())
	
	admin.GET("categories", categoryHandler.GetAllCategories())
	admin.GET("categories/:id", categoryHandler.GetCategory())
	admin.POST("categories", categoryHandler.CreateCategory())
	admin.PUT("categories/:id", categoryHandler.UpdateCategory())
	admin.DELETE("categories/:id", categoryHandler.DeleteCategory())
	
	admin.GET("items", itemHandler.GetAllItems())
	admin.GET("items/:id", itemHandler.GetItem())
	admin.GET("items/category/:category_id", itemHandler.GetItemByCategory())
	admin.GET("items/search", itemHandler.GetItemByName())
	admin.POST("items", itemHandler.CreateItem())
	admin.PUT("items/:id", itemHandler.UpdateItem())
	admin.DELETE("items/:id", itemHandler.DeleteItem())
	return e
}
