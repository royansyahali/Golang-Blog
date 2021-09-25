package routers

import (
	"github.com/julienschmidt/httprouter"
	"github.com/royansyahali/blog/controllers"
	"github.com/royansyahali/blog/middleware"
)

func CommentRouter(r *httprouter.Router, c controllers.CommentController, a middleware.AuthMiddleware) {
	r.POST("/blog/comment", a.Login(c.InsertComment))
}
