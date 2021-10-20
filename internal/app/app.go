package app

import (
	"context"

	"github.com/thewolf27/anti-bruteforce/internal/config"
	"github.com/thewolf27/anti-bruteforce/internal/handler"
	"github.com/thewolf27/anti-bruteforce/internal/server"
	"github.com/thewolf27/anti-bruteforce/internal/storage"
	"github.com/thewolf27/anti-bruteforce/pkg/logger"
)

type App struct {
	Config *config.Config
}

func NewApp() *App {
	config := config.NewConfig()

	return &App{
		Config: config,
	}
}

func (app *App) Run() {
	ctx := context.Background()

	logger := logger.NewLogger(app.Config.LoggerConfig.Level)
	storage := storage.NewStorage(app.Config.StorageConfig.Dsn)
	storage.Connect(ctx)

	handler := handler.NewHandler(storage, logger)

	server := server.NewServer(app.Config.ServerConfig.Address, handler)
	server.Serve()
}