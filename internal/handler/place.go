package handler

import (
	"github.com/gin-gonic/gin"
	"kursachDB/internal/domain/models"
	"kursachDB/internal/handler/responses"
	"kursachDB/internal/services"
	"net/http"
	"strconv"
)

const (
	ErrInvalidArguments = "invalid arguments"
	ErrAlreadyExists    = "already exists"
	ErrRecordNotFound   = "record not found"
	ErrInternalServer   = "internal server error"
)

type Place interface {
	Add(place services.AddPlace) error
	Delete(id int) error
	GetAll() []models.Place
}

// @Summary AddPlace
// @Tags place
// @Description Add place
// @ID add-place
// @Accept json
// @Produce json
// @Router /api/place/add [post]
func (h *Handler) AddPlace(ctx *gin.Context) {
	var input services.AddPlace
	if err := ctx.ShouldBindJSON(&input); err != nil {
		responses.NewErrorResponse(ctx, http.StatusBadRequest, ErrInvalidArguments)
		return
	}

	err := validateAddPlace(input)
	if err != nil {
		responses.NewErrorResponse(ctx, http.StatusBadRequest, ErrInvalidArguments)
		return
	}

	err = h.services.Place.Add(input)
	if err != nil {
		responses.NewErrorResponse(ctx, http.StatusInternalServerError, ErrInternalServer)
		return
	}

	ctx.JSON(http.StatusOK,
		gin.H{
			"status": "success",
		},
	)
}

func validateAddPlace(place services.AddPlace) error {
	panic("implement me")
}

// @Summary DeletePlace
// @Tags place
// @Description Delete place
// @ID delete-place
// @Accept json
// @Produce json
// @Router /api/place/delete [delete]
func (h *Handler) DeletePlace(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Query("id"))
	if err != nil {
		responses.NewErrorResponse(ctx, http.StatusBadRequest, ErrInvalidArguments)
		return
	}

	err = h.services.Place.Delete(id)
	if err != nil {
		responses.NewErrorResponse(ctx, http.StatusInternalServerError, ErrInternalServer)
		return
	}

	ctx.JSON(http.StatusOK,
		gin.H{
			"status": "success",
		},
	)
}

// @Summary GetAllPlaces
// @Tags place
// @Description Get all places
// @ID get-all-places
// @Accept json
// @Produce json
// @Router /api/place/ [get]
func (h *Handler) GetAllPlaces(ctx *gin.Context) {
	res := h.services.Place.GetAll()
	ctx.JSON(http.StatusOK, res)
}
