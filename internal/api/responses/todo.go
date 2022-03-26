package responses

import (
	"time"

	"github.com/google/uuid"
)

type Todo struct {
	Code        uuid.UUID `json:"code"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"created_at"`
}
