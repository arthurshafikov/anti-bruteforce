package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestResetBucket(t *testing.T) {
	w, c, h, mockBag := getWriterContextAndHandlerWithMocks(t)
	c.Request = httptest.NewRequest(http.MethodPost, "/bucket/reset", nil)
	gomock.InOrder(
		mockBag.bucket.EXPECT().ResetBucket(),
	)

	h.ResetBucket(c)

	require.Equal(t, http.StatusOK, w.Code)
}
