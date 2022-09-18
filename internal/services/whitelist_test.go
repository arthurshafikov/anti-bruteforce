package services

import (
	"testing"

	mock_repository "github.com/arthurshafikov/anti-bruteforce/internal/repository/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func getWhitelistService(t *testing.T) (*WhitelistService, *mock_repository.MockWhitelist) {
	t.Helper()

	ctrl := gomock.NewController(t)
	whitelistRepoMock := mock_repository.NewMockWhitelist(ctrl)

	return NewWhitelistService(whitelistRepoMock), whitelistRepoMock
}

func TestAddToWhitelist(t *testing.T) {
	bsService, repo := getWhitelistService(t)
	gomock.InOrder(
		repo.EXPECT().AddToWhitelist(subnetInput.Subnet).Return(nil),
	)

	err := bsService.AddToWhitelist(subnetInput)
	require.NoError(t, err)
}

func TestCheckIfIPInWhitelist(t *testing.T) {
	bsService, repo := getWhitelistService(t)
	gomock.InOrder(
		repo.EXPECT().CheckIfIPInWhitelist(ip).Return(false, nil),
	)

	res, err := bsService.CheckIfIPInWhitelist(ip)
	require.NoError(t, err)
	require.False(t, res)
}

func TestRemoveFromWhitelist(t *testing.T) {
	bsService, repo := getWhitelistService(t)
	gomock.InOrder(
		repo.EXPECT().RemoveFromWhitelist(subnetInput.Subnet).Return(nil),
	)

	err := bsService.RemoveFromWhitelist(subnetInput)
	require.NoError(t, err)
}
