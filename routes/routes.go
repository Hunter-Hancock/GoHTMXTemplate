package routes

import (
	"fmt"

	"github.com/Hunter-Hancock/dbproject/app"
	"github.com/go-chi/chi/v5"
)

func RegisterRoutes() *chi.Mux {
	r := chi.NewRouter()

	app, err := app.NewApplication()
	if err != nil {
		fmt.Println(err)
	}

	r.Use(app.Middleware.RequireUser)
	r.Get("/order", app.TestHandler.Test)

	return r
}
