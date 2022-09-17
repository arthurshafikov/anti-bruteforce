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

func newWhitelistRepoMock(t *testing.T) (*sql.DB, sqlmock.Sqlmock, *Whitelist) {
	t.Helper()
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	whitelistRepoMock := &Whitelist{
		db: sqlxDB,
	}

	return mockDB, mock, whitelistRepoMock
}

func TestCheckIfIPInWhitelist(t *testing.T) {
	mockDB, mock, whitelistRepoMock := newWhitelistRepoMock(t)
	defer mockDB.Close()

	mock.ExpectExec(fmt.Sprintf("SELECT 1 FROM %s", core.WhitelistIpsTable)).
		WithArgs(testSubnet).
		WillReturnResult(sqlmock.NewResult(1, 1))

	res, err := whitelistRepoMock.CheckIfIPInWhitelist(testSubnet)
	require.NoError(t, err)
	require.True(t, res)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestAddToWhitelist(t *testing.T) {
	mockDB, mock, whitelistRepoMock := newWhitelistRepoMock(t)
	defer mockDB.Close()

	mock.ExpectExec(fmt.Sprintf("INSERT INTO %s", core.WhitelistIpsTable)).
		WithArgs(testSubnet).
		WillReturnResult(sqlmock.NewResult(0, 0))

	err := whitelistRepoMock.AddToWhitelist(testSubnet)
	require.NoError(t, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestRemoveFromWhitelist(t *testing.T) {
	mockDB, mock, whitelistRepoMock := newWhitelistRepoMock(t)
	defer mockDB.Close()

	mock.ExpectExec(fmt.Sprintf("DELETE FROM %s", core.WhitelistIpsTable)).
		WithArgs(testSubnet).
		WillReturnResult(sqlmock.NewResult(0, 0))

	err := whitelistRepoMock.RemoveFromWhitelist(testSubnet)
	require.NoError(t, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}
