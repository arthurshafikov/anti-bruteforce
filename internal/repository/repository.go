package repository

import (
	"github.com/arthurshafikov/anti-bruteforce/internal/repository/repositories"
	"github.com/jmoiron/sqlx"
)

type Blacklist interface {
	AddToBlacklist(subnet string) error
	RemoveFromBlacklist(subnet string) error
	CheckIfIPInBlacklist(ip string) (bool, error)
}

type Whitelist interface {
	AddToWhitelist(subnet string) error
	RemoveFromWhitelist(subnet string) error
	CheckIfIPInWhitelist(ip string) (bool, error)
}

type Repository struct {
	Blacklist
	Whitelist
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Blacklist: repositories.NewBlacklist(db),
		Whitelist: repositories.NewWhitelist(db),
	}
}
