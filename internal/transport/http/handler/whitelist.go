package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) AddToWhitelist(c *gin.Context) {
	subnetInput, err := h.getSubnetInput(c)
	if err != nil {
		return
	}

	if err = h.services.Whitelist.AddToWhitelist(subnetInput); err != nil {
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

	if err = h.services.Whitelist.RemoveFromWhitelist(subnetInput); err != nil {
		h.setUnprocessableEntityJSONResponse(c, err.Error())
		return
	}

	h.setOkJSONResponse(c)
}
