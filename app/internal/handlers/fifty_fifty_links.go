package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/cweiser22/urls-ac/internal/service"
	"github.com/cweiser22/urls-ac/internal/utils"
	"github.com/go-chi/chi/v5"
	"github.com/spf13/viper"
	"net/http"
)

type FiftyFiftyHandler struct {
	fiftyFiftyService *service.FiftyFiftyLinkService
	shortCodeService  *service.ShortCodeService
	redirectProtocol  string
}

func NewFiftyFiftyHandler(service *service.FiftyFiftyLinkService, shortCodeService *service.ShortCodeService) *FiftyFiftyHandler {
	environment := viper.GetString("environment")
	redirectProtocol := "http://"
	if environment == "production" {
		redirectProtocol = "https://"
	}
	return &FiftyFiftyHandler{fiftyFiftyService: service, shortCodeService: shortCodeService, redirectProtocol: redirectProtocol}
}

type CreateFiftyFiftyLinkResponse struct {
	ShortURL string `json:"shortUrl"`
	URLa     string `json:"urlA"`
	URLb     string `json:"urlB"`
}

type createRequest struct {
	Probability float64 `json:"probability"`
	URLa        string  `json:"urlA"`
	URLb        string  `json:"urlB"`
}

// Create handles POST /create
func (h *FiftyFiftyHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req createRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if req.Probability < 0 || req.Probability > 1 {
		http.Error(w, "Probability must be between 0 and 1", http.StatusBadRequest)
		return
	}

	shortCode := h.shortCodeService.GenerateShortcode(fmt.Sprintf("%f-%s-%s", req.Probability, req.URLa, req.URLb), 6)

	URLa, err := utils.ValidateAndFixURL(req.URLa)
	if err != nil {
		http.Error(w, "Invalid URL A: "+err.Error(), http.StatusBadRequest)
		return
	}

	URLb, err := utils.ValidateAndFixURL(req.URLb)
	if err != nil {
		http.Error(w, "Invalid URL B: "+err.Error(), http.StatusBadRequest)
		return
	}

	link, err := h.fiftyFiftyService.Create(req.Probability, URLa, URLb, shortCode)
	if err != nil {
		http.Error(w, "Failed to create link: "+err.Error(), http.StatusInternalServerError)
		return
	}

	host := viper.GetString("host")
	if host == "" {
		http.Error(w, "Internal Server Error: Host not configured", http.StatusInternalServerError)
		return
	}

	responseBody := CreateFiftyFiftyLinkResponse{
		ShortURL: h.redirectProtocol + host + "/ff/" + link.ShortCode,
		URLa:     link.URLa,
		URLb:     link.URLb,
	}

	// Respond with the short URL
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(responseBody)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

}

// Redirect handles GET /{shortCode}
func (h *FiftyFiftyHandler) Redirect(w http.ResponseWriter, r *http.Request) {
	shortCode := chi.URLParam(r, "shortCode")
	link, err := h.fiftyFiftyService.GetByShortCode(shortCode)
	if err != nil {
		http.Error(w, "Link not found", http.StatusNotFound)
		return
	}

	target := h.fiftyFiftyService.GetLink(link)
	http.Redirect(w, r, target, http.StatusFound)
}
