package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"jdnielss.dev/cats-social-app/usecase"
	"jdnielss.dev/cats-social-app/utils/common"
)

type catController struct {
	uc usecase.CatUseCase
	rg *gin.RouterGroup
}

func (c *catController) getHandler(ctx *gin.Context) {
	q := ctx.Request.URL.Query()

	queryString := q.Encode()

	// Call the repository method to get cats
	cat, err := c.uc.Find(queryString)
	if err != nil {
		println(err)
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	// Send response
	common.SendListResponse(ctx, "Ok", cat)
}

func (e *catController) Route() {
	e.rg.GET("/cat", e.getHandler)
}

func NewCatController(uc usecase.CatUseCase, rg *gin.RouterGroup) *catController {
	return &catController{uc: uc, rg: rg}
}
