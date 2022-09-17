package repositories

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/arthurshafikov/anti-bruteforce/internal/core"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
)

var testSubnet = "194.20.10.0/24"

func newBlacklistRepoMock(t *testing.T) (*sql.DB, sqlmock.Sqlmock, *Blacklist) {
	t.Helper()
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	blacklistRepoMock := &Blacklist{
		db: sqlxDB,
	}

	return mockDB, mock, blacklistRepoMock
}

func TestAddToBlacklist(t *testing.T) {
	mockDB, mock, blacklistRepoMock := newBlacklistRepoMock(t)
	defer mockDB.Close()

	mock.ExpectExec(fmt.Sprintf("INSERT INTO %s", core.BlacklistIpsTable)).
		WithArgs(testSubnet).
		WillReturnResult(sqlmock.NewResult(0, 0))

	err := blacklistRepoMock.AddToBlacklist(testSubnet)
	require.NoError(t, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestRemoveFromBlacklist(t *testing.T) {
	mockDB, mock, blacklistRepoMock := newBlacklistRepoMock(t)
	defer mockDB.Close()

	mock.ExpectExec(fmt.Sprintf("DELETE FROM %s", core.BlacklistIpsTable)).
		WithArgs(testSubnet).
		WillReturnResult(sqlmock.NewResult(0, 0))

	err := blacklistRepoMock.RemoveFromBlacklist(testSubnet)
	require.NoError(t, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestCheckIfIPInBlacklist(t *testing.T) {
	mockDB, mock, blacklistRepoMock := newBlacklistRepoMock(t)
	defer mockDB.Close()

	mock.ExpectExec(fmt.Sprintf("SELECT 1 FROM %s", core.BlacklistIpsTable)).
		WithArgs(testSubnet).
		WillReturnResult(sqlmock.NewResult(1, 1))

	res, err := blacklistRepoMock.CheckIfIPInBlacklist(testSubnet)
	require.NoError(t, err)
	require.True(t, res)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}
