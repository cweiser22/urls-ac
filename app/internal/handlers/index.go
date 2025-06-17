package handlers

import (
	"fmt"
	"net/http"
)

type IndexHandler struct{}

func NewIndexHandler() *IndexHandler {
	return &IndexHandler{}
}

func (h *IndexHandler) AppHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("AppHandler called")
	http.ServeFile(w, r, "web/static/index.html")
}
