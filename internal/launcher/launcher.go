package launcher

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/thewolf27/anti-bruteforce/internal/app"
	"github.com/thewolf27/anti-bruteforce/internal/bucket"
	"github.com/thewolf27/anti-bruteforce/internal/config"
	"github.com/thewolf27/anti-bruteforce/internal/models"
	grcpapi "github.com/thewolf27/anti-bruteforce/internal/server/grpc/api"
	"github.com/thewolf27/anti-bruteforce/internal/server/http"
	"github.com/thewolf27/anti-bruteforce/internal/storage"
	"github.com/thewolf27/anti-bruteforce/pkg/logger"
)

func Run(config *config.Config) {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	logger := logger.NewLogger(config.LoggerConfig.Level)
	storage := storage.NewStorage(config.StorageConfig.Dsn)
	storage.Connect(ctx)

	bucket := bucket.NewLeakyBucket(ctx, models.AuthorizeLimits{
		LimitAttemptsForLogin:    config.NumberOfAttemptsForLogin,
		LimitAttemptsForPassword: config.NumberOfAttemptsForPassword,
		LimitAttemptsForIP:       config.NumberOfAttemptsForIP,
	})

	app := app.NewApp(ctx, logger, storage, bucket)

	go grcpapi.RunGrpcServer(ctx, config.GrpcServerConfig.Address, app)

	handler := http.NewHandler(app)

	server := http.NewServer(config.ServerConfig.Address, handler)
	server.Serve(ctx)
}
