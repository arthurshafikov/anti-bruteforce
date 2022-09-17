package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) initBucketRoutes(engine *gin.Engine) {
	engine.POST("/bucket/reset", h.ResetBucket)
}

func (h *Handler) ResetBucket(c *gin.Context) {
	h.services.Bucket.ResetBucket()
	c.JSON(http.StatusOK, ServerResponse{OkResponseMessage})
}
