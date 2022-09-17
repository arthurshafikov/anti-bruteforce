package bucket

import (
	"sync"
	"testing"

	"github.com/arthurshafikov/anti-bruteforce/internal/core"
	"github.com/stretchr/testify/require"
)

var authInput = core.AuthorizeInput{
	Login:    "testlogin",
	Password: "testpass",
	IP:       "127.0.0.1",
}

func TestAdd(t *testing.T) {
	lb := getLeakyBucketWithLimits(10, 10, 10)

	for i := 0; i < 10; i++ {
		res := lb.Add(authInput)
		require.True(t, res)
	}
}

func TestResetBucket(t *testing.T) {
	lb := getLeakyBucketWithLimits(1, 1, 1)

	res := lb.Add(authInput)
	require.True(t, res)

	res = lb.Add(authInput)
	require.False(t, res)

	lb.resetBucket()

	res = lb.Add(authInput)
	require.True(t, res)
}

func TestConcurrentAddLogin(t *testing.T) {
	lb := getLeakyBucketWithLimits(10000, 1, 1)
	login := "testlogin"

	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 1000; j++ {
				res := lb.addLogin(login)
				require.True(t, res)
			}
			wg.Done()
		}()
	}

	wg.Wait()
	require.Equal(t, int64(3000), lb.buckets[loginBucketKey].ValuesMap[login])
}

func getLeakyBucketWithLimits(limitLogin, limitPass, limitIP int64) *LeakyBucket {
	return NewLeakyBucket(nil, core.AuthorizeLimits{
		LimitAttemptsForLogin:    limitLogin,
		LimitAttemptsForPassword: limitPass,
		LimitAttemptsForIP:       limitIP,
	})
}
