package http

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"github.com/thewolf27/anti-bruteforce/internal/mocks"
	"github.com/thewolf27/anti-bruteforce/internal/models"
)

func TestAuthorize(t *testing.T) {
	t.Run("without json body", func(t *testing.T) {
		w, c, h := getWriterContextAndHandler(t)
		c.Request = httptest.NewRequest(http.MethodPost, "/authorize", nil)

		h.Authorize(c)

		require.Equal(t, http.StatusUnprocessableEntity, w.Code)
	})

	t.Run("with json body", func(t *testing.T) {
		w, c, h := getWriterContextAndHandler(t)
		jsonBody := getAuthorizeJsonBody(t)
		c.Request = httptest.NewRequest(http.MethodPost, "/authorize", bytes.NewBuffer(jsonBody))

		h.Authorize(c)

		require.Equal(t, http.StatusOK, w.Code)
	})
}

func TestResetBucket(t *testing.T) {
	w, c, h := getWriterContextAndHandler(t)
	c.Request = httptest.NewRequest(http.MethodPost, "/bucket/reset", nil)

	h.ResetBucket(c)

	require.Equal(t, http.StatusOK, w.Code)
}

func TestAddToWhiteList(t *testing.T) {
	t.Run("without json body", func(t *testing.T) {
		w, c, h := getWriterContextAndHandler(t)
		c.Request = httptest.NewRequest(http.MethodPost, "/whitelist/add", nil)

		h.AddToWhiteList(c)

		require.Equal(t, http.StatusUnprocessableEntity, w.Code)
	})

	t.Run("with json body", func(t *testing.T) {
		w, c, h := getWriterContextAndHandler(t)
		jsonBody := getSubnetJsonBody(t)
		c.Request = httptest.NewRequest(http.MethodPost, "/whitelist/add", bytes.NewBuffer(jsonBody))

		h.AddToWhiteList(c)

		require.Equal(t, http.StatusCreated, w.Code)
	})
}

func TestRemoveFromWhiteList(t *testing.T) {
	t.Run("without json body", func(t *testing.T) {
		w, c, h := getWriterContextAndHandler(t)
		c.Request = httptest.NewRequest(http.MethodDelete, "/whitelist/remove", nil)

		h.RemoveFromWhiteList(c)

		require.Equal(t, http.StatusUnprocessableEntity, w.Code)
	})

	t.Run("with json body", func(t *testing.T) {
		w, c, h := getWriterContextAndHandler(t)
		jsonBody := getSubnetJsonBody(t)
		c.Request = httptest.NewRequest(http.MethodDelete, "/whitelist/remove", bytes.NewBuffer(jsonBody))

		h.RemoveFromWhiteList(c)

		require.Equal(t, http.StatusOK, w.Code)
	})
}

func TestAddToBlackList(t *testing.T) {
	t.Run("without json body", func(t *testing.T) {
		w, c, h := getWriterContextAndHandler(t)
		c.Request = httptest.NewRequest(http.MethodPost, "/blacklist/add", nil)

		h.AddToBlackList(c)

		require.Equal(t, http.StatusUnprocessableEntity, w.Code)
	})

	t.Run("with json body", func(t *testing.T) {
		w, c, h := getWriterContextAndHandler(t)
		jsonBody := getSubnetJsonBody(t)
		c.Request = httptest.NewRequest(http.MethodPost, "/blacklist/add", bytes.NewBuffer(jsonBody))

		h.AddToBlackList(c)

		require.Equal(t, http.StatusCreated, w.Code)
	})
}

func TestRemoveFromBlackList(t *testing.T) {
	t.Run("without json body", func(t *testing.T) {
		w, c, h := getWriterContextAndHandler(t)
		c.Request = httptest.NewRequest(http.MethodDelete, "/blacklist/remove", nil)

		h.RemoveFromBlackList(c)

		require.Equal(t, http.StatusUnprocessableEntity, w.Code)
	})

	t.Run("with json body", func(t *testing.T) {
		w, c, h := getWriterContextAndHandler(t)
		jsonBody := getSubnetJsonBody(t)
		c.Request = httptest.NewRequest(http.MethodDelete, "/blacklist/remove", bytes.NewBuffer(jsonBody))

		h.RemoveFromBlackList(c)

		require.Equal(t, http.StatusOK, w.Code)
	})
}

func getWriterContextAndHandler(t *testing.T) (*httptest.ResponseRecorder, *gin.Context, *Handler) {
	gin.SetMode(gin.TestMode)
	h := NewHandler(&mocks.App{})
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	return w, c, h
}

func getSubnetJsonBody(t *testing.T) []byte {
	jsonBody, err := json.Marshal(models.SubnetInput{
		Subnet: "198.24.15.0/24",
	})
	require.NoError(t, err)

	return jsonBody
}

func getAuthorizeJsonBody(t *testing.T) []byte {
	jsonBody, err := json.Marshal(models.AuthorizeInput{
		Login:    "testlogin",
		Password: "testpass",
		IP:       "198.24.15.10",
	})
	require.NoError(t, err)

	return jsonBody
}
