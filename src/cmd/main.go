package main

import (
	"api"
	"net/http"
)

func main() {
	http.HandleFunc("/greeting", api.HelloPost)
	http.ListenAndServe(":9999", nil)
}
