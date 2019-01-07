package main

import (
	"net/http"

	"leeif.me/Go-test/go-web/controller"
)

func main() {
	controller.Startup()
	http.ListenAndServe(":8888", nil)
}
