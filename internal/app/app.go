package app

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/arthurshafikov/anti-bruteforce/internal/bucket"
	"github.com/arthurshafikov/anti-bruteforce/internal/config"
	"github.com/arthurshafikov/anti-bruteforce/internal/core"
	"github.com/arthurshafikov/anti-bruteforce/internal/repository"
	grcpapi "github.com/arthurshafikov/anti-bruteforce/internal/server/grpc/api"
	"github.com/arthurshafikov/anti-bruteforce/internal/server/http"
	"github.com/arthurshafikov/anti-bruteforce/internal/services"
	"github.com/arthurshafikov/anti-bruteforce/pkg/logger"
	"github.com/arthurshafikov/anti-bruteforce/pkg/postgres"
	"golang.org/x/sync/errgroup"
)

func Run(config *config.Config) {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()
	group, ctx := errgroup.WithContext(ctx)

	logger := logger.NewLogger(config.LoggerConfig.Level)
	db := postgres.NewSqlxDB(ctx, group, config.StorageConfig.Dsn)
	repository := repository.NewRepository(db)

	bucket := bucket.NewLeakyBucket(ctx, core.AuthorizeLimits{
		LimitAttemptsForLogin:    config.NumberOfAttemptsForLogin,
		LimitAttemptsForPassword: config.NumberOfAttemptsForPassword,
		LimitAttemptsForIP:       config.NumberOfAttemptsForIP,
	})
	group.Go(func() error {
		bucket.Leak()

		return nil
	})

	services := services.NewServices(&services.Dependencies{
		Logger:      logger,
		LeakyBucket: bucket,
		Repository:  repository,
	})

	group.Go(func() error {
		grcpapi.RunGrpcServer(ctx, config.GrpcServerConfig.Address, services)

		return nil
	})

	handler := http.NewHandler(services)

	server := http.NewServer(config.ServerConfig.Address, handler)
	server.Serve(ctx)

	if err := group.Wait(); err != nil {
		logger.Error(err.Error())
	}
}
