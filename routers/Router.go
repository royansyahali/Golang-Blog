package routers

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"blog/configs"
	"blog/controllers/impl"
	middleware "blog/middleware/impl"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

func SetupRouter() {
	v := validator.New()
	db, err := configs.Connection()
	if err != nil {
		log.Fatal(err)
	}
	AuthMiddleware := middleware.SetupAuth(db)
	UserController := impl.SetupUser(db, v)
	PostController := impl.SetupPost(db, v)
	TagController := impl.SetupTag(db, v)
	CategoryController := impl.SetupCategory(db, v)
	CommentController := impl.SetupComment(db, v)

	router := httprouter.New()
	UserRouter(router, UserController, AuthMiddleware)
	PostRouter(router, PostController, AuthMiddleware)
	TagRouter(router, TagController, AuthMiddleware)
	CategoryRouter(router, CategoryController, AuthMiddleware)
	CommentRouter(router, CommentController, AuthMiddleware)
	router.GET("/", impl.Welcome)
	router.NotFound = http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, http.StatusText(http.StatusNotFound))
	})
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	server := http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	log.Fatal(server.ListenAndServe())
}
