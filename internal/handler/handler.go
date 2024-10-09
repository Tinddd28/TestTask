package handler

import (
	"github.com/Tinddd28/TestTask/internal/handler/middleware"
	"log/slog"
	"net/http"
)

type Handler struct {
	logger *slog.Logger
}

func NewHandler(logger *slog.Logger) *Handler {
	return &Handler{logger: logger}
}

func (h *Handler) InitRoutes() *http.ServeMux {
	mainRouter := http.NewServeMux()
	h.logger.Debug("Initializing routes...")
	songRouter := http.NewServeMux()
	songRouter.HandleFunc("GET /all", h.GetAllSongs)
	songRouter.HandleFunc("GET /{id}", h.GetSong)
	songRouter.HandleFunc("DELETE /{id}", h.DeleteSong)
	songRouter.HandleFunc("PUT /{id}", h.UpdateSong)
	songRouter.HandleFunc("POST /", h.CreateSong)

	mainRouter.Handle("/songs/", http.StripPrefix("/songs", middleware.Logging(songRouter, h.logger)))

	return mainRouter
}
