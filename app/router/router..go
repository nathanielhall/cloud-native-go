package router

import (
	"github.com/go-chi/chi"
	"github.com/nathanielhall/cloud-native-go/app/app"
	"github.com/nathanielhall/cloud-native-go/app/requestlog"
	"github.com/nathanielhall/cloud-native-go/app/router/middleware"
)

func New(a *app.App) *chi.Mux {
	l := a.Logger()
	r := chi.NewRouter()

	r.Method("GET", "/", requestlog.NewHandler(a.HandleIndex, l))
	// r.Get("/healthz/liveness", app.HandleLive)
	// r.Method("GET", "/healthz/readiness", requestlog.NewHandler(a.HandleReady, l))

	// Routes for APIs
	r.Route("/api/v1", func(r chi.Router) {
		r.Use(middleware.ContentTypeJson)

		// Routes for todos
		r.Method("GET", "/todos", requestlog.NewHandler(a.HandleListTodos, l))
		// r.Method("POST", "/todos", handler.NewHandler(a.HandleCreateBook, l))
		// r.Method("GET", "/todos/{id}", handler.NewHandler(a.HandleReadBook, l))
		// r.Method("PUT", "/todos/{id}", handler.NewHandler(a.HandleUpdateBook, l))
		// r.Method("DELETE", "/todos/{id}", handler.NewHandler(a.HandleDeleteBook, l))
	})

	r.Method("GET", "/", requestlog.NewHandler(a.HandleIndex, l))
	return r
}