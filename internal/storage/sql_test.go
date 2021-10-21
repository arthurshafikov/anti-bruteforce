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

func TestAddToWhiteList(t *testing.T) {
	mockDB, mock, mockStorage := newSQLStorageMock(t)
	defer mockDB.Close()

	mock.ExpectExec(fmt.Sprintf("INSERT INTO %s", WhiteListIpsTable)).
		WithArgs(testSubnet).
		WillReturnResult(sqlmock.NewResult(0, 0))

	err := mockStorage.AddToWhiteList(testSubnet)
	require.NoError(t, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestRemoveFromWhiteList(t *testing.T) {
	mockDB, mock, mockStorage := newSQLStorageMock(t)
	defer mockDB.Close()

	mock.ExpectExec(fmt.Sprintf("DELETE FROM %s", WhiteListIpsTable)).
		WithArgs(testSubnet).
		WillReturnResult(sqlmock.NewResult(0, 0))

	err := mockStorage.RemoveFromWhiteList(testSubnet)
	require.NoError(t, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestAddToBlackList(t *testing.T) {
	mockDB, mock, mockStorage := newSQLStorageMock(t)
	defer mockDB.Close()

	mock.ExpectExec(fmt.Sprintf("INSERT INTO %s", BlackListIpsTable)).
		WithArgs(testSubnet).
		WillReturnResult(sqlmock.NewResult(0, 0))

	err := mockStorage.AddToBlackList(testSubnet)
	require.NoError(t, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestRemoveFromBlackList(t *testing.T) {
	mockDB, mock, mockStorage := newSQLStorageMock(t)
	defer mockDB.Close()

	mock.ExpectExec(fmt.Sprintf("DELETE FROM %s", BlackListIpsTable)).
		WithArgs(testSubnet).
		WillReturnResult(sqlmock.NewResult(0, 0))

	err := mockStorage.RemoveFromBlackList(testSubnet)
	require.NoError(t, err)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestCheckIfIPInWhiteList(t *testing.T) {
	mockDB, mock, mockStorage := newSQLStorageMock(t)
	defer mockDB.Close()

	mock.ExpectExec(fmt.Sprintf("SELECT 1 FROM %s", WhiteListIpsTable)).
		WithArgs(testSubnet).
		WillReturnResult(sqlmock.NewResult(1, 1))

	res, err := mockStorage.CheckIfIPInWhiteList(testSubnet)
	require.NoError(t, err)
	require.True(t, res)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestCheckIfIPInBlackList(t *testing.T) {
	mockDB, mock, mockStorage := newSQLStorageMock(t)
	defer mockDB.Close()

	mock.ExpectExec(fmt.Sprintf("SELECT 1 FROM %s", BlackListIpsTable)).
		WithArgs(testSubnet).
		WillReturnResult(sqlmock.NewResult(1, 1))

	res, err := mockStorage.CheckIfIPInBlackList(testSubnet)
	require.NoError(t, err)
	require.True(t, res)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}
