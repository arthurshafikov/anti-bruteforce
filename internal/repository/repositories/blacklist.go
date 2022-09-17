package repositories

import (
	"fmt"

	"github.com/arthurshafikov/anti-bruteforce/internal/core"
	"github.com/jmoiron/sqlx"
)

type Blacklist struct {
	db *sqlx.DB
}

func NewBlacklist(db *sqlx.DB) *Blacklist {
	return &Blacklist{
		db: db,
	}
}

func (b *Blacklist) AddToBlacklist(subnet string) error {
	_, err := b.db.Exec(
		fmt.Sprintf("INSERT INTO %s (%s) VALUES ($1);", core.BlacklistIpsTable, core.SubnetColumnName), subnet,
	)
	if err != nil {
		return err
	}

	return nil
}

func (b *Blacklist) RemoveFromBlacklist(subnet string) error {
	_, err := b.db.Exec(
		fmt.Sprintf("DELETE FROM %s WHERE %s = $1;", core.BlacklistIpsTable, core.SubnetColumnName), subnet,
	)
	if err != nil {
		return err
	}

	return nil
}

func (b *Blacklist) CheckIfIPInBlacklist(ip string) (bool, error) {
	res, err := b.db.Exec(
		fmt.Sprintf("SELECT 1 FROM %s WHERE %s >>= $1;", core.BlacklistIpsTable, core.SubnetColumnName), ip,
	)
	if err != nil {
		return false, err
	}
	if n, _ := res.RowsAffected(); n > 0 {
		return true, nil
	}

	return false, nil
}
