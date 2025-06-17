package handlers

import (
	"context"
	"net/http"

	"github.com/adrianozp/go-plateful/app/place/entities"

	"github.com/gin-gonic/gin"
)

type Finder interface {
	Find(ctx context.Context, id string) (entities.Place, error)
}

type FindPlaceHandler struct {
	finder Finder
}

func NewFindPlaceHandler(finder Finder) FindPlaceHandler {
	return FindPlaceHandler{
		finder: finder,
	}
}

func (h FindPlaceHandler) GetById(c *gin.Context) {
	id := c.Param("id")
	place, err := h.finder.Find(c, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Place not found"})
		return
	}
	c.JSON(http.StatusOK, toDTO(place))
}

func RegisterFindPlaceRoutes(router *gin.Engine, handler FindPlaceHandler) {
	router.GET("/places/:id", handler.GetById)
}
