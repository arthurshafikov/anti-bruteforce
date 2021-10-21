package bucket

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

var authInput = AuthorizeInput{
	Login:    "testlogin",
	Password: "testpass",
	IP:       "127.0.0.1",
}
var authLimits = AuthorizeLimits{10, 10, 10}

func TestAdd(t *testing.T) {
	lb := NewLeakyBucket(context.Background(), authLimits)

	for i := 0; i < 10; i++ {
		res := lb.Add(authInput)
		require.True(t, res)
	}

	res := lb.Add(authInput)
	require.False(t, res)
}

func TestResetBucket(t *testing.T) {
	lb := NewLeakyBucket(context.Background(), authLimits)

	for i := 0; i < 10; i++ {
		res := lb.Add(authInput)
		require.True(t, res)
	}

	lb.resetBucket()

	for i := 0; i < 10; i++ {
		res := lb.Add(authInput)
		require.True(t, res)
	}
}
