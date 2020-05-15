package httputil

import "github.com/gin-gonic/gin"

func NewError(ctx *gin.Context, status int, err error) {
	er := HTTPError{
		Code:    status,
		Message: err.Error(),
	}
	ctx.JSON(status, er)
}

type HTTPError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type HTTPError400 struct {
	HTTPError
	Code int `json:"code" example:"400"`
}

type HTTPError404 struct {
	HTTPError
	Code int `json:"code" example:"404"`
}

type HTTPError502 struct {
	HTTPError
	Code int `json:"code" example:"502"`
}
