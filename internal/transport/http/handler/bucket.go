package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) ResetBucket(c *gin.Context) {
	h.services.Bucket.ResetBucket()
	c.JSON(http.StatusOK, ServerResponse{OkResponseMessage})
}
