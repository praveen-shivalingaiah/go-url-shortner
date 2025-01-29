package http

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/praveen-shivalingaiah/go-url-shortner/app"
)

type Handler struct {
	shortner *app.ShortnerService
}

type ShortnerRequest struct {
	URL string `json:url`
}

type ShortnerResponse struct {
	ShortID string `json:"short_id"`
}

func NewHandlerService(shortner *app.ShortnerService) *Handler {
	return &Handler{shortner: shortner}
}

func (h *Handler) RegisterRoutes(r *mux.Router) {

	r.HandleFunc("/shorten", h.ShortnerURL)
	r.HandleFunc("/{shortID}", h.ResolveURL)
}

func (h *Handler) ShortnerURL(w http.ResponseWriter, r *http.Request) {
	var req ShortnerRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	shortID, err := h.shortner.ShortenURL(req.URL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	response := ShortnerResponse{ShortID: shortID}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) ResolveURL(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortID := vars["shortID"]

	url, err := h.shortner.ResolveURL(shortID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, url, http.StatusOK)
}
