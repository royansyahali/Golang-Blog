package routers

import (
	"github.com/julienschmidt/httprouter"
	"github.com/royansyahali/blog/controllers"
	"github.com/royansyahali/blog/middleware"
)

func CategoryRouter(r *httprouter.Router, c controllers.CategoryController, a middleware.AuthMiddleware) {
	r.POST("/blog/category", a.Login(c.InsertCategory))
	r.GET("/blog/category/:id", c.FindByIdCategory)
	r.GET("/blog/categories/all", c.GetAllCategory)
}
