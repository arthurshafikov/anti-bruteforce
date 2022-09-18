package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestAddToBlacklist(t *testing.T) {
	t.Run("without json body", func(t *testing.T) {
		w, c, h, _ := getWriterContextAndHandlerWithMocks(t)
		c.Request = httptest.NewRequest(http.MethodPost, "/blacklist/add", nil)

		h.AddToBlacklist(c)

		require.Equal(t, http.StatusUnprocessableEntity, w.Code)
	})

	t.Run("with json body", func(t *testing.T) {
		w, c, h, mockBag := getWriterContextAndHandlerWithMocks(t)
		jsonBody := getSubnetJSONBody(t)
		c.Request = httptest.NewRequest(http.MethodPost, "/blacklist/add", bytes.NewBuffer(jsonBody))
		gomock.InOrder(
			mockBag.blacklist.EXPECT().AddToBlacklist(subnetInput).Return(nil),
		)

		h.AddToBlacklist(c)

		require.Equal(t, http.StatusCreated, w.Code)
	})
}

func TestRemoveFromBlacklist(t *testing.T) {
	t.Run("without json body", func(t *testing.T) {
		w, c, h, _ := getWriterContextAndHandlerWithMocks(t)
		c.Request = httptest.NewRequest(http.MethodDelete, "/blacklist/remove", nil)

		h.RemoveFromBlacklist(c)

		require.Equal(t, http.StatusUnprocessableEntity, w.Code)
	})

	t.Run("with json body", func(t *testing.T) {
		w, c, h, mockBag := getWriterContextAndHandlerWithMocks(t)
		jsonBody := getSubnetJSONBody(t)
		c.Request = httptest.NewRequest(http.MethodDelete, "/blacklist/remove", bytes.NewBuffer(jsonBody))
		gomock.InOrder(
			mockBag.blacklist.EXPECT().RemoveFromBlacklist(subnetInput).Return(nil),
		)

		h.RemoveFromBlacklist(c)

		require.Equal(t, http.StatusOK, w.Code)
	})
}
