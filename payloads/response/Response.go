package response

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func ResponseError(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)

	res := Response{
		Status:  status,
		Message: err.Error(),
		Data:    nil,
	}

	err = json.NewEncoder(w).Encode(res)

	if err != nil {
		log.Fatal(err)
	}
}

func ResponseMessage(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Add("content-type", "application/json")
	res := Response{
		Status:  status,
		Data:    data,
		Message: http.StatusText(status),
	}

	err := json.NewEncoder(w).Encode(res)

	if err != nil {
		log.Fatal(err)
	}
}
