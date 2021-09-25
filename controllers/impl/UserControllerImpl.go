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
	updaterequest "blog/payloads/request/updateRequest"
	"blog/payloads/response"
	implRepo "blog/repositories/impl"
	"blog/services"
	implSer "blog/services/impl"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

type UserControllerImpl struct {
	UserService services.UserService
}

func SetupUser(db *sql.DB, v *validator.Validate) controllers.UserController {
	repo := implRepo.NewUserRepository()
	service := implSer.NewUserService(repo, db, v)
	controller := UserControllerImpl{UserService: service}
	return &controller
}

func (u *UserControllerImpl) InsertUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Header.Get("content-type") == "application/json" {
		reqUser := &request.UserRequest{}
		decode := json.NewDecoder(r.Body)
		if err := decode.Decode(reqUser); err != nil {
			response.ResponseError(w, http.StatusUnprocessableEntity, err)
			return
		}
		reqUser.Prepare()
		if err := u.UserService.InsertUser(r.Context(), reqUser); err != nil {
			response.ResponseError(w, http.StatusInternalServerError, err)
			return
		}
		response.ResponseMessage(w, http.StatusText(http.StatusOK), http.StatusOK)
		return
	}
	response.ResponseError(w, http.StatusNotFound, errors.New("must content be application/json"))
}

func (u *UserControllerImpl) UpdateUser(w http.ResponseWriter, r *http.Request, id_user httprouter.Params) {
	if r.Header.Get("content-type") == "application/json" {
		reqUser := &updaterequest.UserUpdate{}
		id, err := strconv.Atoi(id_user.ByName("id"))
		if err != nil {
			response.ResponseError(w, http.StatusInternalServerError, err)
			return
		}
		decode := json.NewDecoder(r.Body)
		if err := decode.Decode(reqUser); err != nil {
			response.ResponseError(w, http.StatusUnprocessableEntity, err)
			return
		}
		reqUser.Prepare()
		if err := u.UserService.UpdateUser(r.Context(), reqUser, id); err != nil {
			response.ResponseError(w, http.StatusInternalServerError, err)
			return
		}
		response.ResponseMessage(w, http.StatusText(http.StatusOK), http.StatusOK)
		return
	}
	response.ResponseError(w, http.StatusNotFound, errors.New("must content be application/json"))
}

func (u *UserControllerImpl) DeleteUser(w http.ResponseWriter, r *http.Request, id_user httprouter.Params) {
	if r.Header.Get("content-type") == "application/json" {
		id, err := strconv.Atoi(id_user.ByName("id"))
		if err != nil {
			response.ResponseError(w, http.StatusInternalServerError, err)
			return
		}
		if err := u.UserService.DeleteUser(r.Context(), id); err != nil {
			response.ResponseError(w, http.StatusInternalServerError, err)
			return
		}
		response.ResponseMessage(w, http.StatusText(http.StatusOK), http.StatusOK)
		return
	}
	response.ResponseError(w, http.StatusNotFound, errors.New("must content be application/json"))
}

func (u *UserControllerImpl) FindByIdUser(w http.ResponseWriter, r *http.Request, id_user httprouter.Params) {
	if r.Header.Get("content-type") == "application/json" {
		var user entities.User
		id, err := strconv.Atoi(id_user.ByName("id"))
		if err != nil {
			response.ResponseError(w, http.StatusInternalServerError, err)
			return
		}
		user, err = u.UserService.FindByIdUser(r.Context(), id)
		if err != nil {
			response.ResponseError(w, http.StatusInternalServerError, err)
			return
		}
		response.ResponseMessage(w, user, http.StatusOK)
		return
	}
	response.ResponseError(w, http.StatusNotFound, errors.New("must content be application/json"))
}

func (u *UserControllerImpl) GetAllUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Header.Get("content-type") == "application/json" {
		var user []entities.User
		user, err := u.UserService.GetAllUser(r.Context())
		if err != nil {
			response.ResponseError(w, http.StatusInternalServerError, err)
			return
		}
		response.ResponseMessage(w, user, http.StatusOK)
		return
	}
	response.ResponseError(w, http.StatusNotFound, errors.New("must content be application/json"))
}
