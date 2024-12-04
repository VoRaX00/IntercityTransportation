package app

import (
	"context"
	"fmt"
	"kursachDB/internal/app/server"
	"kursachDB/internal/config"
	"kursachDB/internal/handler"
	"kursachDB/internal/storage/postgres"
	"log/slog"
)

type App struct {
	Server  *server.Server
	Storage *postgres.Storage
}

func New(log *slog.Logger, storagePath string, serverConfig config.ServerConfig) *App {
	storage, err := connectDB(storagePath)
	if err != nil {
		panic(err)
	}

	service := handler.NewService()
	handlers := handler.NewHandler(service)
	serv := server.New(log, serverConfig.Port, serverConfig.Timeout, handlers.InitRoutes())

	return &App{
		Storage: storage,
		Server:  serv,
	}
}

func (a *App) Stop(ctx context.Context) {
	err := a.Storage.Close()
	if err != nil {
		panic(err)
	}
	a.Server.Shutdown(ctx)
}

func connectDB(storagePath string) (*postgres.Storage, error) {
	db, err := postgres.New(storagePath)
	if err != nil {
		return nil, fmt.Errorf("could not connect to database: %w", err)
	}

	return db, nil
}
