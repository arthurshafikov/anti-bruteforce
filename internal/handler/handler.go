package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Storage interface {
	AddToWhiteList(string) error
	RemoveFromWhiteList(string) error
	AddToBlackList(string) error
	RemoveFromBlackList(string) error
}

type Logger interface {
	Warn(string)
	Info(string)
	Error(string)
}

type Handler struct {
	Storage Storage
	Logger  Logger
}

func NewHandler(storage Storage, logger Logger) *Handler {
	return &Handler{
		Storage: storage,
		Logger:  logger,
	}
}

func (h *Handler) Home(c *gin.Context) {
	c.JSON(http.StatusOK, "OK")
}

func (h *Handler) Authorize(c *gin.Context) {
	panic("Implement me")
}

func (h *Handler) ResetBucket(c *gin.Context) {
	panic("Implement me")
}

func (h *Handler) AddToWhiteList(c *gin.Context) {
	panic("Implement me")
}

func (h *Handler) RemoveFromWhiteList(c *gin.Context) {
	panic("Implement me")
}

func (h *Handler) AddToBlackList(c *gin.Context) {
	panic("Implement me")
}

func (h *Handler) RemoveFromBlackList(c *gin.Context) {
	panic("Implement me")
}
