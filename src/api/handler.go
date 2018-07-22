package api

import (
	"encoding/json"
	"greeting"
	"net/http"
)

type ResponseMessage struct {
	Message string `json:"message"`
}

func HelloPost(responseWriter http.ResponseWriter, request *http.Request) {
	greeting := greeting.Greeting{
		Prefix:             "hello",
		DefaultEmptyString: "thailand",
	}
	resposeMessage := ResponseMessage{
		Message: greeting.Message("world"),
	}
	encoder := json.NewEncoder(responseWriter)
	encoder.Encode(resposeMessage)
}
