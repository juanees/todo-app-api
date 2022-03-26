package main

import (
	"github.com/juanees/todo-app-api/internal/api"
)

func main() {
	api := api.New()
	api.Start()
}
