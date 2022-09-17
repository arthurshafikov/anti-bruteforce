package launcher

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/arthurshafikov/anti-bruteforce/internal/app"
	"github.com/arthurshafikov/anti-bruteforce/internal/bucket"
	"github.com/arthurshafikov/anti-bruteforce/internal/config"
	"github.com/arthurshafikov/anti-bruteforce/internal/core"
	grcpapi "github.com/arthurshafikov/anti-bruteforce/internal/server/grpc/api"
	"github.com/arthurshafikov/anti-bruteforce/internal/server/http"
	"github.com/arthurshafikov/anti-bruteforce/internal/storage"
	"github.com/arthurshafikov/anti-bruteforce/pkg/logger"
)

func Run(config *config.Config) {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	logger := logger.NewLogger(config.LoggerConfig.Level)
	storage := storage.NewStorage(config.StorageConfig.Dsn)
	storage.Connect(ctx)

	bucket := bucket.NewLeakyBucket(ctx, core.AuthorizeLimits{
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
