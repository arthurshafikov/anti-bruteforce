package handler

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

func (h *Handler) InitRoutes(engine *gin.Engine) {
	engine.GET("/", h.Home)

	h.initAuthRoutes(engine)
	h.initBlacklistRoutes(engine)
	h.initWhitelistRoutes(engine)
	h.initBucketRoutes(engine)
}

func (h *Handler) Home(c *gin.Context) {
	h.setOkJSONResponse(c)
}

func (h *Handler) getSubnetInput(c *gin.Context) (core.SubnetInput, error) {
	var subnetInput core.SubnetInput

	if err := c.ShouldBindJSON(&subnetInput); err != nil {
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
