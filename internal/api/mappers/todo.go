package mappers

import (
	"github.com/juanees/todo-app-api/internal/api/models"
	"github.com/juanees/todo-app-api/internal/api/requests"
	"github.com/juanees/todo-app-api/internal/api/responses"
)

func MapTodoFromRequest(dto requests.Todo) models.Todo {
	return models.Todo{
		Code:        dto.Code,
		Title:       dto.Title,
		Description: dto.Description,
		Completed:   *dto.Completed,
	}
}

func MapTodoToResponse(todo models.Todo) responses.Todo {
	return responses.Todo{
		Code:        todo.Code,
		Title:       todo.Title,
		Description: todo.Description,
		Completed:   todo.Completed,
		CreatedAt:   todo.CreatedAt,
	}
}
