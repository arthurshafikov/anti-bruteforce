package logger

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	errSome     = fmt.Errorf("some error")
	someWarning = "some warning"
	someInfo    = "some info"
)

func TestError(t *testing.T) {
	result := wrapLogTest(t, func() {
		logger := NewLogger(errorLevel)
		logger.Error(errSome)

		// should not be logged
		logger.Warn(someWarning)
		logger.Info(someInfo)
	})

	require.Contains(t, result, errSome.Error())
	require.NotContains(t, result, someWarning)
	require.NotContains(t, result, someInfo)
}

func TestWarn(t *testing.T) {
	result := wrapLogTest(t, func() {
		logger := NewLogger(warnLevel)
		logger.Warn(someWarning)

		// should not be logged
		logger.Info(someInfo)
	})

	require.Contains(t, result, someWarning)
	require.NotContains(t, result, someInfo)
}

func TestInfo(t *testing.T) {
	result := wrapLogTest(t, func() {
		logger := NewLogger(infoLevel)
		logger.Info(someInfo)
	})

	require.Contains(t, result, someInfo)
}

func TestDebug(t *testing.T) {
	result := wrapLogTest(t, func() {
		logger := NewLogger(debugLevel)
		logger.Error(errSome)
		logger.Warn(someWarning)
		logger.Info(someInfo)
	})

	require.Contains(t, result, errSome.Error())
	require.Contains(t, result, someWarning)
	require.Contains(t, result, someInfo)
}

func wrapLogTest(t *testing.T, callback func()) string {
	t.Helper()

	rescueStdout := os.Stdout
	reader, writer, err := os.Pipe()
	require.NoError(t, err)
	os.Stdout = writer

	callback()

	writer.Close()
	out, err := ioutil.ReadAll(reader)
	require.NoError(t, err)
	os.Stdout = rescueStdout

	return string(out)
}
