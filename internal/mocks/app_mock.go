package mocks

import "github.com/arthurshafikov/anti-bruteforce/internal/core"

type App struct{}

func (app *App) Authorize(core.AuthorizeInput) bool {
	return true
}

func (app *App) ResetBucket() {
}

func (app *App) AddToWhitelist(subnetInput core.SubnetInput) error {
	return nil
}

func (app *App) AddToBlacklist(subnetInput core.SubnetInput) error {
	return nil
}

func (app *App) RemoveFromWhitelist(subnetInput core.SubnetInput) error {
	return nil
}

func (app *App) RemoveFromBlacklist(subnetInput core.SubnetInput) error {
	return nil
}
