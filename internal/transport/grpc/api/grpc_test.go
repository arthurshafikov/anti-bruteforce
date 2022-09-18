package api

import (
	"context"
	"fmt"
	"testing"

	"github.com/arthurshafikov/anti-bruteforce/internal/core"
	"github.com/arthurshafikov/anti-bruteforce/internal/services"
	mock_services "github.com/arthurshafikov/anti-bruteforce/internal/services/mocks"
	"github.com/arthurshafikov/anti-bruteforce/internal/transport/grpc/generated"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

var (
	subnetInput = core.SubnetInput{
		Subnet: "127.0.0.1/12",
	}
	errSome = fmt.Errorf("some error")
)

type mockBag struct {
	Logger           *mock_services.MockLogger
	AuthService      *mock_services.MockAuth
	BucketService    *mock_services.MockBucket
	BlacklistService *mock_services.MockBlacklist
	WhitelistService *mock_services.MockWhitelist
}

func getAppService(t *testing.T) (*AppService, *mockBag) {
	t.Helper()

	ctrl := gomock.NewController(t)
	loggerMock := mock_services.NewMockLogger(ctrl)
	bucketServiceMock := mock_services.NewMockBucket(ctrl)
	authServiceMock := mock_services.NewMockAuth(ctrl)
	blacklistServiceMock := mock_services.NewMockBlacklist(ctrl)
	whitelistServiceMock := mock_services.NewMockWhitelist(ctrl)

	return &AppService{
			services: &services.Services{
				Auth:      authServiceMock,
				Blacklist: blacklistServiceMock,
				Whitelist: whitelistServiceMock,
				Bucket:    bucketServiceMock,
			},
		}, &mockBag{
			Logger:           loggerMock,
			AuthService:      authServiceMock,
			BucketService:    bucketServiceMock,
			BlacklistService: blacklistServiceMock,
			WhitelistService: whitelistServiceMock,
		}
}

func TestResetBucket(t *testing.T) {
	appService, mockBag := getAppService(t)
	gomock.InOrder(
		mockBag.BucketService.EXPECT().ResetBucket(),
	)

	response, err := appService.ResetBucket(context.Background(), &generated.EmptyRequest{})

	require.NoError(t, err)
	require.Equal(t, successResponse, response)
}

func TestAddToWhitelist(t *testing.T) {
	appService, mockBag := getAppService(t)
	gomock.InOrder(
		mockBag.WhitelistService.EXPECT().AddToWhitelist(subnetInput).Return(nil),
	)

	response, err := appService.AddToWhitelist(context.Background(), &generated.SubnetRequest{
		Subnet: subnetInput.Subnet,
	})

	require.NoError(t, err)
	require.Equal(t, successResponse, response)
}

func TestAddToWhitelistReturnsError(t *testing.T) {
	appService, mockBag := getAppService(t)
	gomock.InOrder(
		mockBag.WhitelistService.EXPECT().AddToWhitelist(subnetInput).Return(errSome),
	)

	response, err := appService.AddToWhitelist(context.Background(), &generated.SubnetRequest{
		Subnet: subnetInput.Subnet,
	})

	require.ErrorIs(t, errSome, err)
	require.Nil(t, response)
}

func TestAddToBlacklist(t *testing.T) {
	appService, mockBag := getAppService(t)
	gomock.InOrder(
		mockBag.BlacklistService.EXPECT().AddToBlacklist(subnetInput).Return(nil),
	)

	response, err := appService.AddToBlacklist(context.Background(), &generated.SubnetRequest{
		Subnet: subnetInput.Subnet,
	})

	require.NoError(t, err)
	require.Equal(t, successResponse, response)
}

func TestAddToBlacklistReturnsError(t *testing.T) {
	appService, mockBag := getAppService(t)
	gomock.InOrder(
		mockBag.BlacklistService.EXPECT().AddToBlacklist(subnetInput).Return(errSome),
	)

	response, err := appService.AddToBlacklist(context.Background(), &generated.SubnetRequest{
		Subnet: subnetInput.Subnet,
	})

	require.ErrorIs(t, errSome, err)
	require.Nil(t, response)
}
