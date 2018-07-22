package api

import "net/http"

func HelloPost(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Write([]byte(`{"message":"hello world"}`))
}
