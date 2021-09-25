package routers

import (
	"github.com/julienschmidt/httprouter"
	"github.com/royansyahali/blog/controllers"
	"github.com/royansyahali/blog/middleware"
)

func PostRouter(r *httprouter.Router, c controllers.PostController, a middleware.AuthMiddleware) {
	r.POST("/blog/post", a.Login(c.InsertPost))
	r.PUT("/blog/post/update/:id", a.Login(c.UpdatePost))
	r.GET("/blog/post/:id", c.FindByIdPost)
	r.GET("/blog/posts/all", c.GetAllPost)
	r.DELETE("/blog/post/delete/:id", a.Login(c.DeletePost))
}
