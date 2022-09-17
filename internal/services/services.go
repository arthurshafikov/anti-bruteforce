package services

import (
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

type Auth interface {
	Authorize(input core.AuthorizeInput) bool
}

type Blacklist interface {
	AddToBlacklist(subnetInput core.SubnetInput) error
	CheckIfIPInBlacklist(ip string) (bool, error)
	RemoveFromBlacklist(subnetInput core.SubnetInput) error
}

type Whitelist interface {
	AddToWhitelist(subnetInput core.SubnetInput) error
	CheckIfIPInWhitelist(ip string) (bool, error)
	RemoveFromWhitelist(subnetInput core.SubnetInput) error
}

type Bucket interface {
	ResetBucket()
}

type Services struct {
	Auth
	Blacklist
	Whitelist
	Bucket
}

type Dependencies struct {
	Logger      Logger
	LeakyBucket LeakyBucket
	Repository  *repository.Repository
}

func NewServices(deps *Dependencies) *Services {
	blacklist := NewBlacklistService(deps.Repository.Blacklist)
	whitelist := NewWhitelistService(deps.Repository.Whitelist)
	auth := NewAuthService(deps.Logger, deps.LeakyBucket, blacklist, whitelist)
	bucket := NewBucketService(deps.LeakyBucket)

	return &Services{
		Auth:      auth,
		Blacklist: blacklist,
		Whitelist: whitelist,
		Bucket:    bucket,
	}
}
