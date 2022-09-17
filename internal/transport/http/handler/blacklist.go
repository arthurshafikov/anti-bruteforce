package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) AddToBlacklist(c *gin.Context) {
	subnetInput, err := h.getSubnetInput(c)
	if err != nil {
		return
	}

	if err = h.services.Blacklist.AddToBlacklist(subnetInput); err != nil {
		h.setUnprocessableEntityJSONResponse(c, err.Error())
		return
	}

	h.setJSONResponse(c, http.StatusCreated, OkResponseMessage)
}

func (h *Handler) RemoveFromBlacklist(c *gin.Context) {
	subnetInput, err := h.getSubnetInput(c)
	if err != nil {
		return
	}

	if err = h.services.Blacklist.RemoveFromBlacklist(subnetInput); err != nil {
		h.setUnprocessableEntityJSONResponse(c, err.Error())
		return
	}

	h.setOkJSONResponse(c)
}
