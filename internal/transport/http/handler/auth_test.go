package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

var tooManyRequestsServerResponse = ServerResponse{
	Data: TooManyRequestsErrorMessage,
}

func TestAuthorize(t *testing.T) {
	t.Run("without json body", func(t *testing.T) {
		w, c, h, _ := getWriterContextAndHandlerWithMocks(t)
		c.Request = httptest.NewRequest(http.MethodPost, "/authorize", nil)

		h.Authorize(c)

		require.Equal(t, http.StatusUnprocessableEntity, w.Code)
	})

	t.Run("with json body", func(t *testing.T) {
		w, c, h, mockBag := getWriterContextAndHandlerWithMocks(t)
		jsonBody := getAuthorizeJSONBody(t)
		c.Request = httptest.NewRequest(http.MethodPost, "/authorize", bytes.NewBuffer(jsonBody))
		gomock.InOrder(
			mockBag.auth.EXPECT().Authorize(authorizeInput).Return(true),
		)

		h.Authorize(c)

		require.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("too many requests", func(t *testing.T) {
		w, c, h, mockBag := getWriterContextAndHandlerWithMocks(t)
		jsonBody := getAuthorizeJSONBody(t)
		c.Request = httptest.NewRequest(http.MethodPost, "/authorize", bytes.NewBuffer(jsonBody))
		gomock.InOrder(
			mockBag.auth.EXPECT().Authorize(authorizeInput).Return(false),
		)

		h.Authorize(c)

		var response ServerResponse
		require.NoError(t, json.Unmarshal(w.Body.Bytes(), &response))
		require.Equal(t, tooManyRequestsServerResponse, response)
		require.Equal(t, http.StatusTooManyRequests, w.Code)
	})
}
