package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thewolf27/anti-bruteforce/internal/models"
)

const (
	// QuerySubnetParam is a subnet request parameter key.
	QuerySubnetParam = "subnet"

	// OkResponseMessage is a response ok message.
	OkResponseMessage = "OK"

	// WrongSubnetErrorMessage is a request wrong subnet format error.
	WrongSubnetErrorMessage = "wrong subnet format"

	// WrongAuthorizeInputErrorMessage is a request login, password or ip missing format error.
	WrongAuthorizeInputErrorMessage = "login, password and ip field are required"

	// TooManyRequestsErrorMessage is a 429 text error.
	TooManyRequestsErrorMessage = "too many requests"
)

type ServerResponse struct {
	Data string `json:"data"`
}

type App interface {
	Authorize(models.AuthorizeInput) bool
	ResetBucket()
	AddToWhiteList(models.SubnetInput) error
	AddToBlackList(models.SubnetInput) error
	RemoveFromWhiteList(models.SubnetInput) error
	RemoveFromBlackList(models.SubnetInput) error
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
	c.JSON(http.StatusOK, ServerResponse{OkResponseMessage})
}

func (h *Handler) Authorize(c *gin.Context) {
	var authInput models.AuthorizeInput
	err := c.ShouldBindJSON(&authInput)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, ServerResponse{WrongAuthorizeInputErrorMessage})
		return
	}

	res := h.App.Authorize(authInput)

	if res {
		c.JSON(http.StatusOK, ServerResponse{OkResponseMessage})
		return
	}

	c.JSON(http.StatusTooManyRequests, ServerResponse{TooManyRequestsErrorMessage})
}

func (h *Handler) ResetBucket(c *gin.Context) {
	h.App.ResetBucket()
	c.JSON(http.StatusOK, ServerResponse{OkResponseMessage})
}

func (h *Handler) AddToWhiteList(c *gin.Context) {
	subnetInput, err := h.getSubnetInput(c)
	if err != nil {
		return
	}

	err = h.App.AddToWhiteList(subnetInput)
	if err != nil {
		h.setWrongSubnetErrorMessageResponse(c, err)
		return
	}

	c.JSON(http.StatusCreated, ServerResponse{OkResponseMessage})
}

func (h *Handler) AddToBlackList(c *gin.Context) {
	subnetInput, err := h.getSubnetInput(c)
	if err != nil {
		return
	}

	err = h.App.AddToBlackList(subnetInput)
	if err != nil {
		h.setWrongSubnetErrorMessageResponse(c, err)
		return
	}

	c.JSON(http.StatusCreated, ServerResponse{OkResponseMessage})
}

func (h *Handler) RemoveFromWhiteList(c *gin.Context) {
	subnetInput, err := h.getSubnetInput(c)
	if err != nil {
		return
	}

	err = h.App.RemoveFromWhiteList(subnetInput)
	if err != nil {
		h.setWrongSubnetErrorMessageResponse(c, err)
		return
	}

	c.JSON(http.StatusOK, ServerResponse{OkResponseMessage})
}

func (h *Handler) RemoveFromBlackList(c *gin.Context) {
	subnetInput, err := h.getSubnetInput(c)
	if err != nil {
		return
	}

	err = h.App.RemoveFromBlackList(subnetInput)
	if err != nil {
		h.setWrongSubnetErrorMessageResponse(c, err)
		return
	}

	c.JSON(http.StatusOK, ServerResponse{OkResponseMessage})
}

func (h *Handler) getSubnetInput(c *gin.Context) (models.SubnetInput, error) {
	var subnetInput models.SubnetInput
	err := c.ShouldBindJSON(&subnetInput)
	if err != nil {
		h.setWrongSubnetErrorMessageResponse(c, err)
		return models.SubnetInput{}, err
	}

	return subnetInput, nil
}

func (h *Handler) setWrongSubnetErrorMessageResponse(c *gin.Context, err error) {
	c.JSON(http.StatusUnprocessableEntity, ServerResponse{
		Data: WrongSubnetErrorMessage,
	})
}
