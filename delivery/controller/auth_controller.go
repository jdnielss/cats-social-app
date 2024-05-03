package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"jdnielss.dev/cats-social-app/model/dto"
	"jdnielss.dev/cats-social-app/usecase"
	"jdnielss.dev/cats-social-app/utils/common"
)

type authController struct {
	uc usecase.AuthUseCase
	rg *gin.RouterGroup
}

func (c *authController) LoginHandler(ctx *gin.Context) {
	var payload dto.LoginRequestDTO
	if err := ctx.BindJSON(&payload); err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	auth, err := c.uc.Login(payload)
	if err != nil {
		var statusCode int
		errorMessage := err.Error()
		switch errorMessage {
		case "INVALID_INPUT":
			statusCode = http.StatusBadRequest
		case "NOT_FOUND":
			statusCode = http.StatusNotFound
		default:
			statusCode = http.StatusInternalServerError
		}
		common.SendErrorResponse(ctx, statusCode, errorMessage)
		return
	}

	// Send response
	common.SendListResponse(ctx, "User logged successfully", auth)
}

func (e *authController) Route() {
	e.rg.POST("/login", e.LoginHandler)
}

func NewAuthController(uc usecase.AuthUseCase, rg *gin.RouterGroup) *authController {
	return &authController{uc: uc, rg: rg}
}
