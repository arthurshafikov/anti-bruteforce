package app

import (
	"context"

	"github.com/thewolf27/anti-bruteforce/internal/models"
)

type Storage interface {
	AddToWhiteList(string) error
	AddToBlackList(string) error
	RemoveFromWhiteList(string) error
	RemoveFromBlackList(string) error
	CheckIfIPInWhiteList(string) (bool, error)
	CheckIfIPInBlackList(string) (bool, error)
	ResetDatabase() error
}

type Logger interface {
	Warn(string)
	Info(string)
	Error(string)
}

type LeakyBucket interface {
	Add(models.AuthorizeInput) bool
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

func (app *App) Authorize(input models.AuthorizeInput) bool {
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

func (app *App) AddToWhiteList(subnetInput models.SubnetInput) error {
	err := app.Storage.AddToWhiteList(subnetInput.Subnet)
	if err != nil {
		return err
	}

	return nil
}

func (app *App) AddToBlackList(subnetInput models.SubnetInput) error {
	err := app.Storage.AddToBlackList(subnetInput.Subnet)
	if err != nil {
		return err
	}

	return nil
}

func (app *App) RemoveFromWhiteList(subnetInput models.SubnetInput) error {
	err := app.Storage.RemoveFromWhiteList(subnetInput.Subnet)
	if err != nil {
		return err
	}

	return nil
}

func (app *App) RemoveFromBlackList(subnetInput models.SubnetInput) error {
	err := app.Storage.RemoveFromBlackList(subnetInput.Subnet)
	if err != nil {
		return err
	}

	return nil
}
