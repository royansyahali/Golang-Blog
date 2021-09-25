package impl

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"blog/controllers"
	"blog/entities"
	"blog/payloads/request"
	"blog/payloads/response"
	implRepo "blog/repositories/impl"
	"blog/services"
	implSer "blog/services/impl"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

type CategoryControllerImpl struct {
	CategoryService services.CategoryService
}

func SetupCategory(db *sql.DB, v *validator.Validate) controllers.CategoryController {
	repo := implRepo.NewCategoryRepository()
	service := implSer.NewCategoryService(repo, db, v)
	controller := CategoryControllerImpl{CategoryService: service}
	return &controller
}

func (c *CategoryControllerImpl) InsertCategory(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Header.Get("content-type") == "application/json" {
		reqCategory := &request.CategoryRequest{}
		decode := json.NewDecoder(r.Body)
		if err := decode.Decode(reqCategory); err != nil {
			response.ResponseError(w, http.StatusUnprocessableEntity, err)
			return
		}
		reqCategory.Prepare()
		if err := c.CategoryService.InsertCategory(r.Context(), reqCategory); err != nil {
			response.ResponseError(w, http.StatusInternalServerError, err)
			return
		}
		response.ResponseMessage(w, http.StatusText(http.StatusOK), http.StatusOK)
		return
	}
	response.ResponseError(w, http.StatusNotFound, errors.New("must content be application/json"))
}

func (c *CategoryControllerImpl) FindByIdCategory(w http.ResponseWriter, r *http.Request, id_Category httprouter.Params) {
	if r.Header.Get("content-type") == "application/json" {
		var category entities.Category
		id, err := strconv.Atoi(id_Category.ByName("id"))
		if err != nil {
			response.ResponseError(w, http.StatusInternalServerError, err)
			return
		}
		category, err = c.CategoryService.FindByIdCategory(r.Context(), id)
		if err != nil {
			response.ResponseError(w, http.StatusInternalServerError, err)
			return
		}
		response.ResponseMessage(w, category, http.StatusOK)
		return
	}
	response.ResponseError(w, http.StatusNotFound, errors.New("must content be application/json"))
}

func (c *CategoryControllerImpl) GetAllCategory(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Header.Get("content-type") == "application/json" {
		var category []entities.Category
		category, err := c.CategoryService.GetAllCategory(r.Context())
		if err != nil {
			response.ResponseError(w, http.StatusInternalServerError, err)
			return
		}
		response.ResponseMessage(w, category, http.StatusOK)
		return
	}
	response.ResponseError(w, http.StatusNotFound, errors.New("must content be application/json"))
}
