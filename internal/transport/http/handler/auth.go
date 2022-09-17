package handler

import (
	"net/http"

	"github.com/arthurshafikov/anti-bruteforce/internal/core"
	"github.com/gin-gonic/gin"
)

func (h *Handler) Authorize(c *gin.Context) {
	var authInput core.AuthorizeInput

	if err := c.ShouldBindJSON(&authInput); err != nil {
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
