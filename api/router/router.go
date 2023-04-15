package router

import (
	"github.com/go-chi/chi"
	"github.com/nathanielhall/cloud-native-go/api/requestlog"
	"github.com/nathanielhall/cloud-native-go/api/resource/todo"
	"github.com/nathanielhall/cloud-native-go/api/router/middleware"
	"github.com/nathanielhall/cloud-native-go/util/logger"
	"gorm.io/gorm"
)

func New(l *logger.Logger, db *gorm.DB) *chi.Mux {
	r := chi.NewRouter()

	r.Route("/api/v1", func(r chi.Router) {
		r.Use(middleware.ContentTypeJson)

		todoAPI := todo.New(l, db)
		r.Method("GET", "/todos", requestlog.NewHandler(todoAPI.List, l))
		// r.Method("POST", "/todos", requestlog.NewHandler(todoAPI.Create, l))
		// r.Method("GET", "/todos/{id}", requestlog.NewHandler(todoAPI.Read, l))
		// r.Method("PUT", "/todos/{id}", requestlog.NewHandler(todoAPI.Update, l))
		// r.Method("DELETE", "/todos/{id}", requestlog.NewHandler(todoAPI.Delete, l))
	})

	return r
}
