package impl

import (
	"database/sql"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/royansyahali/blog/entities"
	"github.com/royansyahali/blog/middleware"
	implRepo "github.com/royansyahali/blog/repositories/impl"
	"github.com/royansyahali/blog/services"
	implSer "github.com/royansyahali/blog/services/impl"
)

type AuthMiddlewareImpl struct {
	AuthService services.AuthService
}

func SetupAuth(db *sql.DB) middleware.AuthMiddleware {
	repo := implRepo.NewAuthRepository()
	service := implSer.NewAuthService(repo, db)
	middleware := AuthMiddlewareImpl{AuthService: service}
	return &middleware
}

func (u *AuthMiddlewareImpl) Login(handle httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		// Get the Basic Authentication credentials
		username, password, hasAuth := r.BasicAuth()
		user := &entities.User{
			Username: username,
			Password: password,
		}
		err := u.AuthService.Login(r.Context(), user)
		if hasAuth && err == nil {
			// Delegate request to the given handle
			handle(w, r, p)
		} else {
			// Request Basic Authentication otherwise
			w.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		}
	}
}
