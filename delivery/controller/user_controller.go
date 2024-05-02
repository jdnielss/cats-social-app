package controller

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"jdnielss.dev/cats-social-app/model"
	"jdnielss.dev/cats-social-app/usecase"
	"jdnielss.dev/cats-social-app/utils/common"
)

type userController struct {
	uc usecase.UserUseCase
	rg *gin.RouterGroup
}

func (a *userController) Register(ctx *gin.Context) {
	var payload model.User
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	newRsp, err := a.uc.Register(payload)
	if err != nil {
		if strings.Contains(err.Error(), "email already registered") {
			common.SendErrorResponse(ctx, http.StatusConflict, err.Error()) // Return 409 Conflict status code
		} else {
			common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		}
		return
	}

	common.SendCreateResponse(ctx, "Ok", newRsp)
}

func (a *userController) Route() {
	a.rg.POST("/user/register", a.Register)

}

func NewUserController(uc usecase.UserUseCase, rg *gin.RouterGroup) *userController {
	return &userController{uc: uc, rg: rg}
}
