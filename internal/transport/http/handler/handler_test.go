package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/arthurshafikov/anti-bruteforce/internal/core"
	"github.com/arthurshafikov/anti-bruteforce/internal/services"
	mock_services "github.com/arthurshafikov/anti-bruteforce/internal/services/mocks"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

var (
	authorizeInput = core.AuthorizeInput{
		Login:    "testlogin",
		Password: "testpass",
		IP:       "198.24.15.10",
	}
	subnetInput = core.SubnetInput{
		Subnet: "198.24.15.0/24",
	}
)

type MockBag struct {
	auth      *mock_services.MockAuth
	blacklist *mock_services.MockBlacklist
	bucket    *mock_services.MockBucket
	whitelist *mock_services.MockWhitelist
}

func getWriterContextAndHandlerWithMocks(t *testing.T) (*httptest.ResponseRecorder, *gin.Context, *Handler, *MockBag) {
	t.Helper()
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	mockBag := &MockBag{
		auth:      mock_services.NewMockAuth(ctrl),
		blacklist: mock_services.NewMockBlacklist(ctrl),
		bucket:    mock_services.NewMockBucket(ctrl),
		whitelist: mock_services.NewMockWhitelist(ctrl),
	}
	h := NewHandler(&services.Services{
		Auth:      mockBag.auth,
		Blacklist: mockBag.blacklist,
		Whitelist: mockBag.whitelist,
		Bucket:    mockBag.bucket,
	})
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	return w, c, h, mockBag
}

func getSubnetJSONBody(t *testing.T) []byte {
	t.Helper()
	jsonBody, err := json.Marshal(subnetInput)
	require.NoError(t, err)

	return jsonBody
}

func getAuthorizeJSONBody(t *testing.T) []byte {
	t.Helper()
	jsonBody, err := json.Marshal(authorizeInput)
	require.NoError(t, err)

	return jsonBody
}

func TestHome(t *testing.T) {
	w, c, h, _ := getWriterContextAndHandlerWithMocks(t)
	c.Request = httptest.NewRequest(http.MethodGet, "/", nil)

	h.Home(c)

	require.Equal(t, http.StatusOK, w.Code)
}

func TestGetSubnetInput(t *testing.T) {
	t.Run("without json body", func(t *testing.T) {
		w, c, h, _ := getWriterContextAndHandlerWithMocks(t)
		c.Request = httptest.NewRequest(http.MethodGet, "/", nil)

		_, err := h.getSubnetInput(c)

		require.Error(t, err)
		require.Equal(t, http.StatusUnprocessableEntity, w.Code)
	})

	t.Run("with json body", func(t *testing.T) {
		w, c, h, _ := getWriterContextAndHandlerWithMocks(t)
		jsonBody := getSubnetJSONBody(t)
		c.Request = httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(jsonBody))

		result, err := h.getSubnetInput(c)

		require.NoError(t, err)
		require.Equal(t, subnetInput, result)
		require.Equal(t, http.StatusOK, w.Code)
	})
}
