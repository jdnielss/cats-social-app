package controller

import (
	"net/http"

	"enigmacamp.com/be-lms-university/model/dto"
	"enigmacamp.com/be-lms-university/usecase"
	"enigmacamp.com/be-lms-university/utils/common"
	"github.com/gin-gonic/gin"
)

type enrollmentController struct {
	uc usecase.EnrollmentUseCase
	rg *gin.RouterGroup
}

func (e *enrollmentController) createHandler(ctx *gin.Context) {
	var payload dto.EnrollmentRequestDto
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	payloadResponse, err := e.uc.RegisterNewEnrollment(payload)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	common.SendCreateResponse(ctx, "Ok", payloadResponse)
}

func (e *enrollmentController) getHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	enrollment, err := e.uc.FindById(id)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusNotFound, err.Error())
		return
	}
	common.SendSingleResponse(ctx, "Ok", enrollment)

}

func (e *enrollmentController) Route() {
	e.rg.POST("/enrollments", e.createHandler)
	e.rg.GET("/enrollments", e.createHandler)
	e.rg.GET("/enrollments/:id", e.getHandler)
}

func NewEnrollmentController(uc usecase.EnrollmentUseCase, rg *gin.RouterGroup) *enrollmentController {
	return &enrollmentController{uc: uc, rg: rg}
}
