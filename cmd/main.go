package main

import (
	"context"
	"errors"
	"github.com/Tinddd28/TestTask/internal/config"
	"github.com/Tinddd28/TestTask/internal/handler"
	"github.com/Tinddd28/TestTask/internal/repository"
	"github.com/Tinddd28/TestTask/internal/repository/postgres"
	"github.com/Tinddd28/TestTask/internal/service"
	"github.com/Tinddd28/TestTask/pkg/mylog"
	"github.com/joho/godotenv"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// @title TestTask API
// @version 1
// @description This is a simple API for songs and verses
// @host 0.0.0.0:52352
// @BasePath /
func main() {
	logger := mylog.SetupLogger()

	err := godotenv.Load()
	if err != nil {
		logger.Warn("Error loading .env file: %s", err)
	}
	cfg := config.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
	}

	logger.Debug("Config: ", slog.Any("Postgres config", cfg))

	db, err := postgres.NewPostgres(cfg)
	if err != nil {
		logger.Error("Error connecting to db: %s", err)
		return
	}
	repo := repository.NewRepository(db)
	s := service.NewService(repo)
	h := handler.NewHandler(s, logger)
	logger.Info("Successfully connected to db")
	defer db.Close()
	defer logger.Info("DB connection closed")

	server := &http.Server{
		Addr:    ":52352",
		Handler: h.InitGinRoutes(),
	}
	logger.Info("Starting server...")
	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Error("Error starting server: %s", err)
		}
	}()

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGTERM)

	<-done

	logger.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	time.Sleep(4 * time.Second)
	if err = server.Shutdown(ctx); err != nil {
		logger.Error("Server forced to shutdown: %s", err)
	}

	logger.Info("Server stopped")

	//err = http.ListenAndServe(":52352", h.InitRoutes())
	//if err != nil {
	//	logger.Error("Error starting server: %s", err)
	//}
}
