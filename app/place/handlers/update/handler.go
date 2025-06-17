package handlers

import (
	"context"
	"net/http"

	"github.com/adrianozp/go-plateful/app/place/entities"

	"github.com/gin-gonic/gin"
)

type Updater interface {
	Update(ctx context.Context, id string, place entities.Place) error
}

type UpdatePlaceHandler struct {
	updater Updater
}

func NewUpdatePlaceHandler(updater Updater) UpdatePlaceHandler {
	return UpdatePlaceHandler{
		updater: updater,
	}
}

func (h UpdatePlaceHandler) Update(c *gin.Context) {
	var payload dto
	if err := c.Bind(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")
	if err := h.updater.Update(c, id, payload.toEntity()); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func RegisterUpdatePlaceRoutes(router *gin.Engine, handler UpdatePlaceHandler) {
	router.PUT("/places/:id", handler.Update)
}
