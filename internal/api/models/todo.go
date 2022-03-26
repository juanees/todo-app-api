package models

import (
	"gorm.io/gorm"

	"github.com/google/uuid"
)

type Todo struct {
	gorm.Model
	Code        uuid.UUID `gorm:"uniqueIndex"`
	Title       string
	Description string
	Completed   bool
}
