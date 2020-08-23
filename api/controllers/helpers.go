package controllers

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gnt-cc/model"
)

func createErrorBody(msg string) model.ErrorResponse {
	return model.ErrorResponse{Message: msg}
}

func createInternalServerErrorBody() model.ErrorResponse {
	return createErrorBody(MsgInternalServerError)
}

func abortWithInternalServerError(c *gin.Context, err error) {
	log.Error(err)
	c.AbortWithStatusJSON(500, createInternalServerErrorBody())
}
