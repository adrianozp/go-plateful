package handlers

import (
	"context"
	"net/http"

	"github.com/adrianozp/go-plateful/app/review/entities"

	"github.com/gin-gonic/gin"
)

type Creator interface {
	Create(ctx context.Context, review entities.Review) (string, error)
}

type CreateReviewHandler struct {
	creator Creator
}

func NewCreateReviewHandler(creator Creator) CreateReviewHandler {
	return CreateReviewHandler{
		creator: creator,
	}
}

func (h CreateReviewHandler) Create(c *gin.Context) {
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
func RegisterCreateReviewRoutes(router *gin.Engine, handler CreateReviewHandler) {
	router.POST("/reviews", handler.Create)
}
