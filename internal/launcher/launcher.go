package launcher

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/thewolf27/anti-bruteforce/internal/app"
	"github.com/thewolf27/anti-bruteforce/internal/config"
	"github.com/thewolf27/anti-bruteforce/internal/handler"
	"github.com/thewolf27/anti-bruteforce/internal/server"
	"github.com/thewolf27/anti-bruteforce/internal/storage"
	"github.com/thewolf27/anti-bruteforce/pkg/logger"
)

func Run() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	config := config.NewConfig()
	logger := logger.NewLogger(config.LoggerConfig.Level)
	storage := storage.NewStorage(config.StorageConfig.Dsn)
	storage.Connect(ctx)

	app := app.NewApp(ctx, config, logger, storage)
	handler := handler.NewHandler(app)

	server := server.NewServer(app.Config.ServerConfig.Address, handler)
	server.Serve(ctx)
}
