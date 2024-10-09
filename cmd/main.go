package main

import (
	"github.com/Tinddd28/TestTask/internal/config"
	"github.com/Tinddd28/TestTask/internal/handler"
	"github.com/Tinddd28/TestTask/internal/repository/postgres"
	"github.com/Tinddd28/TestTask/pkg/mylog"
	"github.com/joho/godotenv"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	logger := mylog.SetupLogger()
	logger.Info("Starting server...")

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
	h := handler.NewHandler(logger)
	logger.Info("Successfully connected to db")
	defer db.Close()
	// TODO: Добавить запусук сервера в горутине
	err = http.ListenAndServe(":52352", h.InitRoutes())
	if err != nil {
		logger.Error("Error starting server: %s", err)
	}
}

// Done: Добавить запуск БД (написать докеркомпоз, пока что только для БД)
// Done:  Добавить .env и его считывание
// Done: Файл миграций
// TODO: Написать слой работы с БД, дописать функции для работы с БД в repository/postgres/song.go
// TODO: Написать сервсиный слой
// TODO: Написать тесты с моками

// TODO: Добавить работу с АПИ
