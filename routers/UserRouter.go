package routers

import (
	"blog/controllers"
	"blog/middleware"

	"github.com/julienschmidt/httprouter"
)

func UserRouter(r *httprouter.Router, c controllers.UserController, a middleware.AuthMiddleware) {
	r.POST("/blog/user", c.InsertUser)
	r.PUT("/blog/user/update/:id", a.Login(c.UpdateUser))
	r.GET("/blog/user/:id", c.FindByIdUser)
	r.GET("/blog/users/all", c.GetAllUser)
	r.DELETE("/blog/user/delete/:id", a.Login(c.DeleteUser))
}
