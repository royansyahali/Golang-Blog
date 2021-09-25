package controllers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type CategoryController interface {
	InsertCategory(http.ResponseWriter, *http.Request, httprouter.Params)
	FindByIdCategory(http.ResponseWriter, *http.Request, httprouter.Params)
	GetAllCategory(http.ResponseWriter, *http.Request, httprouter.Params)
}
