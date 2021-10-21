package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thewolf27/anti-bruteforce/internal/bucket"
	"github.com/thewolf27/anti-bruteforce/internal/config"
)

const (
	// QuerySubnetParam is a subnet request parameter key.
	QuerySubnetParam = "subnet"

	// WrongSubnetErrorMessage is a request wrong subnet format error.
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
	Storage     Storage
	Logger      Logger
	LeakyBucket *bucket.LeakyBucket
}

func NewHandler(ctx context.Context, storage Storage, logger Logger, appConfig config.AppConfig) *Handler {
	leakyBucket := bucket.NewLeakyBucket(ctx, bucket.AuthorizeLimits{
		LimitAttemptsForLogin:    appConfig.NumberOfAttemptsForLogin,
		LimitAttemptsForPassword: appConfig.NumberOfAttemptsForPassword,
		LimitAttemptsForIP:       appConfig.NumberOfAttemptsForIP,
	})
	go leakyBucket.Leak()

	return &Handler{
		Storage:     storage,
		Logger:      logger,
		LeakyBucket: leakyBucket,
	}
}

func (h *Handler) Home(c *gin.Context) {
	c.JSON(http.StatusOK, "OK")
}

func (h *Handler) Authorize(c *gin.Context) {
	login := c.DefaultQuery("login", "testlogin")
	password := c.DefaultQuery("password", "testpassword")
	ip := c.DefaultQuery("ip", "129.25.10.0")

	res, err := h.Storage.CheckIfIPInBlackList(ip)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if res {
		c.JSON(http.StatusForbidden, "Your IP is in blacklist")
		return
	}

	res, err = h.Storage.CheckIfIPInWhiteList(ip)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if res {
		c.JSON(http.StatusOK, "OK")
		return
	}

	result := h.LeakyBucket.Add(bucket.AuthorizeInput{
		Login:    login,
		Password: password,
		IP:       ip,
	})

	if !result {
		c.JSON(http.StatusTooManyRequests, "Too many requests")
		return
	}

	c.JSON(http.StatusOK, "OK")
}

func (h *Handler) ResetBucket(c *gin.Context) {
	h.LeakyBucket.ResetResetBucketTicker()
	c.JSON(http.StatusNoContent, "")
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
