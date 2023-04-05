package routes

import (
	"code_structure/controllers"
	"code_structure/constants"
	m "code_structure/middlewares"

	"github.com/labstack/echo"
	mid "github.com/labstack/echo/middleware"
)

func New() *echo.Echo{
	
	e := echo.New()
	// Route / to handler function
	// users
	m.LogMiddleware(e)
	
	eJWT := e.Group("/")
	eJWT.Use(mid.JWT(constants.SECRET_JWT))
	// e.GET("/users", controllers.GetUsersController)
	eJWT.GET("users", controllers.GetUsersController)
	eJWT.GET("users/:id", controllers.GetUserController)
	e.POST("users", controllers.CreateUserController)
	e.POST("users/login", controllers.LoginUserController)
	eJWT.DELETE("users/:id", controllers.DeleteUserController)
	eJWT.PUT("users/:id", controllers.UpdateUserController)
	// books
	eJWT.GET("books", controllers.GetBooksController)
	eJWT.GET("books/:id", controllers.GetBookController)
	eJWT.POST("books", controllers.CreateBookController)
	eJWT.DELETE("books/:id", controllers.DeleteBookController)
	eJWT.PUT("books/:id", controllers.UpdateBookController)
	// books
	e.GET("/blogs", controllers.GetBlogsController)
	e.GET("/blogs/:id", controllers.GetBlogController)
	e.POST("/blogs", controllers.CreateBlogController)
	e.DELETE("/blogs/:id", controllers.DeleteBlogController)
	e.PUT("/blogs/:id", controllers.UpdateBlogController)


	return e
}
