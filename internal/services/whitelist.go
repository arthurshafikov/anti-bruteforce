package services

import (
	"github.com/arthurshafikov/anti-bruteforce/internal/core"
	"github.com/arthurshafikov/anti-bruteforce/internal/repository"
)

type WhitelistService struct {
	repo repository.Whitelist
}

func NewWhitelistService(repo repository.Whitelist) *WhitelistService {
	return &WhitelistService{
		repo: repo,
	}
}

func (ws *WhitelistService) AddToWhitelist(subnetInput core.SubnetInput) error {
	return ws.repo.AddToWhitelist(subnetInput.Subnet)
}

func (ws *WhitelistService) CheckIfIPInWhitelist(ip string) (bool, error) {
	return ws.repo.CheckIfIPInWhitelist(ip)
}

func (ws *WhitelistService) RemoveFromWhitelist(subnetInput core.SubnetInput) error {
	return ws.repo.RemoveFromWhitelist(subnetInput.Subnet)
}
