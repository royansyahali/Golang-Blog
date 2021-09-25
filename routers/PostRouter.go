package routers

import (
	"blog/controllers"
	"blog/middleware"

	"github.com/julienschmidt/httprouter"
)

func PostRouter(r *httprouter.Router, c controllers.PostController, a middleware.AuthMiddleware) {
	r.POST("/blog/post", a.Login(c.InsertPost))
	r.PUT("/blog/post/update/:id", a.Login(c.UpdatePost))
	r.GET("/blog/post/:id", c.FindByIdPost)
	r.GET("/blog/posts/all", c.GetAllPost)
	r.DELETE("/blog/post/delete/:id", a.Login(c.DeletePost))
}
