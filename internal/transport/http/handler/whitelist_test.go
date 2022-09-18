package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestAddToWhitelist(t *testing.T) {
	t.Run("without json body", func(t *testing.T) {
		w, c, h, _ := getWriterContextAndHandlerWithMocks(t)
		c.Request = httptest.NewRequest(http.MethodPost, "/whitelist/add", nil)

		h.AddToWhitelist(c)

		require.Equal(t, http.StatusUnprocessableEntity, w.Code)
	})

	t.Run("with json body", func(t *testing.T) {
		w, c, h, mockBag := getWriterContextAndHandlerWithMocks(t)
		jsonBody := getSubnetJSONBody(t)
		c.Request = httptest.NewRequest(http.MethodPost, "/whitelist/add", bytes.NewBuffer(jsonBody))
		gomock.InOrder(
			mockBag.whitelist.EXPECT().AddToWhitelist(subnetInput).Return(nil),
		)

		h.AddToWhitelist(c)

		require.Equal(t, http.StatusCreated, w.Code)
	})
}

func TestRemoveFromWhitelist(t *testing.T) {
	t.Run("without json body", func(t *testing.T) {
		w, c, h, _ := getWriterContextAndHandlerWithMocks(t)
		c.Request = httptest.NewRequest(http.MethodDelete, "/whitelist/remove", nil)

		h.RemoveFromWhitelist(c)

		require.Equal(t, http.StatusUnprocessableEntity, w.Code)
	})

	t.Run("with json body", func(t *testing.T) {
		w, c, h, mockBag := getWriterContextAndHandlerWithMocks(t)
		jsonBody := getSubnetJSONBody(t)
		c.Request = httptest.NewRequest(http.MethodDelete, "/whitelist/remove", bytes.NewBuffer(jsonBody))
		gomock.InOrder(
			mockBag.whitelist.EXPECT().RemoveFromWhitelist(subnetInput).Return(nil),
		)

		h.RemoveFromWhitelist(c)

		require.Equal(t, http.StatusOK, w.Code)
	})
}
