package api

import (
	"encoding/json"
	"greeting"
	"net/http"
)

type RequestMessage struct {
	Message string `json:"message"`
}

type ResponseMessage struct {
	Message string `json:"message"`
}

func HelloPost(responseWriter http.ResponseWriter, request *http.Request) {
	var requestMessage RequestMessage
	body := json.NewDecoder(request.Body)
	body.Decode(&requestMessage)
	greeting := greeting.Greeting{
		Prefix:             "hello",
		DefaultEmptyString: "thailand",
	}
	resposeMessage := ResponseMessage{
		Message: greeting.Message(requestMessage.Message),
	}
	encoder := json.NewEncoder(responseWriter)
	encoder.Encode(resposeMessage)
}
