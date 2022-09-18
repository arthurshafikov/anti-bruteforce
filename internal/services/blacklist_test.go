package services

import (
	"testing"

	"github.com/arthurshafikov/anti-bruteforce/internal/core"
	mock_repository "github.com/arthurshafikov/anti-bruteforce/internal/repository/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

var (
	subnetInput = core.SubnetInput{
		Subnet: "192.0.0.16/12",
	}
	ip = "127.0.0.1"
)

func getBlacklistService(t *testing.T) (*BlacklistService, *mock_repository.MockBlacklist) {
	t.Helper()

	ctrl := gomock.NewController(t)
	blacklistRepoMock := mock_repository.NewMockBlacklist(ctrl)

	return NewBlacklistService(blacklistRepoMock), blacklistRepoMock
}

func TestAddToBlacklist(t *testing.T) {
	bsService, repo := getBlacklistService(t)
	gomock.InOrder(
		repo.EXPECT().AddToBlacklist(subnetInput.Subnet).Return(nil),
	)

	err := bsService.AddToBlacklist(subnetInput)
	require.NoError(t, err)
}

func TestCheckIfIPInBlacklist(t *testing.T) {
	bsService, repo := getBlacklistService(t)
	gomock.InOrder(
		repo.EXPECT().CheckIfIPInBlacklist(ip).Return(false, nil),
	)

	res, err := bsService.CheckIfIPInBlacklist(ip)
	require.NoError(t, err)
	require.False(t, res)
}

func TestRemoveFromBlacklist(t *testing.T) {
	bsService, repo := getBlacklistService(t)
	gomock.InOrder(
		repo.EXPECT().RemoveFromBlacklist(subnetInput.Subnet).Return(nil),
	)

	err := bsService.RemoveFromBlacklist(subnetInput)
	require.NoError(t, err)
}
