package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
	"jdnielss.dev/cats-social-app/model/dto"
	"jdnielss.dev/cats-social-app/usecase"
	"jdnielss.dev/cats-social-app/utils/common"
)

var validate = validator.New()

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

func (e *catController) createHandler(ctx *gin.Context) {
	var payload dto.CatRequestDTO

	// Bind the JSON payload to the struct
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	// Validate the struct
	if valid, errs := validateStruct(payload); !valid {
		// Values not valid, deal with errors here
		common.SendErrorResponse(ctx, http.StatusBadRequest, fmt.Sprintf("validation failed: %v", errs))
		return
	}

	payloadResponse, err := e.uc.Create(payload)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	common.SendCreateResponse(ctx, "Ok", payloadResponse)
}

func validateStruct(payload dto.CatRequestDTO) (bool, []string) {
	var validationErrors []string
	// Perform your custom validation logic here
	// For example, if you're using the validator package, you can do:
	if err := validate.Struct(payload); err != nil {
		// Validation failed, append the error message to the slice
		validationErrors = append(validationErrors, err.Error())
	}

	// Check if there are any validation errors
	if len(validationErrors) > 0 {
		// Validation failed, return false and the validation errors
		return false, validationErrors
	}

	// Validation passed, return true and an empty slice of errors
	return true, nil
}

func (e *catController) Route() {
	e.rg.GET("/cat", e.getHandler)
	e.rg.POST("/cat", e.createHandler)
}

func NewCatController(uc usecase.CatUseCase, rg *gin.RouterGroup) *catController {
	return &catController{uc: uc, rg: rg}
}
