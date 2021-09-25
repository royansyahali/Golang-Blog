package routers

import (
	"github.com/julienschmidt/httprouter"
	"github.com/royansyahali/blog/controllers"
	"github.com/royansyahali/blog/middleware"
)

func UserRouter(r *httprouter.Router, c controllers.UserController, a middleware.AuthMiddleware) {
	r.POST("/blog/user", c.InsertUser)
	r.PUT("/blog/user/update/:id", a.Login(c.UpdateUser))
	r.GET("/blog/user/:id", c.FindByIdUser)
	r.GET("/blog/users/all", c.GetAllUser)
	r.DELETE("/blog/user/delete/:id", a.Login(c.DeleteUser))
}
