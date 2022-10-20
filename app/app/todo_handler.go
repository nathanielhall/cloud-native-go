package app

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/nathanielhall/cloud-native-go/repository"
)

func (app *App) HandleListTodos(w http.ResponseWriter, r *http.Request) {
	app.logger.Debug().Msg("HandleListTodos")

	repo := repository.NewTodoRepository(app.db, app.logger)
	todos, err := repo.GetAll()

	if err != nil {
		app.logger.Warn().Err(err).Msg("")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"error": "%v"}`, appErrDataAccessFailure)
		return
	}
	if todos == nil {
		fmt.Fprint(w, "[]")
		return
	}

	app.logger.Log().Msg("TODOS")

	// dtos := todos.ToDto()
	if err := json.NewEncoder(w).Encode(todos); err != nil {
		app.logger.Warn().Err(err).Msg("")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"error": "%v"}`, appErrJsonCreationFailure)
		return
	}
}
