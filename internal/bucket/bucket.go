package bucket

import (
	"context"
	"log"
	"time"
)

const (
	resetBucketInterval = time.Second * 60
)

type AuthorizeInput struct {
	Login, IP, Password string
}

type AuthorizeLimits struct {
	LimitAttemptsForLogin, LimitAttemptsForPassword, LimitAttemptsForIP int
}

type bucketMap map[string]int

type LeakyBucket struct {
	ctx               context.Context
	loginBucket       bucketMap
	loginLimit        int
	ipBucket          bucketMap
	ipLimit           int
	passwordBucket    bucketMap
	passwordLimit     int
	resetBucketTicker *time.Ticker
}

func NewLeakyBucket(ctx context.Context, authLimits AuthorizeLimits) *LeakyBucket {
	resetBucketTicker := time.NewTicker(resetBucketInterval)
	go func() {
		<-ctx.Done()
		log.Println("Ticker stops")
		resetBucketTicker.Stop()
	}()

	lb := &LeakyBucket{
		ctx:               ctx,
		loginBucket:       make(bucketMap),
		loginLimit:        authLimits.LimitAttemptsForLogin,
		passwordBucket:    make(bucketMap),
		passwordLimit:     authLimits.LimitAttemptsForPassword,
		ipBucket:          make(bucketMap),
		ipLimit:           authLimits.LimitAttemptsForIP,
		resetBucketTicker: resetBucketTicker,
	}

	return lb
}

func (lb *LeakyBucket) Add(input AuthorizeInput) bool {
	return lb.addLogin(input.Login) && lb.addIP(input.IP) && lb.addPassword(input.Password)
}

func (lb *LeakyBucket) addLogin(login string) bool {
	lb.loginBucket[login]++
	return lb.loginBucket[login] <= lb.loginLimit
}

func (lb *LeakyBucket) addIP(ip string) bool {
	lb.ipBucket[ip]++
	return lb.ipBucket[ip] <= lb.ipLimit
}

func (lb *LeakyBucket) addPassword(password string) bool {
	lb.passwordBucket[password]++
	return lb.passwordBucket[password] <= lb.passwordLimit
}

func (lb *LeakyBucket) ResetResetBucketTicker() {
	lb.resetBucketTicker.Reset(resetBucketInterval)
	lb.resetBucket()
}

func (lb *LeakyBucket) resetBucket() {
	lb.loginBucket = make(bucketMap)
	lb.ipBucket = make(bucketMap)
	lb.passwordBucket = make(bucketMap)
}

func (lb *LeakyBucket) Leak() {
	for range lb.resetBucketTicker.C {
		select {
		case <-lb.ctx.Done():
			return
		default:
		}

		lb.resetBucket()
		log.Println("Leak")
	}
}
