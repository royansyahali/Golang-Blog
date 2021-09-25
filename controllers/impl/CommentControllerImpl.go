package impl

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	"github.com/royansyahali/blog/controllers"
	"github.com/royansyahali/blog/payloads/request"
	"github.com/royansyahali/blog/payloads/response"
	implRepo "github.com/royansyahali/blog/repositories/impl"
	"github.com/royansyahali/blog/services"
	implSer "github.com/royansyahali/blog/services/impl"
)

type CommentControllerImpl struct {
	CommentService services.CommentService
}

func SetupComment(db *sql.DB, v *validator.Validate) controllers.CommentController {
	repo := implRepo.NewCommentRepository()
	service := implSer.NewCommentService(repo, db, v)
	controller := CommentControllerImpl{CommentService: service}
	return &controller
}

func (c *CommentControllerImpl) InsertComment(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Header.Get("content-type") == "application/json" {
		reqComment := &request.CommentRequest{}
		decode := json.NewDecoder(r.Body)
		if err := decode.Decode(reqComment); err != nil {
			response.ResponseError(w, http.StatusUnprocessableEntity, err)
			return
		}
		reqComment.Prepare()
		if err := c.CommentService.InsertComment(r.Context(), reqComment); err != nil {
			response.ResponseError(w, http.StatusInternalServerError, err)
			return
		}
		response.ResponseMessage(w, http.StatusText(http.StatusOK), http.StatusOK)
		return
	}
	response.ResponseError(w, http.StatusNotFound, errors.New("must content be application/json"))
}
