package handlers

import (
	"fmt"
	"net/http"

	"github.com/juanees/todo-app-api/internal/api/mappers"
	"github.com/juanees/todo-app-api/internal/api/models"
	"github.com/juanees/todo-app-api/internal/api/requests"
	"github.com/juanees/todo-app-api/internal/api/responses"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

// TODO: ADD DEPENDENCY INJECTION OR IOC FOR TESTING
func init() {
	if dbConnection, err := gorm.Open(sqlite.Open("very-production-ready.db"), &gorm.Config{}); err != nil {
		panic("failed to connect database")
	} else {
		db = dbConnection
	}

	// Migrate the schema
	if err := db.AutoMigrate(&models.Todo{}); err != nil {
		panic("failed to migrate table Todo")
	}

	var todos []models.Todo
	if result := db.Find(&todos); result.Error == nil && result.RowsAffected == 0 {
		//seed data
		seedTodos := []models.Todo{
			{
				Model:       gorm.Model{},
				Code:        uuid.New(),
				Title:       "Todo-app api",
				Description: "Finish this project",
				Completed:   false,
			},
			{
				Model:       gorm.Model{},
				Code:        uuid.New(),
				Title:       "Project Idea",
				Description: "Find a project idea to practice and learn",
				Completed:   true,
			},
			{
				Model:       gorm.Model{},
				Code:        uuid.New(),
				Title:       "Todo-app frontend",
				Description: "Start working in the frontend (using React, axios and some UI Lib)",
				Completed:   false,
			},
		}

		for _, todo := range seedTodos {
			if result := db.Create(&todo); result.Error != nil {
				panic("failed to seed data to table Todo")
			}

		}
	}
}

func GetTodos(c *gin.Context) {

	var todos []models.Todo
	db.Find(&todos)

	var todosDto []responses.Todo

	for _, td := range todos {
		todosDto = append(todosDto, mappers.MapTodoToResponse(td))
	}

	c.JSON(http.StatusOK, todosDto)
}

func GetTodo(c *gin.Context) {

	code := c.Param("code")

	if todo, err := getTodoFromDatabase(code); err != nil {
		c.JSON(http.StatusNotFound, responses.NewNotFound(err.Error()))
	} else {
		c.JSON(http.StatusOK, mappers.MapTodoToResponse(todo))
	}
}

func AddTodo(c *gin.Context) {

	if todoDto, err := getTodoFromRequest(c); err != nil {
		c.JSON(http.StatusBadRequest, responses.NewBadRequest(err.Error()))
	} else {
		todo := mappers.MapTodoFromRequest(todoDto)
		todo.Code = uuid.New()
		if result := db.Create(&todo); result.Error != nil {
			c.JSON(http.StatusInternalServerError, responses.NewInternalServerError(result.Error.Error()))
			return
		}
		c.JSON(http.StatusOK, mappers.MapTodoToResponse(todo))
	}
}

func UpdateTodo(c *gin.Context) {

	code := c.Param("code")

	if todo, err := getTodoFromDatabase(code); err != nil {
		c.JSON(http.StatusNotFound, responses.NewNotFound(err.Error()))
	} else {
		if updatedTodo, err := getTodoFromRequest(c); err != nil {
			c.JSON(http.StatusBadRequest, responses.NewBadRequest(err.Error()))
		} else {
			todo.Completed = *updatedTodo.Completed
			todo.Title = updatedTodo.Title
			todo.Description = updatedTodo.Description
			db.Save(&todo)
			c.JSON(http.StatusOK, mappers.MapTodoToResponse(todo))
		}
	}
}

func DeleteTodo(c *gin.Context) {

	code := c.Param("code")

	if result := db.Where("code = ? ", code).Delete(&models.Todo{}); result.Error != nil {
		c.JSON(http.StatusInternalServerError, responses.NewInternalServerError(result.Error.Error()))
		return
	}
	c.JSON(http.StatusOK, true)
}

func getTodoFromDatabase(code string) (models.Todo, error) {

	var todo models.Todo

	if result := db.First(&todo, "code = ? ", code); result.RowsAffected == 0 || result.Error != nil {
		return models.Todo{}, fmt.Errorf("Todo with code %s", code)
	}
	return todo, nil
}

func getTodoFromRequest(c *gin.Context) (requests.Todo, error) {

	var json requests.Todo
	if err := c.ShouldBindJSON(&json); err != nil {
		return requests.Todo{}, err
	}

	return json, nil
}
