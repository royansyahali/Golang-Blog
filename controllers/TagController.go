package controllers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type TagController interface {
	InsertTag(http.ResponseWriter, *http.Request, httprouter.Params)
	FindByIdTag(http.ResponseWriter, *http.Request, httprouter.Params)
	GetAllTag(http.ResponseWriter, *http.Request, httprouter.Params)
}
