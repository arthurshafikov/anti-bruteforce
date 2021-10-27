package storage

import (
	"context"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" //nolint:gci
)

const (
	// WhiteListIpsTable is a table name.
	WhiteListIpsTable = "whitelist_ips"

	// BlackListIpsTable is a table name.
	BlackListIpsTable = "blacklist_ips"

	// SubnetColumnName is a column in WhiteList/BlackList Tables.
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

func (s *Storage) AddToWhiteList(subnet string) error {
	_, err := s.db.Exec(fmt.Sprintf("INSERT INTO %s (%s) VALUES ($1);", WhiteListIpsTable, SubnetColumnName), subnet)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) RemoveFromWhiteList(subnet string) error {
	_, err := s.db.Exec(fmt.Sprintf("DELETE FROM %s WHERE %s = $1;", WhiteListIpsTable, SubnetColumnName), subnet)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) AddToBlackList(subnet string) error {
	_, err := s.db.Exec(fmt.Sprintf("INSERT INTO %s (%s) VALUES ($1);", BlackListIpsTable, SubnetColumnName), subnet)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) RemoveFromBlackList(subnet string) error {
	_, err := s.db.Exec(fmt.Sprintf("DELETE FROM %s WHERE %s = $1;", BlackListIpsTable, SubnetColumnName), subnet)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) CheckIfIPInWhiteList(ip string) (bool, error) {
	res, err := s.db.Exec(fmt.Sprintf("SELECT 1 FROM %s WHERE %s >>= $1;", WhiteListIpsTable, SubnetColumnName), ip)
	if err != nil {
		return false, err
	}
	if n, _ := res.RowsAffected(); n > 0 {
		return true, nil
	}

	return false, nil
}

func (s *Storage) CheckIfIPInBlackList(ip string) (bool, error) {
	res, err := s.db.Exec(fmt.Sprintf("SELECT 1 FROM %s WHERE %s >>= $1;", BlackListIpsTable, SubnetColumnName), ip)
	if err != nil {
		return false, err
	}
	if n, _ := res.RowsAffected(); n > 0 {
		return true, nil
	}

	return false, nil
}

func (s *Storage) ResetDatabase() error {
	var tables = []string{
		WhiteListIpsTable,
		BlackListIpsTable,
	}
	_, err := s.db.Exec(fmt.Sprintf(`TRUNCATE TABLE %s`, strings.Join(tables, ", ")))

	return err
}
