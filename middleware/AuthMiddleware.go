package middleware

import (
	"github.com/julienschmidt/httprouter"
)

type AuthMiddleware interface {
	Login(httprouter.Handle) httprouter.Handle
}
