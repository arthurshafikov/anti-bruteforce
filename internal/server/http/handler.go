package http

import (
	"net/http"

	"github.com/arthurshafikov/anti-bruteforce/internal/models"
	"github.com/gin-gonic/gin"
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
	AddToWhitelist(models.SubnetInput) error
	AddToBlacklist(models.SubnetInput) error
	RemoveFromWhitelist(models.SubnetInput) error
	RemoveFromBlacklist(models.SubnetInput) error
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
	h.setOkJSONResponse(c)
}

func (h *Handler) Authorize(c *gin.Context) {
	var authInput models.AuthorizeInput
	err := c.ShouldBindJSON(&authInput)
	if err != nil {
		h.setUnprocessableEntityJSONResponse(c, WrongAuthorizeInputErrorMessage)
		return
	}

	res := h.App.Authorize(authInput)

	if res {
		h.setOkJSONResponse(c)
		return
	}

	h.setJSONResponse(c, http.StatusTooManyRequests, TooManyRequestsErrorMessage)
}

func (h *Handler) ResetBucket(c *gin.Context) {
	h.App.ResetBucket()
	c.JSON(http.StatusOK, ServerResponse{OkResponseMessage})
}

func (h *Handler) AddToWhitelist(c *gin.Context) {
	subnetInput, err := h.getSubnetInput(c)
	if err != nil {
		return
	}

	err = h.App.AddToWhitelist(subnetInput)
	if err != nil {
		h.setUnprocessableEntityJSONResponse(c, err.Error())
		return
	}

	h.setJSONResponse(c, http.StatusCreated, OkResponseMessage)
}

func (h *Handler) AddToBlacklist(c *gin.Context) {
	subnetInput, err := h.getSubnetInput(c)
	if err != nil {
		return
	}

	err = h.App.AddToBlacklist(subnetInput)
	if err != nil {
		h.setUnprocessableEntityJSONResponse(c, err.Error())
		return
	}

	h.setJSONResponse(c, http.StatusCreated, OkResponseMessage)
}

func (h *Handler) RemoveFromWhitelist(c *gin.Context) {
	subnetInput, err := h.getSubnetInput(c)
	if err != nil {
		return
	}

	err = h.App.RemoveFromWhitelist(subnetInput)
	if err != nil {
		h.setUnprocessableEntityJSONResponse(c, err.Error())
		return
	}

	h.setOkJSONResponse(c)
}

func (h *Handler) RemoveFromBlacklist(c *gin.Context) {
	subnetInput, err := h.getSubnetInput(c)
	if err != nil {
		return
	}

	err = h.App.RemoveFromBlacklist(subnetInput)
	if err != nil {
		h.setUnprocessableEntityJSONResponse(c, err.Error())
		return
	}

	h.setOkJSONResponse(c)
}

func (h *Handler) getSubnetInput(c *gin.Context) (models.SubnetInput, error) {
	var subnetInput models.SubnetInput
	err := c.ShouldBindJSON(&subnetInput)
	if err != nil {
		h.setUnprocessableEntityJSONResponse(c, err.Error())
		return models.SubnetInput{}, err
	}

	return subnetInput, nil
}

func (h *Handler) setUnprocessableEntityJSONResponse(c *gin.Context, data string) {
	h.setJSONResponse(c, http.StatusUnprocessableEntity, data)
}

func (h *Handler) setOkJSONResponse(c *gin.Context) {
	h.setJSONResponse(c, http.StatusOK, OkResponseMessage)
}

func (h *Handler) setJSONResponse(c *gin.Context, code int, data string) {
	c.JSON(code, ServerResponse{
		Data: data,
	})
}
