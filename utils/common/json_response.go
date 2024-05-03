package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
	modelutil "jdnielss.dev/cats-social-app/utils/model_util"
)

func SendCreateResponse(ctx *gin.Context, description string, data any) {
	ctx.JSON(http.StatusOK, &modelutil.SingleResponse{
		Message: `Success`,
		Data:    data,
	})
}

func SendSingleResponse(ctx *gin.Context, description string, data any) {
	ctx.JSON(http.StatusOK, &modelutil.SingleResponse{
		Message: description,
		Data:    data,
	})
}

func SendListResponse(ctx *gin.Context, description string, data interface{}) {
	// Type assert data to []interface{} if it's a slice
	var responseData []interface{}
	if dataList, ok := data.([]interface{}); ok {
		responseData = dataList
	} else {
		// Handle the case where data is not a slice
		// For example, you might construct a slice with a single element
		responseData = []interface{}{data}
	}

	// Construct and send the response
	ctx.JSON(http.StatusOK, &modelutil.PagedResponse{
		Message: description,
		Data:    responseData,
	})
}

func SendErrorResponse(ctx *gin.Context, code int, description string) {
	ctx.AbortWithStatusJSON(code, &modelutil.Status{
		Code:        code,
		Description: description,
	})
}
