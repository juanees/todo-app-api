package api

import (
	"fmt"
	"os"

	"github.com/juanees/todo-app-api/internal/api/routers"
)

type webAPI struct{}

func New() *webAPI {
	return &webAPI{}
}

func (wa webAPI) Start() {
	router := routers.New()

	scope := os.Getenv("SCOPE")
	routers.AddTodos(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	var url string

	if scope == "PRODUCTION" {
		url = fmt.Sprintf(":%s", port)
	} else {
		url = fmt.Sprintf("localhost:%s", port)
	}

	if err := router.Run(url); err != nil {
		panic(err)
	}
}
