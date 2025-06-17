package handlers

import (
	"context"
	"net/http"

	"github.com/adrianozp/go-plateful/app/review/entities"

	"github.com/gin-gonic/gin"
)

type Updater interface {
	Update(ctx context.Context, id string, review entities.Review) error
}

type UpdateReviewHandler struct {
	updater Updater
}

func NewUpdateReviewHandler(updater Updater) UpdateReviewHandler {
	return UpdateReviewHandler{
		updater: updater,
	}
}

func (h UpdateReviewHandler) Update(c *gin.Context) {
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

func RegisterUpdateReviewRoutes(router *gin.Engine, handler UpdateReviewHandler) {
	router.PUT("/reviews/:id", handler.Update)
}
