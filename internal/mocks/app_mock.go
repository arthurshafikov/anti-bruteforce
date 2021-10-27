package mocks

import "github.com/thewolf27/anti-bruteforce/internal/models"

type App struct{}

func (app *App) Authorize(models.AuthorizeInput) bool {
	return true
}

func (app *App) ResetBucket() {
}

func (app *App) AddToWhiteList(subnetInput models.SubnetInput) error {
	return nil
}

func (app *App) AddToBlackList(subnetInput models.SubnetInput) error {
	return nil
}

func (app *App) RemoveFromWhiteList(subnetInput models.SubnetInput) error {
	return nil
}

func (app *App) RemoveFromBlackList(subnetInput models.SubnetInput) error {
	return nil
}
