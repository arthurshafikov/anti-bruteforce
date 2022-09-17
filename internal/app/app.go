package app

import (
	"context"

	"github.com/arthurshafikov/anti-bruteforce/internal/core"
)

type Storage interface {
	AddToWhitelist(string) error
	AddToBlacklist(string) error
	RemoveFromWhitelist(string) error
	RemoveFromBlacklist(string) error
	CheckIfIPInWhitelist(string) (bool, error)
	CheckIfIPInBlacklist(string) (bool, error)
	ResetDatabase() error
}

type Logger interface {
	Warn(string)
	Info(string)
	Error(string)
}

type LeakyBucket interface {
	Add(core.AuthorizeInput) bool
	Leak()
	ResetResetBucketTicker()
}

type App struct {
	Logger      Logger
	Storage     Storage
	LeakyBucket LeakyBucket
}

func NewApp(ctx context.Context, logger Logger, storage Storage, bucket LeakyBucket) *App {
	go bucket.Leak()

	return &App{
		Logger:      logger,
		Storage:     storage,
		LeakyBucket: bucket,
	}
}

func (app *App) Authorize(input core.AuthorizeInput) bool {
	res, err := app.Storage.CheckIfIPInBlacklist(input.IP)
	if err != nil {
		app.Logger.Error(err.Error())
		return false
	}
	if res {
		return false
	}

	res, err = app.Storage.CheckIfIPInWhitelist(input.IP)
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

func (app *App) AddToWhitelist(subnetInput core.SubnetInput) error {
	err := app.Storage.AddToWhitelist(subnetInput.Subnet)
	if err != nil {
		return err
	}

	return nil
}

func (app *App) AddToBlacklist(subnetInput core.SubnetInput) error {
	err := app.Storage.AddToBlacklist(subnetInput.Subnet)
	if err != nil {
		return err
	}

	return nil
}

func (app *App) RemoveFromWhitelist(subnetInput core.SubnetInput) error {
	err := app.Storage.RemoveFromWhitelist(subnetInput.Subnet)
	if err != nil {
		return err
	}

	return nil
}

func (app *App) RemoveFromBlacklist(subnetInput core.SubnetInput) error {
	err := app.Storage.RemoveFromBlacklist(subnetInput.Subnet)
	if err != nil {
		return err
	}

	return nil
}
