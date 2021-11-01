package storage

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
)

var testSubnet = "194.20.10.0/24"

func newSQLStorageMock(t *testing.T) (*sql.DB, sqlmock.Sqlmock, Storage) {
	t.Helper()
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)

	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	mockStorage := Storage{
		dsn: "test dsn",
		db:  sqlxDB,
	}

	return mockDB, mock, mockStorage
}

func TestAddToWhitelist(t *testing.T) {
	mockDB, mock, mockStorage := newSQLStorageMock(t)
	defer mockDB.Close()

	mock.ExpectExec(fmt.Sprintf("INSERT INTO %s", WhitelistIpsTable)).
		WithArgs(testSubnet).
		WillReturnResult(sqlmock.NewResult(0, 0))

	err := mockStorage.AddToWhitelist(testSubnet)
	require.NoError(t, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestRemoveFromWhitelist(t *testing.T) {
	mockDB, mock, mockStorage := newSQLStorageMock(t)
	defer mockDB.Close()

	mock.ExpectExec(fmt.Sprintf("DELETE FROM %s", WhitelistIpsTable)).
		WithArgs(testSubnet).
		WillReturnResult(sqlmock.NewResult(0, 0))

	err := mockStorage.RemoveFromWhitelist(testSubnet)
	require.NoError(t, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestAddToBlacklist(t *testing.T) {
	mockDB, mock, mockStorage := newSQLStorageMock(t)
	defer mockDB.Close()

	mock.ExpectExec(fmt.Sprintf("INSERT INTO %s", BlacklistIpsTable)).
		WithArgs(testSubnet).
		WillReturnResult(sqlmock.NewResult(0, 0))

	err := mockStorage.AddToBlacklist(testSubnet)
	require.NoError(t, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestRemoveFromBlacklist(t *testing.T) {
	mockDB, mock, mockStorage := newSQLStorageMock(t)
	defer mockDB.Close()

	mock.ExpectExec(fmt.Sprintf("DELETE FROM %s", BlacklistIpsTable)).
		WithArgs(testSubnet).
		WillReturnResult(sqlmock.NewResult(0, 0))

	err := mockStorage.RemoveFromBlacklist(testSubnet)
	require.NoError(t, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestCheckIfIPInWhitelist(t *testing.T) {
	mockDB, mock, mockStorage := newSQLStorageMock(t)
	defer mockDB.Close()

	mock.ExpectExec(fmt.Sprintf("SELECT 1 FROM %s", WhitelistIpsTable)).
		WithArgs(testSubnet).
		WillReturnResult(sqlmock.NewResult(1, 1))

	res, err := mockStorage.CheckIfIPInWhitelist(testSubnet)
	require.NoError(t, err)
	require.True(t, res)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestCheckIfIPInBlacklist(t *testing.T) {
	mockDB, mock, mockStorage := newSQLStorageMock(t)
	defer mockDB.Close()

	mock.ExpectExec(fmt.Sprintf("SELECT 1 FROM %s", BlacklistIpsTable)).
		WithArgs(testSubnet).
		WillReturnResult(sqlmock.NewResult(1, 1))

	res, err := mockStorage.CheckIfIPInBlacklist(testSubnet)
	require.NoError(t, err)
	require.True(t, res)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}
