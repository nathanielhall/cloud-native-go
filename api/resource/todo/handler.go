package todo

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	e "github.com/nathanielhall/cloud-native-go/api/resource/common/err"
	"github.com/nathanielhall/cloud-native-go/util/validator"
	"gorm.io/gorm"
)


func (a *API) List(w http.ResponseWriter, r *http.Request) {
	todos, err := a.repository.GetAll()
	if err != nil {
		a.logger.Error().Err(err).Msg("")
		e.ServerError(w, e.DataAccessFailure)
		return
	}

	if todos == nil {
		fmt.Fprint(w, "[]")
		return
	}

	if err := json.NewEncoder(w).Encode(todos.ToDto()); err != nil {
		a.logger.Error().Err(err).Msg("")
		e.ServerError(w, e.JsonEncodingFailure)
		return
	}
}

func (a *API) Create(w http.ResponseWriter, r *http.Request) {
	form := &Form{}
	if err := json.NewDecoder(r.Body).Decode(form); err != nil {
		a.logger.Error().Err(err).Msg("")
		e.BadRequest(w, e.JsonDecodingFailure)
		return
	}

	if err := a.validator.Struct(form); err != nil {
		resp := validator.ToErrResponse(err)
		if resp == nil {
			e.ServerError(w, e.FormErrResponseFailure)
			return
		}

		respBody, err := json.Marshal(resp)
		if err != nil {
			a.logger.Error().Err(err).Msg("")
			e.ServerError(w, e.JsonEncodingFailure)
			return
		}

		e.ValidationErrors(w, respBody)
		return
	}

	todo, err := a.repository.Create(form.ToModel())
	if err != nil {
		a.logger.Error().Err(err).Msg("")
		e.ServerError(w, e.DataCreationFailure)
		return
	}

	a.logger.Info().Str("id", todo.ID.String()).Msg("new todo created")
	w.WriteHeader(http.StatusCreated)
}

func (a *API) Read(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		e.BadRequest(w, e.InvalidIdInUrlParam)
		return
	}

	todo, err := a.repository.GetById(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		a.logger.Error().Err(err).Msg("")
		e.ServerError(w, e.DataAccessFailure)
		return
	}

	dto := todo.ToDto()
	if err := json.NewEncoder(w).Encode(dto); err != nil {
		a.logger.Error().Err(err).Msg("")
		e.ServerError(w, e.JsonEncodingFailure)
		return
	}
}
func (a *API) Update(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		e.BadRequest(w, e.InvalidIdInUrlParam)
		return
	}

	form := &Form{}
	if err := json.NewDecoder(r.Body).Decode(form); err != nil {
		a.logger.Error().Err(err).Msg("")
		e.BadRequest(w, e.JsonDecodingFailure)
		return
	}

	if err := a.validator.Struct(form); err != nil {
		resp := validator.ToErrResponse(err)
		if resp == nil {
			e.ServerError(w, e.FormErrResponseFailure)
			return
		}

		respBody, err := json.Marshal(resp)
		if err != nil {
			a.logger.Error().Err(err).Msg("")
			e.ServerError(w, e.JsonEncodingFailure)
			return
		}

		e.ValidationErrors(w, respBody)
		return
	}

	todoModel := form.ToModel()
	todoModel.ID = id

	if err := a.repository.Update(todoModel); err != nil {
		if err == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		a.logger.Error().Err(err).Msg("")
		e.ServerError(w, e.DataUpdateFailure)
		return
	}

	a.logger.Info().Str("id", id.String()).Msg("todo updated")
}

func (a *API) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		e.BadRequest(w, e.InvalidIdInUrlParam)
		return
	}

	if err := a.repository.Delete(id); err != nil {
		a.logger.Error().Err(err).Msg("")
		e.ServerError(w, e.DataDeletionFailure)
		return
	}

	a.logger.Info().Str("id", id.String()).Msg("todo deleted")
}
