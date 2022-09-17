package http

import (
	"net/http"

	"github.com/arthurshafikov/anti-bruteforce/internal/core"
	"github.com/arthurshafikov/anti-bruteforce/internal/services"
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

type Handler struct {
	services *services.Services
}

func NewHandler(services *services.Services) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) Home(c *gin.Context) {
	h.setOkJSONResponse(c)
}

func (h *Handler) Authorize(c *gin.Context) {
	var authInput core.AuthorizeInput
	err := c.ShouldBindJSON(&authInput)
	if err != nil {
		h.setUnprocessableEntityJSONResponse(c, WrongAuthorizeInputErrorMessage)
		return
	}

	res := h.services.Auth.Authorize(authInput)

	if res {
		h.setOkJSONResponse(c)
		return
	}

	h.setJSONResponse(c, http.StatusTooManyRequests, TooManyRequestsErrorMessage)
}

func (h *Handler) ResetBucket(c *gin.Context) {
	h.services.Bucket.ResetBucket()
	c.JSON(http.StatusOK, ServerResponse{OkResponseMessage})
}

func (h *Handler) AddToWhitelist(c *gin.Context) {
	subnetInput, err := h.getSubnetInput(c)
	if err != nil {
		return
	}

	err = h.services.Whitelist.AddToWhitelist(subnetInput)
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

	err = h.services.Blacklist.AddToBlacklist(subnetInput)
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

	err = h.services.Whitelist.RemoveFromWhitelist(subnetInput)
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

	err = h.services.Blacklist.RemoveFromBlacklist(subnetInput)
	if err != nil {
		h.setUnprocessableEntityJSONResponse(c, err.Error())
		return
	}

	h.setOkJSONResponse(c)
}

func (h *Handler) getSubnetInput(c *gin.Context) (core.SubnetInput, error) {
	var subnetInput core.SubnetInput
	err := c.ShouldBindJSON(&subnetInput)
	if err != nil {
		h.setUnprocessableEntityJSONResponse(c, err.Error())
		return core.SubnetInput{}, err
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
