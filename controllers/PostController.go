package controllers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type PostController interface {
	InsertPost(http.ResponseWriter, *http.Request, httprouter.Params)
	UpdatePost(http.ResponseWriter, *http.Request, httprouter.Params)
	DeletePost(http.ResponseWriter, *http.Request, httprouter.Params)
	FindByIdPost(http.ResponseWriter, *http.Request, httprouter.Params)
	GetAllPost(http.ResponseWriter, *http.Request, httprouter.Params)
}
