package handler

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"kursachDB/internal/domain/models"
	"kursachDB/internal/handler/responses"
	"kursachDB/internal/services"
	"kursachDB/internal/services/place"
	"net/http"
	"regexp"
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
	GetAll() ([]models.Place, error)
}

// @Summary AddPlace
// @Tags place
// @Description Add place
// @ID add-place
// @Accept json
// @Produce json
// @Param input body services.AddPlace true "Place info for add"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/place/add [post]
func (h *Handler) AddPlace(ctx *gin.Context) {
	var input services.AddPlace
	if err := ctx.ShouldBindJSON(&input); err != nil {
		responses.NewErrorResponse(ctx, http.StatusBadRequest, ErrInvalidArguments)
		return
	}

	err := validateAddPlace(input)
	if err != nil {
		responses.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
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

func validateCityName(fl validator.FieldLevel) bool {
	regex := regexp.MustCompile(`^([А-ЯЁ][а-яё]*)([ -][А-ЯЁ][а-яё]*)*$`)
	return regex.MatchString(fl.Field().String())
}

func validateAddPlace(place services.AddPlace) error {
	valid := validator.New()
	err := valid.RegisterValidation("city_name", validateCityName)
	if err != nil {
		return err
	}

	if valid.Struct(&place) != nil {
		return fmt.Errorf("wrong name place")
	}
	return nil
}

// @Summary DeletePlace
// @Tags place
// @Description Delete place
// @ID delete-place
// @Accept json
// @Produce json
// @Param id path int64 true "ID of the place"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/place/{id} [delete]
func (h *Handler) DeletePlace(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Query("id"))
	if err != nil {
		responses.NewErrorResponse(ctx, http.StatusBadRequest, ErrInvalidArguments)
		return
	}

	err = h.services.Place.Delete(id)
	if err != nil {
		if errors.Is(err, place.ErrPlaceNotFound) {
			responses.NewErrorResponse(ctx, http.StatusNotFound, ErrRecordNotFound)
			return
		}
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
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/place/ [get]
func (h *Handler) GetAllPlaces(ctx *gin.Context) {
	res, err := h.services.Place.GetAll()
	if err != nil {
		responses.NewErrorResponse(ctx, http.StatusInternalServerError, ErrInternalServer)
		return
	}

	ctx.JSON(http.StatusOK, res)
}
