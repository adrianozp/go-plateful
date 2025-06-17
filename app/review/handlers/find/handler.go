package handlers

import (
	"context"
	"net/http"

	"github.com/adrianozp/go-plateful/app/review/entities"

	"github.com/gin-gonic/gin"
)

type Finder interface {
	FindByID(ctx context.Context, id string) (entities.Review, error)
	FindByFilter(ctxx context.Context, filter entities.ReviewFilter) (entities.Reviews, error)
}

type FindReviewHandler struct {
	finder Finder
}

func NewFindReviewHandler(finder Finder) FindReviewHandler {
	return FindReviewHandler{
		finder: finder,
	}
}

func (h FindReviewHandler) GetById(c *gin.Context) {
	id := c.Param("id")
	review, err := h.finder.FindByID(c, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Review not found"})
		return
	}
	c.JSON(http.StatusOK, toDTO(review))
}

func (h FindReviewHandler) FindByFilter(c *gin.Context) {
	var payload filterDTO
	if err := c.Bind(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	reviews, err := h.finder.FindByFilter(c, payload.toEntity())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"reviews": toReviewsDTO(reviews)})
}

func RegisterFindReviewRoutes(router *gin.Engine, handler FindReviewHandler) {
	router.GET("/reviews/:id", handler.GetById)
	router.GET("/reviews", handler.FindByFilter)
}
