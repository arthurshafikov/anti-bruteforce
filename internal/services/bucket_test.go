package services

import (
	"testing"

	mock_services "github.com/arthurshafikov/anti-bruteforce/internal/services/mocks"
	"github.com/golang/mock/gomock"
)

func getBucketService(t *testing.T) (*BucketService, *mock_services.MockLeakyBucket) {
	t.Helper()

	ctrl := gomock.NewController(t)
	leakyBucketMock := mock_services.NewMockLeakyBucket(ctrl)

	return NewBucketService(leakyBucketMock), leakyBucketMock
}

func TestResetBucket(t *testing.T) {
	bService, repo := getBucketService(t)
	gomock.InOrder(
		repo.EXPECT().ResetResetBucketTicker(),
	)

	bService.ResetBucket()
}
