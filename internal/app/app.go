package app

import (
	"context"

	"github.com/thewolf27/anti-bruteforce/internal/bucket"
	"github.com/thewolf27/anti-bruteforce/internal/config"
)

type Storage interface {
	AddToWhiteList(string) error
	AddToBlackList(string) error
	RemoveFromWhiteList(string) error
	RemoveFromBlackList(string) error
	CheckIfIPInWhiteList(string) (bool, error)
	CheckIfIPInBlackList(string) (bool, error)
}

type Logger interface {
	Warn(string)
	Info(string)
	Error(string)
}

type App struct {
	Config      *config.Config
	Logger      Logger
	Storage     Storage
	LeakyBucket *bucket.LeakyBucket
}

func NewApp(ctx context.Context, config *config.Config, logger Logger, storage Storage) *App {
	leakyBucket := bucket.NewLeakyBucket(ctx, bucket.AuthorizeLimits{
		LimitAttemptsForLogin:    config.AppConfig.NumberOfAttemptsForLogin,
		LimitAttemptsForPassword: config.AppConfig.NumberOfAttemptsForPassword,
		LimitAttemptsForIP:       config.AppConfig.NumberOfAttemptsForIP,
	})
	go leakyBucket.Leak()

	return &App{
		Config:      config,
		Logger:      logger,
		Storage:     storage,
		LeakyBucket: leakyBucket,
	}
}

func (app *App) Authorize(input bucket.AuthorizeInput) bool {
	res, err := app.Storage.CheckIfIPInBlackList(input.IP)
	if err != nil {
		app.Logger.Error(err.Error())
		return false
	}
	if res {
		return false
	}

	res, err = app.Storage.CheckIfIPInWhiteList(input.IP)
	if err != nil {
		app.Logger.Error(err.Error())
		return false
	}
	if res {
		return true
	}

	return app.LeakyBucket.Add(input)
}

func (app *App) ResetBucket() {
	app.LeakyBucket.ResetResetBucketTicker()
}

func (app *App) AddToWhiteList(subnet string) error {
	err := app.Storage.AddToWhiteList(subnet)
	if err != nil {
		return err
	}

	return nil
}

func (app *App) AddToBlackList(subnet string) error {
	err := app.Storage.AddToBlackList(subnet)
	if err != nil {
		return err
	}

	return nil
}

func (app *App) RemoveFromWhiteList(subnet string) error {
	err := app.Storage.RemoveFromWhiteList(subnet)
	if err != nil {
		return err
	}

	return nil
}

func (app *App) RemoveFromBlackList(subnet string) error {
	err := app.Storage.RemoveFromBlackList(subnet)
	if err != nil {
		return err
	}

	return nil
}

// func (app *App) Run() {
// 	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
// 	defer cancel()

// 	logger := logger.NewLogger(app.Config.LoggerConfig.Level)
// 	storage := storage.NewStorage(app.Config.StorageConfig.Dsn)
// 	storage.Connect(ctx)

// 	handler := handler.NewHandler(ctx, storage, logger, app.Config.AppConfig)

// 	server := server.NewServer(app.Config.ServerConfig.Address, handler)
// 	server.Serve(ctx)
// }
