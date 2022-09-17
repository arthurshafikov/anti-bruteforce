package services

import (
	"github.com/arthurshafikov/anti-bruteforce/internal/core"
	"github.com/arthurshafikov/anti-bruteforce/internal/repository"
)

type BlacklistService struct {
	repo repository.Blacklist
}

func NewBlacklistService(repo repository.Blacklist) *BlacklistService {
	return &BlacklistService{
		repo: repo,
	}
}

func (bs *BlacklistService) AddToBlacklist(subnetInput core.SubnetInput) error {
	return bs.repo.AddToBlacklist(subnetInput.Subnet)
}

func (bs *BlacklistService) CheckIfIPInBlacklist(ip string) (bool, error) {
	return bs.repo.CheckIfIPInBlacklist(ip)
}

func (bs *BlacklistService) RemoveFromBlacklist(subnetInput core.SubnetInput) error {
	return bs.repo.RemoveFromBlacklist(subnetInput.Subnet)
}
