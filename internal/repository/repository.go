package repository

import (
	"github.com/arthurshafikov/anti-bruteforce/internal/repository/repositories"
	"github.com/jmoiron/sqlx"
)

type Blacklist interface {
	AddToBlacklist(string) error
	RemoveFromBlacklist(string) error
	CheckIfIPInBlacklist(string) (bool, error)
}

type Whitelist interface {
	AddToWhitelist(string) error
	RemoveFromWhitelist(string) error
	CheckIfIPInWhitelist(string) (bool, error)
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
