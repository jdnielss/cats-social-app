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
		println(err)
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	auth, err := c.uc.Login(payload)
	if err != nil {
		println(err)
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	// Send response
	common.SendListResponse(ctx, "Ok", auth)
}

func (e *authController) Route() {
	e.rg.POST("/login", e.LoginHandler)
}

func NewAuthController(uc usecase.AuthUseCase, rg *gin.RouterGroup) *authController {
	return &authController{uc: uc, rg: rg}
}

