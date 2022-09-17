package services

type BucketService struct {
	leakyBucket LeakyBucket
}

func NewBucketService(leakyBucket LeakyBucket) *BucketService {
	return &BucketService{
		leakyBucket: leakyBucket,
	}
}

func (bs *BucketService) ResetBucket() {
	bs.leakyBucket.ResetResetBucketTicker()
}
