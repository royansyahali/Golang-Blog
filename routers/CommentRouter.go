package routers

import (
	"blog/controllers"
	"blog/middleware"

	"github.com/julienschmidt/httprouter"
)

func CommentRouter(r *httprouter.Router, c controllers.CommentController, a middleware.AuthMiddleware) {
	r.POST("/blog/comment", a.Login(c.InsertComment))
}
