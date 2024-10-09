package handler

import "net/http"

func (h *Handler) GetAllSongs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("All songs"))
}

func (h *Handler) GetSong(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) DeleteSong(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) UpdateSong(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) CreateSong(w http.ResponseWriter, r *http.Request) {

}
