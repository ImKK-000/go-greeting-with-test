package api

import (
	"encoding/json"
	"net/http"
)

type ResponseMessage struct {
	Message string `json:"message"`
}

func HelloPost(responseWriter http.ResponseWriter, request *http.Request) {
	resposeMessage := ResponseMessage{
		Message: "hello world",
	}
	encoder := json.NewEncoder(responseWriter)
	encoder.Encode(resposeMessage)
}
