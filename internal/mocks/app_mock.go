package mocks

import "github.com/arthurshafikov/anti-bruteforce/internal/models"

type App struct{}

func (app *App) Authorize(models.AuthorizeInput) bool {
	return true
}

func (app *App) ResetBucket() {
}

func (app *App) AddToWhitelist(subnetInput models.SubnetInput) error {
	return nil
}

func (app *App) AddToBlacklist(subnetInput models.SubnetInput) error {
	return nil
}

func (app *App) RemoveFromWhitelist(subnetInput models.SubnetInput) error {
	return nil
}

func (app *App) RemoveFromBlacklist(subnetInput models.SubnetInput) error {
	return nil
}
