package tests

import (
	"encoding/json"
	"testing"

	"github.com/arthurshafikov/anti-bruteforce/internal/core"
	"github.com/stretchr/testify/require"
)

var (
	successResponse               = "{\"data\":\"OK\"}"
	duplicateSubnetResponseFormat = "{\"data\":\"pq: invalid input syntax for type inet: \\\"%s\\\"\"}"
	tooManyRequestsResponse       = "{\"data\":\"too many requests\"}"

	limitAttemptsForLogin    = 10
	limitAttemptsForPassword = 10
	limitAttemptsForIP       = 10

	testLogin    = "testLogin"
	testPassword = "testPass"
	testSubnet   = "198.24.15.0/24"

	authorizeInput = core.AuthorizeInput{
		Login:    testLogin,
		Password: testPassword,
		IP:       testSubnet,
	}

	subnetInput = core.SubnetInput{
		Subnet: testSubnet,
	}
)

func getJSONBody(t *testing.T, v interface{}) []byte {
	t.Helper()
	jsonBody, err := json.Marshal(v)
	require.NoError(t, err)

	return jsonBody
}
