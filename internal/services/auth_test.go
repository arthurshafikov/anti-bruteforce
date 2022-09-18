package services

import (
	"fmt"
	"testing"

	"github.com/arthurshafikov/anti-bruteforce/internal/core"
	mock_services "github.com/arthurshafikov/anti-bruteforce/internal/services/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

var (
	input = core.AuthorizeInput{
		Login:    "username",
		Password: "somePass",
		IP:       "127.0.0.1",
	}
	errSome = fmt.Errorf("some error")
)

type mockBag struct {
	Logger           *mock_services.MockLogger
	LeakyBucket      *mock_services.MockLeakyBucket
	BlacklistService *mock_services.MockBlacklist
	WhitelistService *mock_services.MockWhitelist
}

func getAuthService(t *testing.T) (*AuthService, *mockBag) {
	t.Helper()

	ctrl := gomock.NewController(t)
	loggerMock := mock_services.NewMockLogger(ctrl)
	leakyBucketMock := mock_services.NewMockLeakyBucket(ctrl)
	blacklistServiceMock := mock_services.NewMockBlacklist(ctrl)
	whitelistServiceMock := mock_services.NewMockWhitelist(ctrl)

	return NewAuthService(loggerMock, leakyBucketMock, blacklistServiceMock, whitelistServiceMock), &mockBag{
		Logger:           loggerMock,
		LeakyBucket:      leakyBucketMock,
		BlacklistService: blacklistServiceMock,
		WhitelistService: whitelistServiceMock,
	}
}

func TestAuthorize(t *testing.T) {
	aService, mockBag := getAuthService(t)
	gomock.InOrder(
		mockBag.BlacklistService.EXPECT().CheckIfIPInBlacklist(input.IP).Return(false, nil),
		mockBag.WhitelistService.EXPECT().CheckIfIPInWhitelist(input.IP).Return(false, nil),
		mockBag.LeakyBucket.EXPECT().Add(input).Return(true),
	)

	res := aService.Authorize(input)
	require.True(t, res)
}

func TestAuthorizeBlacklistReturnsError(t *testing.T) {
	aService, mockBag := getAuthService(t)
	gomock.InOrder(
		mockBag.BlacklistService.EXPECT().CheckIfIPInBlacklist(input.IP).Return(false, errSome),
		mockBag.Logger.EXPECT().Error(errSome),
	)

	res := aService.Authorize(input)
	require.False(t, res)
}

func TestAuthorizeIPInBlacklist(t *testing.T) {
	aService, mockBag := getAuthService(t)
	gomock.InOrder(
		mockBag.BlacklistService.EXPECT().CheckIfIPInBlacklist(input.IP).Return(true, nil),
	)

	res := aService.Authorize(input)
	require.False(t, res)
}

func TestAuthorizeWhitelistReturnsError(t *testing.T) {
	aService, mockBag := getAuthService(t)
	gomock.InOrder(
		mockBag.BlacklistService.EXPECT().CheckIfIPInBlacklist(input.IP).Return(false, nil),
		mockBag.WhitelistService.EXPECT().CheckIfIPInWhitelist(input.IP).Return(false, errSome),
		mockBag.Logger.EXPECT().Error(errSome),
	)

	res := aService.Authorize(input)
	require.False(t, res)
}

func TestAuthorizeIPInWhitelist(t *testing.T) {
	aService, mockBag := getAuthService(t)
	gomock.InOrder(
		mockBag.BlacklistService.EXPECT().CheckIfIPInBlacklist(input.IP).Return(false, nil),
		mockBag.WhitelistService.EXPECT().CheckIfIPInWhitelist(input.IP).Return(true, nil),
	)

	res := aService.Authorize(input)
	require.True(t, res)
}
