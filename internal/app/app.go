package app

import (
	"context"
	"os/signal"
	"syscall"

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
	return &App{
		Config: config.NewConfig(),
	}
}

func (app *App) Run() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	logger := logger.NewLogger(app.Config.LoggerConfig.Level)
	storage := storage.NewStorage(app.Config.StorageConfig.Dsn)
	storage.Connect(ctx)

	handler := handler.NewHandler(ctx, storage, logger, app.Config.AppConfig)

	server := server.NewServer(app.Config.ServerConfig.Address, handler)
	server.Serve(ctx)
}
