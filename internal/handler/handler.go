package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thewolf27/anti-bruteforce/internal/bucket"
)

const (
	// QuerySubnetParam is a subnet request parameter key.
	QuerySubnetParam = "subnet"

	// WrongSubnetErrorMessage is a request wrong subnet format error.
	WrongSubnetErrorMessage = "wrong subnet format"
)

type App interface {
	Authorize(bucket.AuthorizeInput) bool
	ResetBucket()
	AddToWhiteList(string) error
	AddToBlackList(string) error
	RemoveFromWhiteList(string) error
	RemoveFromBlackList(string) error
}

type Handler struct {
	App App
}

func NewHandler(app App) *Handler {
	return &Handler{
		App: app,
	}
}

func (h *Handler) Home(c *gin.Context) {
	c.JSON(http.StatusOK, "OK")
}

func (h *Handler) Authorize(c *gin.Context) {
	login := c.DefaultQuery("login", "testlogin")
	password := c.DefaultQuery("password", "testpassword")
	ip := c.DefaultQuery("ip", "129.25.10.0")

	res := h.App.Authorize(bucket.AuthorizeInput{
		Login:    login,
		Password: password,
		IP:       ip,
	})

	if res {
		c.JSON(http.StatusOK, "OK")
		return
	}

	c.JSON(http.StatusTooManyRequests, "Too many requests")
}

func (h *Handler) ResetBucket(c *gin.Context) {
	h.App.ResetBucket()
	c.JSON(http.StatusNoContent, "")
}

func (h *Handler) AddToWhiteList(c *gin.Context) {
	subnet := c.Query(QuerySubnetParam)

	err := h.App.AddToWhiteList(subnet)
	if err != nil {
		h.setWrongSubnetErrorMessageResponse(c, err)
		return
	}

	c.JSON(http.StatusCreated, "")
}

func (h *Handler) AddToBlackList(c *gin.Context) {
	subnet := c.Query(QuerySubnetParam)

	err := h.App.AddToBlackList(subnet)
	if err != nil {
		h.setWrongSubnetErrorMessageResponse(c, err)
		return
	}

	c.JSON(http.StatusCreated, "")
}

func (h *Handler) RemoveFromWhiteList(c *gin.Context) {
	subnet := c.Query(QuerySubnetParam)

	err := h.App.RemoveFromWhiteList(subnet)
	if err != nil {
		h.setWrongSubnetErrorMessageResponse(c, err)
		return
	}

	c.JSON(http.StatusOK, "")
}

func (h *Handler) RemoveFromBlackList(c *gin.Context) {
	subnet := c.Query(QuerySubnetParam)

	err := h.App.RemoveFromBlackList(subnet)
	if err != nil {
		h.setWrongSubnetErrorMessageResponse(c, err)
		return
	}

	c.JSON(http.StatusOK, "")
}

func (h *Handler) setWrongSubnetErrorMessageResponse(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, fmt.Errorf("%s %w", WrongSubnetErrorMessage, err).Error())
}
