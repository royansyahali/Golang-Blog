package impl

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	"github.com/royansyahali/blog/controllers"
	"github.com/royansyahali/blog/entities"
	"github.com/royansyahali/blog/payloads/request"
	"github.com/royansyahali/blog/payloads/response"
	implRepo "github.com/royansyahali/blog/repositories/impl"
	"github.com/royansyahali/blog/services"
	implSer "github.com/royansyahali/blog/services/impl"
)

type TagControllerImpl struct {
	TagService services.TagService
}

func SetupTag(db *sql.DB, v *validator.Validate) controllers.TagController {
	repo := implRepo.NewTagRepository()
	service := implSer.NewTagService(repo, db, v)
	controller := TagControllerImpl{TagService: service}
	return &controller
}

func (t *TagControllerImpl) InsertTag(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Header.Get("content-type") == "application/json" {
		reqTag := &request.TagRequest{}
		decode := json.NewDecoder(r.Body)
		if err := decode.Decode(reqTag); err != nil {
			response.ResponseError(w, http.StatusUnprocessableEntity, err)
			return
		}
		reqTag.Prepare()
		if err := t.TagService.InsertTag(r.Context(), reqTag); err != nil {
			response.ResponseError(w, http.StatusInternalServerError, err)
			return
		}
		response.ResponseMessage(w, http.StatusText(http.StatusOK), http.StatusOK)
		return
	}
	response.ResponseError(w, http.StatusNotFound, errors.New("must content be application/json"))
}

func (t *TagControllerImpl) FindByIdTag(w http.ResponseWriter, r *http.Request, id_Tag httprouter.Params) {
	if r.Header.Get("content-type") == "application/json" {
		var tag entities.Tag
		id, err := strconv.Atoi(id_Tag.ByName("id"))
		if err != nil {
			response.ResponseError(w, http.StatusInternalServerError, err)
			return
		}
		tag, err = t.TagService.FindByIdTag(r.Context(), id)
		if err != nil {
			response.ResponseError(w, http.StatusInternalServerError, err)
			return
		}
		response.ResponseMessage(w, tag, http.StatusOK)
		return
	}
	response.ResponseError(w, http.StatusNotFound, errors.New("must content be application/json"))
}

func (t *TagControllerImpl) GetAllTag(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Header.Get("content-type") == "application/json" {
		var tag []entities.Tag
		tag, err := t.TagService.GetAllTag(r.Context())
		if err != nil {
			response.ResponseError(w, http.StatusInternalServerError, err)
			return
		}
		response.ResponseMessage(w, tag, http.StatusOK)
		return
	}
	response.ResponseError(w, http.StatusNotFound, errors.New("must content be application/json"))
}
