package impl

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"blog/controllers"
	"blog/payloads/request"
	deleterequest "blog/payloads/request/deleteRequest"
	updaterequest "blog/payloads/request/updateRequest"
	"blog/payloads/response"
	implRepo "blog/repositories/impl"
	"blog/services"
	implSer "blog/services/impl"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

type PostControllerImpl struct {
	PostService services.PostService
}

func SetupPost(db *sql.DB, v *validator.Validate) controllers.PostController {
	repo := implRepo.NewPostRepository()
	service := implSer.NewPostService(repo, db, v)
	controller := PostControllerImpl{PostService: service}
	return &controller
}

func (p *PostControllerImpl) InsertPost(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Header.Get("content-type") == "application/json" {
		reqPost := &request.PostRequest{}
		decode := json.NewDecoder(r.Body)
		if err := decode.Decode(reqPost); err != nil {
			response.ResponseError(w, http.StatusUnprocessableEntity, err)
			return
		}
		reqPost.Prepare()
		if err := p.PostService.InsertPost(r.Context(), reqPost); err != nil {
			response.ResponseError(w, http.StatusInternalServerError, err)
			return
		}
		response.ResponseMessage(w, http.StatusText(http.StatusOK), http.StatusOK)
		return
	}
	response.ResponseError(w, http.StatusNotFound, errors.New("must content be application/json"))
}

func (p *PostControllerImpl) UpdatePost(w http.ResponseWriter, r *http.Request, id_Post httprouter.Params) {
	if r.Header.Get("content-type") == "application/json" {
		reqPost := &updaterequest.PostUpdate{}
		id, err := strconv.Atoi(id_Post.ByName("id"))
		if err != nil {
			response.ResponseError(w, http.StatusInternalServerError, err)
			return
		}
		decode := json.NewDecoder(r.Body)
		if err := decode.Decode(reqPost); err != nil {
			response.ResponseError(w, http.StatusUnprocessableEntity, err)
			return
		}
		reqPost.Prepare()
		if err := p.PostService.UpdatePost(r.Context(), reqPost, id); err != nil {
			response.ResponseError(w, http.StatusInternalServerError, err)
			return
		}
		response.ResponseMessage(w, http.StatusText(http.StatusOK), http.StatusOK)
		return
	}
	response.ResponseError(w, http.StatusNotFound, errors.New("must content be application/json"))
}

func (p *PostControllerImpl) DeletePost(w http.ResponseWriter, r *http.Request, id_Post httprouter.Params) {
	if r.Header.Get("content-type") == "application/json" {
		id, err := strconv.Atoi(id_Post.ByName("id"))
		if err != nil {
			response.ResponseError(w, http.StatusInternalServerError, err)
			return
		}
		reqPost := &deleterequest.PostDelete{}
		decode := json.NewDecoder(r.Body)
		if err := decode.Decode(reqPost); err != nil {
			response.ResponseError(w, http.StatusUnprocessableEntity, err)
			return
		}
		if err := p.PostService.DeletePost(r.Context(), reqPost.UserId, id); err != nil {
			response.ResponseError(w, http.StatusInternalServerError, err)
			return
		}
		response.ResponseMessage(w, http.StatusText(http.StatusOK), http.StatusOK)
		return
	}
	response.ResponseError(w, http.StatusNotFound, errors.New("must content be application/json"))
}

func (p *PostControllerImpl) FindByIdPost(w http.ResponseWriter, r *http.Request, id_Post httprouter.Params) {
	if r.Header.Get("content-type") == "application/json" {
		var Post response.PostResponse
		id, err := strconv.Atoi(id_Post.ByName("id"))
		if err != nil {
			response.ResponseError(w, http.StatusInternalServerError, err)
			return
		}
		Post, err = p.PostService.FindByIdPost(r.Context(), id)
		if err != nil {
			response.ResponseError(w, http.StatusInternalServerError, err)
			return
		}
		response.ResponseMessage(w, Post, http.StatusOK)
		return
	}
	response.ResponseError(w, http.StatusNotFound, errors.New("must content be application/json"))
}

func (p *PostControllerImpl) GetAllPost(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Header.Get("content-type") == "application/json" {
		var post []response.PostResponse
		post, err := p.PostService.GetAllPost(r.Context())
		if err != nil {
			response.ResponseError(w, http.StatusInternalServerError, err)
			return
		}
		response.ResponseMessage(w, post, http.StatusOK)
		return
	}
	response.ResponseError(w, http.StatusNotFound, errors.New("must content be application/json"))
}
