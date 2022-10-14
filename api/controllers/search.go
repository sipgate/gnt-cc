package controllers

import (
	"github.com/gin-gonic/gin"
)

type SearchController struct {
	SearchService searchService
}

// Search godoc
// @Summary Search in all known resources
// @Description ...
// @Produce json
// @Success 200 {object} model.SearchResultsResponse
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /search [get]
// TODO: Add Tests
func (controller *SearchController) Search(c *gin.Context) {
	query, exists := c.GetQuery("query")
	if !exists {
		c.AbortWithStatusJSON(400, createErrorBody("query parameter is required"))
		return
	}

	if len(query) == 0 {
		c.AbortWithStatusJSON(400, createErrorBody("query parameter must not be empty"))
		return
	}

	results := controller.SearchService.Search(query)

	c.JSON(200, results)
}
