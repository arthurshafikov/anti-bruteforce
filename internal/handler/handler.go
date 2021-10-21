package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	// Subnet request parameter key
	QuerySubnetParam = "subnet"

	// Request wrong subnet format error
	WrongSubnetErrorMessage = "wrong subnet format"
)

type Storage interface {
	AddToWhiteList(string) error
	AddToBlackList(string) error
	RemoveFromWhiteList(string) error
	RemoveFromBlackList(string) error
	CheckIfIPInWhiteList(string) (bool, error)
	CheckIfIPInBlackList(string) (bool, error)
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
	subnet := c.Query(QuerySubnetParam)

	err := h.Storage.AddToWhiteList(subnet)
	if err != nil {
		c.JSON(http.StatusInternalServerError, WrongSubnetErrorMessage)
		return
	}

	c.JSON(http.StatusCreated, "")
}

func (h *Handler) AddToBlackList(c *gin.Context) {
	subnet := c.Query(QuerySubnetParam)

	err := h.Storage.AddToBlackList(subnet)
	if err != nil {
		c.JSON(http.StatusInternalServerError, WrongSubnetErrorMessage)
		return
	}

	c.JSON(http.StatusCreated, "")
}

func (h *Handler) RemoveFromWhiteList(c *gin.Context) {
	subnet := c.Query(QuerySubnetParam)

	err := h.Storage.RemoveFromWhiteList(subnet)
	if err != nil {
		c.JSON(http.StatusInternalServerError, WrongSubnetErrorMessage)
		return
	}

	c.JSON(http.StatusOK, "")
}

func (h *Handler) RemoveFromBlackList(c *gin.Context) {
	subnet := c.Query(QuerySubnetParam)

	err := h.Storage.RemoveFromBlackList(subnet)
	if err != nil {
		c.JSON(http.StatusInternalServerError, WrongSubnetErrorMessage)
		return
	}

	c.JSON(http.StatusOK, "")
}
