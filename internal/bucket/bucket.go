package bucket

import (
	"context"
	"sync"
	"time"

	"github.com/arthurshafikov/anti-bruteforce/internal/core"
)

const (
	loginBucketKey    = "login"
	passwordBucketKey = "password"
	ipBucketKey       = "ip"
)

type bucketValuesMap map[string]int64

type bucket struct {
	mu        *sync.Mutex
	limit     int64
	ValuesMap bucketValuesMap
}

type bucketsMap map[string]bucket

type LeakyBucket struct {
	buckets            bucketsMap
	resetBucketsTicker *time.Ticker
}

func NewLeakyBucket(resetBucketsTicker *time.Ticker, authLimits core.AuthorizeLimits) *LeakyBucket {
	buckets := bucketsMap{
		loginBucketKey: bucket{
			mu:        &sync.Mutex{},
			limit:     authLimits.LimitAttemptsForLogin,
			ValuesMap: make(bucketValuesMap),
		},
		passwordBucketKey: bucket{
			mu:        &sync.Mutex{},
			limit:     authLimits.LimitAttemptsForPassword,
			ValuesMap: make(bucketValuesMap),
		},
		ipBucketKey: bucket{
			mu:        &sync.Mutex{},
			limit:     authLimits.LimitAttemptsForIP,
			ValuesMap: make(bucketValuesMap),
		},
	}

	return &LeakyBucket{
		buckets:            buckets,
		resetBucketsTicker: resetBucketsTicker,
	}
}

func (lb *LeakyBucket) Add(input core.AuthorizeInput) bool {
	return lb.addLogin(input.Login) && lb.addIP(input.IP) && lb.addPassword(input.Password)
}

func (lb *LeakyBucket) addLogin(login string) bool {
	return lb.addInBucket(loginBucketKey, login)
}

func (lb *LeakyBucket) addIP(ip string) bool {
	return lb.addInBucket(ipBucketKey, ip)
}

func (lb *LeakyBucket) addPassword(password string) bool {
	return lb.addInBucket(passwordBucketKey, password)
}

func (lb *LeakyBucket) addInBucket(bucketName, value string) bool {
	bucket, ok := lb.buckets[bucketName]
	if !ok {
		return false
	}

	lb.buckets[bucketName].mu.Lock()
	defer lb.buckets[bucketName].mu.Unlock()
	if bucket.ValuesMap[value] >= bucket.limit {
		return false
	}

	lb.buckets[bucketName].ValuesMap[value]++
	return true
}

func (lb *LeakyBucket) ResetResetBucketTicker() {
	lb.resetBucketsTicker.Reset(time.Second)
	lb.resetBucket()
}

func (lb *LeakyBucket) resetBucket() {
	for _, bucketName := range []string{loginBucketKey, passwordBucketKey, ipBucketKey} {
		if bucket, ok := lb.buckets[bucketName]; ok {
			bucket.ValuesMap = make(bucketValuesMap)

			lb.buckets[bucketName] = bucket
		}
	}
}

func (lb *LeakyBucket) Leak(ctx context.Context) {
	for range lb.resetBucketsTicker.C {
		select {
		case <-ctx.Done():
			return
		default:
		}

		lb.resetBucket()
	}
}
