package services

import "github.com/arthurshafikov/anti-bruteforce/internal/core"

type AuthService struct {
	logger           Logger
	leakyBucket      LeakyBucket
	blacklistService Blacklist
	whitelistService Whitelist
}

func NewAuthService(
	logger Logger,
	leakyBucket LeakyBucket,
	blacklistService Blacklist,
	whitelistService Whitelist,
) *AuthService {
	return &AuthService{
		logger:           logger,
		leakyBucket:      leakyBucket,
		blacklistService: blacklistService,
		whitelistService: whitelistService,
	}
}

func (as *AuthService) Authorize(input core.AuthorizeInput) bool {
	res, err := as.blacklistService.CheckIfIPInBlacklist(input.IP)
	if err != nil {
		as.logger.Error(err)
		return false
	}
	if res {
		return false
	}

	res, err = as.whitelistService.CheckIfIPInWhitelist(input.IP)
	if err != nil {
		as.logger.Error(err)
		return false
	}
	if res {
		return true
	}

	return as.leakyBucket.Add(input)
}
