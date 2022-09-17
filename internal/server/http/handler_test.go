package http

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/arthurshafikov/anti-bruteforce/internal/core"
	"github.com/arthurshafikov/anti-bruteforce/internal/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestAuthorize(t *testing.T) {
	t.Run("without json body", func(t *testing.T) {
		w, c, h := getWriterContextAndHandler()
		c.Request = httptest.NewRequest(http.MethodPost, "/authorize", nil)

		h.Authorize(c)

		require.Equal(t, http.StatusUnprocessableEntity, w.Code)
	})

	t.Run("with json body", func(t *testing.T) {
		w, c, h := getWriterContextAndHandler()
		jsonBody := getAuthorizeJSONBody(t)
		c.Request = httptest.NewRequest(http.MethodPost, "/authorize", bytes.NewBuffer(jsonBody))

		h.Authorize(c)

		require.Equal(t, http.StatusOK, w.Code)
	})
}

func TestResetBucket(t *testing.T) {
	w, c, h := getWriterContextAndHandler()
	c.Request = httptest.NewRequest(http.MethodPost, "/bucket/reset", nil)

	h.ResetBucket(c)

	require.Equal(t, http.StatusOK, w.Code)
}

func TestAddToWhitelist(t *testing.T) {
	t.Run("without json body", func(t *testing.T) {
		w, c, h := getWriterContextAndHandler()
		c.Request = httptest.NewRequest(http.MethodPost, "/whitelist/add", nil)

		h.AddToWhitelist(c)

		require.Equal(t, http.StatusUnprocessableEntity, w.Code)
	})

	t.Run("with json body", func(t *testing.T) {
		w, c, h := getWriterContextAndHandler()
		jsonBody := getSubnetJSONBody(t)
		c.Request = httptest.NewRequest(http.MethodPost, "/whitelist/add", bytes.NewBuffer(jsonBody))

		h.AddToWhitelist(c)

		require.Equal(t, http.StatusCreated, w.Code)
	})
}

func TestRemoveFromWhitelist(t *testing.T) {
	t.Run("without json body", func(t *testing.T) {
		w, c, h := getWriterContextAndHandler()
		c.Request = httptest.NewRequest(http.MethodDelete, "/whitelist/remove", nil)

		h.RemoveFromWhitelist(c)

		require.Equal(t, http.StatusUnprocessableEntity, w.Code)
	})

	t.Run("with json body", func(t *testing.T) {
		w, c, h := getWriterContextAndHandler()
		jsonBody := getSubnetJSONBody(t)
		c.Request = httptest.NewRequest(http.MethodDelete, "/whitelist/remove", bytes.NewBuffer(jsonBody))

		h.RemoveFromWhitelist(c)

		require.Equal(t, http.StatusOK, w.Code)
	})
}

func TestAddToBlacklist(t *testing.T) {
	t.Run("without json body", func(t *testing.T) {
		w, c, h := getWriterContextAndHandler()
		c.Request = httptest.NewRequest(http.MethodPost, "/blacklist/add", nil)

		h.AddToBlacklist(c)

		require.Equal(t, http.StatusUnprocessableEntity, w.Code)
	})

	t.Run("with json body", func(t *testing.T) {
		w, c, h := getWriterContextAndHandler()
		jsonBody := getSubnetJSONBody(t)
		c.Request = httptest.NewRequest(http.MethodPost, "/blacklist/add", bytes.NewBuffer(jsonBody))

		h.AddToBlacklist(c)

		require.Equal(t, http.StatusCreated, w.Code)
	})
}

func TestRemoveFromBlacklist(t *testing.T) {
	t.Run("without json body", func(t *testing.T) {
		w, c, h := getWriterContextAndHandler()
		c.Request = httptest.NewRequest(http.MethodDelete, "/blacklist/remove", nil)

		h.RemoveFromBlacklist(c)

		require.Equal(t, http.StatusUnprocessableEntity, w.Code)
	})

	t.Run("with json body", func(t *testing.T) {
		w, c, h := getWriterContextAndHandler()
		jsonBody := getSubnetJSONBody(t)
		c.Request = httptest.NewRequest(http.MethodDelete, "/blacklist/remove", bytes.NewBuffer(jsonBody))

		h.RemoveFromBlacklist(c)

		require.Equal(t, http.StatusOK, w.Code)
	})
}

func getWriterContextAndHandler() (*httptest.ResponseRecorder, *gin.Context, *Handler) {
	gin.SetMode(gin.TestMode)
	h := NewHandler(&mocks.App{})
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	return w, c, h
}

func getSubnetJSONBody(t *testing.T) []byte {
	t.Helper()
	jsonBody, err := json.Marshal(core.SubnetInput{
		Subnet: "198.24.15.0/24",
	})
	require.NoError(t, err)

	return jsonBody
}

func getAuthorizeJSONBody(t *testing.T) []byte {
	t.Helper()
	jsonBody, err := json.Marshal(core.AuthorizeInput{
		Login:    "testlogin",
		Password: "testpass",
		IP:       "198.24.15.10",
	})
	require.NoError(t, err)

	return jsonBody
}
