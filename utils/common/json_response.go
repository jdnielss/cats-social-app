package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
	modelutil "jdnielss.dev/cats-social-app/utils/model_util"
)

func SendCreateResponse(ctx *gin.Context, description string, data any) {
	ctx.JSON(http.StatusCreated, &modelutil.SingleResponse{
		Status: modelutil.Status{
			Code:        http.StatusCreated,
			Description: description,
		},
		Data: data,
	})
}

func SendSingleResponse(ctx *gin.Context, description string, data any) {
	ctx.JSON(http.StatusOK, &modelutil.SingleResponse{
		Status: modelutil.Status{
			Code:        http.StatusOK,
			Description: description,
		},
		Data: data,
	})
}

func SendErrorResponse(ctx *gin.Context, code int, description string) {
	ctx.AbortWithStatusJSON(code, &modelutil.Status{
		Code:        code,
		Description: description,
	})
}