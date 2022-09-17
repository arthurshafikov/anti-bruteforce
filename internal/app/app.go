package app

import (
	"context"

	"github.com/arthurshafikov/anti-bruteforce/internal/core"
	"github.com/arthurshafikov/anti-bruteforce/internal/repository"
)

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
	Repository  *repository.Repository
	LeakyBucket LeakyBucket
}

func NewApp(ctx context.Context, logger Logger, repository *repository.Repository, bucket LeakyBucket) *App {
	go bucket.Leak()

	return &App{
		Logger:      logger,
		Repository:  repository,
		LeakyBucket: bucket,
	}
}

func (app *App) Authorize(input core.AuthorizeInput) bool {
	res, err := app.Repository.Blacklist.CheckIfIPInBlacklist(input.IP)
	if err != nil {
		app.Logger.Error(err.Error())
		return false
	}
	if res {
		return false
	}

	res, err = app.Repository.Whitelist.CheckIfIPInWhitelist(input.IP)
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
	err := app.Repository.Whitelist.AddToWhitelist(subnetInput.Subnet)
	if err != nil {
		return err
	}

	return nil
}

func (app *App) AddToBlacklist(subnetInput core.SubnetInput) error {
	err := app.Repository.Blacklist.AddToBlacklist(subnetInput.Subnet)
	if err != nil {
		return err
	}

	return nil
}

func (app *App) RemoveFromWhitelist(subnetInput core.SubnetInput) error {
	err := app.Repository.Whitelist.RemoveFromWhitelist(subnetInput.Subnet)
	if err != nil {
		return err
	}

	return nil
}

func (app *App) RemoveFromBlacklist(subnetInput core.SubnetInput) error {
	err := app.Repository.Blacklist.RemoveFromBlacklist(subnetInput.Subnet)
	if err != nil {
		return err
	}

	return nil
}
