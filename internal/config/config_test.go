package config

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	configFolderPath  = "./fakeconfigs"
	configFileContent = []byte(`AppConfig:
    NumberOfAttemptsForLogin: 10
    NumberOfAttemptsForPassword: 100
    NumberOfAttemptsForIP: 1000
LoggerConfig:
    Level: "DEBUG"
ServerConfig:
    Address: ":8123"
GrpcServerConfig:
    Address: ":8765"
DatabaseConfig:
    DSN: "host=localhost user=homestead password=secret dbname=homestead sslmode=disable"`)

	expectedAppConfig = AppConfig{
		NumberOfAttemptsForLogin:    10,
		NumberOfAttemptsForPassword: 100,
		NumberOfAttemptsForIP:       1000,
	}
	expectedLoggerConfig = LoggerConfig{
		Level: "DEBUG",
	}
	expectedHTTPServerConfig = ServerConfig{
		Address: ":8123",
	}
	expectedGRPCServerConfig = ServerConfig{
		Address: ":8765",
	}
	expectedDatabaseConfig = DatabaseConfig{
		DSN: "host=localhost user=homestead password=secret dbname=homestead sslmode=disable",
	}
)

func TestNewConfigFromEnvFile(t *testing.T) {
	createFakeConfigFileAndFolder(t)
	defer deleteFakeConfigFileAndFolder(t)

	config := NewConfig(configFolderPath)

	require.Equal(t, expectedAppConfig, config.AppConfig)
	require.Equal(t, expectedLoggerConfig, config.LoggerConfig)
	require.Equal(t, expectedHTTPServerConfig, config.ServerConfig)
	require.Equal(t, expectedGRPCServerConfig, config.GrpcServerConfig)
	require.Equal(t, expectedDatabaseConfig, config.DatabaseConfig)
}

func createFakeConfigFileAndFolder(t *testing.T) {
	t.Helper()
	if err := os.Mkdir(configFolderPath, 0777); err != nil { //nolint:gofumpt
		t.Fatal(err)
	}
	err := os.WriteFile(fmt.Sprintf("%s/config.yml", configFolderPath), configFileContent, 0600) //nolint:gofumpt
	if err != nil {
		t.Fatal(err)
	}
}

func deleteFakeConfigFileAndFolder(t *testing.T) {
	t.Helper()
	if err := os.RemoveAll(configFolderPath); err != nil {
		t.Fatal(err)
	}
}
