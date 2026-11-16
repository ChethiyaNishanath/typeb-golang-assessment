package server

import (
	"net/http"

	"github.com/ChethiyaNishanath/typeb-golang-assessment/internal/handlers"
)

func RegisterRoutes(hwh *handlers.HelloWorldHandler) {
	http.HandleFunc("GET /hello-world", hwh.Greet)
}
