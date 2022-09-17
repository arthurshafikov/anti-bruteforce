package repositories

import (
	"fmt"

	"github.com/arthurshafikov/anti-bruteforce/internal/core"
	"github.com/jmoiron/sqlx"
)

type Whitelist struct {
	db *sqlx.DB
}

func NewWhitelist(db *sqlx.DB) *Whitelist {
	return &Whitelist{
		db: db,
	}
}

func (w *Whitelist) AddToWhitelist(subnet string) error {
	_, err := w.db.Exec(fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES ($1);", core.WhitelistIpsTable, core.SubnetColumnName), subnet,
	)
	if err != nil {
		return err
	}

	return nil
}

func (w *Whitelist) RemoveFromWhitelist(subnet string) error {
	_, err := w.db.Exec(fmt.Sprintf(
		"DELETE FROM %s WHERE %s = $1;", core.WhitelistIpsTable, core.SubnetColumnName), subnet,
	)
	if err != nil {
		return err
	}

	return nil
}

func (w *Whitelist) CheckIfIPInWhitelist(ip string) (bool, error) {
	res, err := w.db.Exec(fmt.Sprintf(
		"SELECT 1 FROM %s WHERE %s >>= $1;", core.WhitelistIpsTable, core.SubnetColumnName), ip,
	)
	if err != nil {
		return false, err
	}
	n, _ := res.RowsAffected()

	return n > 0, nil
}
