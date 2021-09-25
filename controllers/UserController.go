package controllers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type UserController interface {
	InsertUser(http.ResponseWriter, *http.Request, httprouter.Params)
	UpdateUser(http.ResponseWriter, *http.Request, httprouter.Params)
	DeleteUser(http.ResponseWriter, *http.Request, httprouter.Params)
	FindByIdUser(http.ResponseWriter, *http.Request, httprouter.Params)
	GetAllUser(http.ResponseWriter, *http.Request, httprouter.Params)
}
