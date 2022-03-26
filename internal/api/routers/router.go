package routers

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/juanees/todo-app-api/internal/api/handlers"
	"github.com/juanees/todo-app-api/internal/api/responses"
)

func New() *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(gzip.Gzip(gzip.DefaultCompression))
	router.Use(cors.Default())

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, responses.NewNotFound(c.Request.URL.Path))
	})
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	return router
}

func AddTodos(router *gin.Engine) {
	todosGroup := router.Group("api/v1/todos")

	todosGroup.GET("/", handlers.GetTodos)
	todosGroup.GET("/:code", handlers.GetTodo)
	todosGroup.POST("/", handlers.AddTodo)
	todosGroup.PATCH("/:code", handlers.UpdateTodo)
	todosGroup.DELETE("/:code", handlers.DeleteTodo)
}
