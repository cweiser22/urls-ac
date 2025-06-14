package handlers

import "net/http"

type IndexHandler struct{}

func NewIndexHandler() *IndexHandler {
	return &IndexHandler{}
}

func (h *IndexHandler) IndexHandler(w http.ResponseWriter, r *http.Request) {
	// serve static/index.html file
	http.ServeFile(w, r, "static/index.html")

}
