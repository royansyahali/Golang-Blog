package routers

import (
	"github.com/julienschmidt/httprouter"
	"github.com/royansyahali/blog/controllers"
	"github.com/royansyahali/blog/middleware"
)

func TagRouter(r *httprouter.Router, c controllers.TagController, a middleware.AuthMiddleware) {
	r.POST("/blog/tag", a.Login(c.InsertTag))
	r.GET("/blog/tag/:id", c.FindByIdTag)
	r.GET("/blog/tags/all", c.GetAllTag)
}
