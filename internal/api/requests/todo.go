package requests

import "github.com/google/uuid"

type Todo struct {
	Code        uuid.UUID `json:"code"`
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Completed   *bool     `json:"completed" binding:"required"`
}
