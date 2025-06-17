package handlers

import (
	"context"
	"net/http"

	"github.com/adrianozp/go-plateful/app/place/entities"

	"github.com/gin-gonic/gin"
)

type Creator interface {
	Create(ctx context.Context, place entities.Place) (string, error)
}

type CreatePlaceHandler struct {
	creator Creator
}

func NewCreatePlaceHandler(creator Creator) CreatePlaceHandler {
	return CreatePlaceHandler{
		creator: creator,
	}
}

func (h CreatePlaceHandler) Create(c *gin.Context) {
	var payload dto
	if err := c.Bind(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.creator.Create(c, payload.toEntity())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}
func RegisterCreatePlaceRoutes(router *gin.Engine, handler CreatePlaceHandler) {
	router.POST("/places", handler.Create)
}
