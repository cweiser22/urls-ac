package handlers

import (
	"encoding/json"
	"github.com/cweiser22/urls-ac/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/spf13/viper"

	"log/slog"
	"net/http"
)

type URLHandler struct {
	ShortCodeService *service.ShortCodeService
	RedirectProtocol string
}

func NewURLHandler(service *service.ShortCodeService) *URLHandler {
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
}

func (h *URLHandler) CreateShortURL(w http.ResponseWriter, r *http.Request) {

	// Parse the request body to get the long URL
	var requestBody CreateShortURLRequest
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// if long url does not start with http or https, prepend http://
	if !((len(requestBody.LongURL) > 4 && requestBody.LongURL[:4] == "http") || (len(requestBody.LongURL) > 5 && requestBody.LongURL[:5] == "https")) {
		requestBody.LongURL = "http://" + requestBody.LongURL
	}

	// Create a new short URL mapping
	mapping, err := h.ShortCodeService.GetOrCreateMapping(requestBody.LongURL)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	host := viper.GetString("host")
	if host == "" {
		http.Error(w, "Internal Server Error: Host not configured", http.StatusInternalServerError)
		return
	}

	slog.Info(h.RedirectProtocol)

	responseBody := CreateShortURLResponse{
		ShortURL: h.RedirectProtocol + host + "/" + mapping.ShortCode,
	}

	// Respond with the short URL
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(responseBody)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
