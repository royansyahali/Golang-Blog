package controllers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type CommentController interface {
	InsertComment(http.ResponseWriter, *http.Request, httprouter.Params)
}
