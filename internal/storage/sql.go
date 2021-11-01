package storage

import (
	"context"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" //nolint:gci
)

const (
	// WhitelistIpsTable is a table name.
	WhitelistIpsTable = "whitelist_ips"

	// BlacklistIpsTable is a table name.
	BlacklistIpsTable = "blacklist_ips"

	// SubnetColumnName is a column in Whitelist/Blacklist Tables.
	SubnetColumnName = "subnet"
)

type Storage struct {
	db  *sqlx.DB
	dsn string
}

func NewStorage(dsn string) *Storage {
	return &Storage{
		dsn: dsn,
	}
}

func (s *Storage) Connect(ctx context.Context) error {
	db, err := sqlx.Connect("postgres", s.dsn)
	go func() {
		<-ctx.Done()
		s.Close()
	}()
	if err != nil {
		panic(err)
	}
	s.db = db
	return nil
}

func (s *Storage) Close() error {
	err := s.db.Close()
	return err
}

func (s *Storage) AddToWhitelist(subnet string) error {
	_, err := s.db.Exec(fmt.Sprintf("INSERT INTO %s (%s) VALUES ($1);", WhitelistIpsTable, SubnetColumnName), subnet)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) RemoveFromWhitelist(subnet string) error {
	_, err := s.db.Exec(fmt.Sprintf("DELETE FROM %s WHERE %s = $1;", WhitelistIpsTable, SubnetColumnName), subnet)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) AddToBlacklist(subnet string) error {
	_, err := s.db.Exec(fmt.Sprintf("INSERT INTO %s (%s) VALUES ($1);", BlacklistIpsTable, SubnetColumnName), subnet)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) RemoveFromBlacklist(subnet string) error {
	_, err := s.db.Exec(fmt.Sprintf("DELETE FROM %s WHERE %s = $1;", BlacklistIpsTable, SubnetColumnName), subnet)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) CheckIfIPInWhitelist(ip string) (bool, error) {
	res, err := s.db.Exec(fmt.Sprintf("SELECT 1 FROM %s WHERE %s >>= $1;", WhitelistIpsTable, SubnetColumnName), ip)
	if err != nil {
		return false, err
	}
	if n, _ := res.RowsAffected(); n > 0 {
		return true, nil
	}

	return false, nil
}

func (s *Storage) CheckIfIPInBlacklist(ip string) (bool, error) {
	res, err := s.db.Exec(fmt.Sprintf("SELECT 1 FROM %s WHERE %s >>= $1;", BlacklistIpsTable, SubnetColumnName), ip)
	if err != nil {
		return false, err
	}
	if n, _ := res.RowsAffected(); n > 0 {
		return true, nil
	}

	return false, nil
}

func (s *Storage) ResetDatabase() error {
	tables := []string{
		WhitelistIpsTable,
		BlacklistIpsTable,
	}
	_, err := s.db.Exec(fmt.Sprintf(`TRUNCATE TABLE %s`, strings.Join(tables, ", ")))

	return err
}
