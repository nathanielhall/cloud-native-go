package todo

import (
	"encoding/json"
	"fmt"
	"net/http"

	e "github.com/nathanielhall/cloud-native-go/api/resource/common/err"
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

// Create godoc
//
//	@summary		Create book
//	@description	Create book
//	@tags			books
//	@accept			json
//	@produce		json
//	@param			body	body	Form	true	"Book form"
//	@success		201
//	@failure		400	{object}	err.Error
//	@failure		422	{object}	err.Errors
//	@failure		500	{object}	err.Error
//	@router			/books [post]
// func (a *API) Create(w http.ResponseWriter, r *http.Request) {
// 	form := &Form{}
// 	if err := json.NewDecoder(r.Body).Decode(form); err != nil {
// 		a.logger.Error().Err(err).Msg("")
// 		e.BadRequest(w, e.JsonDecodingFailure)
// 		return
// 	}

// 	if err := a.validator.Struct(form); err != nil {
// 		resp := validator.ToErrResponse(err)
// 		if resp == nil {
// 			e.ServerError(w, e.FormErrResponseFailure)
// 			return
// 		}

// 		respBody, err := json.Marshal(resp)
// 		if err != nil {
// 			a.logger.Error().Err(err).Msg("")
// 			e.ServerError(w, e.JsonEncodingFailure)
// 			return
// 		}

// 		e.ValidationErrors(w, respBody)
// 		return
// 	}

// 	book, err := a.repository.CreateBook(form.ToModel())
// 	if err != nil {
// 		a.logger.Error().Err(err).Msg("")
// 		e.ServerError(w, e.DataCreationFailure)
// 		return
// 	}

// 	a.logger.Info().Str("id", book.ID.String()).Msg("new book created")
// 	w.WriteHeader(http.StatusCreated)
// }
