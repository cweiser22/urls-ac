package handlers

import (
	"encoding/json"
	"github.com/cweiser22/urls-ac/internal/service"
	"github.com/cweiser22/urls-ac/internal/utils"
	"github.com/go-chi/chi/v5"
	"github.com/spf13/viper"

	"net/http"
)

type URLHandler struct {
	ShortCodeService *service.ShortenService
	RedirectProtocol string
}

func NewURLHandler(service *service.ShortenService) *URLHandler {
	environment := viper.GetString("environment")
	redirectProtocol := "http://"
	if environment == "production" {
		redirectProtocol = "https://"
	}
	return &URLHandler{
		ShortCodeService: service,
		RedirectProtocol: redirectProtocol,
	}
}

func (h *URLHandler) RedirectFromMapping(w http.ResponseWriter, r *http.Request) {
	// Extract the short URL mapping from the request
	shortCode := chi.URLParam(r, "shortCode")

	// Query the database for the original URL
	longURL, err := h.ShortCodeService.GetLongURL(shortCode)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	if longURL == "" {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	// Redirect to the original URL
	http.Redirect(w, r, longURL, http.StatusFound)
}

type CreateShortURLRequest struct {
	LongURL string `json:"longUrl"`
}

type CreateShortURLResponse struct {
	ShortURL string `json:"shortUrl"`
	LongURL  string `json:"longUrl"`
}

func (h *URLHandler) CreateShortURL(w http.ResponseWriter, r *http.Request) {

	// Parse the request body to get the long URL
	var requestBody CreateShortURLRequest
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	longURL, err := utils.ValidateAndFixURL(requestBody.LongURL)
	if err != nil {
		http.Error(w, "Invalid URL: "+requestBody.LongURL, http.StatusBadRequest)
		return
	}

	// Create a new short URL mapping
	mapping, err := h.ShortCodeService.CreateURLMapping(longURL)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	host := viper.GetString("host")
	if host == "" {
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Invalid URL",
		})
		http.Error(w, "Internal Server Error: Host not configured", http.StatusInternalServerError)
		return
	}

	responseBody := CreateShortURLResponse{
		ShortURL: h.RedirectProtocol + host + "/" + mapping.ShortCode,
		LongURL:  mapping.LongURL,
	}

	// Respond with the short URL
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(responseBody)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
