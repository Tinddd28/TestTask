package handler

import (
	_ "github.com/Tinddd28/TestTask/docs"
	"github.com/Tinddd28/TestTask/internal/handler/middleware"
	"github.com/Tinddd28/TestTask/internal/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log/slog"
)

type Handler struct {
	logger   *slog.Logger
	services *service.Service
}

func NewHandler(serv *service.Service, logger *slog.Logger) *Handler {
	return &Handler{logger: logger, services: serv}
}

//func (h *Handler) InitRoutes() *http.ServeMux {
//	mainRouter := http.NewServeMux()
//	h.logger.Debug("Initializing routes...")
//
//	songRouter := http.NewServeMux()
//	songRouter.HandleFunc("GET /all", h.GetAllSongs)
//	songRouter.HandleFunc("GET /{id}", h.GetSong)
//	songRouter.HandleFunc("DELETE /{id}", h.DeleteSong)
//	songRouter.HandleFunc("PUT /{id}", h.UpdateSong)
//	songRouter.HandleFunc("POST /", h.CreateSong)
//
//	mainRouter.Handle("/songs/", http.StripPrefix("/songs", middleware.Logging(songRouter, h.logger)))
//	mainRouter.Handle("/swagger/", httpSwagger.WrapHandler)
//	return mainRouter
//}

func (h *Handler) InitGinRoutes() *gin.Engine {
	router := gin.Default()
	h.logger.Debug("Initializing routes...")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	songRouter := router.Group("/songs")
	{
		songRouter.GET("/all", h.GetAllSongs)
		songRouter.GET("/:id", h.GetSong)
		songRouter.DELETE("/:id", h.DeleteSong)
		songRouter.PUT("/:id", h.UpdateSong)
		songRouter.POST("/", h.CreateSong)
	}

	router.Use(middleware.LoggingMiddleware(h.logger))

	return router
}
