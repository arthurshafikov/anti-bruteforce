package app

import (
	"context"
	"os/signal"
	"syscall"
	"time"

	"github.com/arthurshafikov/anti-bruteforce/internal/bucket"
	"github.com/arthurshafikov/anti-bruteforce/internal/config"
	"github.com/arthurshafikov/anti-bruteforce/internal/core"
	"github.com/arthurshafikov/anti-bruteforce/internal/repository"
	"github.com/arthurshafikov/anti-bruteforce/internal/services"
	grcpapi "github.com/arthurshafikov/anti-bruteforce/internal/transport/grpc/api"
	"github.com/arthurshafikov/anti-bruteforce/internal/transport/http"
	"github.com/arthurshafikov/anti-bruteforce/pkg/logger"
	"github.com/arthurshafikov/anti-bruteforce/pkg/postgres"
	"golang.org/x/sync/errgroup"
)

func Run(config *config.Config) {
	logger := logger.NewLogger(config.LoggerConfig.Level)

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	group, ctx := errgroup.WithContext(ctx)
	defer func() {
		if err := group.Wait(); err != nil {
			logger.Error(err.Error())
		}
	}()

	db := postgres.NewSqlxDB(ctx, group, config.DatabaseConfig.DSN)
	repository := repository.NewRepository(db)

	resetBucketsTicker := time.NewTicker(time.Second * 60) // todo config
	group.Go(func() error {
		<-ctx.Done()
		resetBucketsTicker.Stop()

		return nil
	})
	bucket := bucket.NewLeakyBucket(resetBucketsTicker, core.AuthorizeLimits{
		LimitAttemptsForLogin:    config.NumberOfAttemptsForLogin,
		LimitAttemptsForPassword: config.NumberOfAttemptsForPassword,
		LimitAttemptsForIP:       config.NumberOfAttemptsForIP,
	})
	defer bucket.ResetResetBucketTicker()
	group.Go(func() error {
		bucket.Leak(ctx)

		return nil
	})

	services := services.NewServices(&services.Dependencies{
		Logger:      logger,
		LeakyBucket: bucket,
		Repository:  repository,
	})

	grcpapi.RunGrpcServer(ctx, group, config.GrpcServerConfig.Address, services)

	handler := http.NewHandler(services)

	http.NewServer(handler).Serve(ctx, group, config.ServerConfig.Address)
}
