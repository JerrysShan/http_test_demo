package main

import (
	"http_test_demo/handlers"
	"net/http"
)

func main() {
	handlers.Routes()
	http.ListenAndServe(":4000", nil)
}
