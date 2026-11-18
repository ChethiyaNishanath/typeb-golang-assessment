package handlers

import (
	"net/http"
	"strings"
)

type Greeter interface {
	Greet(name string) (string, error)
}

type HelloWorldHandler struct {
	greeter Greeter
}

func NewHelloWorldHandler(greeter Greeter) *HelloWorldHandler {
	return &HelloWorldHandler{greeter: greeter}
}

func (h *HelloWorldHandler) Greet(w http.ResponseWriter, r *http.Request) {
	name := strings.TrimSpace(r.URL.Query().Get("name"))

	message, err := h.greeter.Greet(name)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid Input"})
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{"message": message})
}
