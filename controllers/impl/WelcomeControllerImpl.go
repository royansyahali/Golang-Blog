package impl

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Welcome(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintln(w, "Welcome to Restful API Blog")
}
